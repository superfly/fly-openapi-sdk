# ApiMachineMetrics


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**path** | **str** |  | [optional] 
**port** | **int** |  | [optional] 

## Example

```python
from fly_sdk.models.api_machine_metrics import ApiMachineMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineMetrics from a JSON string
api_machine_metrics_instance = ApiMachineMetrics.from_json(json)
# print the JSON string representation of the object
print ApiMachineMetrics.to_json()

# convert the object into a dict
api_machine_metrics_dict = api_machine_metrics_instance.to_dict()
# create an instance of ApiMachineMetrics from a dict
api_machine_metrics_form_dict = api_machine_metrics.from_dict(api_machine_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


