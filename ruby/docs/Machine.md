# FlySDK::Machine

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **checks** | [**Array&lt;CheckStatus&gt;**](CheckStatus.md) |  | [optional] |
| **config** | [**ApiMachineConfig**](ApiMachineConfig.md) |  | [optional] |
| **created_at** | **String** |  | [optional] |
| **events** | [**Array&lt;MachineEvent&gt;**](MachineEvent.md) |  | [optional] |
| **id** | **String** |  | [optional] |
| **image_ref** | [**ImageRef**](ImageRef.md) |  | [optional] |
| **instance_id** | **String** | InstanceID is unique for each version of the machine | [optional] |
| **name** | **String** |  | [optional] |
| **nonce** | **String** | Nonce is only every returned on machine creation if a lease_duration was provided. | [optional] |
| **private_ip** | **String** | PrivateIP is the internal 6PN address of the machine. | [optional] |
| **region** | **String** |  | [optional] |
| **state** | **String** |  | [optional] |
| **updated_at** | **String** |  | [optional] |

## Example

```ruby
require 'fly-sdk-ruby'

instance = FlySDK::Machine.new(
  checks: null,
  config: null,
  created_at: null,
  events: null,
  id: null,
  image_ref: null,
  instance_id: null,
  name: null,
  nonce: null,
  private_ip: null,
  region: null,
  state: null,
  updated_at: null
)
```

