# FlyApi::ApiMachineRestart

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **max_retries** | **Integer** | MaxRetries is only relevant with the on-failure policy. | [optional] |
| **policy** | **String** |  | [optional] |

## Example

```ruby
require 'fly-api-ruby'

instance = FlyApi::ApiMachineRestart.new(
  max_retries: null,
  policy: null
)
```

