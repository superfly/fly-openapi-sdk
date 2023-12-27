# FlySDK::ApiHTTPOptions

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **compress** | **Boolean** |  | [optional] |
| **h2_backend** | **Boolean** |  | [optional] |
| **response** | [**ApiHTTPResponseOptions**](ApiHTTPResponseOptions.md) |  | [optional] |

## Example

```ruby
require 'fly-sdk-ruby'

instance = FlySDK::ApiHTTPOptions.new(
  compress: null,
  h2_backend: null,
  response: null
)
```

