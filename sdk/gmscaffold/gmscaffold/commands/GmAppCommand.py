# -*- encoding: utf-8 -*-
"""
@文件        :GmAppCommand.py
@说明        :
@时间        :2024/11/27 18:37:49
@作者        :Zack
@版本        :1.0
"""

from cleo.helpers import argument, option

from gmscaffold.base.command import BaseCommand, Command
from gmscaffold.entities.GmApp import GmAppData
from gmscaffold.utils import signal as _signal


class RemoveAppCommand(BaseCommand):
    """ """

    name = "gm_app remove_app"
    description = "Remove applications"
    arguments = [
        argument(
            "name",
            description="Application name?",
            optional=True,
        )
    ]
    options = [
        option(
            "remove_app",
            description="Remove gm app",
            flag=True,
        )
    ]

    def remove_app(self, gmapp_name: str = None):
        """Remove applications"""
        print("remove_app", gmapp_name)

    def handle(self):
        self.remove_app()


class CreateAppCommand(BaseCommand):
    """ """

    name = "gm_app create_app"
    description = "Create Applications"

    def create_app(self, *args, **kwargs):
        """Create gm applications"""
        gmapp_data = GmAppData()
        self._input_project_name(gmapp_data)._input_project_author(
            gmapp_data
        )._input_project_email(gmapp_data)._input_project_version(
            gmapp_data
        )._input_project_desc(
            gmapp_data
        )._input_project_date(
            gmapp_data
        )
        self.Gm._signals.send(signal=_signal.create_gmapp_project, data=gmapp_data)

    def handle(self):
        self.create_app()


class GmAppCommand(BaseCommand):
    """ """

    name = "gm_app"
    description = "GM applications operator"

    commands = [RemoveAppCommand, CreateAppCommand]
