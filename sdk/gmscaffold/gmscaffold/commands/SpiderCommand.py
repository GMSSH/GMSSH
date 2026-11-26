# -*- encoding: utf-8 -*-
"""
@文件        :SpiderCommands.py
@说明        :
@时间        :2024/11/25 19:51:23
@作者        :Zack
@版本        :1.0
"""

from gmscaffold.base.command import BaseCommand
from gmscaffold.entities.spider import SpiderData
from gmscaffold.utils import signal as _signal


class CreateSpiderCommand(BaseCommand):
    """ """

    name = "spider create_spider"
    description = "Crawler script creation"

    def create_spider(self, *args, **kwargs):
        """Create spider scripts"""
        spider_data = SpiderData()
        self._input_project_name(spider_data)._input_project_author(
            spider_data
        )._input_project_email(spider_data)._input_project_version(
            spider_data
        )._input_project_desc(
            spider_data
        )._input_project_date(
            spider_data
        )
        self.Gm._signals.send(signal=_signal.create_spider_project, data=spider_data)

    def handle(self):
        self.create_spider()


class RemoveSpiderCommand(BaseCommand):

    name = "spider remove_spider"
    description = "Crawler script removal"

    def remove_spider(self, spider_name: str):
        """Remove spider name

        :param spider_name: Project name
        """
        print("remove_spider ", spider_name)

    def handle(self):
        self.remove_spider()


class SpiderCommand(BaseCommand):
    """ """

    name = "spider"
    description = "Crawler script project scaffolding"
    commands = [CreateSpiderCommand, RemoveSpiderCommand]
