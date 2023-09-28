# ListenSocket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Proto** | Pointer to **string** |  | [optional] 

## Methods

### NewListenSocket

`func NewListenSocket() *ListenSocket`

NewListenSocket instantiates a new ListenSocket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListenSocketWithDefaults

`func NewListenSocketWithDefaults() *ListenSocket`

NewListenSocketWithDefaults instantiates a new ListenSocket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListenSocket) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListenSocket) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListenSocket) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ListenSocket) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetProto

`func (o *ListenSocket) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *ListenSocket) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *ListenSocket) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *ListenSocket) HasProto() bool`

HasProto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


