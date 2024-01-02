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
from pydantic import Field
from fly_sdk.models.api_machine_config import ApiMachineConfig
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class UpdateMachineRequest(BaseModel):
    """
    UpdateMachineRequest
    """ # noqa: E501
    config: Optional[ApiMachineConfig] = None
    current_version: Optional[StrictStr] = None
    lease_ttl: Optional[StrictInt] = None
    lsvd: Optional[StrictBool] = None
    name: Optional[StrictStr] = Field(default=None, description="Unique name for this Machine. If omitted, one is generated for you")
    region: Optional[StrictStr] = Field(default=None, description="The target region. Omitting this param launches in the same region as your WireGuard peer connection (somewhere near you).")
    skip_launch: Optional[StrictBool] = None
    skip_service_registration: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["config", "current_version", "lease_ttl", "lsvd", "name", "region", "skip_launch", "skip_service_registration"]

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
        """Create an instance of UpdateMachineRequest from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of config
        if self.config:
            _dict['config'] = self.config.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of UpdateMachineRequest from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "config": ApiMachineConfig.from_dict(obj.get("config")) if obj.get("config") is not None else None,
            "current_version": obj.get("current_version"),
            "lease_ttl": obj.get("lease_ttl"),
            "lsvd": obj.get("lsvd"),
            "name": obj.get("name"),
            "region": obj.get("region"),
            "skip_launch": obj.get("skip_launch"),
            "skip_service_registration": obj.get("skip_service_registration")
        })
        return _obj


