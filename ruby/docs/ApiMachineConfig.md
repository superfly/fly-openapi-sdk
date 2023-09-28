# FlyApi::ApiMachineConfig

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **auto_destroy** | **Boolean** |  | [optional] |
| **checks** | [**Hash&lt;String, ApiMachineCheck&gt;**](ApiMachineCheck.md) |  | [optional] |
| **disable_machine_autostart** | **Boolean** | Deprecated: use Service.Autostart instead | [optional] |
| **dns** | [**ApiDNSConfig**](ApiDNSConfig.md) |  | [optional] |
| **env** | **Hash&lt;String, String&gt;** | Fields managed from fly.toml If you add anything here, ensure appconfig.Config.ToMachine() is updated | [optional] |
| **files** | [**Array&lt;ApiFile&gt;**](ApiFile.md) |  | [optional] |
| **guest** | [**ApiMachineGuest**](ApiMachineGuest.md) |  | [optional] |
| **host_dedication_id** | **String** |  | [optional] |
| **image** | **String** | Set by fly deploy or fly machines commands | [optional] |
| **init** | [**ApiMachineInit**](ApiMachineInit.md) |  | [optional] |
| **metadata** | **Hash&lt;String, String&gt;** |  | [optional] |
| **metrics** | [**ApiMachineMetrics**](ApiMachineMetrics.md) |  | [optional] |
| **mounts** | [**Array&lt;ApiMachineMount&gt;**](ApiMachineMount.md) |  | [optional] |
| **processes** | [**Array&lt;ApiMachineProcess&gt;**](ApiMachineProcess.md) |  | [optional] |
| **restart** | [**ApiMachineRestart**](ApiMachineRestart.md) |  | [optional] |
| **schedule** | **String** | The following fields can only be set or updated by &#x60;fly machines run|update&#x60; commands \&quot;fly deploy\&quot; must preserve them, if you add anything here, ensure it is propagated on deploys | [optional] |
| **services** | [**Array&lt;ApiMachineService&gt;**](ApiMachineService.md) |  | [optional] |
| **size** | **String** | Deprecated: use Guest instead | [optional] |
| **standbys** | **Array&lt;String&gt;** | Standbys enable a machine to be a standby for another. In the event of a hardware failure, the standby machine will be started. | [optional] |
| **statics** | [**Array&lt;ApiStatic&gt;**](ApiStatic.md) |  | [optional] |
| **stop_config** | [**ApiStopConfig**](ApiStopConfig.md) |  | [optional] |

## Example

```ruby
require 'fly-api-ruby'

instance = FlyApi::ApiMachineConfig.new(
  auto_destroy: null,
  checks: null,
  disable_machine_autostart: null,
  dns: null,
  env: null,
  files: null,
  guest: null,
  host_dedication_id: null,
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

