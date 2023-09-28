# ApiMachineProcess


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cmd** | **List[str]** |  | [optional] 
**entrypoint** | **List[str]** |  | [optional] 
**env** | **Dict[str, str]** |  | [optional] 
**var_exec** | **List[str]** |  | [optional] 
**user** | **str** |  | [optional] 

## Example

```python
from fly-sdk.models.api_machine_process import ApiMachineProcess

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineProcess from a JSON string
api_machine_process_instance = ApiMachineProcess.from_json(json)
# print the JSON string representation of the object
print ApiMachineProcess.to_json()

# convert the object into a dict
api_machine_process_dict = api_machine_process_instance.to_dict()
# create an instance of ApiMachineProcess from a dict
api_machine_process_form_dict = api_machine_process.from_dict(api_machine_process_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


