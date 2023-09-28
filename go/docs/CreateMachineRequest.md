# CreateMachineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ApiMachineConfig**](ApiMachineConfig.md) |  | [optional] 
**LeaseTtl** | Pointer to **int32** |  | [optional] 
**Lsvd** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SkipLaunch** | Pointer to **bool** |  | [optional] 
**SkipServiceRegistration** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateMachineRequest

`func NewCreateMachineRequest() *CreateMachineRequest`

NewCreateMachineRequest instantiates a new CreateMachineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMachineRequestWithDefaults

`func NewCreateMachineRequestWithDefaults() *CreateMachineRequest`

NewCreateMachineRequestWithDefaults instantiates a new CreateMachineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CreateMachineRequest) GetConfig() ApiMachineConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateMachineRequest) GetConfigOk() (*ApiMachineConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateMachineRequest) SetConfig(v ApiMachineConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateMachineRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLeaseTtl

`func (o *CreateMachineRequest) GetLeaseTtl() int32`

GetLeaseTtl returns the LeaseTtl field if non-nil, zero value otherwise.

### GetLeaseTtlOk

`func (o *CreateMachineRequest) GetLeaseTtlOk() (*int32, bool)`

GetLeaseTtlOk returns a tuple with the LeaseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTtl

`func (o *CreateMachineRequest) SetLeaseTtl(v int32)`

SetLeaseTtl sets LeaseTtl field to given value.

### HasLeaseTtl

`func (o *CreateMachineRequest) HasLeaseTtl() bool`

HasLeaseTtl returns a boolean if a field has been set.

### GetLsvd

`func (o *CreateMachineRequest) GetLsvd() bool`

GetLsvd returns the Lsvd field if non-nil, zero value otherwise.

### GetLsvdOk

`func (o *CreateMachineRequest) GetLsvdOk() (*bool, bool)`

GetLsvdOk returns a tuple with the Lsvd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLsvd

`func (o *CreateMachineRequest) SetLsvd(v bool)`

SetLsvd sets Lsvd field to given value.

### HasLsvd

`func (o *CreateMachineRequest) HasLsvd() bool`

HasLsvd returns a boolean if a field has been set.

### GetName

`func (o *CreateMachineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMachineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMachineRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateMachineRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *CreateMachineRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateMachineRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateMachineRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateMachineRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSkipLaunch

`func (o *CreateMachineRequest) GetSkipLaunch() bool`

GetSkipLaunch returns the SkipLaunch field if non-nil, zero value otherwise.

### GetSkipLaunchOk

`func (o *CreateMachineRequest) GetSkipLaunchOk() (*bool, bool)`

GetSkipLaunchOk returns a tuple with the SkipLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipLaunch

`func (o *CreateMachineRequest) SetSkipLaunch(v bool)`

SetSkipLaunch sets SkipLaunch field to given value.

### HasSkipLaunch

`func (o *CreateMachineRequest) HasSkipLaunch() bool`

HasSkipLaunch returns a boolean if a field has been set.

### GetSkipServiceRegistration

`func (o *CreateMachineRequest) GetSkipServiceRegistration() bool`

GetSkipServiceRegistration returns the SkipServiceRegistration field if non-nil, zero value otherwise.

### GetSkipServiceRegistrationOk

`func (o *CreateMachineRequest) GetSkipServiceRegistrationOk() (*bool, bool)`

GetSkipServiceRegistrationOk returns a tuple with the SkipServiceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipServiceRegistration

`func (o *CreateMachineRequest) SetSkipServiceRegistration(v bool)`

SetSkipServiceRegistration sets SkipServiceRegistration field to given value.

### HasSkipServiceRegistration

`func (o *CreateMachineRequest) HasSkipServiceRegistration() bool`

HasSkipServiceRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


