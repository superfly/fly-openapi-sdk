# \VolumesAPI

All URIs are relative to *https://api.machines.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VolumeDelete**](VolumesAPI.md#VolumeDelete) | **Delete** /apps/{app_name}/volumes/{volume_id} | 
[**VolumesCreate**](VolumesAPI.md#VolumesCreate) | **Post** /apps/{app_name}/volumes | 
[**VolumesExtend**](VolumesAPI.md#VolumesExtend) | **Put** /apps/{app_name}/volumes/{volume_id}/extend | 
[**VolumesGetById**](VolumesAPI.md#VolumesGetById) | **Get** /apps/{app_name}/volumes/{volume_id} | 
[**VolumesGetSnapshots**](VolumesAPI.md#VolumesGetSnapshots) | **Get** /apps/{app_name}/volumes/{volume_id}/snapshots | 
[**VolumesList**](VolumesAPI.md#VolumesList) | **Get** /apps/{app_name}/volumes | 
[**VolumesUpdate**](VolumesAPI.md#VolumesUpdate) | **Post** /apps/{app_name}/volumes/{volume_id} | 



## VolumeDelete

> Volume VolumeDelete(ctx, appName, volumeId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    appName := "appName_example" // string | Fly App Name
    volumeId := "volumeId_example" // string | Volume ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.VolumeDelete(context.Background(), appName, volumeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumeDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VolumeDelete`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumeDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesCreate

> Volume VolumesCreate(ctx, appName).Request(request).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    appName := "appName_example" // string | Fly App Name
    request := *openapiclient.NewCreateVolumeRequest() // CreateVolumeRequest | Request body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.VolumesCreate(context.Background(), appName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VolumesCreate`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**CreateVolumeRequest**](CreateVolumeRequest.md) | Request body | 

### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesExtend

> ExtendVolumeResponse VolumesExtend(ctx, appName, volumeId).Request(request).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    appName := "appName_example" // string | Fly App Name
    volumeId := "volumeId_example" // string | Volume ID
    request := *openapiclient.NewExtendVolumeRequest() // ExtendVolumeRequest | Request body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.VolumesExtend(context.Background(), appName, volumeId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesExtend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VolumesExtend`: ExtendVolumeResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesExtend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesExtendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**ExtendVolumeRequest**](ExtendVolumeRequest.md) | Request body | 

### Return type

[**ExtendVolumeResponse**](ExtendVolumeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesGetById

> Volume VolumesGetById(ctx, appName, volumeId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    appName := "appName_example" // string | Fly App Name
    volumeId := "volumeId_example" // string | Volume ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.VolumesGetById(context.Background(), appName, volumeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VolumesGetById`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesGetSnapshots

> []VolumeSnapshot VolumesGetSnapshots(ctx, appName, volumeId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    appName := "appName_example" // string | Fly App Name
    volumeId := "volumeId_example" // string | Volume ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.VolumesGetSnapshots(context.Background(), appName, volumeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesGetSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VolumesGetSnapshots`: []VolumeSnapshot
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesGetSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesGetSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]VolumeSnapshot**](VolumeSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesList

> []Volume VolumesList(ctx, appName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    appName := "appName_example" // string | Fly App Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.VolumesList(context.Background(), appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VolumesList`: []Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesUpdate

> Volume VolumesUpdate(ctx, appName, volumeId).Request(request).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    appName := "appName_example" // string | Fly App Name
    volumeId := "volumeId_example" // string | Volume ID
    request := *openapiclient.NewUpdateVolumeRequest() // UpdateVolumeRequest | Request body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.VolumesUpdate(context.Background(), appName, volumeId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VolumesUpdate`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | Fly App Name | 
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateVolumeRequest**](UpdateVolumeRequest.md) | Request body | 

### Return type

[**Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

