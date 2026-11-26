# -*- encoding: utf-8 -*-
from abc import ABCMeta, abstractmethod
from typing import Any, Dict, List, Optional, Union

from wtforms import Field, Form


class BaseValidator:
    """ """

    def __init__(self, err_message=None):
        self.err_message = err_message

    def __call__(self, form: Form, field: Field):
        """ """
        return self.validator(form, field)

    @abstractmethod
    def validator(self, form, field):
        """ """


class BaseServer(metaclass=ABCMeta):
    """ """

    @abstractmethod
    async def run(self): ...


class RPCMiddleware:
    """中间件基类"""

    def __init__(self, app=None):
        self.app = app

    @abstractmethod
    def process_request(self, request: str, context: Dict) -> Dict:
        """处理请求"""

    @abstractmethod
    def process_response(self, response: Union[Any, str], context: Dict) -> Any:
        """处理请求"""


class ClientTransport:
    """ """

    @abstractmethod
    async def send_message(self, message: Union[Dict[str, Any], str]): ...

    @abstractmethod
    async def close(self): ...

    @abstractmethod
    async def connect(self): ...
