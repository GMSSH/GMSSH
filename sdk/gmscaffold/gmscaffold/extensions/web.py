# -*- encoding: utf-8 -*-
"""
@文件        :web.py
@说明        :
@时间        :2024/11/27 18:45:12
@作者        :Zack
@版本        :1.0
"""

from gmscaffold.base.command import BaseBuilder
from gmscaffold.entities.web import WebData
from gmscaffold.utils import builder
from gmscaffold.utils import signal as _signal
from gmscaffold.utils import template


class Web(BaseBuilder):
    """ """

    component_name = "sanic"

    def __init__(self, Gm):
        """ """
        self.Gm = Gm
        signals_callbacks = {
            _signal.create_web_project: self.start_project,
            _signal.remove_web_project: self.remove_project,
            _signal.create_web_app: self.start_app,
            _signal.remove_web_app: self.remove_app,
        }

        super().__init__(signals_callbacks)

    def start_project(self, sender, **kwargs):
        """ """
        self.check_data(kwargs)
        project_item: WebData = kwargs["data"]
        tpl_data = builder.make_template_data(self, project_item)

        template.build_project(tpl_data)

    def remove_project(self):
        """ """

    def start_app(self, sender, **kwargs):
        """ """
        self.check_data(kwargs)
        project_item: WebData = kwargs["data"]
        tpl_data = builder.make_template_app(self, project_item)

        template.build_app(tpl_data)

    def remove_app(self):
        """ """
