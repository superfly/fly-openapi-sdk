# ApiMachineConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auto_destroy** | **bool** | Optional boolean telling the Machine to destroy itself once it’s complete (default false) | [optional] 
**checks** | [**Dict[str, ApiMachineCheck]**](ApiMachineCheck.md) |  | [optional] 
**disable_machine_autostart** | **bool** | Deprecated: use Service.Autostart instead | [optional] 
**dns** | [**ApiDNSConfig**](ApiDNSConfig.md) |  | [optional] 
**env** | **Dict[str, str]** | An object filled with key/value pairs to be set as environment variables | [optional] 
**files** | [**List[ApiFile]**](ApiFile.md) |  | [optional] 
**guest** | [**ApiMachineGuest**](ApiMachineGuest.md) |  | [optional] 
**image** | **str** | The docker image to run | [optional] 
**init** | [**ApiMachineInit**](ApiMachineInit.md) |  | [optional] 
**metadata** | **Dict[str, str]** |  | [optional] 
**metrics** | [**ApiMachineMetrics**](ApiMachineMetrics.md) |  | [optional] 
**mounts** | [**List[ApiMachineMount]**](ApiMachineMount.md) |  | [optional] 
**processes** | [**List[ApiMachineProcess]**](ApiMachineProcess.md) |  | [optional] 
**restart** | [**ApiMachineRestart**](ApiMachineRestart.md) |  | [optional] 
**schedule** | **str** |  | [optional] 
**services** | [**List[ApiMachineService]**](ApiMachineService.md) |  | [optional] 
**size** | **str** | Deprecated: use Guest instead | [optional] 
**standbys** | **List[str]** | Standbys enable a machine to be a standby for another. In the event of a hardware failure, the standby machine will be started. | [optional] 
**statics** | [**List[ApiStatic]**](ApiStatic.md) |  | [optional] 
**stop_config** | [**ApiStopConfig**](ApiStopConfig.md) |  | [optional] 

## Example

```python
from fly-sdk.models.api_machine_config import ApiMachineConfig

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineConfig from a JSON string
api_machine_config_instance = ApiMachineConfig.from_json(json)
# print the JSON string representation of the object
print ApiMachineConfig.to_json()

# convert the object into a dict
api_machine_config_dict = api_machine_config_instance.to_dict()
# create an instance of ApiMachineConfig from a dict
api_machine_config_form_dict = api_machine_config.from_dict(api_machine_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


