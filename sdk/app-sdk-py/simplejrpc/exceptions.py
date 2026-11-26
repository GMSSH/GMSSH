# -*- encoding: utf-8 -*-
import http
import json
from typing import Optional

from simplejrpc.response import _jsonify


class RPCException(Exception):
    """基础RPC异常"""

    def __init__(
        self,
        message=http.HTTPStatus.BAD_REQUEST.description,
        code=http.HTTPStatus.BAD_REQUEST.value,
        data: Optional[str] = None,
        *args,
        **kwargs,
    ):
        super().__init__(*args, **kwargs)
        self.message = message
        self.code = code
        self.data = data

    def __str__(self):
        """ """
        from simplejrpc._text import TextMessageDecoder

        data = _jsonify(code=self.code, data=self.data, msg=self.message)
        return json.dumps(data, cls=TextMessageDecoder)


class UnauthorizedError(RPCException):
    """未授权异常"""


class ValidationError(RPCException):
    """验证异常"""


class FileNotFoundError(RPCException):
    """ """


class ValueError(RPCException):
    """ """


class RuntimeError(RPCException):
    """ """


class AttributeError(RPCException):
    """ """


class TypeError(RPCException):
    """ """
