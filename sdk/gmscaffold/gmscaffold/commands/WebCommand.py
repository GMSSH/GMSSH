# -*- encoding: utf-8 -*-
"""
@文件        :WebCommand.py
@说明        :
@时间        :2024/11/27 18:44:40
@作者        :Zack
@版本        :1.0
"""

from typing import Any, Dict, Optional, TypeVar

from gmscaffold.base.command import BaseCommand
from gmscaffold.entities.web import WebData
from gmscaffold.utils import signal as _signal

ProjectObject = TypeVar("ProjectObject", bound=WebData)


class AppCommand(object):
    """ """

    def _input_app_name(self, project_object: ProjectObject):
        """ """
        project_name = input("请输入应用名称: ")
        project_object.name = project_name
        return project_object


class CreateWebCommand(BaseCommand):
    """ """

    name = "web start_project"
    description = "Creating a traditional web service based on the sanic framework"

    def start_project(self, *args, **kwargs):
        """Create sanic web server project"""
        web_data = WebData()
        self._input_project_name(web_data)._input_project_author(web_data)._input_project_email(
            web_data
        )._input_project_version(web_data)._input_project_desc(web_data)._input_project_date(web_data)
        self.Gm._signals.send(signal=_signal.create_web_project, data=web_data)

    def handle(self):
        self.start_project()


class RemoveWebCommand(BaseCommand):
    """ """

    name = "web remove_project"
    description = "Remove the traditional web service based on scaffolding to create sanic framework"

    def remove_project(self, *args, **kwargs):
        """ """
        web_data = WebData()
        # TODO: input remove project name
        self.Gm._signals.send(signal=_signal.remove_web_project)

    def handle(self):
        self.remove_project()


class StartAppWebCommand(BaseCommand):
    """ """

    name = "web start_app"
    description = "Create a web service internal sub-application"

    def start_app(self, *args, **kwargs):
        """Create sanic web app"""
        web_data = WebData()
        AppCommand()._input_app_name(web_data)
        self.Gm._signals.send(signal=_signal.create_web_app, data=web_data)

    def handle(self):
        self.start_app()


class RemoveAppWebCommand(BaseCommand):
    """ """

    name = "web remove_app"
    description = "Remove the creation of internal sub-applications of web services based on scaffolding"

    def remove_app(self, *args, **kwargs):
        """Remove sanic web app"""
        web_data = WebData()
        # TODO: input remove app name
        self.Gm._signals.send(signal=_signal.remove_web_app)

    def handle(self):
        self.remove_app()


class WebCommand(BaseCommand):
    """ """

    name = "web"
    description = "Web service scaffolding based on Sanic"
    commands = [
        CreateWebCommand,
        RemoveWebCommand,
        StartAppWebCommand,
        RemoveAppWebCommand,
    ]
