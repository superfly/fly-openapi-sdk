# ListAppsResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**apps** | [**List[ListApp]**](ListApp.md) |  | [optional] 
**total_apps** | **int** |  | [optional] 

## Example

```python
from fly_sdk.models.list_apps_response import ListAppsResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ListAppsResponse from a JSON string
list_apps_response_instance = ListAppsResponse.from_json(json)
# print the JSON string representation of the object
print ListAppsResponse.to_json()

# convert the object into a dict
list_apps_response_dict = list_apps_response_instance.to_dict()
# create an instance of ListAppsResponse from a dict
list_apps_response_form_dict = list_apps_response.from_dict(list_apps_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


