# -*- encoding: utf-8 -*-
"""
@文件        :conf.py
@说明        :
@时间        :2024/11/26 14:33:13
@作者        :Zack
@版本        :1.0
"""
import importlib
from typing import Any, Dict, Iterator

from gmscaffold.log import logger  # type: ignore


class ProcessFinder:
    """ """

    def __init__(self, loader):
        """ """
        self._loader = loader

    def find_extensions(self, listdir: Iterator[Dict[str, Any]]):
        """ """
        for module_name in listdir:
            """ """
            module = self._loader.load_module(module_name)
            module.from_settings(self._loader.Gm)


class ProcessLoader:
    """ """

    def __init__(self, Gm):
        """ """
        self.Gm = Gm
        self._finder = ProcessFinder(self)
        self.load_extensions()

    def load_module(self, module_name: str):
        """ """
        module_dirname, basename = module_name.rsplit(".", maxsplit=1)
        module = importlib.__import__(module_dirname, globals(), locals(), fromlist=["*"])
        return getattr(module, basename)

    def load_extensions(self):
        """ """
        self._finder.find_extensions(self.Gm._settings.EXTENSIONS_BASE)
