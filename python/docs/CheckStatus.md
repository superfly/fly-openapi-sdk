# CheckStatus


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | [optional] 
**output** | **str** |  | [optional] 
**status** | **str** |  | [optional] 
**updated_at** | **str** |  | [optional] 

## Example

```python
from fly-sdk.models.check_status import CheckStatus

# TODO update the JSON string below
json = "{}"
# create an instance of CheckStatus from a JSON string
check_status_instance = CheckStatus.from_json(json)
# print the JSON string representation of the object
print CheckStatus.to_json()

# convert the object into a dict
check_status_dict = check_status_instance.to_dict()
# create an instance of CheckStatus from a dict
check_status_form_dict = check_status.from_dict(check_status_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


