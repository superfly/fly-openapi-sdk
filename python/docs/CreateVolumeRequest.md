# CreateVolumeRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compute** | [**ApiMachineGuest**](ApiMachineGuest.md) |  | [optional] 
**encrypted** | **bool** |  | [optional] 
**fstype** | **str** |  | [optional] 
**machines_only** | **bool** |  | [optional] 
**name** | **str** |  | [optional] 
**region** | **str** |  | [optional] 
**require_unique_zone** | **bool** |  | [optional] 
**size_gb** | **int** |  | [optional] 
**snapshot_id** | **str** | restore from snapshot | [optional] 
**snapshot_retention** | **int** |  | [optional] 
**source_volume_id** | **str** | fork from remote volume | [optional] 

## Example

```python
from fly-sdk.models.create_volume_request import CreateVolumeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateVolumeRequest from a JSON string
create_volume_request_instance = CreateVolumeRequest.from_json(json)
# print the JSON string representation of the object
print CreateVolumeRequest.to_json()

# convert the object into a dict
create_volume_request_dict = create_volume_request_instance.to_dict()
# create an instance of CreateVolumeRequest from a dict
create_volume_request_form_dict = create_volume_request.from_dict(create_volume_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


