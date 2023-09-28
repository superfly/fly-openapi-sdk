# Machine


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**checks** | [**List[CheckStatus]**](CheckStatus.md) |  | [optional] 
**config** | [**ApiMachineConfig**](ApiMachineConfig.md) |  | [optional] 
**created_at** | **str** |  | [optional] 
**events** | [**List[MachineEvent]**](MachineEvent.md) |  | [optional] 
**id** | **str** |  | [optional] 
**image_ref** | [**ImageRef**](ImageRef.md) |  | [optional] 
**instance_id** | **str** | InstanceID is unique for each version of the machine | [optional] 
**name** | **str** |  | [optional] 
**nonce** | **str** | Nonce is only every returned on machine creation if a lease_duration was provided. | [optional] 
**private_ip** | **str** | PrivateIP is the internal 6PN address of the machine. | [optional] 
**region** | **str** |  | [optional] 
**state** | **str** |  | [optional] 
**updated_at** | **str** |  | [optional] 

## Example

```python
from fly-sdk.models.machine import Machine

# TODO update the JSON string below
json = "{}"
# create an instance of Machine from a JSON string
machine_instance = Machine.from_json(json)
# print the JSON string representation of the object
print Machine.to_json()

# convert the object into a dict
machine_dict = machine_instance.to_dict()
# create an instance of Machine from a dict
machine_form_dict = machine.from_dict(machine_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


