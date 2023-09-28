# ProcessStat


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**command** | **str** |  | [optional] 
**cpu** | **int** |  | [optional] 
**directory** | **str** |  | [optional] 
**listen_sockets** | [**List[ListenSocket]**](ListenSocket.md) |  | [optional] 
**pid** | **int** |  | [optional] 
**rss** | **int** |  | [optional] 
**rtime** | **int** |  | [optional] 
**stime** | **int** |  | [optional] 

## Example

```python
from fly-sdk.models.process_stat import ProcessStat

# TODO update the JSON string below
json = "{}"
# create an instance of ProcessStat from a JSON string
process_stat_instance = ProcessStat.from_json(json)
# print the JSON string representation of the object
print ProcessStat.to_json()

# convert the object into a dict
process_stat_dict = process_stat_instance.to_dict()
# create an instance of ProcessStat from a dict
process_stat_form_dict = process_stat.from_dict(process_stat_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


