# -*- encoding: utf-8 -*-
"""
@文件        :cmdline.py
@说明        :
@时间        :2024/11/25 10:54:59
@作者        :Zack
@版本        :1.0
"""
from __future__ import annotations

import os
import sys
from importlib import import_module
from typing import TYPE_CHECKING

import pyfiglet
from cleo.application import Application
from cleo.exceptions import CleoLogicError
from cleo.loaders.factory_command_loader import FactoryCommandLoader

sys.path.insert(0, os.path.dirname(os.path.dirname(__file__)))

from gmscaffold.signalmanager import SignalManager
from gmscaffold.utils.conf import build_component_list

application = Application()


if TYPE_CHECKING:
    from collections.abc import Callable

    from cleo.commands.command import Command


class CommandLoader(FactoryCommandLoader):
    def register_factory(
        self, command_name: str, factory: Callable[[], Command]
    ) -> None:
        if command_name in self._factories:
            raise CleoLogicError(f'The command "{command_name}" already exists.')

        self._factories[command_name] = factory


def load_command(name: str, self) -> Callable[[], Command]:
    def _load() -> Command:
        package, cls = COMMANDS[name].rsplit(".", maxsplit=1)
        module = import_module(package)
        command_class = getattr(module, cls)
        command: Command = command_class(self)
        return command

    return _load


COMMANDS = {
    # gm_app commands
    "gm_app create_app": "gmscaffold.commands.GmAppCommand.CreateAppCommand",
    # "gm_app remove_app": "gmscaffold.commands.GmAppCommand.RemoveAppCommand",
    # spider commands
    "spider create_spider": "gmscaffold.commands.SpiderCommand.CreateSpiderCommand",
    # "spider remove_spider": "gmscaffold.commands.SpiderCommand.RemoveSpiderCommand",
    # web commands
    "web start_project": "gmscaffold.commands.WebCommand.CreateWebCommand",
    "web remove_project": "gmscaffold.commands.WebCommand.RemoveWebCommand",
    "web start_app": "gmscaffold.commands.WebCommand.StartAppWebCommand",
    "web remove_app": "gmscaffold.commands.WebCommand.RemoveAppWebCommand",
}


class GmScaffoldMain(Application):
    """ """

    def __init__(self):
        """ """
        self._signals = SignalManager()
        build_component_list(self)
        super().__init__()
        command_loader = CommandLoader(
            {name: load_command(name, self) for name in COMMANDS}
        )
        self.set_command_loader(command_loader)

    def _from_settings(self, settings):
        """ """
        self._settings = settings

    @property
    def command_loader(self) -> CommandLoader:
        command_loader = self._command_loader
        assert isinstance(command_loader, CommandLoader)
        return command_loader


def execute(argv=None, settings=None):
    """命令行交互处理"""
    welcome_txt = pyfiglet.figlet_format("GMSSH", font="lean")
    print(welcome_txt)
    GmScaffoldMain().run()


if __name__ == "__main__":

    execute()
