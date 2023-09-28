#!/bin/sh

set -e

rm -rf go ruby

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i https://docs.machines.dev/swagger/doc.json  -g ruby -o /local/ruby --additional-properties gemName=fly-api-ruby --additional-properties moduleName=FlyApi

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i https://docs.machines.dev/swagger/doc.json  -g go -o /local/go --additional-properties packageName=flyapi
