# FlySDK::Volume

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **attached_alloc_id** | **String** |  | [optional] |
| **attached_machine_id** | **String** |  | [optional] |
| **block_size** | **Integer** |  | [optional] |
| **blocks** | **Integer** |  | [optional] |
| **blocks_avail** | **Integer** |  | [optional] |
| **blocks_free** | **Integer** |  | [optional] |
| **created_at** | **String** |  | [optional] |
| **encrypted** | **Boolean** |  | [optional] |
| **fstype** | **String** |  | [optional] |
| **id** | **String** |  | [optional] |
| **name** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **size_gb** | **Integer** |  | [optional] |
| **snapshot_retention** | **Integer** |  | [optional] |
| **state** | **String** |  | [optional] |
| **zone** | **String** |  | [optional] |

## Example

```ruby
require 'fly-sdk-ruby'

instance = FlySDK::Volume.new(
  attached_alloc_id: null,
  attached_machine_id: null,
  block_size: null,
  blocks: null,
  blocks_avail: null,
  blocks_free: null,
  created_at: null,
  encrypted: null,
  fstype: null,
  id: null,
  name: null,
  region: null,
  size_gb: null,
  snapshot_retention: null,
  state: null,
  zone: null
)
```

