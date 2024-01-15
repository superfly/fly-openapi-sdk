# StopRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**signal** | **str** |  | [optional] 
**timeout** | **str** |  | [optional] 

## Example

```python
from fly_sdk.models.stop_request import StopRequest

# TODO update the JSON string below
json = "{}"
# create an instance of StopRequest from a JSON string
stop_request_instance = StopRequest.from_json(json)
# print the JSON string representation of the object
print StopRequest.to_json()

# convert the object into a dict
stop_request_dict = stop_request_instance.to_dict()
# create an instance of StopRequest from a dict
stop_request_form_dict = stop_request.from_dict(stop_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


