# ApiMachineHTTPHeader

For http checks, an array of objects with string field Name and array of strings field Values. The key/value pairs specify header and header values that will get passed with the check call.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The header name | [optional] 
**values** | **List[str]** | The header value | [optional] 

## Example

```python
from fly_sdk.models.api_machine_http_header import ApiMachineHTTPHeader

# TODO update the JSON string below
json = "{}"
# create an instance of ApiMachineHTTPHeader from a JSON string
api_machine_http_header_instance = ApiMachineHTTPHeader.from_json(json)
# print the JSON string representation of the object
print ApiMachineHTTPHeader.to_json()

# convert the object into a dict
api_machine_http_header_dict = api_machine_http_header_instance.to_dict()
# create an instance of ApiMachineHTTPHeader from a dict
api_machine_http_header_form_dict = api_machine_http_header.from_dict(api_machine_http_header_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


