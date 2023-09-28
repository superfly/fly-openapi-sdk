# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachedAllocId** | Pointer to **string** |  | [optional] 
**AttachedMachineId** | Pointer to **string** |  | [optional] 
**BlockSize** | Pointer to **int32** |  | [optional] 
**Blocks** | Pointer to **int32** |  | [optional] 
**BlocksAvail** | Pointer to **int32** |  | [optional] 
**BlocksFree** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Encrypted** | Pointer to **bool** |  | [optional] 
**Fstype** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SizeGb** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewVolume

`func NewVolume() *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachedAllocId

`func (o *Volume) GetAttachedAllocId() string`

GetAttachedAllocId returns the AttachedAllocId field if non-nil, zero value otherwise.

### GetAttachedAllocIdOk

`func (o *Volume) GetAttachedAllocIdOk() (*string, bool)`

GetAttachedAllocIdOk returns a tuple with the AttachedAllocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedAllocId

`func (o *Volume) SetAttachedAllocId(v string)`

SetAttachedAllocId sets AttachedAllocId field to given value.

### HasAttachedAllocId

`func (o *Volume) HasAttachedAllocId() bool`

HasAttachedAllocId returns a boolean if a field has been set.

### GetAttachedMachineId

`func (o *Volume) GetAttachedMachineId() string`

GetAttachedMachineId returns the AttachedMachineId field if non-nil, zero value otherwise.

### GetAttachedMachineIdOk

`func (o *Volume) GetAttachedMachineIdOk() (*string, bool)`

GetAttachedMachineIdOk returns a tuple with the AttachedMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedMachineId

`func (o *Volume) SetAttachedMachineId(v string)`

SetAttachedMachineId sets AttachedMachineId field to given value.

### HasAttachedMachineId

`func (o *Volume) HasAttachedMachineId() bool`

HasAttachedMachineId returns a boolean if a field has been set.

### GetBlockSize

`func (o *Volume) GetBlockSize() int32`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *Volume) GetBlockSizeOk() (*int32, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *Volume) SetBlockSize(v int32)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *Volume) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetBlocks

`func (o *Volume) GetBlocks() int32`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *Volume) GetBlocksOk() (*int32, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *Volume) SetBlocks(v int32)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *Volume) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetBlocksAvail

`func (o *Volume) GetBlocksAvail() int32`

GetBlocksAvail returns the BlocksAvail field if non-nil, zero value otherwise.

### GetBlocksAvailOk

`func (o *Volume) GetBlocksAvailOk() (*int32, bool)`

GetBlocksAvailOk returns a tuple with the BlocksAvail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksAvail

`func (o *Volume) SetBlocksAvail(v int32)`

SetBlocksAvail sets BlocksAvail field to given value.

### HasBlocksAvail

`func (o *Volume) HasBlocksAvail() bool`

HasBlocksAvail returns a boolean if a field has been set.

### GetBlocksFree

`func (o *Volume) GetBlocksFree() int32`

GetBlocksFree returns the BlocksFree field if non-nil, zero value otherwise.

### GetBlocksFreeOk

`func (o *Volume) GetBlocksFreeOk() (*int32, bool)`

GetBlocksFreeOk returns a tuple with the BlocksFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksFree

`func (o *Volume) SetBlocksFree(v int32)`

SetBlocksFree sets BlocksFree field to given value.

### HasBlocksFree

`func (o *Volume) HasBlocksFree() bool`

HasBlocksFree returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Volume) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Volume) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Volume) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Volume) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEncrypted

`func (o *Volume) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *Volume) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *Volume) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *Volume) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetFstype

`func (o *Volume) GetFstype() string`

GetFstype returns the Fstype field if non-nil, zero value otherwise.

### GetFstypeOk

`func (o *Volume) GetFstypeOk() (*string, bool)`

GetFstypeOk returns a tuple with the Fstype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFstype

`func (o *Volume) SetFstype(v string)`

SetFstype sets Fstype field to given value.

### HasFstype

`func (o *Volume) HasFstype() bool`

HasFstype returns a boolean if a field has been set.

### GetId

`func (o *Volume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Volume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Volume) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Volume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Volume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Volume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *Volume) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Volume) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Volume) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Volume) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSizeGb

`func (o *Volume) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *Volume) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *Volume) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.

### HasSizeGb

`func (o *Volume) HasSizeGb() bool`

HasSizeGb returns a boolean if a field has been set.

### GetState

`func (o *Volume) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Volume) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Volume) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Volume) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZone

`func (o *Volume) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Volume) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Volume) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Volume) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


