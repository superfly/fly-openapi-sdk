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

// checks if the ApiMachinePort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMachinePort{}

// ApiMachinePort struct for ApiMachinePort
type ApiMachinePort struct {
	EndPort *int32 `json:"end_port,omitempty"`
	ForceHttps *bool `json:"force_https,omitempty"`
	Handlers []string `json:"handlers,omitempty"`
	HttpOptions *ApiHTTPOptions `json:"http_options,omitempty"`
	Port *int32 `json:"port,omitempty"`
	ProxyProtoOptions *ApiProxyProtoOptions `json:"proxy_proto_options,omitempty"`
	StartPort *int32 `json:"start_port,omitempty"`
	TlsOptions *ApiTLSOptions `json:"tls_options,omitempty"`
}

// NewApiMachinePort instantiates a new ApiMachinePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMachinePort() *ApiMachinePort {
	this := ApiMachinePort{}
	return &this
}

// NewApiMachinePortWithDefaults instantiates a new ApiMachinePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMachinePortWithDefaults() *ApiMachinePort {
	this := ApiMachinePort{}
	return &this
}

// GetEndPort returns the EndPort field value if set, zero value otherwise.
func (o *ApiMachinePort) GetEndPort() int32 {
	if o == nil || IsNil(o.EndPort) {
		var ret int32
		return ret
	}
	return *o.EndPort
}

// GetEndPortOk returns a tuple with the EndPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachinePort) GetEndPortOk() (*int32, bool) {
	if o == nil || IsNil(o.EndPort) {
		return nil, false
	}
	return o.EndPort, true
}

// HasEndPort returns a boolean if a field has been set.
func (o *ApiMachinePort) HasEndPort() bool {
	if o != nil && !IsNil(o.EndPort) {
		return true
	}

	return false
}

// SetEndPort gets a reference to the given int32 and assigns it to the EndPort field.
func (o *ApiMachinePort) SetEndPort(v int32) {
	o.EndPort = &v
}

// GetForceHttps returns the ForceHttps field value if set, zero value otherwise.
func (o *ApiMachinePort) GetForceHttps() bool {
	if o == nil || IsNil(o.ForceHttps) {
		var ret bool
		return ret
	}
	return *o.ForceHttps
}

// GetForceHttpsOk returns a tuple with the ForceHttps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachinePort) GetForceHttpsOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceHttps) {
		return nil, false
	}
	return o.ForceHttps, true
}

// HasForceHttps returns a boolean if a field has been set.
func (o *ApiMachinePort) HasForceHttps() bool {
	if o != nil && !IsNil(o.ForceHttps) {
		return true
	}

	return false
}

// SetForceHttps gets a reference to the given bool and assigns it to the ForceHttps field.
func (o *ApiMachinePort) SetForceHttps(v bool) {
	o.ForceHttps = &v
}

// GetHandlers returns the Handlers field value if set, zero value otherwise.
func (o *ApiMachinePort) GetHandlers() []string {
	if o == nil || IsNil(o.Handlers) {
		var ret []string
		return ret
	}
	return o.Handlers
}

// GetHandlersOk returns a tuple with the Handlers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachinePort) GetHandlersOk() ([]string, bool) {
	if o == nil || IsNil(o.Handlers) {
		return nil, false
	}
	return o.Handlers, true
}

// HasHandlers returns a boolean if a field has been set.
func (o *ApiMachinePort) HasHandlers() bool {
	if o != nil && !IsNil(o.Handlers) {
		return true
	}

	return false
}

// SetHandlers gets a reference to the given []string and assigns it to the Handlers field.
func (o *ApiMachinePort) SetHandlers(v []string) {
	o.Handlers = v
}

// GetHttpOptions returns the HttpOptions field value if set, zero value otherwise.
func (o *ApiMachinePort) GetHttpOptions() ApiHTTPOptions {
	if o == nil || IsNil(o.HttpOptions) {
		var ret ApiHTTPOptions
		return ret
	}
	return *o.HttpOptions
}

// GetHttpOptionsOk returns a tuple with the HttpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachinePort) GetHttpOptionsOk() (*ApiHTTPOptions, bool) {
	if o == nil || IsNil(o.HttpOptions) {
		return nil, false
	}
	return o.HttpOptions, true
}

// HasHttpOptions returns a boolean if a field has been set.
func (o *ApiMachinePort) HasHttpOptions() bool {
	if o != nil && !IsNil(o.HttpOptions) {
		return true
	}

	return false
}

