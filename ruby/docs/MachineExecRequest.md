# FlySDK::MachineExecRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cmd** | **String** | Deprecated: use Command instead | [optional] |
| **command** | **Array&lt;String&gt;** |  | [optional] |
| **timeout** | **Integer** |  | [optional] |

## Example

```ruby
require 'fly-sdk-ruby'

instance = FlySDK::MachineExecRequest.new(
  cmd: null,
  command: null,
  timeout: null
)
```

