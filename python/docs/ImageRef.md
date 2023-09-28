# ImageRef


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**digest** | **str** |  | [optional] 
**labels** | **Dict[str, str]** |  | [optional] 
**registry** | **str** |  | [optional] 
**repository** | **str** |  | [optional] 
**tag** | **str** |  | [optional] 

## Example

```python
from fly-sdk.models.image_ref import ImageRef

# TODO update the JSON string below
json = "{}"
# create an instance of ImageRef from a JSON string
image_ref_instance = ImageRef.from_json(json)
# print the JSON string representation of the object
print ImageRef.to_json()

# convert the object into a dict
image_ref_dict = image_ref_instance.to_dict()
# create an instance of ImageRef from a dict
image_ref_form_dict = image_ref.from_dict(image_ref_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


