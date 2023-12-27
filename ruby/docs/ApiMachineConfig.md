# FlySDK::ApiMachineConfig

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **auto_destroy** | **Boolean** | Optional boolean telling the Machine to destroy itself once it’s complete (default false) | [optional] |
| **checks** | [**Hash&lt;String, ApiMachineCheck&gt;**](ApiMachineCheck.md) |  | [optional] |
| **disable_machine_autostart** | **Boolean** | Deprecated: use Service.Autostart instead | [optional] |
| **dns** | [**ApiDNSConfig**](ApiDNSConfig.md) |  | [optional] |
| **env** | **Hash&lt;String, String&gt;** | An object filled with key/value pairs to be set as environment variables | [optional] |
| **files** | [**Array&lt;ApiFile&gt;**](ApiFile.md) |  | [optional] |
| **guest** | [**ApiMachineGuest**](ApiMachineGuest.md) |  | [optional] |
| **image** | **String** | The docker image to run | [optional] |
| **init** | [**ApiMachineInit**](ApiMachineInit.md) |  | [optional] |
| **metadata** | **Hash&lt;String, String&gt;** |  | [optional] |
| **metrics** | [**ApiMachineMetrics**](ApiMachineMetrics.md) |  | [optional] |
| **mounts** | [**Array&lt;ApiMachineMount&gt;**](ApiMachineMount.md) |  | [optional] |
| **processes** | [**Array&lt;ApiMachineProcess&gt;**](ApiMachineProcess.md) |  | [optional] |
| **restart** | [**ApiMachineRestart**](ApiMachineRestart.md) |  | [optional] |
| **schedule** | **String** |  | [optional] |
| **services** | [**Array&lt;ApiMachineService&gt;**](ApiMachineService.md) |  | [optional] |
| **size** | **String** | Deprecated: use Guest instead | [optional] |
| **standbys** | **Array&lt;String&gt;** | Standbys enable a machine to be a standby for another. In the event of a hardware failure, the standby machine will be started. | [optional] |
| **statics** | [**Array&lt;ApiStatic&gt;**](ApiStatic.md) |  | [optional] |
| **stop_config** | [**ApiStopConfig**](ApiStopConfig.md) |  | [optional] |

## Example

```ruby
require 'fly-sdk-ruby'

instance = FlySDK::ApiMachineConfig.new(
  auto_destroy: null,
  checks: null,
  disable_machine_autostart: null,
  dns: null,
  env: null,
  files: null,
  guest: null,
  image: null,
  init: null,
  metadata: null,
  metrics: null,
  mounts: null,
  processes: null,
  restart: null,
  schedule: null,
  services: null,
  size: null,
  standbys: null,
  statics: null,
  stop_config: null
)
```

