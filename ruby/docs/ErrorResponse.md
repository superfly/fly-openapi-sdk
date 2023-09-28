# FlySDK::ErrorResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **details** | **Object** | Deprecated | [optional] |
| **error** | **String** |  | [optional] |
| **status** | [**MainStatusCode**](MainStatusCode.md) |  | [optional] |

## Example

```ruby
require 'fly-sdk-ruby'

instance = FlySDK::ErrorResponse.new(
  details: null,
  error: null,
  status: null
)
```

