/*
Fly Machines API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flyapi

import (
	"encoding/json"
)

// checks if the ApiProxyProtoOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiProxyProtoOptions{}

// ApiProxyProtoOptions struct for ApiProxyProtoOptions
type ApiProxyProtoOptions struct {
	Version *string `json:"version,omitempty"`
}

// NewApiProxyProtoOptions instantiates a new ApiProxyProtoOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiProxyProtoOptions() *ApiProxyProtoOptions {
	this := ApiProxyProtoOptions{}
	return &this
}

// NewApiProxyProtoOptionsWithDefaults instantiates a new ApiProxyProtoOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiProxyProtoOptionsWithDefaults() *ApiProxyProtoOptions {
	this := ApiProxyProtoOptions{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApiProxyProtoOptions) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiProxyProtoOptions) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApiProxyProtoOptions) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApiProxyProtoOptions) SetVersion(v string) {
	o.Version = &v
}

func (o ApiProxyProtoOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiProxyProtoOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableApiProxyProtoOptions struct {
	value *ApiProxyProtoOptions
	isSet bool
}

func (v NullableApiProxyProtoOptions) Get() *ApiProxyProtoOptions {
	return v.value
}

func (v *NullableApiProxyProtoOptions) Set(val *ApiProxyProtoOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableApiProxyProtoOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableApiProxyProtoOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiProxyProtoOptions(val *ApiProxyProtoOptions) *NullableApiProxyProtoOptions {
	return &NullableApiProxyProtoOptions{value: val, isSet: true}
}

func (v NullableApiProxyProtoOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiProxyProtoOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


