# SignalRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**signal** | **str** |  | [optional] 

## Example

```python
from fly_sdk.models.signal_request import SignalRequest

# TODO update the JSON string below
json = "{}"
# create an instance of SignalRequest from a JSON string
signal_request_instance = SignalRequest.from_json(json)
# print the JSON string representation of the object
print SignalRequest.to_json()

# convert the object into a dict
signal_request_dict = signal_request_instance.to_dict()
# create an instance of SignalRequest from a dict
signal_request_form_dict = signal_request.from_dict(signal_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


