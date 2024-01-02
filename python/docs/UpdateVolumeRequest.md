# UpdateVolumeRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**snapshot_retention** | **int** |  | [optional] 

## Example

```python
from fly_sdk.models.update_volume_request import UpdateVolumeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of UpdateVolumeRequest from a JSON string
update_volume_request_instance = UpdateVolumeRequest.from_json(json)
# print the JSON string representation of the object
print UpdateVolumeRequest.to_json()

# convert the object into a dict
update_volume_request_dict = update_volume_request_instance.to_dict()
# create an instance of UpdateVolumeRequest from a dict
update_volume_request_form_dict = update_volume_request.from_dict(update_volume_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


