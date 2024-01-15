# ApiMachineMount


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**add_size_gb** | **int** |  | [optional] 
**encrypted** | **bool** |  | [optional] 
**extend_threshold_percent** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**path** | **str** |  | [optional] 
**size_gb** | **int** |  | [optional] 
**size_gb_limit** | **int** |  | [optional] 
**volume** | **str** |  | [optional] 

## Example

```python
from fly_sdk.models.api_machine_mount import ApiMachineMount

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineMount from a JSON string
api_machine_mount_instance = ApiMachineMount.from_json(json)
# print the JSON string representation of the object
print ApiMachineMount.to_json()

# convert the object into a dict
api_machine_mount_dict = api_machine_mount_instance.to_dict()
# create an instance of ApiMachineMount from a dict
api_machine_mount_form_dict = api_machine_mount.from_dict(api_machine_mount_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


