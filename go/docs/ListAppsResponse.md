# ListAppsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]ListApp**](ListApp.md) |  | [optional] 
**TotalApps** | Pointer to **int32** |  | [optional] 

## Methods

### NewListAppsResponse

`func NewListAppsResponse() *ListAppsResponse`

NewListAppsResponse instantiates a new ListAppsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAppsResponseWithDefaults

`func NewListAppsResponseWithDefaults() *ListAppsResponse`

NewListAppsResponseWithDefaults instantiates a new ListAppsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *ListAppsResponse) GetApps() []ListApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ListAppsResponse) GetAppsOk() (*[]ListApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ListAppsResponse) SetApps(v []ListApp)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ListAppsResponse) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetTotalApps

`func (o *ListAppsResponse) GetTotalApps() int32`

GetTotalApps returns the TotalApps field if non-nil, zero value otherwise.

### GetTotalAppsOk

`func (o *ListAppsResponse) GetTotalAppsOk() (*int32, bool)`

GetTotalAppsOk returns a tuple with the TotalApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApps

`func (o *ListAppsResponse) SetTotalApps(v int32)`

SetTotalApps sets TotalApps field to given value.

### HasTotalApps

`func (o *ListAppsResponse) HasTotalApps() bool`

HasTotalApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


