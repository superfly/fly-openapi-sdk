# fly-sdk.VolumesApi

All URIs are relative to *https://api.machines.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_volume_snapshot**](VolumesApi.md#create_volume_snapshot) | **POST** /apps/{app_name}/volumes/{volume_id}/snapshots | 
[**volume_delete**](VolumesApi.md#volume_delete) | **DELETE** /apps/{app_name}/volumes/{volume_id} | 
[**volumes_create**](VolumesApi.md#volumes_create) | **POST** /apps/{app_name}/volumes | 
[**volumes_extend**](VolumesApi.md#volumes_extend) | **PUT** /apps/{app_name}/volumes/{volume_id}/extend | 
[**volumes_get_by_id**](VolumesApi.md#volumes_get_by_id) | **GET** /apps/{app_name}/volumes/{volume_id} | 
[**volumes_list**](VolumesApi.md#volumes_list) | **GET** /apps/{app_name}/volumes | 
[**volumes_list_snapshots**](VolumesApi.md#volumes_list_snapshots) | **GET** /apps/{app_name}/volumes/{volume_id}/snapshots | 
[**volumes_update**](VolumesApi.md#volumes_update) | **POST** /apps/{app_name}/volumes/{volume_id} | 


# **create_volume_snapshot**
> create_volume_snapshot(app_name, volume_id)



### Example

```python
import time
import os
import fly-sdk
from fly-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.machines.dev/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = fly-sdk.Configuration(
    host = "https://api.machines.dev/v1"
)


# Enter a context with an instance of the API client
with fly-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = fly-sdk.VolumesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    volume_id = 'volume_id_example' # str | Volume ID

    try:
        api_instance.create_volume_snapshot(app_name, volume_id)
    except Exception as e:
        print("Exception when calling VolumesApi->create_volume_snapshot: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **volume_id** | **str**| Volume ID | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **volume_delete**
> Volume volume_delete(app_name, volume_id)



### Example

```python
import time
import os
import fly-sdk
from fly-sdk.models.volume import Volume
from fly-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.machines.dev/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = fly-sdk.Configuration(
    host = "https://api.machines.dev/v1"
)


# Enter a context with an instance of the API client
with fly-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = fly-sdk.VolumesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    volume_id = 'volume_id_example' # str | Volume ID

    try:
        api_response = api_instance.volume_delete(app_name, volume_id)
        print("The response of VolumesApi->volume_delete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling VolumesApi->volume_delete: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **volume_id** | **str**| Volume ID | 

### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **volumes_create**
> Volume volumes_create(app_name, request)



### Example

```python
import time
import os
import fly-sdk
from fly-sdk.models.create_volume_request import CreateVolumeRequest
from fly-sdk.models.volume import Volume
from fly-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.machines.dev/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = fly-sdk.Configuration(
    host = "https://api.machines.dev/v1"
)


# Enter a context with an instance of the API client
with fly-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = fly-sdk.VolumesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    request = fly-sdk.CreateVolumeRequest() # CreateVolumeRequest | Request body

    try:
        api_response = api_instance.volumes_create(app_name, request)
        print("The response of VolumesApi->volumes_create:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling VolumesApi->volumes_create: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **request** | [**CreateVolumeRequest**](CreateVolumeRequest.md)| Request body | 

### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **volumes_extend**
> ExtendVolumeResponse volumes_extend(app_name, volume_id, request)



### Example

```python
import time
import os
import fly-sdk
from fly-sdk.models.extend_volume_request import ExtendVolumeRequest
from fly-sdk.models.extend_volume_response import ExtendVolumeResponse
from fly-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.machines.dev/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = fly-sdk.Configuration(
    host = "https://api.machines.dev/v1"
)


# Enter a context with an instance of the API client
with fly-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = fly-sdk.VolumesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    volume_id = 'volume_id_example' # str | Volume ID
    request = fly-sdk.ExtendVolumeRequest() # ExtendVolumeRequest | Request body

    try:
        api_response = api_instance.volumes_extend(app_name, volume_id, request)
        print("The response of VolumesApi->volumes_extend:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling VolumesApi->volumes_extend: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **volume_id** | **str**| Volume ID | 
 **request** | [**ExtendVolumeRequest**](ExtendVolumeRequest.md)| Request body | 

### Return type

[**ExtendVolumeResponse**](ExtendVolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **volumes_get_by_id**
> Volume volumes_get_by_id(app_name, volume_id)



### Example

```python
import time
import os
import fly-sdk
from fly-sdk.models.volume import Volume
from fly-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.machines.dev/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = fly-sdk.Configuration(
    host = "https://api.machines.dev/v1"
)


# Enter a context with an instance of the API client
with fly-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = fly-sdk.VolumesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    volume_id = 'volume_id_example' # str | Volume ID

    try:
        api_response = api_instance.volumes_get_by_id(app_name, volume_id)
        print("The response of VolumesApi->volumes_get_by_id:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling VolumesApi->volumes_get_by_id: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **volume_id** | **str**| Volume ID | 

### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **volumes_list**
> List[Volume] volumes_list(app_name)



### Example

```python
import time
import os
import fly-sdk
from fly-sdk.models.volume import Volume
from fly-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.machines.dev/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = fly-sdk.Configuration(
    host = "https://api.machines.dev/v1"
)


# Enter a context with an instance of the API client
with fly-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = fly-sdk.VolumesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name

    try:
        api_response = api_instance.volumes_list(app_name)
        print("The response of VolumesApi->volumes_list:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling VolumesApi->volumes_list: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 

### Return type

[**List[Volume]**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **volumes_list_snapshots**
> List[VolumeSnapshot] volumes_list_snapshots(app_name, volume_id)



### Example

```python
import time
import os
import fly-sdk
from fly-sdk.models.volume_snapshot import VolumeSnapshot
from fly-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.machines.dev/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = fly-sdk.Configuration(
    host = "https://api.machines.dev/v1"
)


# Enter a context with an instance of the API client
with fly-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = fly-sdk.VolumesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    volume_id = 'volume_id_example' # str | Volume ID

    try:
        api_response = api_instance.volumes_list_snapshots(app_name, volume_id)
        print("The response of VolumesApi->volumes_list_snapshots:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling VolumesApi->volumes_list_snapshots: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **volume_id** | **str**| Volume ID | 

### Return type

[**List[VolumeSnapshot]**](VolumeSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **volumes_update**
> Volume volumes_update(app_name, volume_id, request)



### Example

```python
import time
import os
import fly-sdk
from fly-sdk.models.update_volume_request import UpdateVolumeRequest
from fly-sdk.models.volume import Volume
from fly-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.machines.dev/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = fly-sdk.Configuration(
    host = "https://api.machines.dev/v1"
)


# Enter a context with an instance of the API client
with fly-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = fly-sdk.VolumesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    volume_id = 'volume_id_example' # str | Volume ID
    request = fly-sdk.UpdateVolumeRequest() # UpdateVolumeRequest | Request body

    try:
        api_response = api_instance.volumes_update(app_name, volume_id, request)
        print("The response of VolumesApi->volumes_update:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling VolumesApi->volumes_update: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **volume_id** | **str**| Volume ID | 
 **request** | [**UpdateVolumeRequest**](UpdateVolumeRequest.md)| Request body | 

### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

