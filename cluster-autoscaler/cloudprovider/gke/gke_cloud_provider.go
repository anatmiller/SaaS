/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gke

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/klog"
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/gce"
	"k8s.io/autoscaler/cluster-autoscaler/config"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	"k8s.io/autoscaler/cluster-autoscaler/utils/gpu"
)

const (
	// ProviderNameGKE is the name of GKE cloud provider.
	ProviderNameGKE = "gke"
)

const (
	maxAutoprovisionedSize = 1000
	minAutoprovisionedSize = 0
)

// Big machines are temporarily commented out.
// TODO(mwielgus): get this list programmatically
var autoprovisionedMachineTypes = []string{
	"n1-standard-1",
	"n1-standard-2",
	"n1-standard-4",
	"n1-standard-8",
	"n1-standard-16",
	//"n1-standard-32",
	//"n1-standard-64",
	"n1-highcpu-2",
	"n1-highcpu-4",
	"n1-highcpu-8",
	"n1-highcpu-16",
	//"n1-highcpu-32",
	// "n1-highcpu-64",
	"n1-highmem-2",
	"n1-highmem-4",
	"n1-highmem-8",
	"n1-highmem-16",
	//"n1-highmem-32",
	//"n1-highmem-64",
}

// GkeCloudProvider implements CloudProvider interface.
type GkeCloudProvider struct {
	gkeManager GkeManager
	// This resource limiter is used if resource limits are not defined through cloud API.
	resourceLimiterFromFlags *cloudprovider.ResourceLimiter
}

// BuildGkeCloudProvider builds CloudProvider implementation for GKE.
func BuildGkeCloudProvider(gkeManager GkeManager, resourceLimiter *cloudprovider.ResourceLimiter) (*GkeCloudProvider, error) {
	return &GkeCloudProvider{gkeManager: gkeManager, resourceLimiterFromFlags: resourceLimiter}, nil
}

// Cleanup cleans up all resources before the cloud provider is removed
func (gke *GkeCloudProvider) Cleanup(ctx context.Context) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeCloudProvider.Cleanup")
	defer span.Finish()

	gke.gkeManager.Cleanup(ctx)
	return nil
}

// Name returns name of the cloud provider.
func (gke *GkeCloudProvider) Name(ctx context.Context) string {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeCloudProvider.Name")
	defer span.Finish()

	return ProviderNameGKE
}

// NodeGroups returns all node groups configured for this cloud provider.
func (gke *GkeCloudProvider) NodeGroups(ctx context.Context) []cloudprovider.NodeGroup {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeCloudProvider.NodeGroups")
	defer span.Finish()

	migs := gke.gkeManager.GetMigs(ctx)
	result := make([]cloudprovider.NodeGroup, 0, len(migs))
	for _, mig := range migs {
		result = append(result, mig.Config)
	}
	return result
}

// NodeGroupForNode returns the node group for the given node.
func (gke *GkeCloudProvider) NodeGroupForNode(ctx context.Context, node *apiv1.Node) (cloudprovider.NodeGroup, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeCloudProvider.NodeGroupForNode")
	defer span.Finish()

	ref, err := gce.GceRefFromProviderId(node.Spec.ProviderID)
	if err != nil {
		return nil, err
	}
	mig, err := gke.gkeManager.GetMigForInstance(ctx, ref)
	return mig, err
}

// Pricing returns pricing model for this cloud provider or error if not available.
func (gke *GkeCloudProvider) Pricing(ctx context.Context) (cloudprovider.PricingModel, errors.AutoscalerError) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeCloudProvider.Pricing")
	defer span.Finish()

	return &gce.GcePriceModel{}, nil
}

// GetAvailableMachineTypes get all machine types that can be requested from the cloud provider.
func (gke *GkeCloudProvider) GetAvailableMachineTypes(ctx context.Context) ([]string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeCloudProvider.GetAvailableMachineTypes")
	defer span.Finish()

	return autoprovisionedMachineTypes, nil
}

