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
from fly_sdk.models.api_machine_check import ApiMachineCheck
from fly_sdk.models.api_machine_port import ApiMachinePort
from fly_sdk.models.api_machine_service_concurrency import ApiMachineServiceConcurrency
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class ApiMachineService(BaseModel):
    """
    ApiMachineService
    """ # noqa: E501
    autostart: Optional[StrictBool] = None
    autostop: Optional[StrictBool] = None
    checks: Optional[List[ApiMachineCheck]] = None
    concurrency: Optional[ApiMachineServiceConcurrency] = None
    force_instance_description: Optional[StrictStr] = None
    force_instance_key: Optional[StrictStr] = None
    internal_port: Optional[StrictInt] = None
    min_machines_running: Optional[StrictInt] = None
    ports: Optional[List[ApiMachinePort]] = None
    protocol: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["autostart", "autostop", "checks", "concurrency", "force_instance_description", "force_instance_key", "internal_port", "min_machines_running", "ports", "protocol"]

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
        """Create an instance of ApiMachineService from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in checks (list)
        _items = []
        if self.checks:
            for _item in self.checks:
                if _item:
                    _items.append(_item.to_dict())
            _dict['checks'] = _items
        # override the default output from pydantic by calling `to_dict()` of concurrency
        if self.concurrency:
            _dict['concurrency'] = self.concurrency.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in ports (list)
        _items = []
        if self.ports:
            for _item in self.ports:
                if _item:
                    _items.append(_item.to_dict())
            _dict['ports'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of ApiMachineService from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "autostart": obj.get("autostart"),
            "autostop": obj.get("autostop"),
            "checks": [ApiMachineCheck.from_dict(_item) for _item in obj.get("checks")] if obj.get("checks") is not None else None,
            "concurrency": ApiMachineServiceConcurrency.from_dict(obj.get("concurrency")) if obj.get("concurrency") is not None else None,
            "force_instance_description": obj.get("force_instance_description"),
            "force_instance_key": obj.get("force_instance_key"),
            "internal_port": obj.get("internal_port"),
            "min_machines_running": obj.get("min_machines_running"),
            "ports": [ApiMachinePort.from_dict(_item) for _item in obj.get("ports")] if obj.get("ports") is not None else None,
            "protocol": obj.get("protocol")
        })
        return _obj


