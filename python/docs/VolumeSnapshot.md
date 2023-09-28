# VolumeSnapshot


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**created_at** | **str** |  | [optional] 
**digest** | **str** |  | [optional] 
**id** | **str** |  | [optional] 
**size** | **int** |  | [optional] 

## Example

```python
from fly-sdk.models.volume_snapshot import VolumeSnapshot

# TODO update the JSON string below
json = "{}"
# create an instance of VolumeSnapshot from a JSON string
volume_snapshot_instance = VolumeSnapshot.from_json(json)
# print the JSON string representation of the object
print VolumeSnapshot.to_json()

# convert the object into a dict
volume_snapshot_dict = volume_snapshot_instance.to_dict()
# create an instance of VolumeSnapshot from a dict
volume_snapshot_form_dict = volume_snapshot.from_dict(volume_snapshot_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