// NewNodeGroup builds a theoretical node group based on the node definition provided. The node group is not automatically
// created on the cloud provider side. The node group is not returned by NodeGroups() until it is created.
func (gke *GkeCloudProvider) NewNodeGroup(ctx context.Context, machineType string, labels map[string]string, systemLabels map[string]string,
	taints []apiv1.Taint, extraResources map[string]resource.Quantity) (cloudprovider.NodeGroup, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeCloudProvider.NewNodeGroup")
	defer span.Finish()

	nodePoolName := fmt.Sprintf("%s-%s-%d", nodeAutoprovisioningPrefix, machineType, time.Now().Unix())
	zone, found := systemLabels[apiv1.LabelZoneFailureDomain]
	if !found {
		return nil, cloudprovider.ErrIllegalConfiguration
	}

	if gpuRequest, found := extraResources[gpu.ResourceNvidiaGPU]; found {
		gpuType, found := systemLabels[gpu.GPULabel]
		if !found {
			return nil, cloudprovider.ErrIllegalConfiguration
		}
		gpuCount, err := getNormalizedGpuCount(gpuRequest.Value())
		if err != nil {
			return nil, err
		}
		extraResources[gpu.ResourceNvidiaGPU] = *resource.NewQuantity(gpuCount, resource.DecimalSI)
		err = validateGpuConfig(gpuType, gpuCount, zone, machineType)
		if err != nil {
			return nil, err
		}
		nodePoolName = fmt.Sprintf("%s-%s-gpu-%d", nodeAutoprovisioningPrefix, machineType, time.Now().Unix())
		labels[gpu.GPULabel] = gpuType

		taint := apiv1.Taint{
			Effect: apiv1.TaintEffectNoSchedule,
			Key:    gpu.ResourceNvidiaGPU,
			Value:  "present",
		}
		taints = append(taints, taint)
	}

	mig := &GkeMig{
		gceRef: gce.GceRef{
			Project: gke.gkeManager.GetProjectId(),
			Zone:    zone,
			Name:    nodePoolName + "-temporary-mig",
		},
		gkeManager:      gke.gkeManager,
		autoprovisioned: true,
		exist:           false,
		nodePoolName:    nodePoolName,
		minSize:         minAutoprovisionedSize,
		maxSize:         maxAutoprovisionedSize,
		spec: &MigSpec{
			MachineType:    machineType,
			Labels:         labels,
			Taints:         taints,
			ExtraResources: extraResources,
		},
	}

	// Try to build a node from autoprovisioning spec. We don't need one right now,
	// but if it fails later, we'd end up with a node group we can't scale anyway,
	// so there's no point creating it.
	if _, err := gke.gkeManager.GetMigTemplateNode(ctx, mig); err != nil {
		return nil, fmt.Errorf("failed to build node from spec: %v", err)
	}

	return mig, nil
}

// GetResourceLimiter returns struct containing limits (max, min) for resources (cores, memory etc.).
func (gke *GkeCloudProvider) GetResourceLimiter(ctx context.Context) (*cloudprovider.ResourceLimiter, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeCloudProvider.GetResourceLimiter")
	defer span.Finish()

	resourceLimiter, err := gke.gkeManager.GetResourceLimiter(ctx)
	if err != nil {
		return nil, err
	}
	if resourceLimiter != nil {
		return resourceLimiter, nil
	}
	return gke.resourceLimiterFromFlags, nil
}

// Refresh is called before every main loop and can be used to dynamically update cloud provider state.
// In particular the list of node groups returned by NodeGroups can change as a result of CloudProvider.Refresh(ctx).
func (gke *GkeCloudProvider) Refresh(ctx context.Context) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeCloudProvider.Refresh")
	defer span.Finish()

	return gke.gkeManager.Refresh(ctx)
}

// GetClusterInfo returns the project id, location and cluster name.
func (gke *GkeCloudProvider) GetClusterInfo() (projectId, location, clusterName string) {
	return gke.gkeManager.GetProjectId(), gke.gkeManager.GetLocation(), gke.gkeManager.GetClusterName()
}

// GetNodeLocations returns the list of zones in which the cluster has nodes.
func (gke *GkeCloudProvider) GetNodeLocations() []string {
	return gke.gkeManager.GetNodeLocations()
}

// MigSpec contains information about what machines in a MIG look like.
type MigSpec struct {
	MachineType    string
	Labels         map[string]string
	Taints         []apiv1.Taint
	ExtraResources map[string]resource.Quantity
}

