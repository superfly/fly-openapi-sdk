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

// checks if the ApiMachineProcess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMachineProcess{}

// ApiMachineProcess struct for ApiMachineProcess
type ApiMachineProcess struct {
	Cmd []string `json:"cmd,omitempty"`
	Entrypoint []string `json:"entrypoint,omitempty"`
	Env *map[string]string `json:"env,omitempty"`
	Exec []string `json:"exec,omitempty"`
	User *string `json:"user,omitempty"`
}

// NewApiMachineProcess instantiates a new ApiMachineProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMachineProcess() *ApiMachineProcess {
	this := ApiMachineProcess{}
	return &this
}

// NewApiMachineProcessWithDefaults instantiates a new ApiMachineProcess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMachineProcessWithDefaults() *ApiMachineProcess {
	this := ApiMachineProcess{}
	return &this
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *ApiMachineProcess) GetCmd() []string {
	if o == nil || IsNil(o.Cmd) {
		var ret []string
		return ret
	}
	return o.Cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachineProcess) GetCmdOk() ([]string, bool) {
	if o == nil || IsNil(o.Cmd) {
		return nil, false
	}
	return o.Cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *ApiMachineProcess) HasCmd() bool {
	if o != nil && !IsNil(o.Cmd) {
		return true
	}

	return false
}

// SetCmd gets a reference to the given []string and assigns it to the Cmd field.
func (o *ApiMachineProcess) SetCmd(v []string) {
	o.Cmd = v
}

// GetEntrypoint returns the Entrypoint field value if set, zero value otherwise.
func (o *ApiMachineProcess) GetEntrypoint() []string {
	if o == nil || IsNil(o.Entrypoint) {
		var ret []string
		return ret
	}
	return o.Entrypoint
}

// GetEntrypointOk returns a tuple with the Entrypoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachineProcess) GetEntrypointOk() ([]string, bool) {
	if o == nil || IsNil(o.Entrypoint) {
		return nil, false
	}
	return o.Entrypoint, true
}

// HasEntrypoint returns a boolean if a field has been set.
func (o *ApiMachineProcess) HasEntrypoint() bool {
	if o != nil && !IsNil(o.Entrypoint) {
		return true
	}

	return false
}

// SetEntrypoint gets a reference to the given []string and assigns it to the Entrypoint field.
func (o *ApiMachineProcess) SetEntrypoint(v []string) {
	o.Entrypoint = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *ApiMachineProcess) GetEnv() map[string]string {
	if o == nil || IsNil(o.Env) {
		var ret map[string]string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachineProcess) GetEnvOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *ApiMachineProcess) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given map[string]string and assigns it to the Env field.
func (o *ApiMachineProcess) SetEnv(v map[string]string) {
	o.Env = &v
}

// GetExec returns the Exec field value if set, zero value otherwise.
func (o *ApiMachineProcess) GetExec() []string {
	if o == nil || IsNil(o.Exec) {
		var ret []string
		return ret
	}
	return o.Exec
}

// GetExecOk returns a tuple with the Exec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachineProcess) GetExecOk() ([]string, bool) {
	if o == nil || IsNil(o.Exec) {
		return nil, false
	}
	return o.Exec, true
}

// HasExec returns a boolean if a field has been set.
func (o *ApiMachineProcess) HasExec() bool {
	if o != nil && !IsNil(o.Exec) {
		return true
	}

	return false
}

// SetExec gets a reference to the given []string and assigns it to the Exec field.
func (o *ApiMachineProcess) SetExec(v []string) {
	o.Exec = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ApiMachineProcess) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMachineProcess) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ApiMachineProcess) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ApiMachineProcess) SetUser(v string) {
	o.User = &v
}

func (o ApiMachineProcess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMachineProcess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cmd) {
		toSerialize["cmd"] = o.Cmd
	}
	if !IsNil(o.Entrypoint) {
		toSerialize["entrypoint"] = o.Entrypoint
	}
	if !IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !IsNil(o.Exec) {
		toSerialize["exec"] = o.Exec
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableApiMachineProcess struct {
	value *ApiMachineProcess
	isSet bool
}

func (v NullableApiMachineProcess) Get() *ApiMachineProcess {
	return v.value
}

func (v *NullableApiMachineProcess) Set(val *ApiMachineProcess) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMachineProcess) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMachineProcess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMachineProcess(val *ApiMachineProcess) *NullableApiMachineProcess {
	return &NullableApiMachineProcess{value: val, isSet: true}
}

func (v NullableApiMachineProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMachineProcess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


