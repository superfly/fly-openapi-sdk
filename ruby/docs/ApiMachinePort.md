# FlyApi::ApiMachinePort

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **end_port** | **Integer** |  | [optional] |
| **force_https** | **Boolean** |  | [optional] |
| **handlers** | **Array&lt;String&gt;** |  | [optional] |
| **http_options** | [**ApiHTTPOptions**](ApiHTTPOptions.md) |  | [optional] |
| **port** | **Integer** |  | [optional] |
| **proxy_proto_options** | [**ApiProxyProtoOptions**](ApiProxyProtoOptions.md) |  | [optional] |
| **start_port** | **Integer** |  | [optional] |
| **tls_options** | [**ApiTLSOptions**](ApiTLSOptions.md) |  | [optional] |

## Example

```ruby
require 'fly-api-ruby'

instance = FlyApi::ApiMachinePort.new(
  end_port: null,
  force_https: null,
  handlers: null,
  http_options: null,
  port: null,
  proxy_proto_options: null,
  start_port: null,
  tls_options: null
)
```

