# -*- encoding: utf-8 -*-
"""
@文件        :command.py
@说明        :
@时间        :2024/11/25 13:57:46
@作者        :Zack
@版本        :1.0
"""
import datetime
import subprocess
from abc import ABCMeta
from dataclasses import asdict
from typing import Any, Dict, Optional, TypeVar

from cleo.commands.command import Command

from gmscaffold.base.dataclass import BaseData, TemplateData
from gmscaffold.exceptions import NoDataError
from gmscaffold.utils import path as _path
from gmscaffold.utils import signal as _signal
from gmscaffold.utils.conf import get_module_name


class _ProjectObject:
    """ """

    name: Optional[str]
    author: Optional[str]
    email: Optional[str]
    description: Optional[str]
    version: Optional[str]
    date: Optional[str]


ProjectObject = TypeVar("ProjectObject", bound=_ProjectObject)


class BaseCommand(Command, metaclass=ABCMeta):
    """ """

    def __init__(self, Gm):
        """ """
        self.Gm = Gm
        super().__init__()

    def _input_project_name(self, project_object: ProjectObject):
        """ """
        project_name = input("请输入项目名称: ")
        project_object.name = project_name
        return self

    def _input_project_author(self, project_object: ProjectObject):
        """ """
        author_name = input("请输入作者名称: ")
        project_object.author = author_name
        return self

    def _input_project_email(self, project_object: ProjectObject):
        """ """
        email = input("请输入联系邮箱：")
        project_object.email = email
        return self

    def _input_project_desc(self, project_object: ProjectObject):
        """ """
        description = input("请输入项目描述信息：")
        project_object.description = description
        return self

    def _input_project_version(self, project_object: ProjectObject):
        """ """
        version = input("请输入项目版本号：")
        project_object.version = version
        return self

    def _input_project_date(self, project_object: ProjectObject):
        """ """
        project_object.date = datetime.datetime.now().strftime("%Y/%m/%d %H:%M:%S")
        return self

    def check_call(self, *args, **kwargs):
        """Execute the command"""
        return subprocess.check_call(*args, **kwargs)


class BaseBuilder(object):
    """ """

    def __init__(self, signals_callbacks):
        """ """
        self.register_signals(signals_callbacks)

    def check_data(self, kwargs: Dict[str, Any]):
        """ """
        if not kwargs.get("data"):
            raise NoDataError("No data to create spider project")

    def register_signals(self, signals_callbacks):
        """ """
        for signal, callback in signals_callbacks.items():
            self.Gm._signals.connect(callback, signal)

    def open_signals(self, signals_callbacks):
        """ """
        for signal in signals_callbacks:
            """ """
            self.Gm._signals.send(signal, data=_signal.OPEN_CONNECT)

    @classmethod
    def from_settings(cls, settings):
        """ """
        return cls(settings)
