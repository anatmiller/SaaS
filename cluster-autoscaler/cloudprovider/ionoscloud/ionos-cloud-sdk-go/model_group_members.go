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

// GroupMembers struct for GroupMembers
type GroupMembers struct {
	// The resource's unique identifier
	Id *string `json:"id,omitempty"`
	// The type of object that has been created
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path)
	Href *string `json:"href,omitempty"`
	// Array of items in that collection
	Items *[]User `json:"items,omitempty"`
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupMembers) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupMembers) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *GroupMembers) SetId(v string) {
	o.Id = &v
}

// HasId returns a boolean if a field has been set.
func (o *GroupMembers) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *GroupMembers) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupMembers) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *GroupMembers) SetType(v Type) {
	o.Type = &v
}

// HasType returns a boolean if a field has been set.
func (o *GroupMembers) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GroupMembers) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupMembers) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Href, true
}

// SetHref sets field value
func (o *GroupMembers) SetHref(v string) {
	o.Href = &v
}

// HasHref returns a boolean if a field has been set.
func (o *GroupMembers) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []User will be returned
func (o *GroupMembers) GetItems() *[]User {
	if o == nil {
		return nil
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupMembers) GetItemsOk() (*[]User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *GroupMembers) SetItems(v []User) {
	o.Items = &v
}

// HasItems returns a boolean if a field has been set.
func (o *GroupMembers) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o GroupMembers) MarshalJSON() ([]byte, error) {
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

type NullableGroupMembers struct {
	value *GroupMembers
	isSet bool
}

func (v NullableGroupMembers) Get() *GroupMembers {
	return v.value
}

func (v *NullableGroupMembers) Set(val *GroupMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMembers(val *GroupMembers) *NullableGroupMembers {
	return &NullableGroupMembers{value: val, isSet: true}
}

func (v NullableGroupMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
