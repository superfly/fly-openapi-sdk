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

from fly-sdk.models.process_stat import ProcessStat  # noqa: E501

class TestProcessStat(unittest.TestCase):
    """ProcessStat unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ProcessStat:
        """Test ProcessStat
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ProcessStat`
        """
        model = ProcessStat()  # noqa: E501
        if include_optional:
            return ProcessStat(
                command = '',
                cpu = 56,
                directory = '',
                listen_sockets = [
                    fly-sdk.models.listen_socket.ListenSocket(
                        address = '', 
                        proto = '', )
                    ],
                pid = 56,
                rss = 56,
                rtime = 56,
                stime = 56
            )
        else:
            return ProcessStat(
        )
        """

    def testProcessStat(self):
        """Test ProcessStat"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
