# MachineExecRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cmd** | Pointer to **string** | Deprecated: use Command instead | [optional] 
**Command** | Pointer to **[]string** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewMachineExecRequest

`func NewMachineExecRequest() *MachineExecRequest`

NewMachineExecRequest instantiates a new MachineExecRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineExecRequestWithDefaults

`func NewMachineExecRequestWithDefaults() *MachineExecRequest`

NewMachineExecRequestWithDefaults instantiates a new MachineExecRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCmd

`func (o *MachineExecRequest) GetCmd() string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *MachineExecRequest) GetCmdOk() (*string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *MachineExecRequest) SetCmd(v string)`

SetCmd sets Cmd field to given value.

### HasCmd

`func (o *MachineExecRequest) HasCmd() bool`

HasCmd returns a boolean if a field has been set.

### GetCommand

`func (o *MachineExecRequest) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *MachineExecRequest) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *MachineExecRequest) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *MachineExecRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetTimeout

`func (o *MachineExecRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *MachineExecRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *MachineExecRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *MachineExecRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


