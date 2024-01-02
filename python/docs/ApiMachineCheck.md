# ApiMachineCheck

An optional object that defines one or more named checks. The key for each check is the check name.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**grace_period** | **str** | The time to wait after a VM starts before checking its health | [optional] 
**headers** | [**List[ApiMachineHTTPHeader]**](ApiMachineHTTPHeader.md) |  | [optional] 
**interval** | **str** | The time between connectivity checks | [optional] 
**method** | **str** | For http checks, the HTTP method to use to when making the request | [optional] 
**path** | **str** | For http checks, the path to send the request to | [optional] 
**port** | **int** | The port to connect to, often the same as internal_port | [optional] 
**protocol** | **str** | For http checks, whether to use http or https | [optional] 
**timeout** | **str** | The maximum time a connection can take before being reported as failing its health check | [optional] 
**tls_server_name** | **str** | If the protocol is https, the hostname to use for TLS certificate validation | [optional] 
**tls_skip_verify** | **bool** | For http checks with https protocol, whether or not to verify the TLS certificate | [optional] 
**type** | **str** | tcp or http | [optional] 

## Example

```python
from fly_sdk.models.api_machine_check import ApiMachineCheck

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineCheck from a JSON string
api_machine_check_instance = ApiMachineCheck.from_json(json)
# print the JSON string representation of the object
print ApiMachineCheck.to_json()

# convert the object into a dict
api_machine_check_dict = api_machine_check_instance.to_dict()
# create an instance of ApiMachineCheck from a dict
api_machine_check_form_dict = api_machine_check.from_dict(api_machine_check_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


