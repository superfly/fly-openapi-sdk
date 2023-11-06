# ApiMachineConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDestroy** | Pointer to **bool** |  | [optional] 
**Checks** | Pointer to [**map[string]ApiMachineCheck**](ApiMachineCheck.md) |  | [optional] 
**DisableMachineAutostart** | Pointer to **bool** | Deprecated: use Service.Autostart instead | [optional] 
**Dns** | Pointer to [**ApiDNSConfig**](ApiDNSConfig.md) |  | [optional] 
**Env** | Pointer to **map[string]string** | Fields managed from fly.toml If you add anything here, ensure appconfig.Config.ToMachine() is updated | [optional] 
**Files** | Pointer to [**[]ApiFile**](ApiFile.md) |  | [optional] 
**Guest** | Pointer to [**ApiMachineGuest**](ApiMachineGuest.md) |  | [optional] 
**Image** | Pointer to **string** | Set by fly deploy or fly machines commands | [optional] 
**Init** | Pointer to [**ApiMachineInit**](ApiMachineInit.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Metrics** | Pointer to [**ApiMachineMetrics**](ApiMachineMetrics.md) |  | [optional] 
**Mounts** | Pointer to [**[]ApiMachineMount**](ApiMachineMount.md) |  | [optional] 
**Processes** | Pointer to [**[]ApiMachineProcess**](ApiMachineProcess.md) |  | [optional] 
**Restart** | Pointer to [**ApiMachineRestart**](ApiMachineRestart.md) |  | [optional] 
**Schedule** | Pointer to **string** | The following fields can only be set or updated by &#x60;fly machines run|update&#x60; commands \&quot;fly deploy\&quot; must preserve them, if you add anything here, ensure it is propagated on deploys | [optional] 
**Services** | Pointer to [**[]ApiMachineService**](ApiMachineService.md) |  | [optional] 
**Size** | Pointer to **string** | Deprecated: use Guest instead | [optional] 
**Standbys** | Pointer to **[]string** | Standbys enable a machine to be a standby for another. In the event of a hardware failure, the standby machine will be started. | [optional] 
**Statics** | Pointer to [**[]ApiStatic**](ApiStatic.md) |  | [optional] 
**StopConfig** | Pointer to [**ApiStopConfig**](ApiStopConfig.md) |  | [optional] 

## Methods

### NewApiMachineConfig

`func NewApiMachineConfig() *ApiMachineConfig`

NewApiMachineConfig instantiates a new ApiMachineConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMachineConfigWithDefaults

`func NewApiMachineConfigWithDefaults() *ApiMachineConfig`

NewApiMachineConfigWithDefaults instantiates a new ApiMachineConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDestroy

`func (o *ApiMachineConfig) GetAutoDestroy() bool`

GetAutoDestroy returns the AutoDestroy field if non-nil, zero value otherwise.

### GetAutoDestroyOk

`func (o *ApiMachineConfig) GetAutoDestroyOk() (*bool, bool)`

GetAutoDestroyOk returns a tuple with the AutoDestroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDestroy

`func (o *ApiMachineConfig) SetAutoDestroy(v bool)`

SetAutoDestroy sets AutoDestroy field to given value.

### HasAutoDestroy

`func (o *ApiMachineConfig) HasAutoDestroy() bool`

HasAutoDestroy returns a boolean if a field has been set.

### GetChecks

`func (o *ApiMachineConfig) GetChecks() map[string]ApiMachineCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ApiMachineConfig) GetChecksOk() (*map[string]ApiMachineCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ApiMachineConfig) SetChecks(v map[string]ApiMachineCheck)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ApiMachineConfig) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetDisableMachineAutostart

`func (o *ApiMachineConfig) GetDisableMachineAutostart() bool`

GetDisableMachineAutostart returns the DisableMachineAutostart field if non-nil, zero value otherwise.

### GetDisableMachineAutostartOk

`func (o *ApiMachineConfig) GetDisableMachineAutostartOk() (*bool, bool)`

GetDisableMachineAutostartOk returns a tuple with the DisableMachineAutostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableMachineAutostart

`func (o *ApiMachineConfig) SetDisableMachineAutostart(v bool)`

SetDisableMachineAutostart sets DisableMachineAutostart field to given value.

### HasDisableMachineAutostart

`func (o *ApiMachineConfig) HasDisableMachineAutostart() bool`

HasDisableMachineAutostart returns a boolean if a field has been set.

### GetDns

`func (o *ApiMachineConfig) GetDns() ApiDNSConfig`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *ApiMachineConfig) GetDnsOk() (*ApiDNSConfig, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *ApiMachineConfig) SetDns(v ApiDNSConfig)`

SetDns sets Dns field to given value.

### HasDns

`func (o *ApiMachineConfig) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetEnv

`func (o *ApiMachineConfig) GetEnv() map[string]string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ApiMachineConfig) GetEnvOk() (*map[string]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ApiMachineConfig) SetEnv(v map[string]string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ApiMachineConfig) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetFiles

