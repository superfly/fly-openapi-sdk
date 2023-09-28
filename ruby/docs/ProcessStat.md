# FlyApi::ProcessStat

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **command** | **String** |  | [optional] |
| **cpu** | **Integer** |  | [optional] |
| **directory** | **String** |  | [optional] |
| **listen_sockets** | [**Array&lt;ListenSocket&gt;**](ListenSocket.md) |  | [optional] |
| **pid** | **Integer** |  | [optional] |
| **rss** | **Integer** |  | [optional] |
| **rtime** | **Integer** |  | [optional] |
| **stime** | **Integer** |  | [optional] |

## Example

```ruby
require 'fly-api-ruby'

instance = FlyApi::ProcessStat.new(
  command: null,
  cpu: null,
  directory: null,
  listen_sockets: null,
  pid: null,
  rss: null,
  rtime: null,
  stime: null
)
```

