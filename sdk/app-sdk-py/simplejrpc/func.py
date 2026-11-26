# -*- encoding: utf-8 -*-
import inspect
from collections import OrderedDict
from typing import Any, Dict, List

from simplejrpc.exceptions import TypeError


def make_signature(fields: List[Any]):
    """
    The function `make_signature` creates a signature object for a function based on a list of field
    names.

    :param fields: A list of any type of objects
    :type fields: List[Any]
    :return: an instance of the `inspect.Signature` class.
    """
    """ """
    params = []
    for name, required in fields:
        if required:
            params.append(inspect.Parameter(name, inspect.Parameter.KEYWORD_ONLY))
        else:
            params.append(
                inspect.Parameter(name, inspect.Parameter.KEYWORD_ONLY, default=None)
            )
    return inspect.Signature(params)


def str2int(value: str | int, name: str):
    """ """
    if isinstance(value, int):
        return value
    if not value.isdigit():
        raise TypeError(f"Field {name}, expected integer")

    return int(value)


def order_dict(data: Dict[str, Any], field: str = "lang") -> OrderedDict:
    """Extract the key value pairs corresponding to the fields and place them at the index position at the beginning of the ordered dictionary
    1. Extract the key value pairs corresponding to the fields
    2. Place the extracted key value pairs at the beginning of the ordered dictionary
    3. Place the remaining key value pairs of the ordered dictionary after the ordered dictionary
    4. Return an ordered dictionary

    :param data: Sort dict data
    :param field: Sort dict data by field
    :return: OrderedDict
    """

    if field not in data:
        return OrderedDict(data)
    ordered_dict = OrderedDict()
    ordered_dict[field] = data[field]
    for key, value in data.items():
        if key != field:
            ordered_dict[key] = value
    return ordered_dict
