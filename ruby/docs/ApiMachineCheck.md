# FlySDK::ApiMachineCheck

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **grace_period** | **String** | The time to wait after a VM starts before checking its health | [optional] |
| **headers** | [**Array&lt;ApiMachineHTTPHeader&gt;**](ApiMachineHTTPHeader.md) |  | [optional] |
| **interval** | **String** | The time between connectivity checks | [optional] |
| **method** | **String** | For http checks, the HTTP method to use to when making the request | [optional] |
| **path** | **String** | For http checks, the path to send the request to | [optional] |
| **port** | **Integer** | The port to connect to, often the same as internal_port | [optional] |
| **protocol** | **String** | For http checks, whether to use http or https | [optional] |
| **timeout** | **String** | The maximum time a connection can take before being reported as failing its health check | [optional] |
| **tls_server_name** | **String** | If the protocol is https, the hostname to use for TLS certificate validation | [optional] |
| **tls_skip_verify** | **Boolean** | For http checks with https protocol, whether or not to verify the TLS certificate | [optional] |
| **type** | **String** | tcp or http | [optional] |

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

