=begin
#Fly Machines API

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 1.0

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.1.0-SNAPSHOT

=end

require 'date'
require 'time'

module FlyApi
  class MainStatusCode
    unknown = "unknown".freeze
    capacityErr = "insufficient_capacity".freeze

    def self.all_vars
      @all_vars ||= [unknown, capacityErr].freeze
    end

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def self.build_from_hash(value)
      new.build_from_hash(value)
    end

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def build_from_hash(value)
      return value if MainStatusCode.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #MainStatusCode"
    end
  end
end
