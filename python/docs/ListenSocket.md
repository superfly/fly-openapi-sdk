# ListenSocket


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** |  | [optional] 
**proto** | **str** |  | [optional] 

## Example

```python
from fly-sdk.models.listen_socket import ListenSocket

# TODO update the JSON string below
json = "{}"
# create an instance of ListenSocket from a JSON string
listen_socket_instance = ListenSocket.from_json(json)
# print the JSON string representation of the object
print ListenSocket.to_json()

# convert the object into a dict
listen_socket_dict = listen_socket_instance.to_dict()
# create an instance of ListenSocket from a dict
listen_socket_form_dict = listen_socket.from_dict(listen_socket_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


