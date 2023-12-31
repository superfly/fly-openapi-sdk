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


from typing import Optional
from pydantic import BaseModel, StrictBool
from fly-sdk.models.api_http_response_options import ApiHTTPResponseOptions

class ApiHTTPOptions(BaseModel):
    """
    ApiHTTPOptions
    """
    compress: Optional[StrictBool] = None
    h2_backend: Optional[StrictBool] = None
    response: Optional[ApiHTTPResponseOptions] = None
    __properties = ["compress", "h2_backend", "response"]

    class Config:
        """Pydantic configuration"""
        allow_population_by_field_name = True
        validate_assignment = True

    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.dict(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> ApiHTTPOptions:
        """Create an instance of ApiHTTPOptions from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of response
        if self.response:
            _dict['response'] = self.response.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> ApiHTTPOptions:
        """Create an instance of ApiHTTPOptions from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return ApiHTTPOptions.parse_obj(obj)

        _obj = ApiHTTPOptions.parse_obj({
            "compress": obj.get("compress"),
            "h2_backend": obj.get("h2_backend"),
            "response": ApiHTTPResponseOptions.from_dict(obj.get("response")) if obj.get("response") is not None else None
        })
        return _obj


