# ApiMachineService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autostart** | Pointer to **bool** |  | [optional] 
**Autostop** | Pointer to **bool** |  | [optional] 
**Checks** | Pointer to [**[]ApiMachineCheck**](ApiMachineCheck.md) |  | [optional] 
**Concurrency** | Pointer to [**ApiMachineServiceConcurrency**](ApiMachineServiceConcurrency.md) |  | [optional] 
**ForceInstanceDescription** | Pointer to **string** |  | [optional] 
**ForceInstanceKey** | Pointer to **string** |  | [optional] 
**InternalPort** | Pointer to **int32** |  | [optional] 
**MinMachinesRunning** | Pointer to **int32** |  | [optional] 
**Ports** | Pointer to [**[]ApiMachinePort**](ApiMachinePort.md) |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 

## Methods

### NewApiMachineService

`func NewApiMachineService() *ApiMachineService`

NewApiMachineService instantiates a new ApiMachineService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMachineServiceWithDefaults

`func NewApiMachineServiceWithDefaults() *ApiMachineService`

NewApiMachineServiceWithDefaults instantiates a new ApiMachineService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutostart

`func (o *ApiMachineService) GetAutostart() bool`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *ApiMachineService) GetAutostartOk() (*bool, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *ApiMachineService) SetAutostart(v bool)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *ApiMachineService) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetAutostop

`func (o *ApiMachineService) GetAutostop() bool`

GetAutostop returns the Autostop field if non-nil, zero value otherwise.

### GetAutostopOk

`func (o *ApiMachineService) GetAutostopOk() (*bool, bool)`

GetAutostopOk returns a tuple with the Autostop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostop

`func (o *ApiMachineService) SetAutostop(v bool)`

SetAutostop sets Autostop field to given value.

### HasAutostop

`func (o *ApiMachineService) HasAutostop() bool`

HasAutostop returns a boolean if a field has been set.

### GetChecks

`func (o *ApiMachineService) GetChecks() []ApiMachineCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ApiMachineService) GetChecksOk() (*[]ApiMachineCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ApiMachineService) SetChecks(v []ApiMachineCheck)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ApiMachineService) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetConcurrency

`func (o *ApiMachineService) GetConcurrency() ApiMachineServiceConcurrency`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *ApiMachineService) GetConcurrencyOk() (*ApiMachineServiceConcurrency, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *ApiMachineService) SetConcurrency(v ApiMachineServiceConcurrency)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *ApiMachineService) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetForceInstanceDescription

`func (o *ApiMachineService) GetForceInstanceDescription() string`

GetForceInstanceDescription returns the ForceInstanceDescription field if non-nil, zero value otherwise.

### GetForceInstanceDescriptionOk

`func (o *ApiMachineService) GetForceInstanceDescriptionOk() (*string, bool)`

GetForceInstanceDescriptionOk returns a tuple with the ForceInstanceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceInstanceDescription

`func (o *ApiMachineService) SetForceInstanceDescription(v string)`

SetForceInstanceDescription sets ForceInstanceDescription field to given value.

### HasForceInstanceDescription

`func (o *ApiMachineService) HasForceInstanceDescription() bool`

HasForceInstanceDescription returns a boolean if a field has been set.

### GetForceInstanceKey

`func (o *ApiMachineService) GetForceInstanceKey() string`

GetForceInstanceKey returns the ForceInstanceKey field if non-nil, zero value otherwise.

### GetForceInstanceKeyOk

`func (o *ApiMachineService) GetForceInstanceKeyOk() (*string, bool)`

GetForceInstanceKeyOk returns a tuple with the ForceInstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceInstanceKey

`func (o *ApiMachineService) SetForceInstanceKey(v string)`

SetForceInstanceKey sets ForceInstanceKey field to given value.

### HasForceInstanceKey

`func (o *ApiMachineService) HasForceInstanceKey() bool`

HasForceInstanceKey returns a boolean if a field has been set.

### GetInternalPort

`func (o *ApiMachineService) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *ApiMachineService) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *ApiMachineService) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.

### HasInternalPort

`func (o *ApiMachineService) HasInternalPort() bool`

HasInternalPort returns a boolean if a field has been set.

### GetMinMachinesRunning

`func (o *ApiMachineService) GetMinMachinesRunning() int32`

GetMinMachinesRunning returns the MinMachinesRunning field if non-nil, zero value otherwise.

### GetMinMachinesRunningOk

`func (o *ApiMachineService) GetMinMachinesRunningOk() (*int32, bool)`

GetMinMachinesRunningOk returns a tuple with the MinMachinesRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMachinesRunning

`func (o *ApiMachineService) SetMinMachinesRunning(v int32)`

SetMinMachinesRunning sets MinMachinesRunning field to given value.

### HasMinMachinesRunning

`func (o *ApiMachineService) HasMinMachinesRunning() bool`

HasMinMachinesRunning returns a boolean if a field has been set.

### GetPorts

`func (o *ApiMachineService) GetPorts() []ApiMachinePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ApiMachineService) GetPortsOk() (*[]ApiMachinePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ApiMachineService) SetPorts(v []ApiMachinePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ApiMachineService) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProtocol

`func (o *ApiMachineService) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApiMachineService) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApiMachineService) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApiMachineService) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


