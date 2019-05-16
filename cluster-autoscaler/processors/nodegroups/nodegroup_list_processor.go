/*
Copyright 2018 The Kubernetes Authors.

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

package nodegroups

import (
	"context"

	"github.com/opentracing/opentracing-go"
	apiv1 "k8s.io/api/core/v1"
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	autoscalingcontext "k8s.io/autoscaler/cluster-autoscaler/context"
)

// NodeGroupListProcessor processes lists of NodeGroups considered in scale-up.
type NodeGroupListProcessor interface {
	Process(ctx context.Context, context *autoscalingcontext.AutoscalingContext, nodeGroups []cloudprovider.NodeGroup,
		nodeInfos map[string]*schedulernodeinfo.NodeInfo,
		unschedulablePods []*apiv1.Pod) ([]cloudprovider.NodeGroup, map[string]*schedulernodeinfo.NodeInfo, error)
	CleanUp(ctx context.Context)
}

// NoOpNodeGroupListProcessor is returning pod lists without processing them.
type NoOpNodeGroupListProcessor struct {
}

// NewDefaultNodeGroupListProcessor creates an instance of NodeGroupListProcessor.
func NewDefaultNodeGroupListProcessor() NodeGroupListProcessor {
	return &NoOpNodeGroupListProcessor{}
}

// Process processes lists of unschedulable and scheduled pods before scaling of the cluster.
func (p *NoOpNodeGroupListProcessor) Process(ctx context.Context, context *autoscalingcontext.AutoscalingContext, nodeGroups []cloudprovider.NodeGroup, nodeInfos map[string]*schedulernodeinfo.NodeInfo,
	unschedulablePods []*apiv1.Pod) ([]cloudprovider.NodeGroup, map[string]*schedulernodeinfo.NodeInfo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "NoOpNodeGroupListProcessor.Process")
	defer span.Finish()

	return nodeGroups, nodeInfos, nil
}

// CleanUp cleans up the processor's internal structures.
func (p *NoOpNodeGroupListProcessor) CleanUp(ctx context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "NoOpNodeGroupListProcessor.CleanUp")
	defer span.Finish()
}
