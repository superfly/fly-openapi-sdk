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

// checks if the ListenSocket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListenSocket{}

// ListenSocket struct for ListenSocket
type ListenSocket struct {
	Address *string `json:"address,omitempty"`
	Proto *string `json:"proto,omitempty"`
}

// NewListenSocket instantiates a new ListenSocket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListenSocket() *ListenSocket {
	this := ListenSocket{}
	return &this
}

// NewListenSocketWithDefaults instantiates a new ListenSocket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListenSocketWithDefaults() *ListenSocket {
	this := ListenSocket{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ListenSocket) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenSocket) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ListenSocket) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ListenSocket) SetAddress(v string) {
	o.Address = &v
}

// GetProto returns the Proto field value if set, zero value otherwise.
func (o *ListenSocket) GetProto() string {
	if o == nil || IsNil(o.Proto) {
		var ret string
		return ret
	}
	return *o.Proto
}

// GetProtoOk returns a tuple with the Proto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenSocket) GetProtoOk() (*string, bool) {
	if o == nil || IsNil(o.Proto) {
		return nil, false
	}
	return o.Proto, true
}

// HasProto returns a boolean if a field has been set.
func (o *ListenSocket) HasProto() bool {
	if o != nil && !IsNil(o.Proto) {
		return true
	}

	return false
}

// SetProto gets a reference to the given string and assigns it to the Proto field.
func (o *ListenSocket) SetProto(v string) {
	o.Proto = &v
}

func (o ListenSocket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListenSocket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Proto) {
		toSerialize["proto"] = o.Proto
	}
	return toSerialize, nil
}

type NullableListenSocket struct {
	value *ListenSocket
	isSet bool
}

func (v NullableListenSocket) Get() *ListenSocket {
	return v.value
}

func (v *NullableListenSocket) Set(val *ListenSocket) {
	v.value = val
	v.isSet = true
}

func (v NullableListenSocket) IsSet() bool {
	return v.isSet
}

func (v *NullableListenSocket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListenSocket(val *ListenSocket) *NullableListenSocket {
	return &NullableListenSocket{value: val, isSet: true}
}

func (v NullableListenSocket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListenSocket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