`func (o *ApiMachineConfig) GetFiles() []ApiFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ApiMachineConfig) GetFilesOk() (*[]ApiFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ApiMachineConfig) SetFiles(v []ApiFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ApiMachineConfig) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetGuest

`func (o *ApiMachineConfig) GetGuest() ApiMachineGuest`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *ApiMachineConfig) GetGuestOk() (*ApiMachineGuest, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *ApiMachineConfig) SetGuest(v ApiMachineGuest)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *ApiMachineConfig) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetImage

`func (o *ApiMachineConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ApiMachineConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ApiMachineConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ApiMachineConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetInit

`func (o *ApiMachineConfig) GetInit() ApiMachineInit`

GetInit returns the Init field if non-nil, zero value otherwise.

### GetInitOk

`func (o *ApiMachineConfig) GetInitOk() (*ApiMachineInit, bool)`

GetInitOk returns a tuple with the Init field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInit

`func (o *ApiMachineConfig) SetInit(v ApiMachineInit)`

SetInit sets Init field to given value.

### HasInit

`func (o *ApiMachineConfig) HasInit() bool`

HasInit returns a boolean if a field has been set.

### GetMetadata

`func (o *ApiMachineConfig) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApiMachineConfig) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApiMachineConfig) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApiMachineConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMetrics

`func (o *ApiMachineConfig) GetMetrics() ApiMachineMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ApiMachineConfig) GetMetricsOk() (*ApiMachineMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ApiMachineConfig) SetMetrics(v ApiMachineMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ApiMachineConfig) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMounts

`func (o *ApiMachineConfig) GetMounts() []ApiMachineMount`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *ApiMachineConfig) GetMountsOk() (*[]ApiMachineMount, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *ApiMachineConfig) SetMounts(v []ApiMachineMount)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *ApiMachineConfig) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### GetProcesses

`func (o *ApiMachineConfig) GetProcesses() []ApiMachineProcess`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *ApiMachineConfig) GetProcessesOk() (*[]ApiMachineProcess, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *ApiMachineConfig) SetProcesses(v []ApiMachineProcess)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *ApiMachineConfig) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### GetRestart

`func (o *ApiMachineConfig) GetRestart() ApiMachineRestart`

GetRestart returns the Restart field if non-nil, zero value otherwise.

### GetRestartOk

`func (o *ApiMachineConfig) GetRestartOk() (*ApiMachineRestart, bool)`

GetRestartOk returns a tuple with the Restart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestart

`func (o *ApiMachineConfig) SetRestart(v ApiMachineRestart)`

SetRestart sets Restart field to given value.

### HasRestart

`func (o *ApiMachineConfig) HasRestart() bool`

HasRestart returns a boolean if a field has been set.

### GetSchedule

`func (o *ApiMachineConfig) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ApiMachineConfig) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ApiMachineConfig) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ApiMachineConfig) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServices

`func (o *ApiMachineConfig) GetServices() []ApiMachineService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ApiMachineConfig) GetServicesOk() (*[]ApiMachineService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ApiMachineConfig) SetServices(v []ApiMachineService)`

SetServices sets Services field to given value.

### HasServices

`func (o *ApiMachineConfig) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSize

`func (o *ApiMachineConfig) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ApiMachineConfig) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ApiMachineConfig) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ApiMachineConfig) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStandbys

`func (o *ApiMachineConfig) GetStandbys() []string`

GetStandbys returns the Standbys field if non-nil, zero value otherwise.

### GetStandbysOk

`func (o *ApiMachineConfig) GetStandbysOk() (*[]string, bool)`

GetStandbysOk returns a tuple with the Standbys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbys

`func (o *ApiMachineConfig) SetStandbys(v []string)`

SetStandbys sets Standbys field to given value.

### HasStandbys

`func (o *ApiMachineConfig) HasStandbys() bool`

HasStandbys returns a boolean if a field has been set.

### GetStatics

`func (o *ApiMachineConfig) GetStatics() []ApiStatic`

GetStatics returns the Statics field if non-nil, zero value otherwise.

### GetStaticsOk

`func (o *ApiMachineConfig) GetStaticsOk() (*[]ApiStatic, bool)`

GetStaticsOk returns a tuple with the Statics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatics

`func (o *ApiMachineConfig) SetStatics(v []ApiStatic)`

SetStatics sets Statics field to given value.

### HasStatics

`func (o *ApiMachineConfig) HasStatics() bool`

HasStatics returns a boolean if a field has been set.

### GetStopConfig

`func (o *ApiMachineConfig) GetStopConfig() ApiStopConfig`

GetStopConfig returns the StopConfig field if non-nil, zero value otherwise.

### GetStopConfigOk

`func (o *ApiMachineConfig) GetStopConfigOk() (*ApiStopConfig, bool)`

GetStopConfigOk returns a tuple with the StopConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopConfig

`func (o *ApiMachineConfig) SetStopConfig(v ApiStopConfig)`

SetStopConfig sets StopConfig field to given value.

### HasStopConfig

`func (o *ApiMachineConfig) HasStopConfig() bool`

HasStopConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


