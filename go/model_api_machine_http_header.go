/*
Fly Machines API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fly-sdk

import (
	"encoding/json"
)

// checks if the ApiMachineHTTPHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMachineHTTPHeader{}

// ApiMachineHTTPHeader For http checks, an array of objects with string field Name and array of strings field Values. The key/value pairs specify header and header values that will get passed with the check call.
type ApiMachineHTTPHeader struct {
	// The header name
	Name *string `json:"name,omitempty"`
	// The header value
	Values []string `json:"values,omitempty"`
}

// NewApiMachineHTTPHeader instantiates a new ApiMachineHTTPHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMachineHTTPHeader() *ApiMachineHTTPHeader {
	this := ApiMachineHTTPHeader{}
	return &this
}

// NewApiMachineHTTPHeaderWithDefaults instantiates a new ApiMachineHTTPHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMachineHTTPHeaderWithDefaults() *ApiMachineHTTPHeader {
	this := ApiMachineHTTPHeader{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiMachineHTTPHeader) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachineHTTPHeader) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiMachineHTTPHeader) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiMachineHTTPHeader) SetName(v string) {
	o.Name = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ApiMachineHTTPHeader) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachineHTTPHeader) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ApiMachineHTTPHeader) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ApiMachineHTTPHeader) SetValues(v []string) {
	o.Values = v
}

func (o ApiMachineHTTPHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMachineHTTPHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableApiMachineHTTPHeader struct {
	value *ApiMachineHTTPHeader
	isSet bool
}

func (v NullableApiMachineHTTPHeader) Get() *ApiMachineHTTPHeader {
	return v.value
}

func (v *NullableApiMachineHTTPHeader) Set(val *ApiMachineHTTPHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMachineHTTPHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMachineHTTPHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMachineHTTPHeader(val *ApiMachineHTTPHeader) *NullableApiMachineHTTPHeader {
	return &NullableApiMachineHTTPHeader{value: val, isSet: true}
}

func (v NullableApiMachineHTTPHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMachineHTTPHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


