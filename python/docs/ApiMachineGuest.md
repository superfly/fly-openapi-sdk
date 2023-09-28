# ApiMachineGuest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpu_kind** | **str** |  | [optional] 
**cpus** | **int** |  | [optional] 
**gpu_kind** | **str** |  | [optional] 
**kernel_args** | **List[str]** |  | [optional] 
**memory_mb** | **int** |  | [optional] 

## Example

```python
from fly-sdk.models.api_machine_guest import ApiMachineGuest

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineGuest from a JSON string
api_machine_guest_instance = ApiMachineGuest.from_json(json)
# print the JSON string representation of the object
print ApiMachineGuest.to_json()

# convert the object into a dict
api_machine_guest_dict = api_machine_guest_instance.to_dict()
# create an instance of ApiMachineGuest from a dict
api_machine_guest_form_dict = api_machine_guest.from_dict(api_machine_guest_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


