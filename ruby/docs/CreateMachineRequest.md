# FlySDK::CreateMachineRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **config** | [**ApiMachineConfig**](ApiMachineConfig.md) |  | [optional] |
| **lease_ttl** | **Integer** |  | [optional] |
| **lsvd** | **Boolean** |  | [optional] |
| **name** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **skip_launch** | **Boolean** |  | [optional] |
| **skip_service_registration** | **Boolean** |  | [optional] |

## Example

```ruby
require 'fly-sdk-ruby'

instance = FlySDK::CreateMachineRequest.new(
  config: null,
  lease_ttl: null,
  lsvd: null,
  name: null,
  region: null,
  skip_launch: null,
  skip_service_registration: null
)
```

