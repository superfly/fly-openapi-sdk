# ApiHTTPOptions


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compress** | **bool** |  | [optional] 
**h2_backend** | **bool** |  | [optional] 
**response** | [**ApiHTTPResponseOptions**](ApiHTTPResponseOptions.md) |  | [optional] 

## Example

```python
from fly_sdk.models.api_http_options import ApiHTTPOptions

# TODO update the JSON string below
json = "{}"
# create an instance of ApiHTTPOptions from a JSON string
api_http_options_instance = ApiHTTPOptions.from_json(json)
# print the JSON string representation of the object
print ApiHTTPOptions.to_json()

# convert the object into a dict
api_http_options_dict = api_http_options_instance.to_dict()
# create an instance of ApiHTTPOptions from a dict
api_http_options_form_dict = api_http_options.from_dict(api_http_options_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


