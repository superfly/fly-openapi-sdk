# ImageRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Registry** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewImageRef

`func NewImageRef() *ImageRef`

NewImageRef instantiates a new ImageRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageRefWithDefaults

`func NewImageRefWithDefaults() *ImageRef`

NewImageRefWithDefaults instantiates a new ImageRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *ImageRef) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ImageRef) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ImageRef) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ImageRef) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetLabels

`func (o *ImageRef) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ImageRef) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ImageRef) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ImageRef) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetRegistry

`func (o *ImageRef) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *ImageRef) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *ImageRef) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *ImageRef) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRepository

`func (o *ImageRef) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ImageRef) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ImageRef) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ImageRef) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetTag

`func (o *ImageRef) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ImageRef) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ImageRef) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ImageRef) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


