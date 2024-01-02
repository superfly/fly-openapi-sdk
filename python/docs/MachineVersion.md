# MachineVersion


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_config** | [**ApiMachineConfig**](ApiMachineConfig.md) |  | [optional] 
**version** | **str** |  | [optional] 

## Example

```python
from fly_sdk.models.machine_version import MachineVersion

# TODO update the JSON string below
json = "{}"
# create an instance of MachineVersion from a JSON string
machine_version_instance = MachineVersion.from_json(json)
# print the JSON string representation of the object
print MachineVersion.to_json()

# convert the object into a dict
machine_version_dict = machine_version_instance.to_dict()
# create an instance of MachineVersion from a dict
machine_version_form_dict = machine_version.from_dict(machine_version_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


