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

// checks if the MachineEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineEvent{}

// MachineEvent struct for MachineEvent
type MachineEvent struct {
	Id *string `json:"id,omitempty"`
	Request map[string]interface{} `json:"request,omitempty"`
	Source *string `json:"source,omitempty"`
	Status *string `json:"status,omitempty"`
	Timestamp *int32 `json:"timestamp,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewMachineEvent instantiates a new MachineEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineEvent() *MachineEvent {
	this := MachineEvent{}
	return &this
}

// NewMachineEventWithDefaults instantiates a new MachineEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineEventWithDefaults() *MachineEvent {
	this := MachineEvent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MachineEvent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineEvent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MachineEvent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MachineEvent) SetId(v string) {
	o.Id = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *MachineEvent) GetRequest() map[string]interface{} {
	if o == nil || IsNil(o.Request) {
		var ret map[string]interface{}
		return ret
	}
	return o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineEvent) GetRequestOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Request) {
		return map[string]interface{}{}, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *MachineEvent) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given map[string]interface{} and assigns it to the Request field.
func (o *MachineEvent) SetRequest(v map[string]interface{}) {
	o.Request = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *MachineEvent) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineEvent) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MachineEvent) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *MachineEvent) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MachineEvent) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineEvent) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MachineEvent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MachineEvent) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MachineEvent) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineEvent) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MachineEvent) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *MachineEvent) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MachineEvent) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineEvent) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MachineEvent) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MachineEvent) SetType(v string) {
	o.Type = &v
}

func (o MachineEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MachineEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMachineEvent struct {
	value *MachineEvent
	isSet bool
}

func (v NullableMachineEvent) Get() *MachineEvent {
	return v.value
}

func (v *NullableMachineEvent) Set(val *MachineEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineEvent(val *MachineEvent) *NullableMachineEvent {
	return &NullableMachineEvent{value: val, isSet: true}
}

func (v NullableMachineEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


