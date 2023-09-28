# FlySDK::ApiMachineCheck

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **grace_period** | **String** |  | [optional] |
| **headers** | [**Array&lt;ApiMachineHTTPHeader&gt;**](ApiMachineHTTPHeader.md) |  | [optional] |
| **interval** | **String** |  | [optional] |
| **method** | **String** |  | [optional] |
| **path** | **String** |  | [optional] |
| **port** | **Integer** |  | [optional] |
| **protocol** | **String** |  | [optional] |
| **timeout** | **String** |  | [optional] |
| **tls_server_name** | **String** |  | [optional] |
| **tls_skip_verify** | **Boolean** |  | [optional] |
| **type** | **String** |  | [optional] |

## Example

```ruby
require 'fly-sdk-ruby'

instance = FlySDK::ApiMachineCheck.new(
  grace_period: null,
  headers: null,
  interval: null,
  method: null,
  path: null,
  port: null,
  protocol: null,
  timeout: null,
  tls_server_name: null,
  tls_skip_verify: null,
  type: null
)
```

