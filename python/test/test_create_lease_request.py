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

from fly-sdk.models.create_lease_request import CreateLeaseRequest  # noqa: E501

class TestCreateLeaseRequest(unittest.TestCase):
    """CreateLeaseRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> CreateLeaseRequest:
        """Test CreateLeaseRequest
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CreateLeaseRequest`
        """
        model = CreateLeaseRequest()  # noqa: E501
        if include_optional:
            return CreateLeaseRequest(
                description = '',
                ttl = 56
            )
        else:
            return CreateLeaseRequest(
        )
        """

    def testCreateLeaseRequest(self):
        """Test CreateLeaseRequest"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
