# -*- encoding: utf-8 -*-
"""
@文件        :decorators.py
@说明        :
@时间        :2024/11/28 11:57:30
@作者        :Zack
@版本        :1.0
"""
import sys
from functools import partial, wraps


def hide_except(fn=None, exc_type: BaseException = None):
    """ """
    if fn is None:
        return partial(hide_except, exc_type=exc_type)

    @wraps(fn)
    def inner(*args, **kwargs):
        """ """
        try:
            return fn(*args, **kwargs)
        except exc_type as e:
            print("\n [gm] 退出...")
            sys.exit()

    return inner
