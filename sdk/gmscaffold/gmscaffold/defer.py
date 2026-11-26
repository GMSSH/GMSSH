# -*- encoding: utf-8 -*-
"""
@文件        :defer.py
@说明        :
@时间        :2024/11/25 11:08:37
@作者        :Zack
@版本        :1.0
"""
from __future__ import annotations

from typing import Awaitable, TypeVar

_SelfResultT = TypeVar("_SelfResultT")


class Deferred(Awaitable[_SelfResultT]):
    """ """
