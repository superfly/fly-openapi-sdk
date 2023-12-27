# ApiHTTPOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compress** | Pointer to **bool** |  | [optional] 
**H2Backend** | Pointer to **bool** |  | [optional] 
**Response** | Pointer to [**ApiHTTPResponseOptions**](ApiHTTPResponseOptions.md) |  | [optional] 

## Methods

### NewApiHTTPOptions

`func NewApiHTTPOptions() *ApiHTTPOptions`

NewApiHTTPOptions instantiates a new ApiHTTPOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHTTPOptionsWithDefaults

`func NewApiHTTPOptionsWithDefaults() *ApiHTTPOptions`

NewApiHTTPOptionsWithDefaults instantiates a new ApiHTTPOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompress

`func (o *ApiHTTPOptions) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *ApiHTTPOptions) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *ApiHTTPOptions) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *ApiHTTPOptions) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetH2Backend

`func (o *ApiHTTPOptions) GetH2Backend() bool`

GetH2Backend returns the H2Backend field if non-nil, zero value otherwise.

### GetH2BackendOk

`func (o *ApiHTTPOptions) GetH2BackendOk() (*bool, bool)`

GetH2BackendOk returns a tuple with the H2Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH2Backend

`func (o *ApiHTTPOptions) SetH2Backend(v bool)`

SetH2Backend sets H2Backend field to given value.

### HasH2Backend

`func (o *ApiHTTPOptions) HasH2Backend() bool`

HasH2Backend returns a boolean if a field has been set.

### GetResponse

`func (o *ApiHTTPOptions) GetResponse() ApiHTTPResponseOptions`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ApiHTTPOptions) GetResponseOk() (*ApiHTTPResponseOptions, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ApiHTTPOptions) SetResponse(v ApiHTTPResponseOptions)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ApiHTTPOptions) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