// GkeMig represents the GKE Managed Instance Group implementation of a NodeGroup.
type GkeMig struct {
	gceRef gce.GceRef

	gkeManager      GkeManager
	minSize         int
	maxSize         int
	autoprovisioned bool
	exist           bool
	nodePoolName    string
	spec            *MigSpec
}

// GceRef returns Mig's GceRef
func (mig *GkeMig) GceRef() gce.GceRef {
	return mig.gceRef
}

// NodePoolName returns the name of the GKE node pool this Mig belongs to.
func (mig *GkeMig) NodePoolName() string {
	return mig.nodePoolName
}

// Spec returns specification of the Mig.
func (mig *GkeMig) Spec() *MigSpec {
	return mig.spec
}

// MaxSize returns maximum size of the node group.
func (mig *GkeMig) MaxSize() int {
	return mig.maxSize
}

// MinSize returns minimum size of the node group.
func (mig *GkeMig) MinSize() int {
	return mig.minSize
}

// TargetSize returns the current TARGET size of the node group. It is possible that the
// number is different from the number of nodes registered in Kubernetes.
func (mig *GkeMig) TargetSize(ctx context.Context) (int, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeMig.TargetSize")
	defer span.Finish()

	if !mig.exist {
		return 0, nil
	}
	size, err := mig.gkeManager.GetMigSize(ctx, mig)
	return int(size), err
}

// IncreaseSize increases Mig size
func (mig *GkeMig) IncreaseSize(ctx context.Context, delta int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeMig.IncreaseSize")
	defer span.Finish()

	if delta <= 0 {
		return fmt.Errorf("size increase must be positive")
	}
	size, err := mig.gkeManager.GetMigSize(ctx, mig)
	if err != nil {
		return err
	}
	if int(size)+delta > mig.MaxSize() {
		return fmt.Errorf("size increase too large - desired:%d max:%d", int(size)+delta, mig.MaxSize())
	}
	return mig.gkeManager.SetMigSize(ctx, mig, size+int64(delta))
}

// DecreaseTargetSize decreases the target size of the node group. This function
// doesn't permit to delete any existing node and can be used only to reduce the
// request for new nodes that have not been yet fulfilled. Delta should be negative.
func (mig *GkeMig) DecreaseTargetSize(ctx context.Context, delta int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeMig.DecreaseTargetSize")
	defer span.Finish()

	if delta >= 0 {
		return fmt.Errorf("size decrease must be negative")
	}
	size, err := mig.gkeManager.GetMigSize(ctx, mig)
	if err != nil {
		return err
	}
	nodes, err := mig.gkeManager.GetMigNodes(ctx, mig)
	if err != nil {
		return err
	}
	if int(size)+delta < len(nodes) {
		return fmt.Errorf("attempt to delete existing nodes targetSize:%d delta:%d existingNodes: %d",
			size, delta, len(nodes))
	}
	return mig.gkeManager.SetMigSize(ctx, mig, size+int64(delta))
}

// Belongs returns true if the given node belongs to the NodeGroup.
func (mig *GkeMig) Belongs(ctx context.Context, node *apiv1.Node) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeMig.Belongs")
	defer span.Finish()

	ref, err := gce.GceRefFromProviderId(node.Spec.ProviderID)
	if err != nil {
		return false, err
	}
	targetMig, err := mig.gkeManager.GetMigForInstance(ctx, ref)
	if err != nil {
		return false, err
	}
	if targetMig == nil {
		return false, fmt.Errorf("%s doesn't belong to a known mig", node.Name)
	}
	if targetMig.Id() != mig.Id() {
		return false, nil
	}
	return true, nil
}

// DeleteNodes deletes the nodes from the group.
func (mig *GkeMig) DeleteNodes(ctx context.Context, nodes []*apiv1.Node) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeMig.DeleteNodes")
	defer span.Finish()

	size, err := mig.gkeManager.GetMigSize(ctx, mig)
	if err != nil {
		return err
	}
	if int(size) <= mig.MinSize() {
		return fmt.Errorf("min size reached, nodes will not be deleted")
	}
	refs := make([]*gce.GceRef, 0, len(nodes))
	for _, node := range nodes {

		belongs, err := mig.Belongs(ctx, node)
		if err != nil {
			return err
		}
		if !belongs {
			return fmt.Errorf("%s belong to a different mig than %s", node.Name, mig.Id())
		}
		gceref, err := gce.GceRefFromProviderId(node.Spec.ProviderID)
		if err != nil {
			return err
		}
		refs = append(refs, gceref)
	}
	return mig.gkeManager.DeleteInstances(ctx, refs)
}

