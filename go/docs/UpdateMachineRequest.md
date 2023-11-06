# UpdateMachineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ApiMachineConfig**](ApiMachineConfig.md) |  | [optional] 
**CurrentVersion** | Pointer to **string** |  | [optional] 
**LeaseTtl** | Pointer to **int32** |  | [optional] 
**Lsvd** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | Unique name for this Machine. If omitted, one is generated for you | [optional] 
**Region** | Pointer to **string** | The target region. Omitting this param launches in the same region as your WireGuard peer connection (somewhere near you). | [optional] 
**SkipLaunch** | Pointer to **bool** |  | [optional] 
**SkipServiceRegistration** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateMachineRequest

`func NewUpdateMachineRequest() *UpdateMachineRequest`

NewUpdateMachineRequest instantiates a new UpdateMachineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMachineRequestWithDefaults

`func NewUpdateMachineRequestWithDefaults() *UpdateMachineRequest`

NewUpdateMachineRequestWithDefaults instantiates a new UpdateMachineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *UpdateMachineRequest) GetConfig() ApiMachineConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateMachineRequest) GetConfigOk() (*ApiMachineConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateMachineRequest) SetConfig(v ApiMachineConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateMachineRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *UpdateMachineRequest) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *UpdateMachineRequest) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *UpdateMachineRequest) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *UpdateMachineRequest) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetLeaseTtl

`func (o *UpdateMachineRequest) GetLeaseTtl() int32`

GetLeaseTtl returns the LeaseTtl field if non-nil, zero value otherwise.

### GetLeaseTtlOk

`func (o *UpdateMachineRequest) GetLeaseTtlOk() (*int32, bool)`

GetLeaseTtlOk returns a tuple with the LeaseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTtl

`func (o *UpdateMachineRequest) SetLeaseTtl(v int32)`

SetLeaseTtl sets LeaseTtl field to given value.

### HasLeaseTtl

`func (o *UpdateMachineRequest) HasLeaseTtl() bool`

HasLeaseTtl returns a boolean if a field has been set.

### GetLsvd

`func (o *UpdateMachineRequest) GetLsvd() bool`

GetLsvd returns the Lsvd field if non-nil, zero value otherwise.

### GetLsvdOk

`func (o *UpdateMachineRequest) GetLsvdOk() (*bool, bool)`

GetLsvdOk returns a tuple with the Lsvd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLsvd

`func (o *UpdateMachineRequest) SetLsvd(v bool)`

SetLsvd sets Lsvd field to given value.

### HasLsvd

`func (o *UpdateMachineRequest) HasLsvd() bool`

HasLsvd returns a boolean if a field has been set.

### GetName

`func (o *UpdateMachineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMachineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMachineRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMachineRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *UpdateMachineRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UpdateMachineRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UpdateMachineRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UpdateMachineRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSkipLaunch

`func (o *UpdateMachineRequest) GetSkipLaunch() bool`

GetSkipLaunch returns the SkipLaunch field if non-nil, zero value otherwise.

### GetSkipLaunchOk

`func (o *UpdateMachineRequest) GetSkipLaunchOk() (*bool, bool)`

GetSkipLaunchOk returns a tuple with the SkipLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipLaunch

`func (o *UpdateMachineRequest) SetSkipLaunch(v bool)`

SetSkipLaunch sets SkipLaunch field to given value.

### HasSkipLaunch

`func (o *UpdateMachineRequest) HasSkipLaunch() bool`

HasSkipLaunch returns a boolean if a field has been set.

### GetSkipServiceRegistration

`func (o *UpdateMachineRequest) GetSkipServiceRegistration() bool`

GetSkipServiceRegistration returns the SkipServiceRegistration field if non-nil, zero value otherwise.

### GetSkipServiceRegistrationOk

`func (o *UpdateMachineRequest) GetSkipServiceRegistrationOk() (*bool, bool)`

GetSkipServiceRegistrationOk returns a tuple with the SkipServiceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipServiceRegistration

`func (o *UpdateMachineRequest) SetSkipServiceRegistration(v bool)`

SetSkipServiceRegistration sets SkipServiceRegistration field to given value.

### HasSkipServiceRegistration

`func (o *UpdateMachineRequest) HasSkipServiceRegistration() bool`

HasSkipServiceRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


