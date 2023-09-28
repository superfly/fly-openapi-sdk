# FlyApi::MachineExecRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cmd** | **String** | Deprecated: use Command instead | [optional] |
| **command** | **Array&lt;String&gt;** |  | [optional] |
| **timeout** | **Integer** |  | [optional] |

## Example

```ruby
require 'fly-api-ruby'

instance = FlyApi::MachineExecRequest.new(
  cmd: null,
  command: null,
  timeout: null
)
```

