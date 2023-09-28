# FlyApi::AppsApi

All URIs are relative to *https://api.machines.dev/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**apps_create**](AppsApi.md#apps_create) | **POST** /apps |  |
| [**apps_delete**](AppsApi.md#apps_delete) | **DELETE** /apps/{app_name} |  |
| [**apps_list**](AppsApi.md#apps_list) | **GET** /apps |  |
| [**apps_show**](AppsApi.md#apps_show) | **GET** /apps/{app_name} |  |


## apps_create

> apps_create(request)



### Examples

```ruby
require 'time'
require 'fly-api-ruby'

api_instance = FlyApi::AppsApi.new
request = FlyApi::CreateAppRequest.new # CreateAppRequest | App body

begin
  
  api_instance.apps_create(request)
rescue FlyApi::ApiError => e
  puts "Error when calling AppsApi->apps_create: #{e}"
end
```

#### Using the apps_create_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> apps_create_with_http_info(request)

```ruby
begin
  
  data, status_code, headers = api_instance.apps_create_with_http_info(request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue FlyApi::ApiError => e
  puts "Error when calling AppsApi->apps_create_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request** | [**CreateAppRequest**](CreateAppRequest.md) | App body |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## apps_delete

> apps_delete(app_name)



### Examples

```ruby
require 'time'
require 'fly-api-ruby'

api_instance = FlyApi::AppsApi.new
app_name = 'app_name_example' # String | Fly App Name

begin
  
  api_instance.apps_delete(app_name)
rescue FlyApi::ApiError => e
  puts "Error when calling AppsApi->apps_delete: #{e}"
end
```

#### Using the apps_delete_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> apps_delete_with_http_info(app_name)

```ruby
begin
  
  data, status_code, headers = api_instance.apps_delete_with_http_info(app_name)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue FlyApi::ApiError => e
  puts "Error when calling AppsApi->apps_delete_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **app_name** | **String** | Fly App Name |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## apps_list

> <ListAppsResponse> apps_list(opts)



### Examples

```ruby
require 'time'
require 'fly-api-ruby'

api_instance = FlyApi::AppsApi.new
opts = {
  org_slug: 'org_slug_example' # String | The org slug, or 'personal', to filter apps
}

begin
  
  result = api_instance.apps_list(opts)
  p result
rescue FlyApi::ApiError => e
  puts "Error when calling AppsApi->apps_list: #{e}"
end
```

#### Using the apps_list_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ListAppsResponse>, Integer, Hash)> apps_list_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.apps_list_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ListAppsResponse>
rescue FlyApi::ApiError => e
  puts "Error when calling AppsApi->apps_list_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **org_slug** | **String** | The org slug, or &#39;personal&#39;, to filter apps | [optional] |

### Return type

[**ListAppsResponse**](ListAppsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## apps_show

> <App> apps_show(app_name)



### Examples

```ruby
require 'time'
require 'fly-api-ruby'

api_instance = FlyApi::AppsApi.new
app_name = 'app_name_example' # String | Fly App Name

begin
  
  result = api_instance.apps_show(app_name)
  p result
rescue FlyApi::ApiError => e
  puts "Error when calling AppsApi->apps_show: #{e}"
end
```

#### Using the apps_show_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<App>, Integer, Hash)> apps_show_with_http_info(app_name)

```ruby
begin
  
  data, status_code, headers = api_instance.apps_show_with_http_info(app_name)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <App>
rescue FlyApi::ApiError => e
  puts "Error when calling AppsApi->apps_show_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **app_name** | **String** | Fly App Name |  |

### Return type

[**App**](App.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

