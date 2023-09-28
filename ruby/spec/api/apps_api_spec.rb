=begin
#Fly Machines API

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 1.0

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.1.0-SNAPSHOT

=end

require 'spec_helper'
require 'json'

# Unit tests for FlySDK::AppsApi
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'AppsApi' do
  before do
    # run before each test
    @api_instance = FlySDK::AppsApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of AppsApi' do
    it 'should create an instance of AppsApi' do
      expect(@api_instance).to be_instance_of(FlySDK::AppsApi)
    end
  end

  # unit tests for apps_create
  # @param request App body
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'apps_create test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for apps_delete
  # @param app_name Fly App Name
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'apps_delete test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for apps_list
  # @param [Hash] opts the optional parameters
  # @option opts [String] :org_slug The org slug, or &#39;personal&#39;, to filter apps
  # @return [ListAppsResponse]
  describe 'apps_list test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for apps_show
  # @param app_name Fly App Name
  # @param [Hash] opts the optional parameters
  # @return [App]
  describe 'apps_show test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

end