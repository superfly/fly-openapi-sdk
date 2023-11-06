# ListApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**MachineCount** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewListApp

`func NewListApp() *ListApp`

NewListApp instantiates a new ListApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAppWithDefaults

`func NewListAppWithDefaults() *ListApp`

NewListAppWithDefaults instantiates a new ListApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListApp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMachineCount

`func (o *ListApp) GetMachineCount() int32`

GetMachineCount returns the MachineCount field if non-nil, zero value otherwise.

### GetMachineCountOk

`func (o *ListApp) GetMachineCountOk() (*int32, bool)`

GetMachineCountOk returns a tuple with the MachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineCount

`func (o *ListApp) SetMachineCount(v int32)`

SetMachineCount sets MachineCount field to given value.

### HasMachineCount

`func (o *ListApp) HasMachineCount() bool`

HasMachineCount returns a boolean if a field has been set.

### GetName

`func (o *ListApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *ListApp) GetNetwork() map[string]interface{}`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ListApp) GetNetworkOk() (*map[string]interface{}, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ListApp) SetNetwork(v map[string]interface{})`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ListApp) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


