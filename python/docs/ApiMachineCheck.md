# ApiMachineCheck


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**grace_period** | **str** |  | [optional] 
**headers** | [**List[ApiMachineHTTPHeader]**](ApiMachineHTTPHeader.md) |  | [optional] 
**interval** | **str** |  | [optional] 
**method** | **str** |  | [optional] 
**path** | **str** |  | [optional] 
**port** | **int** |  | [optional] 
**protocol** | **str** |  | [optional] 
**timeout** | **str** |  | [optional] 
**tls_server_name** | **str** |  | [optional] 
**tls_skip_verify** | **bool** |  | [optional] 
**type** | **str** |  | [optional] 

## Example

```python
from fly-sdk.models.api_machine_check import ApiMachineCheck

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineCheck from a JSON string
api_machine_check_instance = ApiMachineCheck.from_json(json)
# print the JSON string representation of the object
print ApiMachineCheck.to_json()

# convert the object into a dict
api_machine_check_dict = api_machine_check_instance.to_dict()
# create an instance of ApiMachineCheck from a dict
api_machine_check_form_dict = api_machine_check.from_dict(api_machine_check_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


