# fly_sdk.AppsApi

All URIs are relative to *https://api.machines.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**apps_create**](AppsApi.md#apps_create) | **POST** /apps | 
[**apps_delete**](AppsApi.md#apps_delete) | **DELETE** /apps/{app_name} | 
[**apps_list**](AppsApi.md#apps_list) | **GET** /apps | 
[**apps_show**](AppsApi.md#apps_show) | **GET** /apps/{app_name} | 


# **apps_create**
> apps_create(request)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.create_app_request import CreateAppRequest
from fly_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.machines.dev/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = fly_sdk.Configuration(
    host = "https://api.machines.dev/v1"
)


# Enter a context with an instance of the API client
with fly_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = fly_sdk.AppsApi(api_client)
    request = fly_sdk.CreateAppRequest() # CreateAppRequest | App body

    try:
        api_instance.apps_create(request)
    except Exception as e:
        print("Exception when calling AppsApi->apps_create: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**CreateAppRequest**](CreateAppRequest.md)| App body | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Created |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apps_delete**
> apps_delete(app_name)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.machines.dev/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = fly_sdk.Configuration(
    host = "https://api.machines.dev/v1"
)


# Enter a context with an instance of the API client
with fly_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = fly_sdk.AppsApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name

    try:
        api_instance.apps_delete(app_name)
    except Exception as e:
        print("Exception when calling AppsApi->apps_delete: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 

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
**202** | Accepted |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apps_list**
> ListAppsResponse apps_list(org_slug)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.list_apps_response import ListAppsResponse
from fly_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.machines.dev/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = fly_sdk.Configuration(
    host = "https://api.machines.dev/v1"
)


# Enter a context with an instance of the API client
with fly_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = fly_sdk.AppsApi(api_client)
    org_slug = 'org_slug_example' # str | The org slug, or 'personal', to filter apps

    try:
        api_response = api_instance.apps_list(org_slug)
        print("The response of AppsApi->apps_list:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AppsApi->apps_list: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org_slug** | **str**| The org slug, or &#39;personal&#39;, to filter apps | 

### Return type

[**ListAppsResponse**](ListAppsResponse.md)

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

# **apps_show**
> App apps_show(app_name)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.app import App
from fly_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.machines.dev/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = fly_sdk.Configuration(
    host = "https://api.machines.dev/v1"
)


# Enter a context with an instance of the API client
with fly_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = fly_sdk.AppsApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name

    try:
        api_response = api_instance.apps_show(app_name)
        print("The response of AppsApi->apps_show:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AppsApi->apps_show: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 

### Return type

[**App**](App.md)

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

