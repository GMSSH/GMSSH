# -*- encoding: utf-8 -*-
"""
@文件        :exceptions.py
@说明        :
@时间        :2024/11/25 11:55:00
@作者        :Zack
@版本        :1.0
"""


class BaseGmException(Exception):
    """ """


class NoDataError(BaseGmException):
    """ """


class BaseCommandException(BaseGmException):
    """ """


class BuildProjectCommandError(BaseCommandException):
    """ """


class BuildValueError(BaseGmException):
    """ """
