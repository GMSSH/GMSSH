# -*- encoding: utf-8 -*-
import json
from typing import Any, Optional

from jsonrpcserver import Success


def res_success(
    code: int = 200, msg: str = None, data: Any = None, method: str = None
) -> str:
    """返回模板，成功的响应"""
    return _jsonify(code=code, msg=msg, data=data, method=method)


def res_failure(
    code: int = 400, msg: Any = None, data: Any = None, method: str = None
) -> str:
    """ """

    return _jsonify(code=code, msg=str(msg), data=data, method=method)


def raise_exception(
    except_type: Exception,
    msg: str,
    code: Optional[int] = 400,
    method: Optional[str] = None,
):
    """ """
    raise except_type(res_failure(msg=msg, code=code, method=method))


def _jsonify(code: int = 200, msg: str = None, data: Any = None, method: str = None):
    """ """

    res_data = {
        "code": code,
        "meta": {"endpoint": method, "close": 1},
        "data": data,
        "msg": msg,
    }

    return res_data


def jsonify(code: int = 200, msg: str = None, data: Any = None, method: str = None):
    """ """

    res_data = _jsonify(code=code, msg=str(msg), data=data, method=method)
    return Success(res_data)


class Response:
    """ """

    def __init__(self, payload: str):
        """ """
        self.payload = payload

    def raw(self):
        """ """
        return self.payload

    def json(self):
        """ """
        return self.payload

    def to_dict(self):
        """ """
        return (
            json.loads(self.payload) if isinstance(self.payload, str) else self.payload
        )

    def __str__(self):
        """ """
        return self.payload

    def __repr__(self):
        """ """
        return f"<Response payload={self.payload}>"
