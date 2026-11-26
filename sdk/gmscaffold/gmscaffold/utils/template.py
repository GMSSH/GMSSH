# -*- encoding: utf-8 -*-
"""
@文件        :template.py
@说明        :
@时间        :2024/11/27 11:43:41
@作者        :Zack
@版本        :1.0
"""
import os
import sys
from pathlib import Path, PurePosixPath, PureWindowsPath
from typing import Any, Dict

import jinja2

from gmscaffold.base.dataclass import TemplateData
from gmscaffold.exceptions import BuildValueError
from gmscaffold.utils import path as _path


def to_camel_case(s):
    """ """
    return "".join(word.title() for word in s.split("_"))


def to_capitalize(s):
    """ """
    return s.capitalize()


def to_upper(s):
    """ """
    return s.upper()


class Template:
    """Generate project folder"""

    def __init__(self, template_path: str, render: Dict[str, Any], target_dir: str = None):
        """
        :param template_path : Project template path
        :param target_dir : Create folder into target directory
        """
        self._target_dir = target_dir or os.getcwd()
        self._template_path = template_path

        if not isinstance(render, dict):
            """ """
            raise BuildValueError(f"参数类型错误,期望类型: dict; {render}")

        self._render_vars = render
        os.makedirs(self._target_dir, exist_ok=True)

    def exits(self) -> bool:
        """ """
        return os.path.exists(self._template_path)

    def basename(self, path: str) -> str:
        """ """
        _path = Path(path)
        return _path.name

    def to_cwd(self, basename: str):
        """ """
        if "win" in sys.platform:
            pure_path = PureWindowsPath(self._target_dir)
        else:
            pure_path = PurePosixPath(self._target_dir)
        return pure_path.with_name(basename)

    def iter_listdir(self):
        """ """
        if self.exits():
            yield from os.listdir(self._template_path)

    def listdir(self):
        """ """
        if not self.exits():
            return []
        return os.listdir(self._template_path)

    def render(self, file: str):
        """ """
        path = Path(file)
        template_path = os.path.dirname(file)
        template_loader = jinja2.FileSystemLoader(searchpath=template_path)
        template_env = jinja2.Environment(loader=template_loader)
        template_env.filters["to_capitalize"] = to_capitalize
        template = template_env.get_template(path.name)
        rendered_code = template.render(self._render_vars)
        return rendered_code

    def fwrite(self, content: str, file: str, mode="w", encoding="utf-8"):
        """
        :param content: 文件内容
        :param file: 需要写入的文件
        """
        file_folder = os.path.dirname(file)
        new_file = _path.replace_suffix(file, old=".tmpl", new=".py")
        os.makedirs(file_folder, exist_ok=True)
        with open(new_file, mode, encoding=encoding) as f:
            f.write(content)


def build_project(tpl_data: TemplateData):
    """ """
    tpl = Template(tpl_data.template_path, tpl_data.render, tpl_data.target_dir)
    project_folder = os.path.join(tpl._target_dir, tpl_data.template_path_root)
    os.makedirs(project_folder, exist_ok=True)
    for file in tpl_data.files:
        """ """
        rendered_code = tpl.render(file)
        n_path = _path.replace_path(file, tpl_data.template_path, project_folder)
        tpl.fwrite(rendered_code, n_path)


def build_app(tpl_data: TemplateData):
    """ """
    #: 构建
    tpl = Template(tpl_data.template_path, tpl_data.render, tpl_data.target_dir)
    os.makedirs(tpl_data.template_path_root, exist_ok=True)
    for file in tpl_data.files:
        rendered_code = tpl.render(file)
        n_path = _path.replace_path(file, tpl_data.template_path, tpl_data.template_path_root)
        tpl.fwrite(rendered_code, n_path)
