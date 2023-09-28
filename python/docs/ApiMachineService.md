# ApiMachineService


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**autostart** | **bool** |  | [optional] 
**autostop** | **bool** |  | [optional] 
**checks** | [**List[ApiMachineCheck]**](ApiMachineCheck.md) |  | [optional] 
**concurrency** | [**ApiMachineServiceConcurrency**](ApiMachineServiceConcurrency.md) |  | [optional] 
**force_instance_description** | **str** |  | [optional] 
**force_instance_key** | **str** |  | [optional] 
**internal_port** | **int** |  | [optional] 
**min_machines_running** | **int** |  | [optional] 
**ports** | [**List[ApiMachinePort]**](ApiMachinePort.md) |  | [optional] 
**protocol** | **str** |  | [optional] 

## Example

```python
from fly-sdk.models.api_machine_service import ApiMachineService

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineService from a JSON string
api_machine_service_instance = ApiMachineService.from_json(json)
# print the JSON string representation of the object
print ApiMachineService.to_json()

# convert the object into a dict
api_machine_service_dict = api_machine_service_instance.to_dict()
# create an instance of ApiMachineService from a dict
api_machine_service_form_dict = api_machine_service.from_dict(api_machine_service_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


