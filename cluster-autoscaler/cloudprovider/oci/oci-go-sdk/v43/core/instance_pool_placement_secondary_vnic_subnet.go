// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/oci-go-sdk/v43/common"
)

// InstancePoolPlacementSecondaryVnicSubnet The secondary VNIC object for the placement configuration for an instance pool.
type InstancePoolPlacementSecondaryVnicSubnet struct {

	// The subnet OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the secondary VNIC.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The display name of the VNIC. This is also use to match against the instance configuration defined
	// secondary VNIC.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m InstancePoolPlacementSecondaryVnicSubnet) String() string {
	return common.PointerString(m)
}
