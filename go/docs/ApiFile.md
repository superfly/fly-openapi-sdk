# ApiFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuestPath** | Pointer to **string** | GuestPath is the path on the machine where the file will be written and must be an absolute path. For example: /full/path/to/file.json | [optional] 
**RawValue** | Pointer to **string** | The base64 encoded string of the file contents. | [optional] 
**SecretName** | Pointer to **string** | The name of the secret that contains the base64 encoded file contents. | [optional] 

## Methods

### NewApiFile

`func NewApiFile() *ApiFile`

NewApiFile instantiates a new ApiFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFileWithDefaults

`func NewApiFileWithDefaults() *ApiFile`

NewApiFileWithDefaults instantiates a new ApiFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuestPath

`func (o *ApiFile) GetGuestPath() string`

GetGuestPath returns the GuestPath field if non-nil, zero value otherwise.

### GetGuestPathOk

`func (o *ApiFile) GetGuestPathOk() (*string, bool)`

GetGuestPathOk returns a tuple with the GuestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPath

`func (o *ApiFile) SetGuestPath(v string)`

SetGuestPath sets GuestPath field to given value.

### HasGuestPath

`func (o *ApiFile) HasGuestPath() bool`

HasGuestPath returns a boolean if a field has been set.

### GetRawValue

`func (o *ApiFile) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *ApiFile) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *ApiFile) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *ApiFile) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### GetSecretName

`func (o *ApiFile) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *ApiFile) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *ApiFile) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.

### HasSecretName

`func (o *ApiFile) HasSecretName() bool`

HasSecretName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


