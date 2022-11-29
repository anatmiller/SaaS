/*
Copyright 2022 The Kubernetes Authors.

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

package dra

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	kubecontainer "k8s.io/kubernetes/pkg/kubelet/container"
)

// Manager manages all the DRA resource plugins running on a node.
type Manager interface {
	// PrepareResources prepares resources for a container in a pod.
	// It communicates with the DRA resource plugin to prepare resources and
	// returns resource info to trigger CDI injection by the runtime.
	PrepareResources(pod *v1.Pod, container *v1.Container) (*ContainerInfo, error)

	// UnprepareResources calls NodeUnprepareResource GRPC from DRA plugin to unprepare pod resources
	UnprepareResources(pod *v1.Pod) error

	// PodMightNeedToUnprepareResources returns true if the pod with the given UID
	// might need to unprepare resources.
	PodMightNeedToUnprepareResources(UID types.UID) bool
}

// ContainerInfo contains information required by the runtime to consume prepared resources.
type ContainerInfo struct {
	// The Annotations for the container
	Annotations []kubecontainer.Annotation
}
