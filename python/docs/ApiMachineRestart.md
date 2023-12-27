# ApiMachineRestart

The Machine restart policy defines whether and how flyd restarts a Machine after its main process exits. See https://fly.io/docs/machines/guides-examples/machine-restart-policy/.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**max_retries** | **int** | When policy is on-failure, the maximum number of times to attempt to restart the Machine before letting it stop. | [optional] 
**policy** | **str** | * no - Never try to restart a Machine automatically when its main process exits, whether that’s on purpose or on a crash. * always - Always restart a Machine automatically and never let it enter a stopped state, even when the main process exits cleanly. * on-failure - Try up to MaxRetries times to automatically restart the Machine if it exits with a non-zero exit code. Default when no explicit policy is set, and for Machines with schedules. | [optional] 

## Example

```python
from fly-sdk.models.api_machine_restart import ApiMachineRestart

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineRestart from a JSON string
api_machine_restart_instance = ApiMachineRestart.from_json(json)
# print the JSON string representation of the object
print ApiMachineRestart.to_json()

# convert the object into a dict
api_machine_restart_dict = api_machine_restart_instance.to_dict()
# create an instance of ApiMachineRestart from a dict
api_machine_restart_form_dict = api_machine_restart.from_dict(api_machine_restart_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


