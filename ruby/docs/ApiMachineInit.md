# FlyApi::ApiMachineInit

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cmd** | **Array&lt;String&gt;** |  | [optional] |
| **entrypoint** | **Array&lt;String&gt;** |  | [optional] |
| **exec** | **Array&lt;String&gt;** |  | [optional] |
| **kernel_args** | **Array&lt;String&gt;** |  | [optional] |
| **swap_size_mb** | **Integer** |  | [optional] |
| **tty** | **Boolean** |  | [optional] |

## Example

```ruby
require 'fly-api-ruby'

instance = FlyApi::ApiMachineInit.new(
  cmd: null,
  entrypoint: null,
  exec: null,
  kernel_args: null,
  swap_size_mb: null,
  tty: null
)
```

