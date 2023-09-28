# FlyApi::ApiMachineServiceConcurrency

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **hard_limit** | **Integer** |  | [optional] |
| **soft_limit** | **Integer** |  | [optional] |
| **type** | **String** |  | [optional] |

## Example

```ruby
require 'fly-api-ruby'

instance = FlyApi::ApiMachineServiceConcurrency.new(
  hard_limit: null,
  soft_limit: null,
  type: null
)
```

