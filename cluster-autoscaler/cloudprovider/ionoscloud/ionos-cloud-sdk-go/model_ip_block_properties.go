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

// IpBlockProperties struct for IpBlockProperties
type IpBlockProperties struct {
	// A collection of IPs associated with the IP Block
	Ips *[]string `json:"ips,omitempty"`
	// Location of that IP Block. Property cannot be modified after creation (disallowed in update requests)
	Location *string `json:"location"`
	// The size of the IP block
	Size *int32 `json:"size"`
	// A name of that resource
	Name *string `json:"name,omitempty"`
	// Read-Only attribute. Lists consumption detail of an individual ip
	IpConsumers *[]IpConsumer `json:"ipConsumers,omitempty"`
}

// GetIps returns the Ips field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *IpBlockProperties) GetIps() *[]string {
	if o == nil {
		return nil
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpBlockProperties) GetIpsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ips, true
}

// SetIps sets field value
func (o *IpBlockProperties) SetIps(v []string) {
	o.Ips = &v
}

// HasIps returns a boolean if a field has been set.
func (o *IpBlockProperties) HasIps() bool {
	if o != nil && o.Ips != nil {
		return true
	}

	return false
}

// GetLocation returns the Location field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpBlockProperties) GetLocation() *string {
	if o == nil {
		return nil
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpBlockProperties) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location, true
}

// SetLocation sets field value
func (o *IpBlockProperties) SetLocation(v string) {
	o.Location = &v
}

// HasLocation returns a boolean if a field has been set.
func (o *IpBlockProperties) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// GetSize returns the Size field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *IpBlockProperties) GetSize() *int32 {
	if o == nil {
		return nil
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpBlockProperties) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size, true
}

// SetSize sets field value
func (o *IpBlockProperties) SetSize(v int32) {
	o.Size = &v
}

// HasSize returns a boolean if a field has been set.
func (o *IpBlockProperties) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpBlockProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpBlockProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *IpBlockProperties) SetName(v string) {
	o.Name = &v
}

// HasName returns a boolean if a field has been set.
func (o *IpBlockProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetIpConsumers returns the IpConsumers field value
// If the value is explicit nil, the zero value for []IpConsumer will be returned
func (o *IpBlockProperties) GetIpConsumers() *[]IpConsumer {
	if o == nil {
		return nil
	}

	return o.IpConsumers
}

// GetIpConsumersOk returns a tuple with the IpConsumers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpBlockProperties) GetIpConsumersOk() (*[]IpConsumer, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpConsumers, true
}

// SetIpConsumers sets field value
func (o *IpBlockProperties) SetIpConsumers(v []IpConsumer) {
	o.IpConsumers = &v
}

// HasIpConsumers returns a boolean if a field has been set.
func (o *IpBlockProperties) HasIpConsumers() bool {
	if o != nil && o.IpConsumers != nil {
		return true
	}

	return false
}

func (o IpBlockProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Ips != nil {
		toSerialize["ips"] = o.Ips
	}

	if o.Location != nil {
		toSerialize["location"] = o.Location
	}

	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.IpConsumers != nil {
		toSerialize["ipConsumers"] = o.IpConsumers
	}

	return json.Marshal(toSerialize)
}

type NullableIpBlockProperties struct {
	value *IpBlockProperties
	isSet bool
}

func (v NullableIpBlockProperties) Get() *IpBlockProperties {
	return v.value
}

func (v *NullableIpBlockProperties) Set(val *IpBlockProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableIpBlockProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableIpBlockProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpBlockProperties(val *IpBlockProperties) *NullableIpBlockProperties {
	return &NullableIpBlockProperties{value: val, isSet: true}
}

func (v NullableIpBlockProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpBlockProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
