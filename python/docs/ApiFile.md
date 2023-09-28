# ApiFile


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**guest_path** | **str** | GuestPath is the path on the machine where the file will be written and must be an absolute path. i.e. /full/path/to/file.json | [optional] 
**raw_value** | **str** | RawValue containts the base64 encoded string of the file contents. | [optional] 
**secret_name** | **str** | SecretName is the name of the secret that contains the base64 encoded file contents. | [optional] 

## Example

```python
from fly-sdk.models.api_file import ApiFile

# TODO update the JSON string below
json = "{}"
# create an instance of ApiFile from a JSON string
api_file_instance = ApiFile.from_json(json)
# print the JSON string representation of the object
print ApiFile.to_json()

# convert the object into a dict
api_file_dict = api_file_instance.to_dict()
# create an instance of ApiFile from a dict
api_file_form_dict = api_file.from_dict(api_file_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


