# ApiMachineRestart


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**max_retries** | **int** | MaxRetries is only relevant with the on-failure policy. | [optional] 
**policy** | **str** |  | [optional] 

## Example

```python
from fly-sdk.models.api_machine_restart import ApiMachineRestart

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineRestart from a JSON string
api_machine_restart_instance = ApiMachineRestart.from_json(json)
# print the JSON string representation of the object
print ApiMachineRestart.to_json()

# convert the object into a dict
api_machine_restart_dict = api_machine_restart_instance.to_dict()
# create an instance of ApiMachineRestart from a dict
api_machine_restart_form_dict = api_machine_restart.from_dict(api_machine_restart_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


