# -*- encoding: utf-8 -*-
"""
@文件        :GmApp.py
@说明        :
@时间        :2024/11/27 18:40:46
@作者        :Zack
@版本        :1.0
"""
from dataclasses import asdict

from gmscaffold.base.command import BaseBuilder
from gmscaffold.base.dataclass import TemplateData
from gmscaffold.entities.GmApp import GmAppData
from gmscaffold.log import logger
from gmscaffold.utils import path as _path
from gmscaffold.utils import signal as _signal
from gmscaffold.utils import template
from gmscaffold.utils.conf import get_module_name


class GmApp(BaseBuilder):
    """ """

    component_name = "gmapp"

    def __init__(self, Gm):
        """ """
        self.Gm = Gm

        signals_callbacks = {
            _signal.create_gmapp_project: self.create_gm_app,
            _signal.remove_gmapp_project: self.remove_gm_app,
        }

        super().__init__(signals_callbacks)

    def create_gm_app(self, sender, **kwargs):
        """ """
        self.check_data(kwargs)
        project_item: GmAppData = kwargs["data"]
        module_name = get_module_name(self)
        module_tpl = self.Gm._settings.extension_template(module_name)
        files_list = _path.get_files_in_directory(module_tpl)
        tpl_root = project_item.name
        project_data = asdict(project_item)
        tpl_data = TemplateData(files_list, project_data, module_tpl, tpl_root)

        # TODO: 测试
        # test_folder = os.path.dirname(os.path.dirname(__file__))
        # tpl_data.target_dir = os.path.join(test_folder, "test")

        template.build_project(tpl_data)
        logger.debug(f"[{project_item.name}] create_gm_app successfully")

    def remove_gm_app(self, sender, **kwargs):
        """ """
        # TODO: NotImplementedError
        print("[GmApp] remove_gm_app...", sender, kwargs)

    @classmethod
    def from_settings(cls, settings):
        """ """
        return cls(settings)
