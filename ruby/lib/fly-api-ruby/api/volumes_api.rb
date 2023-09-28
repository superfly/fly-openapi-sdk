=begin
#Fly Machines API

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 1.0

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.1.0-SNAPSHOT

=end

require 'cgi'

module FlyApi
  class VolumesApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # @param app_name [String] Fly App Name
    # @param volume_id [String] Volume ID
    # @param [Hash] opts the optional parameters
    # @return [Volume]
    def volume_delete(app_name, volume_id, opts = {})
      data, _status_code, _headers = volume_delete_with_http_info(app_name, volume_id, opts)
      data
    end

    # @param app_name [String] Fly App Name
    # @param volume_id [String] Volume ID
    # @param [Hash] opts the optional parameters
    # @return [Array<(Volume, Integer, Hash)>] Volume data, response status code and response headers
    def volume_delete_with_http_info(app_name, volume_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: VolumesApi.volume_delete ...'
      end
      # verify the required parameter 'app_name' is set
      if @api_client.config.client_side_validation && app_name.nil?
        fail ArgumentError, "Missing the required parameter 'app_name' when calling VolumesApi.volume_delete"
      end
      # verify the required parameter 'volume_id' is set
      if @api_client.config.client_side_validation && volume_id.nil?
        fail ArgumentError, "Missing the required parameter 'volume_id' when calling VolumesApi.volume_delete"
      end
      # resource path
      local_var_path = '/apps/{app_name}/volumes/{volume_id}'.sub('{' + 'app_name' + '}', CGI.escape(app_name.to_s)).sub('{' + 'volume_id' + '}', CGI.escape(volume_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Volume'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"VolumesApi.volume_delete",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:DELETE, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: VolumesApi#volume_delete\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param app_name [String] Fly App Name
    # @param request [CreateVolumeRequest] Request body
    # @param [Hash] opts the optional parameters
    # @return [Volume]
    def volumes_create(app_name, request, opts = {})
      data, _status_code, _headers = volumes_create_with_http_info(app_name, request, opts)
      data
    end

    # @param app_name [String] Fly App Name
    # @param request [CreateVolumeRequest] Request body
    # @param [Hash] opts the optional parameters
    # @return [Array<(Volume, Integer, Hash)>] Volume data, response status code and response headers
    def volumes_create_with_http_info(app_name, request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: VolumesApi.volumes_create ...'
      end
      # verify the required parameter 'app_name' is set
      if @api_client.config.client_side_validation && app_name.nil?
        fail ArgumentError, "Missing the required parameter 'app_name' when calling VolumesApi.volumes_create"
      end
      # verify the required parameter 'request' is set
      if @api_client.config.client_side_validation && request.nil?
        fail ArgumentError, "Missing the required parameter 'request' when calling VolumesApi.volumes_create"
      end
      # resource path
      local_var_path = '/apps/{app_name}/volumes'.sub('{' + 'app_name' + '}', CGI.escape(app_name.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(request)

      # return_type
      return_type = opts[:debug_return_type] || 'Volume'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"VolumesApi.volumes_create",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: VolumesApi#volumes_create\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param app_name [String] Fly App Name
    # @param volume_id [String] Volume ID
    # @param request [ExtendVolumeRequest] Request body
    # @param [Hash] opts the optional parameters
    # @return [ExtendVolumeResponse]
    def volumes_extend(app_name, volume_id, request, opts = {})
      data, _status_code, _headers = volumes_extend_with_http_info(app_name, volume_id, request, opts)
      data
    end

    # @param app_name [String] Fly App Name
    # @param volume_id [String] Volume ID
    # @param request [ExtendVolumeRequest] Request body
    # @param [Hash] opts the optional parameters
    # @return [Array<(ExtendVolumeResponse, Integer, Hash)>] ExtendVolumeResponse data, response status code and response headers
    def volumes_extend_with_http_info(app_name, volume_id, request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: VolumesApi.volumes_extend ...'
      end
      # verify the required parameter 'app_name' is set
      if @api_client.config.client_side_validation && app_name.nil?
        fail ArgumentError, "Missing the required parameter 'app_name' when calling VolumesApi.volumes_extend"
      end
      # verify the required parameter 'volume_id' is set
      if @api_client.config.client_side_validation && volume_id.nil?
        fail ArgumentError, "Missing the required parameter 'volume_id' when calling VolumesApi.volumes_extend"
      end
      # verify the required parameter 'request' is set
      if @api_client.config.client_side_validation && request.nil?
        fail ArgumentError, "Missing the required parameter 'request' when calling VolumesApi.volumes_extend"
      end
      # resource path
      local_var_path = '/apps/{app_name}/volumes/{volume_id}/extend'.sub('{' + 'app_name' + '}', CGI.escape(app_name.to_s)).sub('{' + 'volume_id' + '}', CGI.escape(volume_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(request)

      # return_type
      return_type = opts[:debug_return_type] || 'ExtendVolumeResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"VolumesApi.volumes_extend",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PUT, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: VolumesApi#volumes_extend\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param app_name [String] Fly App Name
    # @param volume_id [String] Volume ID
    # @param [Hash] opts the optional parameters
    # @return [Volume]
    def volumes_get_by_id(app_name, volume_id, opts = {})
      data, _status_code, _headers = volumes_get_by_id_with_http_info(app_name, volume_id, opts)
      data
    end

    # @param app_name [String] Fly App Name
    # @param volume_id [String] Volume ID
    # @param [Hash] opts the optional parameters
    # @return [Array<(Volume, Integer, Hash)>] Volume data, response status code and response headers
    def volumes_get_by_id_with_http_info(app_name, volume_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: VolumesApi.volumes_get_by_id ...'
      end
      # verify the required parameter 'app_name' is set
      if @api_client.config.client_side_validation && app_name.nil?
        fail ArgumentError, "Missing the required parameter 'app_name' when calling VolumesApi.volumes_get_by_id"
      end
      # verify the required parameter 'volume_id' is set
      if @api_client.config.client_side_validation && volume_id.nil?
        fail ArgumentError, "Missing the required parameter 'volume_id' when calling VolumesApi.volumes_get_by_id"
      end
      # resource path
      local_var_path = '/apps/{app_name}/volumes/{volume_id}'.sub('{' + 'app_name' + '}', CGI.escape(app_name.to_s)).sub('{' + 'volume_id' + '}', CGI.escape(volume_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Volume'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"VolumesApi.volumes_get_by_id",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: VolumesApi#volumes_get_by_id\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param app_name [String] Fly App Name
    # @param volume_id [String] Volume ID
    # @param [Hash] opts the optional parameters
    # @return [Array<VolumeSnapshot>]
    def volumes_get_snapshots(app_name, volume_id, opts = {})
      data, _status_code, _headers = volumes_get_snapshots_with_http_info(app_name, volume_id, opts)
      data
    end

    # @param app_name [String] Fly App Name
    # @param volume_id [String] Volume ID
    # @param [Hash] opts the optional parameters
    # @return [Array<(Array<VolumeSnapshot>, Integer, Hash)>] Array<VolumeSnapshot> data, response status code and response headers
    def volumes_get_snapshots_with_http_info(app_name, volume_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: VolumesApi.volumes_get_snapshots ...'
      end
      # verify the required parameter 'app_name' is set
      if @api_client.config.client_side_validation && app_name.nil?
        fail ArgumentError, "Missing the required parameter 'app_name' when calling VolumesApi.volumes_get_snapshots"
      end
      # verify the required parameter 'volume_id' is set
      if @api_client.config.client_side_validation && volume_id.nil?
        fail ArgumentError, "Missing the required parameter 'volume_id' when calling VolumesApi.volumes_get_snapshots"
      end
      # resource path
      local_var_path = '/apps/{app_name}/volumes/{volume_id}/snapshots'.sub('{' + 'app_name' + '}', CGI.escape(app_name.to_s)).sub('{' + 'volume_id' + '}', CGI.escape(volume_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Array<VolumeSnapshot>'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"VolumesApi.volumes_get_snapshots",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: VolumesApi#volumes_get_snapshots\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # @param app_name [String] Fly App Name
    # @param [Hash] opts the optional parameters
    # @return [Array<Volume>]
    def volumes_list(app_name, opts = {})
      data, _status_code, _headers = volumes_list_with_http_info(app_name, opts)
      data
    end

    # @param app_name [String] Fly App Name
    # @param [Hash] opts the optional parameters
    # @return [Array<(Array<Volume>, Integer, Hash)>] Array<Volume> data, response status code and response headers
    def volumes_list_with_http_info(app_name, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: VolumesApi.volumes_list ...'
      end
      # verify the required parameter 'app_name' is set
      if @api_client.config.client_side_validation && app_name.nil?
        fail ArgumentError, "Missing the required parameter 'app_name' when calling VolumesApi.volumes_list"
      end
      # resource path
      local_var_path = '/apps/{app_name}/volumes'.sub('{' + 'app_name' + '}', CGI.escape(app_name.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Array<Volume>'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"VolumesApi.volumes_list",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: VolumesApi#volumes_list\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
