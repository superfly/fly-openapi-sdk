# FlyApi::ApiTLSOptions

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **alpn** | **Array&lt;String&gt;** |  | [optional] |
| **default_self_signed** | **Boolean** |  | [optional] |
| **versions** | **Array&lt;String&gt;** |  | [optional] |

## Example

```ruby
require 'fly-api-ruby'

instance = FlyApi::ApiTLSOptions.new(
  alpn: null,
  default_self_signed: null,
  versions: null
)
```

