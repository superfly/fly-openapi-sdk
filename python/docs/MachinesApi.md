# fly_sdk.MachinesApi

All URIs are relative to *https://api.machines.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**machines_cordon**](MachinesApi.md#machines_cordon) | **POST** /apps/{app_name}/machines/{machine_id}/cordon | 
[**machines_create**](MachinesApi.md#machines_create) | **POST** /apps/{app_name}/machines | 
[**machines_create_lease**](MachinesApi.md#machines_create_lease) | **POST** /apps/{app_name}/machines/{machine_id}/lease | 
[**machines_delete**](MachinesApi.md#machines_delete) | **DELETE** /apps/{app_name}/machines/{machine_id} | 
[**machines_delete_metadata**](MachinesApi.md#machines_delete_metadata) | **DELETE** /apps/{app_name}/machines/{machine_id}/metadata/{key} | 
[**machines_exec**](MachinesApi.md#machines_exec) | **POST** /apps/{app_name}/machines/{machine_id}/exec | 
[**machines_list**](MachinesApi.md#machines_list) | **GET** /apps/{app_name}/machines | 
[**machines_list_events**](MachinesApi.md#machines_list_events) | **GET** /apps/{app_name}/machines/{machine_id}/events | 
[**machines_list_processes**](MachinesApi.md#machines_list_processes) | **GET** /apps/{app_name}/machines/{machine_id}/ps | 
[**machines_list_versions**](MachinesApi.md#machines_list_versions) | **GET** /apps/{app_name}/machines/{machine_id}/versions | 
[**machines_release_lease**](MachinesApi.md#machines_release_lease) | **DELETE** /apps/{app_name}/machines/{machine_id}/lease | 
[**machines_restart**](MachinesApi.md#machines_restart) | **POST** /apps/{app_name}/machines/{machine_id}/restart | 
[**machines_show**](MachinesApi.md#machines_show) | **GET** /apps/{app_name}/machines/{machine_id} | 
[**machines_show_lease**](MachinesApi.md#machines_show_lease) | **GET** /apps/{app_name}/machines/{machine_id}/lease | 
[**machines_show_metadata**](MachinesApi.md#machines_show_metadata) | **GET** /apps/{app_name}/machines/{machine_id}/metadata | 
[**machines_signal**](MachinesApi.md#machines_signal) | **POST** /apps/{app_name}/machines/{machine_id}/signal | 
[**machines_start**](MachinesApi.md#machines_start) | **POST** /apps/{app_name}/machines/{machine_id}/start | 
[**machines_stop**](MachinesApi.md#machines_stop) | **POST** /apps/{app_name}/machines/{machine_id}/stop | 
[**machines_uncordon**](MachinesApi.md#machines_uncordon) | **POST** /apps/{app_name}/machines/{machine_id}/uncordon | 
[**machines_update**](MachinesApi.md#machines_update) | **POST** /apps/{app_name}/machines/{machine_id} | 
[**machines_update_metadata**](MachinesApi.md#machines_update_metadata) | **POST** /apps/{app_name}/machines/{machine_id}/metadata/{key} | 
[**machines_wait**](MachinesApi.md#machines_wait) | **GET** /apps/{app_name}/machines/{machine_id}/wait | 


# **machines_cordon**
> machines_cordon(app_name, machine_id)



“Cordoning” a machine refers to disabling its services, so the proxy won’t route requests to it. In flyctl this is used by blue/green deployments; one set of machines is started up with services disabled, and when they are all healthy, the services are enabled on the new machines and disabled on the old ones.

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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID

    try:
        api_instance.machines_cordon(app_name, machine_id)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_cordon: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 

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

# **machines_create**
> Machine machines_create(app_name, request)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.create_machine_request import CreateMachineRequest
from fly_sdk.models.machine import Machine
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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    request = fly_sdk.CreateMachineRequest() # CreateMachineRequest | Create machine request

    try:
        api_response = api_instance.machines_create(app_name, request)
        print("The response of MachinesApi->machines_create:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_create: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **request** | [**CreateMachineRequest**](CreateMachineRequest.md)| Create machine request | 

### Return type

[**Machine**](Machine.md)

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

# **machines_create_lease**
> Lease machines_create_lease(app_name, machine_id, request)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.create_lease_request import CreateLeaseRequest
from fly_sdk.models.lease import Lease
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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID
    request = fly_sdk.CreateLeaseRequest() # CreateLeaseRequest | Request body

    try:
        api_response = api_instance.machines_create_lease(app_name, machine_id, request)
        print("The response of MachinesApi->machines_create_lease:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_create_lease: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 
 **request** | [**CreateLeaseRequest**](CreateLeaseRequest.md)| Request body | 

### Return type

[**Lease**](Lease.md)

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

# **machines_delete**
> machines_delete(app_name, machine_id)



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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID

    try:
        api_instance.machines_delete(app_name, machine_id)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_delete: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 

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

# **machines_delete_metadata**
> machines_delete_metadata(app_name, machine_id, key)



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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID
    key = 'key_example' # str | Metadata Key

    try:
        api_instance.machines_delete_metadata(app_name, machine_id, key)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_delete_metadata: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 
 **key** | **str**| Metadata Key | 

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
**204** | No Content |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **machines_exec**
> str machines_exec(app_name, machine_id, request)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.machine_exec_request import MachineExecRequest
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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID
    request = fly_sdk.MachineExecRequest() # MachineExecRequest | Request body

    try:
        api_response = api_instance.machines_exec(app_name, machine_id, request)
        print("The response of MachinesApi->machines_exec:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_exec: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 
 **request** | [**MachineExecRequest**](MachineExecRequest.md)| Request body | 

### Return type

**str**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream, application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Raw command output bytes are written back |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **machines_list**
> List[Machine] machines_list(app_name, include_deleted=include_deleted, region=region)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.machine import Machine
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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    include_deleted = True # bool | Include deleted machines (optional)
    region = 'region_example' # str | Region filter (optional)

    try:
        api_response = api_instance.machines_list(app_name, include_deleted=include_deleted, region=region)
        print("The response of MachinesApi->machines_list:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_list: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **include_deleted** | **bool**| Include deleted machines | [optional] 
 **region** | **str**| Region filter | [optional] 

### Return type

[**List[Machine]**](Machine.md)

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

# **machines_list_events**
> List[MachineEvent] machines_list_events(app_name, machine_id)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.machine_event import MachineEvent
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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID

    try:
        api_response = api_instance.machines_list_events(app_name, machine_id)
        print("The response of MachinesApi->machines_list_events:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_list_events: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 

### Return type

[**List[MachineEvent]**](MachineEvent.md)

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

# **machines_list_processes**
> List[ProcessStat] machines_list_processes(app_name, machine_id, sort_by=sort_by, order=order)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.process_stat import ProcessStat
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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID
    sort_by = 'sort_by_example' # str | Sort by (optional)
    order = 'order_example' # str | Order (optional)

    try:
        api_response = api_instance.machines_list_processes(app_name, machine_id, sort_by=sort_by, order=order)
        print("The response of MachinesApi->machines_list_processes:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_list_processes: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 
 **sort_by** | **str**| Sort by | [optional] 
 **order** | **str**| Order | [optional] 

### Return type

[**List[ProcessStat]**](ProcessStat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **machines_list_versions**
> List[MachineVersion] machines_list_versions(app_name, machine_id)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.machine_version import MachineVersion
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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID

    try:
        api_response = api_instance.machines_list_versions(app_name, machine_id)
        print("The response of MachinesApi->machines_list_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_list_versions: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 

### Return type

[**List[MachineVersion]**](MachineVersion.md)

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

# **machines_release_lease**
> machines_release_lease(app_name, machine_id)



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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID

    try:
        api_instance.machines_release_lease(app_name, machine_id)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_release_lease: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 

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

# **machines_restart**
> machines_restart(app_name, machine_id, timeout=timeout)



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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID
    timeout = 'timeout_example' # str | Restart timeout as a Go duration string or number of seconds (optional)

    try:
        api_instance.machines_restart(app_name, machine_id, timeout=timeout)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_restart: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 
 **timeout** | **str**| Restart timeout as a Go duration string or number of seconds | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **machines_show**
> Machine machines_show(app_name, machine_id)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.machine import Machine
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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID

    try:
        api_response = api_instance.machines_show(app_name, machine_id)
        print("The response of MachinesApi->machines_show:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_show: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 

### Return type

[**Machine**](Machine.md)

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

# **machines_show_lease**
> Lease machines_show_lease(app_name, machine_id)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.lease import Lease
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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID

    try:
        api_response = api_instance.machines_show_lease(app_name, machine_id)
        print("The response of MachinesApi->machines_show_lease:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_show_lease: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 

### Return type

[**Lease**](Lease.md)

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

# **machines_show_metadata**
> Dict[str, str] machines_show_metadata(app_name, machine_id)



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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID

    try:
        api_response = api_instance.machines_show_metadata(app_name, machine_id)
        print("The response of MachinesApi->machines_show_metadata:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_show_metadata: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 

### Return type

**Dict[str, str]**

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

# **machines_signal**
> machines_signal(app_name, machine_id, request)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.signal_request import SignalRequest
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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID
    request = fly_sdk.SignalRequest() # SignalRequest | Request body

    try:
        api_instance.machines_signal(app_name, machine_id, request)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_signal: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 
 **request** | [**SignalRequest**](SignalRequest.md)| Request body | 

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
**200** | OK |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **machines_start**
> machines_start(app_name, machine_id)



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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID

    try:
        api_instance.machines_start(app_name, machine_id)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_start: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 

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

# **machines_stop**
> machines_stop(app_name, machine_id, request=request)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.stop_request import StopRequest
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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID
    request = fly_sdk.StopRequest() # StopRequest | Optional request body (optional)

    try:
        api_instance.machines_stop(app_name, machine_id, request=request)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_stop: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 
 **request** | [**StopRequest**](StopRequest.md)| Optional request body | [optional] 

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
**200** | OK |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **machines_uncordon**
> machines_uncordon(app_name, machine_id)



“Cordoning” a machine refers to disabling its services, so the proxy won’t route requests to it. In flyctl this is used by blue/green deployments; one set of machines is started up with services disabled, and when they are all healthy, the services are enabled on the new machines and disabled on the old ones.

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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID

    try:
        api_instance.machines_uncordon(app_name, machine_id)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_uncordon: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 

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

# **machines_update**
> Machine machines_update(app_name, machine_id, request)



### Example

```python
import time
import os
import fly_sdk
from fly_sdk.models.machine import Machine
from fly_sdk.models.update_machine_request import UpdateMachineRequest
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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID
    request = fly_sdk.UpdateMachineRequest() # UpdateMachineRequest | Request body

    try:
        api_response = api_instance.machines_update(app_name, machine_id, request)
        print("The response of MachinesApi->machines_update:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_update: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 
 **request** | [**UpdateMachineRequest**](UpdateMachineRequest.md)| Request body | 

### Return type

[**Machine**](Machine.md)

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

# **machines_update_metadata**
> machines_update_metadata(app_name, machine_id, key)



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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID
    key = 'key_example' # str | Metadata Key

    try:
        api_instance.machines_update_metadata(app_name, machine_id, key)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_update_metadata: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 
 **key** | **str**| Metadata Key | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No Content |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **machines_wait**
> machines_wait(app_name, machine_id, instance_id=instance_id, timeout=timeout, state=state)



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
    api_instance = fly_sdk.MachinesApi(api_client)
    app_name = 'app_name_example' # str | Fly App Name
    machine_id = 'machine_id_example' # str | Machine ID
    instance_id = 'instance_id_example' # str | instance? version? TODO (optional)
    timeout = 56 # int | wait timeout. default 60s (optional)
    state = 'state_example' # str | desired state (optional)

    try:
        api_instance.machines_wait(app_name, machine_id, instance_id=instance_id, timeout=timeout, state=state)
    except Exception as e:
        print("Exception when calling MachinesApi->machines_wait: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_name** | **str**| Fly App Name | 
 **machine_id** | **str**| Machine ID | 
 **instance_id** | **str**| instance? version? TODO | [optional] 
 **timeout** | **int**| wait timeout. default 60s | [optional] 
 **state** | **str**| desired state | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

