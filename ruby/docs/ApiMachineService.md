# FlySDK::ApiMachineService

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **autostart** | **Boolean** |  | [optional] |
| **autostop** | **Boolean** |  | [optional] |
| **checks** | [**Array&lt;ApiMachineCheck&gt;**](ApiMachineCheck.md) |  | [optional] |
| **concurrency** | [**ApiMachineServiceConcurrency**](ApiMachineServiceConcurrency.md) |  | [optional] |
| **force_instance_description** | **String** |  | [optional] |
| **force_instance_key** | **String** |  | [optional] |
| **internal_port** | **Integer** |  | [optional] |
| **min_machines_running** | **Integer** |  | [optional] |
| **ports** | [**Array&lt;ApiMachinePort&gt;**](ApiMachinePort.md) |  | [optional] |
| **protocol** | **String** |  | [optional] |

## Example

```ruby
require 'fly-sdk-ruby'

instance = FlySDK::ApiMachineService.new(
  autostart: null,
  autostop: null,
  checks: null,
  concurrency: null,
  force_instance_description: null,
  force_instance_key: null,
  internal_port: null,
  min_machines_running: null,
  ports: null,
  protocol: null
)
```

