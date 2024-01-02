# ApiMachineInit


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cmd** | **List[str]** |  | [optional] 
**entrypoint** | **List[str]** |  | [optional] 
**var_exec** | **List[str]** |  | [optional] 
**kernel_args** | **List[str]** |  | [optional] 
**swap_size_mb** | **int** |  | [optional] 
**tty** | **bool** |  | [optional] 

## Example

```python
from fly_sdk.models.api_machine_init import ApiMachineInit

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineInit from a JSON string
api_machine_init_instance = ApiMachineInit.from_json(json)
# print the JSON string representation of the object
print ApiMachineInit.to_json()

# convert the object into a dict
api_machine_init_dict = api_machine_init_instance.to_dict()
# create an instance of ApiMachineInit from a dict
api_machine_init_form_dict = api_machine_init.from_dict(api_machine_init_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


