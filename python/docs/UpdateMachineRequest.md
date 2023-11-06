# UpdateMachineRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config** | [**ApiMachineConfig**](ApiMachineConfig.md) |  | [optional] 
**current_version** | **str** |  | [optional] 
**lease_ttl** | **int** |  | [optional] 
**lsvd** | **bool** |  | [optional] 
**name** | **str** | Unique name for this Machine. If omitted, one is generated for you | [optional] 
**region** | **str** | The target region. Omitting this param launches in the same region as your WireGuard peer connection (somewhere near you). | [optional] 
**skip_launch** | **bool** |  | [optional] 
**skip_service_registration** | **bool** |  | [optional] 

## Example

```python
from fly-sdk.models.update_machine_request import UpdateMachineRequest

# TODO update the JSON string below
json = "{}"
# create an instance of UpdateMachineRequest from a JSON string
update_machine_request_instance = UpdateMachineRequest.from_json(json)
# print the JSON string representation of the object
print UpdateMachineRequest.to_json()

# convert the object into a dict
update_machine_request_dict = update_machine_request_instance.to_dict()
# create an instance of UpdateMachineRequest from a dict
update_machine_request_form_dict = update_machine_request.from_dict(update_machine_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


