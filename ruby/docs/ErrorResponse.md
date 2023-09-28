# FlyApi::ErrorResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **details** | **Object** | Deprecated | [optional] |
| **error** | **String** |  | [optional] |
| **status** | [**MainStatusCode**](MainStatusCode.md) |  | [optional] |

## Example

```ruby
require 'fly-api-ruby'

instance = FlyApi::ErrorResponse.new(
  details: null,
  error: null,
  status: null
)
```

