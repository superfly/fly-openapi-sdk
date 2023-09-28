# ApiMachineRestart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxRetries** | Pointer to **int32** | MaxRetries is only relevant with the on-failure policy. | [optional] 
**Policy** | Pointer to **string** |  | [optional] 

## Methods

### NewApiMachineRestart

`func NewApiMachineRestart() *ApiMachineRestart`

NewApiMachineRestart instantiates a new ApiMachineRestart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMachineRestartWithDefaults

`func NewApiMachineRestartWithDefaults() *ApiMachineRestart`

NewApiMachineRestartWithDefaults instantiates a new ApiMachineRestart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxRetries

`func (o *ApiMachineRestart) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *ApiMachineRestart) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *ApiMachineRestart) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *ApiMachineRestart) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetPolicy

`func (o *ApiMachineRestart) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ApiMachineRestart) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ApiMachineRestart) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ApiMachineRestart) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


