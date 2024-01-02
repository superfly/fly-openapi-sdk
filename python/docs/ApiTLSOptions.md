# ApiTLSOptions


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**alpn** | **List[str]** |  | [optional] 
**default_self_signed** | **bool** |  | [optional] 
**versions** | **List[str]** |  | [optional] 

## Example

```python
from fly_sdk.models.api_tls_options import ApiTLSOptions

# TODO update the JSON string below
json = "{}"
# create an instance of ApiTLSOptions from a JSON string
api_tls_options_instance = ApiTLSOptions.from_json(json)
# print the JSON string representation of the object
print ApiTLSOptions.to_json()

# convert the object into a dict
api_tls_options_dict = api_tls_options_instance.to_dict()
# create an instance of ApiTLSOptions from a dict
api_tls_options_form_dict = api_tls_options.from_dict(api_tls_options_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


