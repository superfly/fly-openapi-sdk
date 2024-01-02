# ExtendVolumeRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**size_gb** | **int** |  | [optional] 

## Example

```python
from fly_sdk.models.extend_volume_request import ExtendVolumeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ExtendVolumeRequest from a JSON string
extend_volume_request_instance = ExtendVolumeRequest.from_json(json)
# print the JSON string representation of the object
print ExtendVolumeRequest.to_json()

# convert the object into a dict
extend_volume_request_dict = extend_volume_request_instance.to_dict()
# create an instance of ExtendVolumeRequest from a dict
extend_volume_request_form_dict = extend_volume_request.from_dict(extend_volume_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


