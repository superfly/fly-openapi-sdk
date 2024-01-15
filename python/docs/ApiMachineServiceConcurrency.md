# ApiMachineServiceConcurrency


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**hard_limit** | **int** |  | [optional] 
**soft_limit** | **int** |  | [optional] 
**type** | **str** |  | [optional] 

## Example

```python
from fly_sdk.models.api_machine_service_concurrency import ApiMachineServiceConcurrency

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineServiceConcurrency from a JSON string
api_machine_service_concurrency_instance = ApiMachineServiceConcurrency.from_json(json)
# print the JSON string representation of the object
print ApiMachineServiceConcurrency.to_json()

# convert the object into a dict
api_machine_service_concurrency_dict = api_machine_service_concurrency_instance.to_dict()
# create an instance of ApiMachineServiceConcurrency from a dict
api_machine_service_concurrency_form_dict = api_machine_service_concurrency.from_dict(api_machine_service_concurrency_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


