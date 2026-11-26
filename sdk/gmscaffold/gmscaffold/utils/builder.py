# -*- encoding: utf-8 -*-
"""
@文件        :builder.py
@说明        :
@时间        :2025/01/16 10:43:38
@作者        :Zack
@版本        :1.0
"""
import os
from dataclasses import asdict

from gmscaffold.base.dataclass import BaseData, TemplateData
from gmscaffold.utils import path as _path
from gmscaffold.utils.conf import get_module_basename_with_path, get_module_name

APP_NAME = "app"
APP_SKIP_FOLDER = [APP_NAME]


def make_template_data(ptr, pData: BaseData):
    """ """
    module_name = get_module_name(ptr)
    module_tpl = ptr.Gm._settings.extension_template(module_name)
    files_list = _path.get_files_in_directory(module_tpl, ignore_folders=APP_SKIP_FOLDER)
    tpl_root = pData.name
    project_data = asdict(pData)
    tpl_data = TemplateData(files_list, project_data, module_tpl, tpl_root)
    return tpl_data


def make_template_app(ptr, pData: BaseData):
    """ """
    cwd = os.getcwd()
    tpl_root = os.path.join(cwd, "blueprints", pData.name)
    _module_name = get_module_name(ptr)
    _module_tpl = ptr.Gm._settings.extension_template(_module_name)
    module_tpl = os.path.join(_module_tpl, APP_NAME)
    files_list = _path.get_files_in_directory(module_tpl)
    project_data = asdict(pData)
    project_data["package"] = get_module_basename_with_path(cwd)
    tpl_data = TemplateData(files_list, project_data, module_tpl, tpl_root)
    return tpl_data
