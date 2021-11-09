/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionossdk

import (
	"encoding/json"
)

// KubernetesNodePools struct for KubernetesNodePools
type KubernetesNodePools struct {
	// Unique representation for Kubernetes Node Pool as a collection on a resource.
	Id *string `json:"id,omitempty"`
	// The type of resource within a collection
	Type *string `json:"type,omitempty"`
	// URL to the collection representation (absolute path)
	Href *string `json:"href,omitempty"`
	// Array of items in that collection
	Items *[]KubernetesNodePool `json:"items,omitempty"`
}



// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePools) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePools) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *KubernetesNodePools) SetId(v string) {
	o.Id = &v
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesNodePools) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}



// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePools) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePools) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *KubernetesNodePools) SetType(v string) {
	o.Type = &v
}

// HasType returns a boolean if a field has been set.
func (o *KubernetesNodePools) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}



// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePools) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePools) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Href, true
}

// SetHref sets field value
func (o *KubernetesNodePools) SetHref(v string) {
	o.Href = &v
}

// HasHref returns a boolean if a field has been set.
func (o *KubernetesNodePools) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}



// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []KubernetesNodePool will be returned
func (o *KubernetesNodePools) GetItems() *[]KubernetesNodePool {
	if o == nil {
		return nil
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePools) GetItemsOk() (*[]KubernetesNodePool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *KubernetesNodePools) SetItems(v []KubernetesNodePool) {
	o.Items = &v
}

// HasItems returns a boolean if a field has been set.
func (o *KubernetesNodePools) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}


func (o KubernetesNodePools) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	
	return json.Marshal(toSerialize)
}

type NullableKubernetesNodePools struct {
	value *KubernetesNodePools
	isSet bool
}

func (v NullableKubernetesNodePools) Get() *KubernetesNodePools {
	return v.value
}

func (v *NullableKubernetesNodePools) Set(val *KubernetesNodePools) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodePools) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodePools) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodePools(val *KubernetesNodePools) *NullableKubernetesNodePools {
	return &NullableKubernetesNodePools{value: val, isSet: true}
}

func (v NullableKubernetesNodePools) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodePools) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


