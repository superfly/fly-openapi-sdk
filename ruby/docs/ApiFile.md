# FlySDK::ApiFile

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **guest_path** | **String** | GuestPath is the path on the machine where the file will be written and must be an absolute path. For example: /full/path/to/file.json | [optional] |
| **raw_value** | **String** | The base64 encoded string of the file contents. | [optional] |
| **secret_name** | **String** | The name of the secret that contains the base64 encoded file contents. | [optional] |

## Example

```ruby
require 'fly-sdk-ruby'

instance = FlySDK::ApiFile.new(
  guest_path: null,
  raw_value: null,
  secret_name: null
)
```

