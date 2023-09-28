# CreateAppRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**app_name** | **str** |  | [optional] 
**network** | **str** |  | [optional] 
**org_slug** | **str** |  | [optional] 

## Example

```python
from fly-sdk.models.create_app_request import CreateAppRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateAppRequest from a JSON string
create_app_request_instance = CreateAppRequest.from_json(json)
# print the JSON string representation of the object
print CreateAppRequest.to_json()

# convert the object into a dict
create_app_request_dict = create_app_request_instance.to_dict()
# create an instance of CreateAppRequest from a dict
create_app_request_form_dict = create_app_request.from_dict(create_app_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


