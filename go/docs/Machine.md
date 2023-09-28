# Machine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | Pointer to [**[]CheckStatus**](CheckStatus.md) |  | [optional] 
**Config** | Pointer to [**ApiMachineConfig**](ApiMachineConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Events** | Pointer to [**[]MachineEvent**](MachineEvent.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImageRef** | Pointer to [**ImageRef**](ImageRef.md) |  | [optional] 
**InstanceId** | Pointer to **string** | InstanceID is unique for each version of the machine | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** | Nonce is only every returned on machine creation if a lease_duration was provided. | [optional] 
**PrivateIp** | Pointer to **string** | PrivateIP is the internal 6PN address of the machine. | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewMachine

`func NewMachine() *Machine`

NewMachine instantiates a new Machine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineWithDefaults

`func NewMachineWithDefaults() *Machine`

NewMachineWithDefaults instantiates a new Machine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *Machine) GetChecks() []CheckStatus`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *Machine) GetChecksOk() (*[]CheckStatus, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *Machine) SetChecks(v []CheckStatus)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *Machine) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetConfig

`func (o *Machine) GetConfig() ApiMachineConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Machine) GetConfigOk() (*ApiMachineConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Machine) SetConfig(v ApiMachineConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Machine) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Machine) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Machine) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Machine) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Machine) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvents

`func (o *Machine) GetEvents() []MachineEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Machine) GetEventsOk() (*[]MachineEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Machine) SetEvents(v []MachineEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Machine) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *Machine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Machine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Machine) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Machine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageRef

`func (o *Machine) GetImageRef() ImageRef`

GetImageRef returns the ImageRef field if non-nil, zero value otherwise.

### GetImageRefOk

`func (o *Machine) GetImageRefOk() (*ImageRef, bool)`

GetImageRefOk returns a tuple with the ImageRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRef

`func (o *Machine) SetImageRef(v ImageRef)`

SetImageRef sets ImageRef field to given value.

### HasImageRef

`func (o *Machine) HasImageRef() bool`

HasImageRef returns a boolean if a field has been set.

### GetInstanceId

`func (o *Machine) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Machine) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Machine) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *Machine) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetName

`func (o *Machine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Machine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Machine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Machine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNonce

`func (o *Machine) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *Machine) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *Machine) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *Machine) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetPrivateIp

`func (o *Machine) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *Machine) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *Machine) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *Machine) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetRegion

`func (o *Machine) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Machine) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Machine) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Machine) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetState

`func (o *Machine) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Machine) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Machine) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Machine) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Machine) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Machine) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Machine) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Machine) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


