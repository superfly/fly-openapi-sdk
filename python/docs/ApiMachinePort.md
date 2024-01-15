# ApiMachinePort


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**end_port** | **int** |  | [optional] 
**force_https** | **bool** |  | [optional] 
**handlers** | **List[str]** |  | [optional] 
**http_options** | [**ApiHTTPOptions**](ApiHTTPOptions.md) |  | [optional] 
**port** | **int** |  | [optional] 
**proxy_proto_options** | [**ApiProxyProtoOptions**](ApiProxyProtoOptions.md) |  | [optional] 
**start_port** | **int** |  | [optional] 
**tls_options** | [**ApiTLSOptions**](ApiTLSOptions.md) |  | [optional] 

## Example

```python
from fly_sdk.models.api_machine_port import ApiMachinePort

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachinePort from a JSON string
api_machine_port_instance = ApiMachinePort.from_json(json)
# print the JSON string representation of the object
print ApiMachinePort.to_json()

# convert the object into a dict
api_machine_port_dict = api_machine_port_instance.to_dict()
# create an instance of ApiMachinePort from a dict
api_machine_port_form_dict = api_machine_port.from_dict(api_machine_port_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


