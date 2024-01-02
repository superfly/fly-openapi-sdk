# Volume


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**attached_alloc_id** | **str** |  | [optional] 
**attached_machine_id** | **str** |  | [optional] 
**block_size** | **int** |  | [optional] 
**blocks** | **int** |  | [optional] 
**blocks_avail** | **int** |  | [optional] 
**blocks_free** | **int** |  | [optional] 
**created_at** | **str** |  | [optional] 
**encrypted** | **bool** |  | [optional] 
**fstype** | **str** |  | [optional] 
**id** | **str** |  | [optional] 
**name** | **str** |  | [optional] 
**region** | **str** |  | [optional] 
**size_gb** | **int** |  | [optional] 
**snapshot_retention** | **int** |  | [optional] 
**state** | **str** |  | [optional] 
**zone** | **str** |  | [optional] 

## Example

```python
from fly_sdk.models.volume import Volume

# TODO update the JSON string below
json = "{}"
# create an instance of Volume from a JSON string
volume_instance = Volume.from_json(json)
# print the JSON string representation of the object
print Volume.to_json()

# convert the object into a dict
volume_dict = volume_instance.to_dict()
# create an instance of Volume from a dict
volume_form_dict = volume.from_dict(volume_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


