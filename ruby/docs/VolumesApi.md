# FlyApi::VolumesApi

All URIs are relative to *https://api.machines.dev/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**volume_delete**](VolumesApi.md#volume_delete) | **DELETE** /apps/{app_name}/volumes/{volume_id} |  |
| [**volumes_create**](VolumesApi.md#volumes_create) | **POST** /apps/{app_name}/volumes |  |
| [**volumes_extend**](VolumesApi.md#volumes_extend) | **PUT** /apps/{app_name}/volumes/{volume_id}/extend |  |
| [**volumes_get_by_id**](VolumesApi.md#volumes_get_by_id) | **GET** /apps/{app_name}/volumes/{volume_id} |  |
| [**volumes_get_snapshots**](VolumesApi.md#volumes_get_snapshots) | **GET** /apps/{app_name}/volumes/{volume_id}/snapshots |  |
| [**volumes_list**](VolumesApi.md#volumes_list) | **GET** /apps/{app_name}/volumes |  |


## volume_delete

> <Volume> volume_delete(app_name, volume_id)



### Examples

```ruby
require 'time'
require 'fly-api-ruby'

api_instance = FlyApi::VolumesApi.new
app_name = 'app_name_example' # String | Fly App Name
volume_id = 'volume_id_example' # String | Volume ID

begin
  
  result = api_instance.volume_delete(app_name, volume_id)
  p result
rescue FlyApi::ApiError => e
  puts "Error when calling VolumesApi->volume_delete: #{e}"
end
```

#### Using the volume_delete_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Volume>, Integer, Hash)> volume_delete_with_http_info(app_name, volume_id)

```ruby
begin
  
  data, status_code, headers = api_instance.volume_delete_with_http_info(app_name, volume_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Volume>
rescue FlyApi::ApiError => e
  puts "Error when calling VolumesApi->volume_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **app_name** | **String** | Fly App Name |  |
| **volume_id** | **String** | Volume ID |  |

### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## volumes_create

> <Volume> volumes_create(app_name, request)



### Examples

```ruby
require 'time'
require 'fly-api-ruby'

api_instance = FlyApi::VolumesApi.new
app_name = 'app_name_example' # String | Fly App Name
request = FlyApi::CreateVolumeRequest.new # CreateVolumeRequest | Request body

begin
  
  result = api_instance.volumes_create(app_name, request)
  p result
rescue FlyApi::ApiError => e
  puts "Error when calling VolumesApi->volumes_create: #{e}"
end
```

#### Using the volumes_create_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Volume>, Integer, Hash)> volumes_create_with_http_info(app_name, request)

```ruby
begin
  
  data, status_code, headers = api_instance.volumes_create_with_http_info(app_name, request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Volume>
rescue FlyApi::ApiError => e
  puts "Error when calling VolumesApi->volumes_create_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **app_name** | **String** | Fly App Name |  |
| **request** | [**CreateVolumeRequest**](CreateVolumeRequest.md) | Request body |  |

### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## volumes_extend

> <ExtendVolumeResponse> volumes_extend(app_name, volume_id, request)



### Examples

```ruby
require 'time'
require 'fly-api-ruby'

api_instance = FlyApi::VolumesApi.new
app_name = 'app_name_example' # String | Fly App Name
volume_id = 'volume_id_example' # String | Volume ID
request = FlyApi::ExtendVolumeRequest.new # ExtendVolumeRequest | Request body

begin
  
  result = api_instance.volumes_extend(app_name, volume_id, request)
  p result
rescue FlyApi::ApiError => e
  puts "Error when calling VolumesApi->volumes_extend: #{e}"
end
```

#### Using the volumes_extend_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ExtendVolumeResponse>, Integer, Hash)> volumes_extend_with_http_info(app_name, volume_id, request)

```ruby
begin
  
  data, status_code, headers = api_instance.volumes_extend_with_http_info(app_name, volume_id, request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ExtendVolumeResponse>
rescue FlyApi::ApiError => e
  puts "Error when calling VolumesApi->volumes_extend_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **app_name** | **String** | Fly App Name |  |
| **volume_id** | **String** | Volume ID |  |
| **request** | [**ExtendVolumeRequest**](ExtendVolumeRequest.md) | Request body |  |

### Return type

[**ExtendVolumeResponse**](ExtendVolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## volumes_get_by_id

> <Volume> volumes_get_by_id(app_name, volume_id)



### Examples

```ruby
require 'time'
require 'fly-api-ruby'

api_instance = FlyApi::VolumesApi.new
app_name = 'app_name_example' # String | Fly App Name
volume_id = 'volume_id_example' # String | Volume ID

begin
  
  result = api_instance.volumes_get_by_id(app_name, volume_id)
  p result
rescue FlyApi::ApiError => e
  puts "Error when calling VolumesApi->volumes_get_by_id: #{e}"
end
```

#### Using the volumes_get_by_id_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Volume>, Integer, Hash)> volumes_get_by_id_with_http_info(app_name, volume_id)

```ruby
begin
  
  data, status_code, headers = api_instance.volumes_get_by_id_with_http_info(app_name, volume_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Volume>
rescue FlyApi::ApiError => e
  puts "Error when calling VolumesApi->volumes_get_by_id_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **app_name** | **String** | Fly App Name |  |
| **volume_id** | **String** | Volume ID |  |

### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## volumes_get_snapshots

> <Array<VolumeSnapshot>> volumes_get_snapshots(app_name, volume_id)



### Examples

```ruby
require 'time'
require 'fly-api-ruby'

api_instance = FlyApi::VolumesApi.new
app_name = 'app_name_example' # String | Fly App Name
volume_id = 'volume_id_example' # String | Volume ID

begin
  
  result = api_instance.volumes_get_snapshots(app_name, volume_id)
  p result
rescue FlyApi::ApiError => e
  puts "Error when calling VolumesApi->volumes_get_snapshots: #{e}"
end
```

#### Using the volumes_get_snapshots_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<VolumeSnapshot>>, Integer, Hash)> volumes_get_snapshots_with_http_info(app_name, volume_id)

```ruby
begin
  
  data, status_code, headers = api_instance.volumes_get_snapshots_with_http_info(app_name, volume_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<VolumeSnapshot>>
rescue FlyApi::ApiError => e
  puts "Error when calling VolumesApi->volumes_get_snapshots_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **app_name** | **String** | Fly App Name |  |
| **volume_id** | **String** | Volume ID |  |

### Return type

[**Array&lt;VolumeSnapshot&gt;**](VolumeSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## volumes_list

> <Array<Volume>> volumes_list(app_name)



### Examples

```ruby
require 'time'
require 'fly-api-ruby'

api_instance = FlyApi::VolumesApi.new
app_name = 'app_name_example' # String | Fly App Name

begin
  
  result = api_instance.volumes_list(app_name)
  p result
rescue FlyApi::ApiError => e
  puts "Error when calling VolumesApi->volumes_list: #{e}"
end
```

#### Using the volumes_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Volume>>, Integer, Hash)> volumes_list_with_http_info(app_name)

```ruby
begin
  
  data, status_code, headers = api_instance.volumes_list_with_http_info(app_name)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Volume>>
rescue FlyApi::ApiError => e
  puts "Error when calling VolumesApi->volumes_list_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **app_name** | **String** | Fly App Name |  |

### Return type

[**Array&lt;Volume&gt;**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

