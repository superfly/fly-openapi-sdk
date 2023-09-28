# ApiMachineProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cmd** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **[]string** |  | [optional] 
**Env** | Pointer to **map[string]string** |  | [optional] 
**Exec** | Pointer to **[]string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewApiMachineProcess

`func NewApiMachineProcess() *ApiMachineProcess`

NewApiMachineProcess instantiates a new ApiMachineProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMachineProcessWithDefaults

`func NewApiMachineProcessWithDefaults() *ApiMachineProcess`

NewApiMachineProcessWithDefaults instantiates a new ApiMachineProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCmd

`func (o *ApiMachineProcess) GetCmd() []string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *ApiMachineProcess) GetCmdOk() (*[]string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *ApiMachineProcess) SetCmd(v []string)`

SetCmd sets Cmd field to given value.

### HasCmd

`func (o *ApiMachineProcess) HasCmd() bool`

HasCmd returns a boolean if a field has been set.

### GetEntrypoint

`func (o *ApiMachineProcess) GetEntrypoint() []string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *ApiMachineProcess) GetEntrypointOk() (*[]string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *ApiMachineProcess) SetEntrypoint(v []string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *ApiMachineProcess) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetEnv

`func (o *ApiMachineProcess) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ApiMachineProcess) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ApiMachineProcess) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ApiMachineProcess) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetExec

`func (o *ApiMachineProcess) GetExec() []string`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *ApiMachineProcess) GetExecOk() (*[]string, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *ApiMachineProcess) SetExec(v []string)`

SetExec sets Exec field to given value.

### HasExec

`func (o *ApiMachineProcess) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetUser

`func (o *ApiMachineProcess) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApiMachineProcess) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApiMachineProcess) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ApiMachineProcess) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


