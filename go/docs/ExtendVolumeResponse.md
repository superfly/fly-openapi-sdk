# ExtendVolumeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeedsRestart** | Pointer to **bool** |  | [optional] 
**Volume** | Pointer to [**Volume**](Volume.md) |  | [optional] 

## Methods

### NewExtendVolumeResponse

`func NewExtendVolumeResponse() *ExtendVolumeResponse`

NewExtendVolumeResponse instantiates a new ExtendVolumeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendVolumeResponseWithDefaults

`func NewExtendVolumeResponseWithDefaults() *ExtendVolumeResponse`

NewExtendVolumeResponseWithDefaults instantiates a new ExtendVolumeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeedsRestart

`func (o *ExtendVolumeResponse) GetNeedsRestart() bool`

GetNeedsRestart returns the NeedsRestart field if non-nil, zero value otherwise.

### GetNeedsRestartOk

`func (o *ExtendVolumeResponse) GetNeedsRestartOk() (*bool, bool)`

GetNeedsRestartOk returns a tuple with the NeedsRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRestart

`func (o *ExtendVolumeResponse) SetNeedsRestart(v bool)`

SetNeedsRestart sets NeedsRestart field to given value.

### HasNeedsRestart

`func (o *ExtendVolumeResponse) HasNeedsRestart() bool`

HasNeedsRestart returns a boolean if a field has been set.

### GetVolume

`func (o *ExtendVolumeResponse) GetVolume() Volume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ExtendVolumeResponse) GetVolumeOk() (*Volume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ExtendVolumeResponse) SetVolume(v Volume)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ExtendVolumeResponse) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


