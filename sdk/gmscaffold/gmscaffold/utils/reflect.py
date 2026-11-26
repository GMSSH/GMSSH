# -*- encoding: utf-8 -*-
"""
@文件        :reflect.py
@说明        :
@时间        :2024/11/25 11:13:26
@作者        :Zack
@版本        :1.0
"""

from typing import Callable, Type, TypeVar

_inlineCallbacksExtraneous = []


_T_Callable = TypeVar("_T_Callable", bound=Callable[..., object])


def qual(clazz: Type[object]) -> str:
    """
    Return full import path of a class.
    """
    return clazz.__module__ + "." + clazz.__name__


def _extraneous(f: _T_Callable) -> _T_Callable:
    """
    Mark the given callable as extraneous to inlineCallbacks exception
    reporting; don't show these functions.

    @param f: a function that you NEVER WANT TO SEE AGAIN in ANY TRACEBACK
        reported by Failure.

    @type f: function

    @return: f
    """
    _inlineCallbacksExtraneous.append(f.__code__)
    return f


def declspec_check(inst, declspec: str):
    """Calling specification, print the calling specifications supported by the current object

    :param inst: Reflection Objects
    :param declspec: Reflection Method
    """
    inst_declspec = dir(inst)
    dir_declspec_atts = "\n\t".join(filter(lambda x: not x.startswith("_"), inst_declspec))

    # TODO: 使用logger
    print(f"ERROR : Could not consume arg: {declspec}")
    print("Usage <command>\n    available commands:")
    print("\t{:>10}".format(dir_declspec_atts))


def declspec_revoke(inst, declspec: str, *args, **kwargs):
    """Call

    :param inst: Reflection Objects
    :param declspec: Reflection Method
    :param args: Reflecting method parameters
    :param kwargs: Reflecting method parameters
    """
    if hasattr(inst, declspec):
        method = getattr(inst, declspec)
        if callable(method):
            method(*args, **kwargs)
    else:
        declspec_check(inst, declspec)
