# ApiHTTPResponseOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewApiHTTPResponseOptions

`func NewApiHTTPResponseOptions() *ApiHTTPResponseOptions`

NewApiHTTPResponseOptions instantiates a new ApiHTTPResponseOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHTTPResponseOptionsWithDefaults

`func NewApiHTTPResponseOptionsWithDefaults() *ApiHTTPResponseOptions`

NewApiHTTPResponseOptionsWithDefaults instantiates a new ApiHTTPResponseOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *ApiHTTPResponseOptions) GetHeaders() map[string]map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ApiHTTPResponseOptions) GetHeadersOk() (*map[string]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ApiHTTPResponseOptions) SetHeaders(v map[string]map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ApiHTTPResponseOptions) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


