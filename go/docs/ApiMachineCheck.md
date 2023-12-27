# ApiMachineCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GracePeriod** | Pointer to **string** | The time to wait after a VM starts before checking its health | [optional] 
**Headers** | Pointer to [**[]ApiMachineHTTPHeader**](ApiMachineHTTPHeader.md) |  | [optional] 
**Interval** | Pointer to **string** | The time between connectivity checks | [optional] 
**Method** | Pointer to **string** | For http checks, the HTTP method to use to when making the request | [optional] 
**Path** | Pointer to **string** | For http checks, the path to send the request to | [optional] 
**Port** | Pointer to **int32** | The port to connect to, often the same as internal_port | [optional] 
**Protocol** | Pointer to **string** | For http checks, whether to use http or https | [optional] 
**Timeout** | Pointer to **string** | The maximum time a connection can take before being reported as failing its health check | [optional] 
**TlsServerName** | Pointer to **string** | If the protocol is https, the hostname to use for TLS certificate validation | [optional] 
**TlsSkipVerify** | Pointer to **bool** | For http checks with https protocol, whether or not to verify the TLS certificate | [optional] 
**Type** | Pointer to **string** | tcp or http | [optional] 

## Methods

### NewApiMachineCheck

`func NewApiMachineCheck() *ApiMachineCheck`

NewApiMachineCheck instantiates a new ApiMachineCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMachineCheckWithDefaults

`func NewApiMachineCheckWithDefaults() *ApiMachineCheck`

NewApiMachineCheckWithDefaults instantiates a new ApiMachineCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGracePeriod

`func (o *ApiMachineCheck) GetGracePeriod() string`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *ApiMachineCheck) GetGracePeriodOk() (*string, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *ApiMachineCheck) SetGracePeriod(v string)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *ApiMachineCheck) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetHeaders

`func (o *ApiMachineCheck) GetHeaders() []ApiMachineHTTPHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ApiMachineCheck) GetHeadersOk() (*[]ApiMachineHTTPHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ApiMachineCheck) SetHeaders(v []ApiMachineHTTPHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ApiMachineCheck) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetInterval

`func (o *ApiMachineCheck) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ApiMachineCheck) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ApiMachineCheck) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ApiMachineCheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMethod

`func (o *ApiMachineCheck) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApiMachineCheck) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApiMachineCheck) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ApiMachineCheck) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *ApiMachineCheck) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiMachineCheck) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiMachineCheck) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApiMachineCheck) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *ApiMachineCheck) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApiMachineCheck) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApiMachineCheck) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApiMachineCheck) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *ApiMachineCheck) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApiMachineCheck) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApiMachineCheck) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApiMachineCheck) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTimeout

`func (o *ApiMachineCheck) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ApiMachineCheck) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ApiMachineCheck) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ApiMachineCheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTlsServerName

`func (o *ApiMachineCheck) GetTlsServerName() string`

GetTlsServerName returns the TlsServerName field if non-nil, zero value otherwise.

### GetTlsServerNameOk

`func (o *ApiMachineCheck) GetTlsServerNameOk() (*string, bool)`

GetTlsServerNameOk returns a tuple with the TlsServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerName

`func (o *ApiMachineCheck) SetTlsServerName(v string)`

SetTlsServerName sets TlsServerName field to given value.

### HasTlsServerName

`func (o *ApiMachineCheck) HasTlsServerName() bool`

HasTlsServerName returns a boolean if a field has been set.

### GetTlsSkipVerify

`func (o *ApiMachineCheck) GetTlsSkipVerify() bool`

GetTlsSkipVerify returns the TlsSkipVerify field if non-nil, zero value otherwise.

### GetTlsSkipVerifyOk

`func (o *ApiMachineCheck) GetTlsSkipVerifyOk() (*bool, bool)`

GetTlsSkipVerifyOk returns a tuple with the TlsSkipVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsSkipVerify

`func (o *ApiMachineCheck) SetTlsSkipVerify(v bool)`

SetTlsSkipVerify sets TlsSkipVerify field to given value.

### HasTlsSkipVerify

`func (o *ApiMachineCheck) HasTlsSkipVerify() bool`

HasTlsSkipVerify returns a boolean if a field has been set.

### GetType

`func (o *ApiMachineCheck) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiMachineCheck) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiMachineCheck) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiMachineCheck) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


