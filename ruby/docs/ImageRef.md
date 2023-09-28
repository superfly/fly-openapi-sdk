# FlyApi::ImageRef

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **digest** | **String** |  | [optional] |
| **labels** | **Hash&lt;String, String&gt;** |  | [optional] |
| **registry** | **String** |  | [optional] |
| **repository** | **String** |  | [optional] |
| **tag** | **String** |  | [optional] |

## Example

```ruby
require 'fly-api-ruby'

instance = FlyApi::ImageRef.new(
  digest: null,
  labels: null,
  registry: null,
  repository: null,
  tag: null
)
```

