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

// checks if the UpdateVolumeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVolumeRequest{}

// UpdateVolumeRequest struct for UpdateVolumeRequest
type UpdateVolumeRequest struct {
	SnapshotRetention *int32 `json:"snapshot_retention,omitempty"`
}

// NewUpdateVolumeRequest instantiates a new UpdateVolumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVolumeRequest() *UpdateVolumeRequest {
	this := UpdateVolumeRequest{}
	return &this
}

// NewUpdateVolumeRequestWithDefaults instantiates a new UpdateVolumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVolumeRequestWithDefaults() *UpdateVolumeRequest {
	this := UpdateVolumeRequest{}
	return &this
}

// GetSnapshotRetention returns the SnapshotRetention field value if set, zero value otherwise.
func (o *UpdateVolumeRequest) GetSnapshotRetention() int32 {
	if o == nil || IsNil(o.SnapshotRetention) {
		var ret int32
		return ret
	}
	return *o.SnapshotRetention
}

// GetSnapshotRetentionOk returns a tuple with the SnapshotRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeRequest) GetSnapshotRetentionOk() (*int32, bool) {
	if o == nil || IsNil(o.SnapshotRetention) {
		return nil, false
	}
	return o.SnapshotRetention, true
}

// HasSnapshotRetention returns a boolean if a field has been set.
func (o *UpdateVolumeRequest) HasSnapshotRetention() bool {
	if o != nil && !IsNil(o.SnapshotRetention) {
		return true
	}

	return false
}

// SetSnapshotRetention gets a reference to the given int32 and assigns it to the SnapshotRetention field.
func (o *UpdateVolumeRequest) SetSnapshotRetention(v int32) {
	o.SnapshotRetention = &v
}

func (o UpdateVolumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVolumeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SnapshotRetention) {
		toSerialize["snapshot_retention"] = o.SnapshotRetention
	}
	return toSerialize, nil
}

type NullableUpdateVolumeRequest struct {
	value *UpdateVolumeRequest
	isSet bool
}

func (v NullableUpdateVolumeRequest) Get() *UpdateVolumeRequest {
	return v.value
}

func (v *NullableUpdateVolumeRequest) Set(val *UpdateVolumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVolumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVolumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVolumeRequest(val *UpdateVolumeRequest) *NullableUpdateVolumeRequest {
	return &NullableUpdateVolumeRequest{value: val, isSet: true}
}

func (v NullableUpdateVolumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVolumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

