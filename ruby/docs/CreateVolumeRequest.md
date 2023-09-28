# FlyApi::CreateVolumeRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **compute** | [**ApiMachineGuest**](ApiMachineGuest.md) |  | [optional] |
| **encrypted** | **Boolean** |  | [optional] |
| **fstype** | **String** |  | [optional] |
| **machines_only** | **Boolean** |  | [optional] |
| **name** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **require_unique_zone** | **Boolean** |  | [optional] |
| **size_gb** | **Integer** |  | [optional] |
| **snapshot_id** | **String** | restore from snapshot | [optional] |
| **source_volume_id** | **String** | fork from remote volume | [optional] |

## Example

```ruby
require 'fly-api-ruby'

instance = FlyApi::CreateVolumeRequest.new(
  compute: null,
  encrypted: null,
  fstype: null,
  machines_only: null,
  name: null,
  region: null,
  require_unique_zone: null,
  size_gb: null,
  snapshot_id: null,
  source_volume_id: null
)
```

