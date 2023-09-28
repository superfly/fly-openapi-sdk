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

from fly-sdk.models.api_machine_init import ApiMachineInit  # noqa: E501

class TestApiMachineInit(unittest.TestCase):
    """ApiMachineInit unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ApiMachineInit:
        """Test ApiMachineInit
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ApiMachineInit`
        """
        model = ApiMachineInit()  # noqa: E501
        if include_optional:
            return ApiMachineInit(
                cmd = [
                    ''
                    ],
                entrypoint = [
                    ''
                    ],
                var_exec = [
                    ''
                    ],
                kernel_args = [
                    ''
                    ],
                swap_size_mb = 56,
                tty = True
            )
        else:
            return ApiMachineInit(
        )
        """

    def testApiMachineInit(self):
        """Test ApiMachineInit"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()