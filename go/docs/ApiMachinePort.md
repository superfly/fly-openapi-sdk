# ApiMachinePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndPort** | Pointer to **int32** |  | [optional] 
**ForceHttps** | Pointer to **bool** |  | [optional] 
**Handlers** | Pointer to **[]string** |  | [optional] 
**HttpOptions** | Pointer to [**ApiHTTPOptions**](ApiHTTPOptions.md) |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**ProxyProtoOptions** | Pointer to [**ApiProxyProtoOptions**](ApiProxyProtoOptions.md) |  | [optional] 
**StartPort** | Pointer to **int32** |  | [optional] 
**TlsOptions** | Pointer to [**ApiTLSOptions**](ApiTLSOptions.md) |  | [optional] 

## Methods

### NewApiMachinePort

`func NewApiMachinePort() *ApiMachinePort`

NewApiMachinePort instantiates a new ApiMachinePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMachinePortWithDefaults

`func NewApiMachinePortWithDefaults() *ApiMachinePort`

NewApiMachinePortWithDefaults instantiates a new ApiMachinePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndPort

`func (o *ApiMachinePort) GetEndPort() int32`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *ApiMachinePort) GetEndPortOk() (*int32, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *ApiMachinePort) SetEndPort(v int32)`

SetEndPort sets EndPort field to given value.

### HasEndPort

`func (o *ApiMachinePort) HasEndPort() bool`

HasEndPort returns a boolean if a field has been set.

### GetForceHttps

`func (o *ApiMachinePort) GetForceHttps() bool`

GetForceHttps returns the ForceHttps field if non-nil, zero value otherwise.

### GetForceHttpsOk

`func (o *ApiMachinePort) GetForceHttpsOk() (*bool, bool)`

GetForceHttpsOk returns a tuple with the ForceHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceHttps

`func (o *ApiMachinePort) SetForceHttps(v bool)`

SetForceHttps sets ForceHttps field to given value.

### HasForceHttps

`func (o *ApiMachinePort) HasForceHttps() bool`

HasForceHttps returns a boolean if a field has been set.

### GetHandlers

`func (o *ApiMachinePort) GetHandlers() []string`

GetHandlers returns the Handlers field if non-nil, zero value otherwise.

### GetHandlersOk

`func (o *ApiMachinePort) GetHandlersOk() (*[]string, bool)`

GetHandlersOk returns a tuple with the Handlers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlers

`func (o *ApiMachinePort) SetHandlers(v []string)`

SetHandlers sets Handlers field to given value.

### HasHandlers

`func (o *ApiMachinePort) HasHandlers() bool`

HasHandlers returns a boolean if a field has been set.

### GetHttpOptions

`func (o *ApiMachinePort) GetHttpOptions() ApiHTTPOptions`

GetHttpOptions returns the HttpOptions field if non-nil, zero value otherwise.

### GetHttpOptionsOk

`func (o *ApiMachinePort) GetHttpOptionsOk() (*ApiHTTPOptions, bool)`

GetHttpOptionsOk returns a tuple with the HttpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpOptions

`func (o *ApiMachinePort) SetHttpOptions(v ApiHTTPOptions)`

SetHttpOptions sets HttpOptions field to given value.

### HasHttpOptions

`func (o *ApiMachinePort) HasHttpOptions() bool`

HasHttpOptions returns a boolean if a field has been set.

### GetPort

`func (o *ApiMachinePort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApiMachinePort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApiMachinePort) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApiMachinePort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProxyProtoOptions

`func (o *ApiMachinePort) GetProxyProtoOptions() ApiProxyProtoOptions`

GetProxyProtoOptions returns the ProxyProtoOptions field if non-nil, zero value otherwise.

### GetProxyProtoOptionsOk

`func (o *ApiMachinePort) GetProxyProtoOptionsOk() (*ApiProxyProtoOptions, bool)`

GetProxyProtoOptionsOk returns a tuple with the ProxyProtoOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtoOptions

`func (o *ApiMachinePort) SetProxyProtoOptions(v ApiProxyProtoOptions)`

SetProxyProtoOptions sets ProxyProtoOptions field to given value.

### HasProxyProtoOptions

`func (o *ApiMachinePort) HasProxyProtoOptions() bool`

HasProxyProtoOptions returns a boolean if a field has been set.

### GetStartPort

`func (o *ApiMachinePort) GetStartPort() int32`

GetStartPort returns the StartPort field if non-nil, zero value otherwise.

### GetStartPortOk

`func (o *ApiMachinePort) GetStartPortOk() (*int32, bool)`

GetStartPortOk returns a tuple with the StartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPort

`func (o *ApiMachinePort) SetStartPort(v int32)`

SetStartPort sets StartPort field to given value.

### HasStartPort

`func (o *ApiMachinePort) HasStartPort() bool`

HasStartPort returns a boolean if a field has been set.

### GetTlsOptions

`func (o *ApiMachinePort) GetTlsOptions() ApiTLSOptions`

GetTlsOptions returns the TlsOptions field if non-nil, zero value otherwise.

### GetTlsOptionsOk

`func (o *ApiMachinePort) GetTlsOptionsOk() (*ApiTLSOptions, bool)`

GetTlsOptionsOk returns a tuple with the TlsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptions

`func (o *ApiMachinePort) SetTlsOptions(v ApiTLSOptions)`

SetTlsOptions sets TlsOptions field to given value.

### HasTlsOptions

`func (o *ApiMachinePort) HasTlsOptions() bool`

HasTlsOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


