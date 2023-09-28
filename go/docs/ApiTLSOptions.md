# ApiTLSOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alpn** | Pointer to **[]string** |  | [optional] 
**DefaultSelfSigned** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiTLSOptions

`func NewApiTLSOptions() *ApiTLSOptions`

NewApiTLSOptions instantiates a new ApiTLSOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTLSOptionsWithDefaults

`func NewApiTLSOptionsWithDefaults() *ApiTLSOptions`

NewApiTLSOptionsWithDefaults instantiates a new ApiTLSOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpn

`func (o *ApiTLSOptions) GetAlpn() []string`

GetAlpn returns the Alpn field if non-nil, zero value otherwise.

### GetAlpnOk

`func (o *ApiTLSOptions) GetAlpnOk() (*[]string, bool)`

GetAlpnOk returns a tuple with the Alpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpn

`func (o *ApiTLSOptions) SetAlpn(v []string)`

SetAlpn sets Alpn field to given value.

### HasAlpn

`func (o *ApiTLSOptions) HasAlpn() bool`

HasAlpn returns a boolean if a field has been set.

### GetDefaultSelfSigned

`func (o *ApiTLSOptions) GetDefaultSelfSigned() bool`

GetDefaultSelfSigned returns the DefaultSelfSigned field if non-nil, zero value otherwise.

### GetDefaultSelfSignedOk

`func (o *ApiTLSOptions) GetDefaultSelfSignedOk() (*bool, bool)`

GetDefaultSelfSignedOk returns a tuple with the DefaultSelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSelfSigned

`func (o *ApiTLSOptions) SetDefaultSelfSigned(v bool)`

SetDefaultSelfSigned sets DefaultSelfSigned field to given value.

### HasDefaultSelfSigned

`func (o *ApiTLSOptions) HasDefaultSelfSigned() bool`

HasDefaultSelfSigned returns a boolean if a field has been set.

### GetVersions

`func (o *ApiTLSOptions) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ApiTLSOptions) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ApiTLSOptions) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ApiTLSOptions) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


