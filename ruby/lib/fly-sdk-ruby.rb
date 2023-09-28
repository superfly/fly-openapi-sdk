=begin
#Fly Machines API

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 1.0

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.1.0-SNAPSHOT

=end

# Common files
require 'fly-sdk-ruby/api_client'
require 'fly-sdk-ruby/api_error'
require 'fly-sdk-ruby/version'
require 'fly-sdk-ruby/configuration'

# Models
require 'fly-sdk-ruby/models/api_dns_config'
require 'fly-sdk-ruby/models/api_file'
require 'fly-sdk-ruby/models/api_http_options'
require 'fly-sdk-ruby/models/api_http_response_options'
require 'fly-sdk-ruby/models/api_machine_check'
require 'fly-sdk-ruby/models/api_machine_config'
require 'fly-sdk-ruby/models/api_machine_guest'
require 'fly-sdk-ruby/models/api_machine_http_header'
require 'fly-sdk-ruby/models/api_machine_init'
require 'fly-sdk-ruby/models/api_machine_metrics'
require 'fly-sdk-ruby/models/api_machine_mount'
require 'fly-sdk-ruby/models/api_machine_port'
require 'fly-sdk-ruby/models/api_machine_process'
require 'fly-sdk-ruby/models/api_machine_restart'
require 'fly-sdk-ruby/models/api_machine_service'
require 'fly-sdk-ruby/models/api_machine_service_concurrency'
require 'fly-sdk-ruby/models/api_proxy_proto_options'
require 'fly-sdk-ruby/models/api_static'
require 'fly-sdk-ruby/models/api_stop_config'
require 'fly-sdk-ruby/models/api_tls_options'
require 'fly-sdk-ruby/models/app'
require 'fly-sdk-ruby/models/check_status'
require 'fly-sdk-ruby/models/create_app_request'
require 'fly-sdk-ruby/models/create_lease_request'
require 'fly-sdk-ruby/models/create_machine_request'
require 'fly-sdk-ruby/models/create_volume_request'
require 'fly-sdk-ruby/models/error_response'
require 'fly-sdk-ruby/models/extend_volume_request'
require 'fly-sdk-ruby/models/extend_volume_response'
require 'fly-sdk-ruby/models/image_ref'
require 'fly-sdk-ruby/models/lease'
require 'fly-sdk-ruby/models/list_app'
require 'fly-sdk-ruby/models/list_apps_response'
require 'fly-sdk-ruby/models/listen_socket'
require 'fly-sdk-ruby/models/machine'
require 'fly-sdk-ruby/models/machine_event'
require 'fly-sdk-ruby/models/machine_exec_request'
require 'fly-sdk-ruby/models/machine_version'
require 'fly-sdk-ruby/models/main_status_code'
require 'fly-sdk-ruby/models/organization'
require 'fly-sdk-ruby/models/process_stat'
require 'fly-sdk-ruby/models/signal_request'
require 'fly-sdk-ruby/models/stop_request'
require 'fly-sdk-ruby/models/update_machine_request'
require 'fly-sdk-ruby/models/volume'
require 'fly-sdk-ruby/models/volume_snapshot'

# APIs
require 'fly-sdk-ruby/api/apps_api'
require 'fly-sdk-ruby/api/machines_api'
require 'fly-sdk-ruby/api/volumes_api'

module FlySDK
  class << self
    # Customize default settings for the SDK using block.
    #   FlySDK.configure do |config|
    #     config.username = "xxx"
    #     config.password = "xxx"
    #   end
    # If no block given, return the default Configuration object.
    def configure
      if block_given?
        yield(Configuration.default)
      else
        Configuration.default
      end
    end
  end
end