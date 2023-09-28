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

// checks if the ImageRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageRef{}

// ImageRef struct for ImageRef
type ImageRef struct {
	Digest *string `json:"digest,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
	Registry *string `json:"registry,omitempty"`
	Repository *string `json:"repository,omitempty"`
	Tag *string `json:"tag,omitempty"`
}

// NewImageRef instantiates a new ImageRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageRef() *ImageRef {
	this := ImageRef{}
	return &this
}

// NewImageRefWithDefaults instantiates a new ImageRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageRefWithDefaults() *ImageRef {
	this := ImageRef{}
	return &this
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *ImageRef) GetDigest() string {
	if o == nil || IsNil(o.Digest) {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageRef) GetDigestOk() (*string, bool) {
	if o == nil || IsNil(o.Digest) {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *ImageRef) HasDigest() bool {
	if o != nil && !IsNil(o.Digest) {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *ImageRef) SetDigest(v string) {
	o.Digest = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ImageRef) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageRef) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ImageRef) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *ImageRef) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *ImageRef) GetRegistry() string {
	if o == nil || IsNil(o.Registry) {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageRef) GetRegistryOk() (*string, bool) {
	if o == nil || IsNil(o.Registry) {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *ImageRef) HasRegistry() bool {
	if o != nil && !IsNil(o.Registry) {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *ImageRef) SetRegistry(v string) {
	o.Registry = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ImageRef) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageRef) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ImageRef) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *ImageRef) SetRepository(v string) {
	o.Repository = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ImageRef) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageRef) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ImageRef) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *ImageRef) SetTag(v string) {
	o.Tag = &v
}

func (o ImageRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Digest) {
		toSerialize["digest"] = o.Digest
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Registry) {
		toSerialize["registry"] = o.Registry
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

type NullableImageRef struct {
	value *ImageRef
	isSet bool
}

func (v NullableImageRef) Get() *ImageRef {
	return v.value
}

func (v *NullableImageRef) Set(val *ImageRef) {
	v.value = val
	v.isSet = true
}

func (v NullableImageRef) IsSet() bool {
	return v.isSet
}

func (v *NullableImageRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageRef(val *ImageRef) *NullableImageRef {
	return &NullableImageRef{value: val, isSet: true}
}

func (v NullableImageRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


