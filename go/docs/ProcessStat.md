# ProcessStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Directory** | Pointer to **string** |  | [optional] 
**ListenSockets** | Pointer to [**[]ListenSocket**](ListenSocket.md) |  | [optional] 
**Pid** | Pointer to **int32** |  | [optional] 
**Rss** | Pointer to **int32** |  | [optional] 
**Rtime** | Pointer to **int32** |  | [optional] 
**Stime** | Pointer to **int32** |  | [optional] 

## Methods

### NewProcessStat

`func NewProcessStat() *ProcessStat`

NewProcessStat instantiates a new ProcessStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessStatWithDefaults

`func NewProcessStatWithDefaults() *ProcessStat`

NewProcessStatWithDefaults instantiates a new ProcessStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *ProcessStat) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ProcessStat) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ProcessStat) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ProcessStat) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetCpu

`func (o *ProcessStat) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ProcessStat) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ProcessStat) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ProcessStat) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDirectory

`func (o *ProcessStat) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *ProcessStat) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *ProcessStat) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *ProcessStat) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetListenSockets

`func (o *ProcessStat) GetListenSockets() []ListenSocket`

GetListenSockets returns the ListenSockets field if non-nil, zero value otherwise.

### GetListenSocketsOk

`func (o *ProcessStat) GetListenSocketsOk() (*[]ListenSocket, bool)`

GetListenSocketsOk returns a tuple with the ListenSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenSockets

`func (o *ProcessStat) SetListenSockets(v []ListenSocket)`

SetListenSockets sets ListenSockets field to given value.

### HasListenSockets

`func (o *ProcessStat) HasListenSockets() bool`

HasListenSockets returns a boolean if a field has been set.

### GetPid

`func (o *ProcessStat) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ProcessStat) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ProcessStat) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *ProcessStat) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetRss

`func (o *ProcessStat) GetRss() int32`

GetRss returns the Rss field if non-nil, zero value otherwise.

### GetRssOk

`func (o *ProcessStat) GetRssOk() (*int32, bool)`

GetRssOk returns a tuple with the Rss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRss

`func (o *ProcessStat) SetRss(v int32)`

SetRss sets Rss field to given value.

### HasRss

`func (o *ProcessStat) HasRss() bool`

HasRss returns a boolean if a field has been set.

### GetRtime

`func (o *ProcessStat) GetRtime() int32`

GetRtime returns the Rtime field if non-nil, zero value otherwise.

### GetRtimeOk

`func (o *ProcessStat) GetRtimeOk() (*int32, bool)`

GetRtimeOk returns a tuple with the Rtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtime

`func (o *ProcessStat) SetRtime(v int32)`

SetRtime sets Rtime field to given value.

### HasRtime

`func (o *ProcessStat) HasRtime() bool`

HasRtime returns a boolean if a field has been set.

### GetStime

`func (o *ProcessStat) GetStime() int32`

GetStime returns the Stime field if non-nil, zero value otherwise.

### GetStimeOk

`func (o *ProcessStat) GetStimeOk() (*int32, bool)`

GetStimeOk returns a tuple with the Stime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStime

`func (o *ProcessStat) SetStime(v int32)`

SetStime sets Stime field to given value.

### HasStime

`func (o *ProcessStat) HasStime() bool`

HasStime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


