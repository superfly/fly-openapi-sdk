# ListApp


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**machine_count** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**network** | **object** |  | [optional] 

## Example

```python
from fly-sdk.models.list_app import ListApp

# TODO update the JSON string below
json = "{}"
# create an instance of ListApp from a JSON string
list_app_instance = ListApp.from_json(json)
# print the JSON string representation of the object
print ListApp.to_json()

# convert the object into a dict
list_app_dict = list_app_instance.to_dict()
# create an instance of ListApp from a dict
list_app_form_dict = list_app.from_dict(list_app_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