// Id returns mig url.
func (mig *GkeMig) Id() string {
	return gce.GenerateMigUrl(mig.gceRef)
}

// Debug returns a debug string for the Mig.
func (mig *GkeMig) Debug() string {
	return fmt.Sprintf("%s (%d:%d)", mig.Id(), mig.MinSize(), mig.MaxSize())
}

// Nodes returns a list of all nodes that belong to this node group.
func (mig *GkeMig) Nodes(ctx context.Context) ([]cloudprovider.Instance, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeMig.Nodes")
	defer span.Finish()

	instanceNames, err := mig.gkeManager.GetMigNodes(ctx, mig)
	if err != nil {
		return nil, err
	}
	instances := make([]cloudprovider.Instance, 0, len(instanceNames))
	for _, instanceName := range instanceNames {
		instances = append(instances, cloudprovider.Instance{Id: instanceName})
	}
	return instances, nil
}

// Exist checks if the node group really exists on the cloud provider side. Allows to tell the
// theoretical node group from the real one.
func (mig *GkeMig) Exist() bool {
	return mig.exist
}

// Create creates the node group on the cloud provider side.
func (mig *GkeMig) Create(ctx context.Context) (cloudprovider.NodeGroup, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeMig.Create")
	defer span.Finish()

	if !mig.exist && mig.autoprovisioned {
		return mig.gkeManager.CreateNodePool(ctx, mig)
	}
	return nil, fmt.Errorf("cannot create non-autoprovisioned node group")
}

// Delete deletes the node group on the cloud provider side.
// This will be executed only for autoprovisioned node groups, once their size drops to 0.
func (mig *GkeMig) Delete(ctx context.Context) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeMig.Delete")
	defer span.Finish()

	if mig.exist && mig.autoprovisioned {
		return mig.gkeManager.DeleteNodePool(ctx, mig)
	}
	return fmt.Errorf("cannot delete non-autoprovisioned node group")
}

// Autoprovisioned returns true if the node group is autoprovisioned.
func (mig *GkeMig) Autoprovisioned() bool {
	return mig.autoprovisioned
}

// TemplateNodeInfo returns a node template for this node group.
func (mig *GkeMig) TemplateNodeInfo(ctx context.Context) (*schedulernodeinfo.NodeInfo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GkeMig.TemplateNodeInfo")
	defer span.Finish()

	node, err := mig.gkeManager.GetMigTemplateNode(ctx, mig)
	if err != nil {
		return nil, err
	}
	nodeInfo := schedulernodeinfo.NewNodeInfo(cloudprovider.BuildKubeProxy(mig.Id()))
	nodeInfo.SetNode(node)
	return nodeInfo, nil
}

// BuildGKE builds a new GKE cloud provider, manager etc.
func BuildGKE(ctx context.Context, opts config.AutoscalingOptions, do cloudprovider.NodeGroupDiscoveryOptions, rl *cloudprovider.ResourceLimiter) cloudprovider.CloudProvider {
	if do.DiscoverySpecified() {
		klog.Fatal("GKE gets nodegroup specification via API, command line specs are not allowed")
	}
	var config io.ReadCloser
	if opts.CloudConfig != "" {
		var err error
		config, err = os.Open(opts.CloudConfig)
		if err != nil {
			klog.Fatalf("Couldn't open cloud provider configuration %s: %#v", opts.CloudConfig, err)
		}
		defer config.Close()
	}

	mode := ModeGKE
	if opts.NodeAutoprovisioningEnabled {
		mode = ModeGKENAP
	}
	manager, err := CreateGkeManager(ctx, config, mode, opts.ClusterName, opts.Regional)
	if err != nil {
		klog.Fatalf("Failed to create GKE Manager: %v", err)
	}

	provider, err := BuildGkeCloudProvider(manager, rl)
	if err != nil {
		klog.Fatalf("Failed to create GKE cloud provider: %v", err)
	}
	// Register GKE & GCE API usage metrics.
	registerMetrics()
	gce.RegisterMetrics()
	return provider
}
