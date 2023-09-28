#!/bin/sh

set -e

rm -rf go ruby python

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i https://docs.machines.dev/swagger/doc.json  -g ruby -o /local/ruby --additional-properties gemName=fly-sdk-ruby --additional-properties moduleName=FlySDK

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i https://docs.machines.dev/swagger/doc.json  -g go -o /local/go --additional-properties packageName=fly-sdk

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i https://docs.machines.dev/swagger/doc.json  -g python -o /local/python --additional-properties packageName=fly-sdk --additional-properties projectName=fly-sdk
