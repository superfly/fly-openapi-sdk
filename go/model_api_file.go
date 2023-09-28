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

// checks if the ApiFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiFile{}

// ApiFile struct for ApiFile
type ApiFile struct {
	// GuestPath is the path on the machine where the file will be written and must be an absolute path. i.e. /full/path/to/file.json
	GuestPath *string `json:"guest_path,omitempty"`
	// RawValue containts the base64 encoded string of the file contents.
	RawValue *string `json:"raw_value,omitempty"`
	// SecretName is the name of the secret that contains the base64 encoded file contents.
	SecretName *string `json:"secret_name,omitempty"`
}

// NewApiFile instantiates a new ApiFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiFile() *ApiFile {
	this := ApiFile{}
	return &this
}

// NewApiFileWithDefaults instantiates a new ApiFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiFileWithDefaults() *ApiFile {
	this := ApiFile{}
	return &this
}

// GetGuestPath returns the GuestPath field value if set, zero value otherwise.
func (o *ApiFile) GetGuestPath() string {
	if o == nil || IsNil(o.GuestPath) {
		var ret string
		return ret
	}
	return *o.GuestPath
}

// GetGuestPathOk returns a tuple with the GuestPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFile) GetGuestPathOk() (*string, bool) {
	if o == nil || IsNil(o.GuestPath) {
		return nil, false
	}
	return o.GuestPath, true
}

// HasGuestPath returns a boolean if a field has been set.
func (o *ApiFile) HasGuestPath() bool {
	if o != nil && !IsNil(o.GuestPath) {
		return true
	}

	return false
}

// SetGuestPath gets a reference to the given string and assigns it to the GuestPath field.
func (o *ApiFile) SetGuestPath(v string) {
	o.GuestPath = &v
}

// GetRawValue returns the RawValue field value if set, zero value otherwise.
func (o *ApiFile) GetRawValue() string {
	if o == nil || IsNil(o.RawValue) {
		var ret string
		return ret
	}
	return *o.RawValue
}

// GetRawValueOk returns a tuple with the RawValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFile) GetRawValueOk() (*string, bool) {
	if o == nil || IsNil(o.RawValue) {
		return nil, false
	}
	return o.RawValue, true
}

// HasRawValue returns a boolean if a field has been set.
func (o *ApiFile) HasRawValue() bool {
	if o != nil && !IsNil(o.RawValue) {
		return true
	}

	return false
}

// SetRawValue gets a reference to the given string and assigns it to the RawValue field.
func (o *ApiFile) SetRawValue(v string) {
	o.RawValue = &v
}

// GetSecretName returns the SecretName field value if set, zero value otherwise.
func (o *ApiFile) GetSecretName() string {
	if o == nil || IsNil(o.SecretName) {
		var ret string
		return ret
	}
	return *o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFile) GetSecretNameOk() (*string, bool) {
	if o == nil || IsNil(o.SecretName) {
		return nil, false
	}
	return o.SecretName, true
}

// HasSecretName returns a boolean if a field has been set.
func (o *ApiFile) HasSecretName() bool {
	if o != nil && !IsNil(o.SecretName) {
		return true
	}

	return false
}

// SetSecretName gets a reference to the given string and assigns it to the SecretName field.
func (o *ApiFile) SetSecretName(v string) {
	o.SecretName = &v
}

func (o ApiFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuestPath) {
		toSerialize["guest_path"] = o.GuestPath
	}
	if !IsNil(o.RawValue) {
		toSerialize["raw_value"] = o.RawValue
	}
	if !IsNil(o.SecretName) {
		toSerialize["secret_name"] = o.SecretName
	}
	return toSerialize, nil
}

type NullableApiFile struct {
	value *ApiFile
	isSet bool
}

func (v NullableApiFile) Get() *ApiFile {
	return v.value
}

func (v *NullableApiFile) Set(val *ApiFile) {
	v.value = val
	v.isSet = true
}

func (v NullableApiFile) IsSet() bool {
	return v.isSet
}

func (v *NullableApiFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiFile(val *ApiFile) *NullableApiFile {
	return &NullableApiFile{value: val, isSet: true}
}

func (v NullableApiFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

