# CreateAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**OrgSlug** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateAppRequest

`func NewCreateAppRequest() *CreateAppRequest`

NewCreateAppRequest instantiates a new CreateAppRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppRequestWithDefaults

`func NewCreateAppRequestWithDefaults() *CreateAppRequest`

NewCreateAppRequestWithDefaults instantiates a new CreateAppRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *CreateAppRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *CreateAppRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *CreateAppRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *CreateAppRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetNetwork

`func (o *CreateAppRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateAppRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateAppRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateAppRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOrgSlug

`func (o *CreateAppRequest) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *CreateAppRequest) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *CreateAppRequest) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.

### HasOrgSlug

`func (o *CreateAppRequest) HasOrgSlug() bool`

HasOrgSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


