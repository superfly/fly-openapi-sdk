# MachineExecRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cmd** | **str** | Deprecated: use Command instead | [optional] 
**command** | **List[str]** |  | [optional] 
**timeout** | **int** |  | [optional] 

## Example

```python
from fly_sdk.models.machine_exec_request import MachineExecRequest

# TODO update the JSON string below
json = "{}"
# create an instance of MachineExecRequest from a JSON string
machine_exec_request_instance = MachineExecRequest.from_json(json)
# print the JSON string representation of the object
print MachineExecRequest.to_json()

# convert the object into a dict
machine_exec_request_dict = machine_exec_request_instance.to_dict()
# create an instance of MachineExecRequest from a dict
machine_exec_request_form_dict = machine_exec_request.from_dict(machine_exec_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


