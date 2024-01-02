# ApiStopConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**signal** | **str** |  | [optional] 
**timeout** | **str** |  | [optional] 

## Example

```python
from fly_sdk.models.api_stop_config import ApiStopConfig

# TODO update the JSON string below
json = "{}"
# create an instance of ApiStopConfig from a JSON string
api_stop_config_instance = ApiStopConfig.from_json(json)
# print the JSON string representation of the object
print ApiStopConfig.to_json()

# convert the object into a dict
api_stop_config_dict = api_stop_config_instance.to_dict()
# create an instance of ApiStopConfig from a dict
api_stop_config_form_dict = api_stop_config.from_dict(api_stop_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


