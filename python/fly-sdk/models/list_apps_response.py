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


from typing import List, Optional
from pydantic import BaseModel, StrictInt
from fly-sdk.models.list_app import ListApp

class ListAppsResponse(BaseModel):
    """
    ListAppsResponse
    """
    apps: Optional[List[ListApp]] = None
    total_apps: Optional[StrictInt] = None
    __properties = ["apps", "total_apps"]

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
    def from_json(cls, json_str: str) -> ListAppsResponse:
        """Create an instance of ListAppsResponse from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of each item in apps (list)
        _items = []
        if self.apps:
            for _item in self.apps:
                if _item:
                    _items.append(_item.to_dict())
            _dict['apps'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> ListAppsResponse:
        """Create an instance of ListAppsResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return ListAppsResponse.parse_obj(obj)

        _obj = ListAppsResponse.parse_obj({
            "apps": [ListApp.from_dict(_item) for _item in obj.get("apps")] if obj.get("apps") is not None else None,
            "total_apps": obj.get("total_apps")
        })
        return _obj

