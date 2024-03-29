# coding: utf-8

"""
    Fly Machines API

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

    The version of the OpenAPI document: 1.0
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json


from typing import Any, ClassVar, Dict, List, Optional
from pydantic import BaseModel, StrictBool, StrictInt, StrictStr
from fly_sdk.models.api_http_options import ApiHTTPOptions
from fly_sdk.models.api_proxy_proto_options import ApiProxyProtoOptions
from fly_sdk.models.api_tls_options import ApiTLSOptions
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class ApiMachinePort(BaseModel):
    """
    ApiMachinePort
    """ # noqa: E501
    end_port: Optional[StrictInt] = None
    force_https: Optional[StrictBool] = None
    handlers: Optional[List[StrictStr]] = None
    http_options: Optional[ApiHTTPOptions] = None
    port: Optional[StrictInt] = None
    proxy_proto_options: Optional[ApiProxyProtoOptions] = None
    start_port: Optional[StrictInt] = None
    tls_options: Optional[ApiTLSOptions] = None
    __properties: ClassVar[List[str]] = ["end_port", "force_https", "handlers", "http_options", "port", "proxy_proto_options", "start_port", "tls_options"]

    model_config = {
        "populate_by_name": True,
        "validate_assignment": True
    }


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of ApiMachinePort from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        _dict = self.model_dump(
            by_alias=True,
            exclude={
            },
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of http_options
        if self.http_options:
            _dict['http_options'] = self.http_options.to_dict()
        # override the default output from pydantic by calling `to_dict()` of proxy_proto_options
        if self.proxy_proto_options:
            _dict['proxy_proto_options'] = self.proxy_proto_options.to_dict()
        # override the default output from pydantic by calling `to_dict()` of tls_options
        if self.tls_options:
            _dict['tls_options'] = self.tls_options.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of ApiMachinePort from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "end_port": obj.get("end_port"),
            "force_https": obj.get("force_https"),
            "handlers": obj.get("handlers"),
            "http_options": ApiHTTPOptions.from_dict(obj.get("http_options")) if obj.get("http_options") is not None else None,
            "port": obj.get("port"),
            "proxy_proto_options": ApiProxyProtoOptions.from_dict(obj.get("proxy_proto_options")) if obj.get("proxy_proto_options") is not None else None,
            "start_port": obj.get("start_port"),
            "tls_options": ApiTLSOptions.from_dict(obj.get("tls_options")) if obj.get("tls_options") is not None else None
        })
        return _obj