// SetHttpOptions gets a reference to the given ApiHTTPOptions and assigns it to the HttpOptions field.
func (o *ApiMachinePort) SetHttpOptions(v ApiHTTPOptions) {
	o.HttpOptions = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ApiMachinePort) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachinePort) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ApiMachinePort) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ApiMachinePort) SetPort(v int32) {
	o.Port = &v
}

// GetProxyProtoOptions returns the ProxyProtoOptions field value if set, zero value otherwise.
func (o *ApiMachinePort) GetProxyProtoOptions() ApiProxyProtoOptions {
	if o == nil || IsNil(o.ProxyProtoOptions) {
		var ret ApiProxyProtoOptions
		return ret
	}
	return *o.ProxyProtoOptions
}

// GetProxyProtoOptionsOk returns a tuple with the ProxyProtoOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachinePort) GetProxyProtoOptionsOk() (*ApiProxyProtoOptions, bool) {
	if o == nil || IsNil(o.ProxyProtoOptions) {
		return nil, false
	}
	return o.ProxyProtoOptions, true
}

// HasProxyProtoOptions returns a boolean if a field has been set.
func (o *ApiMachinePort) HasProxyProtoOptions() bool {
	if o != nil && !IsNil(o.ProxyProtoOptions) {
		return true
	}

	return false
}

// SetProxyProtoOptions gets a reference to the given ApiProxyProtoOptions and assigns it to the ProxyProtoOptions field.
func (o *ApiMachinePort) SetProxyProtoOptions(v ApiProxyProtoOptions) {
	o.ProxyProtoOptions = &v
}

// GetStartPort returns the StartPort field value if set, zero value otherwise.
func (o *ApiMachinePort) GetStartPort() int32 {
	if o == nil || IsNil(o.StartPort) {
		var ret int32
		return ret
	}
	return *o.StartPort
}

// GetStartPortOk returns a tuple with the StartPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachinePort) GetStartPortOk() (*int32, bool) {
	if o == nil || IsNil(o.StartPort) {
		return nil, false
	}
	return o.StartPort, true
}

// HasStartPort returns a boolean if a field has been set.
func (o *ApiMachinePort) HasStartPort() bool {
	if o != nil && !IsNil(o.StartPort) {
		return true
	}

	return false
}

// SetStartPort gets a reference to the given int32 and assigns it to the StartPort field.
func (o *ApiMachinePort) SetStartPort(v int32) {
	o.StartPort = &v
}

// GetTlsOptions returns the TlsOptions field value if set, zero value otherwise.
func (o *ApiMachinePort) GetTlsOptions() ApiTLSOptions {
	if o == nil || IsNil(o.TlsOptions) {
		var ret ApiTLSOptions
		return ret
	}
	return *o.TlsOptions
}

// GetTlsOptionsOk returns a tuple with the TlsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachinePort) GetTlsOptionsOk() (*ApiTLSOptions, bool) {
	if o == nil || IsNil(o.TlsOptions) {
		return nil, false
	}
	return o.TlsOptions, true
}

// HasTlsOptions returns a boolean if a field has been set.
func (o *ApiMachinePort) HasTlsOptions() bool {
	if o != nil && !IsNil(o.TlsOptions) {
		return true
	}

	return false
}

// SetTlsOptions gets a reference to the given ApiTLSOptions and assigns it to the TlsOptions field.
func (o *ApiMachinePort) SetTlsOptions(v ApiTLSOptions) {
	o.TlsOptions = &v
}

func (o ApiMachinePort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMachinePort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndPort) {
		toSerialize["end_port"] = o.EndPort
	}
	if !IsNil(o.ForceHttps) {
		toSerialize["force_https"] = o.ForceHttps
	}
	if !IsNil(o.Handlers) {
		toSerialize["handlers"] = o.Handlers
	}
	if !IsNil(o.HttpOptions) {
		toSerialize["http_options"] = o.HttpOptions
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.ProxyProtoOptions) {
		toSerialize["proxy_proto_options"] = o.ProxyProtoOptions
	}
	if !IsNil(o.StartPort) {
		toSerialize["start_port"] = o.StartPort
	}
	if !IsNil(o.TlsOptions) {
		toSerialize["tls_options"] = o.TlsOptions
	}
	return toSerialize, nil
}

type NullableApiMachinePort struct {
	value *ApiMachinePort
	isSet bool
}

func (v NullableApiMachinePort) Get() *ApiMachinePort {
	return v.value
}

func (v *NullableApiMachinePort) Set(val *ApiMachinePort) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMachinePort) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMachinePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMachinePort(val *ApiMachinePort) *NullableApiMachinePort {
	return &NullableApiMachinePort{value: val, isSet: true}
}

func (v NullableApiMachinePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMachinePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


