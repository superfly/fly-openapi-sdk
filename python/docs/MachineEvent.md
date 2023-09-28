# MachineEvent


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | [optional] 
**request** | **object** |  | [optional] 
**source** | **str** |  | [optional] 
**status** | **str** |  | [optional] 
**timestamp** | **int** |  | [optional] 
**type** | **str** |  | [optional] 

## Example

```python
from fly-sdk.models.machine_event import MachineEvent

# TODO update the JSON string below
json = "{}"
# create an instance of MachineEvent from a JSON string
machine_event_instance = MachineEvent.from_json(json)
# print the JSON string representation of the object
print MachineEvent.to_json()

# convert the object into a dict
machine_event_dict = machine_event_instance.to_dict()
# create an instance of MachineEvent from a dict
machine_event_form_dict = machine_event.from_dict(machine_event_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


