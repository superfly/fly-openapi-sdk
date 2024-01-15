# ApiDNSConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**skip_registration** | **bool** |  | [optional] 

## Example

```python
from fly_sdk.models.api_dns_config import ApiDNSConfig

# TODO update the JSON string below
json = "{}"
# create an instance of ApiDNSConfig from a JSON string
api_dns_config_instance = ApiDNSConfig.from_json(json)
# print the JSON string representation of the object
print ApiDNSConfig.to_json()

# convert the object into a dict
api_dns_config_dict = api_dns_config_instance.to_dict()
# create an instance of ApiDNSConfig from a dict
api_dns_config_form_dict = api_dns_config.from_dict(api_dns_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


