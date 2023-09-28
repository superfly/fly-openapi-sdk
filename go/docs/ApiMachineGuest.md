# ApiMachineGuest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuKind** | Pointer to **string** |  | [optional] 
**Cpus** | Pointer to **int32** |  | [optional] 
**GpuKind** | Pointer to **string** |  | [optional] 
**KernelArgs** | Pointer to **[]string** |  | [optional] 
**MemoryMb** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiMachineGuest

`func NewApiMachineGuest() *ApiMachineGuest`

NewApiMachineGuest instantiates a new ApiMachineGuest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMachineGuestWithDefaults

`func NewApiMachineGuestWithDefaults() *ApiMachineGuest`

NewApiMachineGuestWithDefaults instantiates a new ApiMachineGuest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuKind

`func (o *ApiMachineGuest) GetCpuKind() string`

GetCpuKind returns the CpuKind field if non-nil, zero value otherwise.

### GetCpuKindOk

`func (o *ApiMachineGuest) GetCpuKindOk() (*string, bool)`

GetCpuKindOk returns a tuple with the CpuKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuKind

`func (o *ApiMachineGuest) SetCpuKind(v string)`

SetCpuKind sets CpuKind field to given value.

### HasCpuKind

`func (o *ApiMachineGuest) HasCpuKind() bool`

HasCpuKind returns a boolean if a field has been set.

### GetCpus

`func (o *ApiMachineGuest) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *ApiMachineGuest) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *ApiMachineGuest) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *ApiMachineGuest) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetGpuKind

`func (o *ApiMachineGuest) GetGpuKind() string`

GetGpuKind returns the GpuKind field if non-nil, zero value otherwise.

### GetGpuKindOk

`func (o *ApiMachineGuest) GetGpuKindOk() (*string, bool)`

GetGpuKindOk returns a tuple with the GpuKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuKind

`func (o *ApiMachineGuest) SetGpuKind(v string)`

SetGpuKind sets GpuKind field to given value.

### HasGpuKind

`func (o *ApiMachineGuest) HasGpuKind() bool`

HasGpuKind returns a boolean if a field has been set.

### GetKernelArgs

`func (o *ApiMachineGuest) GetKernelArgs() []string`

GetKernelArgs returns the KernelArgs field if non-nil, zero value otherwise.

### GetKernelArgsOk

`func (o *ApiMachineGuest) GetKernelArgsOk() (*[]string, bool)`

GetKernelArgsOk returns a tuple with the KernelArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelArgs

`func (o *ApiMachineGuest) SetKernelArgs(v []string)`

SetKernelArgs sets KernelArgs field to given value.

### HasKernelArgs

`func (o *ApiMachineGuest) HasKernelArgs() bool`

HasKernelArgs returns a boolean if a field has been set.

### GetMemoryMb

`func (o *ApiMachineGuest) GetMemoryMb() int32`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *ApiMachineGuest) GetMemoryMbOk() (*int32, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *ApiMachineGuest) SetMemoryMb(v int32)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *ApiMachineGuest) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


