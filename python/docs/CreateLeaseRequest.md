# CreateLeaseRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** |  | [optional] 
**ttl** | **int** |  | [optional] 

## Example

```python
from fly-sdk.models.create_lease_request import CreateLeaseRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateLeaseRequest from a JSON string
create_lease_request_instance = CreateLeaseRequest.from_json(json)
# print the JSON string representation of the object
print CreateLeaseRequest.to_json()

# convert the object into a dict
create_lease_request_dict = create_lease_request_instance.to_dict()
# create an instance of CreateLeaseRequest from a dict
create_lease_request_form_dict = create_lease_request.from_dict(create_lease_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


