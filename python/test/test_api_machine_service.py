# coding: utf-8

"""
    Fly Machines API

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: 1.0
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest
import datetime

from fly-sdk.models.api_machine_service import ApiMachineService

class TestApiMachineService(unittest.TestCase):
    """ApiMachineService unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ApiMachineService:
        """Test ApiMachineService
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ApiMachineService`
        """
        model = ApiMachineService()
        if include_optional:
            return ApiMachineService(
                autostart = True,
                autostop = True,
                checks = [
                    fly-sdk.models.api/machine_check.api.MachineCheck(
                        grace_period = '', 
                        headers = [
                            fly-sdk.models.api/machine_http_header.api.MachineHTTPHeader(
                                name = '', 
                                values = [
                                    ''
                                    ], )
                            ], 
                        interval = '', 
                        method = '', 
                        path = '', 
                        port = 56, 
                        protocol = '', 
                        timeout = '', 
                        tls_server_name = '', 
                        tls_skip_verify = True, 
                        type = '', )
                    ],
                concurrency = fly-sdk.models.api/machine_service_concurrency.api.MachineServiceConcurrency(
                    hard_limit = 56, 
                    soft_limit = 56, 
                    type = '', ),
                force_instance_description = '',
                force_instance_key = '',
                internal_port = 56,
                min_machines_running = 56,
                ports = [
                    fly-sdk.models.api/machine_port.api.MachinePort(
                        end_port = 56, 
                        force_https = True, 
                        handlers = [
                            ''
                            ], 
                        http_options = fly-sdk.models.api/http_options.api.HTTPOptions(
                            compress = True, 
                            response = fly-sdk.models.api/http_response_options.api.HTTPResponseOptions(
                                headers = {
                                    'key' : None
                                    }, ), ), 
                        port = 56, 
                        proxy_proto_options = fly-sdk.models.api/proxy_proto_options.api.ProxyProtoOptions(
                            version = '', ), 
                        start_port = 56, 
                        tls_options = fly-sdk.models.api/tls_options.api.TLSOptions(
                            alpn = [
                                ''
                                ], 
                            default_self_signed = True, 
                            versions = [
                                ''
                                ], ), )
                    ],
                protocol = ''
            )
        else:
            return ApiMachineService(
        )
        """

    def testApiMachineService(self):
        """Test ApiMachineService"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
