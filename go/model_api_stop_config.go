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

// checks if the ApiStopConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiStopConfig{}

// ApiStopConfig struct for ApiStopConfig
type ApiStopConfig struct {
	Signal *string `json:"signal,omitempty"`
	Timeout *string `json:"timeout,omitempty"`
}

// NewApiStopConfig instantiates a new ApiStopConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStopConfig() *ApiStopConfig {
	this := ApiStopConfig{}
	return &this
}

// NewApiStopConfigWithDefaults instantiates a new ApiStopConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStopConfigWithDefaults() *ApiStopConfig {
	this := ApiStopConfig{}
	return &this
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *ApiStopConfig) GetSignal() string {
	if o == nil || IsNil(o.Signal) {
		var ret string
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStopConfig) GetSignalOk() (*string, bool) {
	if o == nil || IsNil(o.Signal) {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *ApiStopConfig) HasSignal() bool {
	if o != nil && !IsNil(o.Signal) {
		return true
	}

	return false
}

// SetSignal gets a reference to the given string and assigns it to the Signal field.
func (o *ApiStopConfig) SetSignal(v string) {
	o.Signal = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ApiStopConfig) GetTimeout() string {
	if o == nil || IsNil(o.Timeout) {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStopConfig) GetTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ApiStopConfig) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *ApiStopConfig) SetTimeout(v string) {
	o.Timeout = &v
}

func (o ApiStopConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiStopConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Signal) {
		toSerialize["signal"] = o.Signal
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return toSerialize, nil
}

type NullableApiStopConfig struct {
	value *ApiStopConfig
	isSet bool
}

func (v NullableApiStopConfig) Get() *ApiStopConfig {
	return v.value
}

func (v *NullableApiStopConfig) Set(val *ApiStopConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStopConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStopConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStopConfig(val *ApiStopConfig) *NullableApiStopConfig {
	return &NullableApiStopConfig{value: val, isSet: true}
}

func (v NullableApiStopConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStopConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

