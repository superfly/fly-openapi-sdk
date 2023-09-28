# ExtendVolumeResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**needs_restart** | **bool** |  | [optional] 
**volume** | [**Volume**](Volume.md) |  | [optional] 

## Example

```python
from fly-sdk.models.extend_volume_response import ExtendVolumeResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ExtendVolumeResponse from a JSON string
extend_volume_response_instance = ExtendVolumeResponse.from_json(json)
# print the JSON string representation of the object
print ExtendVolumeResponse.to_json()

# convert the object into a dict
extend_volume_response_dict = extend_volume_response_instance.to_dict()
# create an instance of ExtendVolumeResponse from a dict
extend_volume_response_form_dict = extend_volume_response.from_dict(extend_volume_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


