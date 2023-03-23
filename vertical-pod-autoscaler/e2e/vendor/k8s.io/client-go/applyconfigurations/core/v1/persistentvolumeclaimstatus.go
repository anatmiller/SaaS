/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// PersistentVolumeClaimStatusApplyConfiguration represents an declarative configuration of the PersistentVolumeClaimStatus type for use
// with apply.
type PersistentVolumeClaimStatusApplyConfiguration struct {
	Phase              *v1.PersistentVolumeClaimPhase                     `json:"phase,omitempty"`
	AccessModes        []v1.PersistentVolumeAccessMode                    `json:"accessModes,omitempty"`
	Capacity           *v1.ResourceList                                   `json:"capacity,omitempty"`
	Conditions         []PersistentVolumeClaimConditionApplyConfiguration `json:"conditions,omitempty"`
	AllocatedResources *v1.ResourceList                                   `json:"allocatedResources,omitempty"`
	ResizeStatus       *v1.PersistentVolumeClaimResizeStatus              `json:"resizeStatus,omitempty"`
}

// PersistentVolumeClaimStatusApplyConfiguration constructs an declarative configuration of the PersistentVolumeClaimStatus type for use with
// apply.
func PersistentVolumeClaimStatus() *PersistentVolumeClaimStatusApplyConfiguration {
	return &PersistentVolumeClaimStatusApplyConfiguration{}
}

// WithPhase sets the Phase field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Phase field is set to the value of the last call.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithPhase(value v1.PersistentVolumeClaimPhase) *PersistentVolumeClaimStatusApplyConfiguration {
	b.Phase = &value
	return b
}

// WithAccessModes adds the given value to the AccessModes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AccessModes field.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithAccessModes(values ...v1.PersistentVolumeAccessMode) *PersistentVolumeClaimStatusApplyConfiguration {
	for i := range values {
		b.AccessModes = append(b.AccessModes, values[i])
	}
	return b
}

// WithCapacity sets the Capacity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Capacity field is set to the value of the last call.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithCapacity(value v1.ResourceList) *PersistentVolumeClaimStatusApplyConfiguration {
	b.Capacity = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithConditions(values ...*PersistentVolumeClaimConditionApplyConfiguration) *PersistentVolumeClaimStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithAllocatedResources sets the AllocatedResources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllocatedResources field is set to the value of the last call.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithAllocatedResources(value v1.ResourceList) *PersistentVolumeClaimStatusApplyConfiguration {
	b.AllocatedResources = &value
	return b
}

// WithResizeStatus sets the ResizeStatus field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResizeStatus field is set to the value of the last call.
func (b *PersistentVolumeClaimStatusApplyConfiguration) WithResizeStatus(value v1.PersistentVolumeClaimResizeStatus) *PersistentVolumeClaimStatusApplyConfiguration {
	b.ResizeStatus = &value
	return b
}