# -*- encoding: utf-8 -*-
"""
@文件        :cmdline.py
@说明        :
@时间        :2024/11/25 10:54:59
@作者        :Zack
@版本        :1.0
"""
import os

BASE_DIR = os.path.dirname(os.path.dirname(__file__))
os.environ["PYTHONPATH"] = BASE_DIR
STATIC_DIR = os.path.join(BASE_DIR, "templates")

try:
    pass
except Exception:
    pass


class _BaseExtension:
    """ """

    def __init__(self, module, template, weight=0, version=None):
        """ """
        self.module = module
        self.template = template
        self.weight = weight
        self.version = version  # No features added yet

    def __getitem__(self, key: str):
        """ """
        return self.__dict__[key]

    def __missing__(self, key):
        """ """
        return None


def _construct_extension(name: str, template: str, weight=0):
    """Construct extension info

    :param name: Extension module
    :param template: Extension template path basename
    :param weight: Load extension component weight, No features added yet
    """
    return _BaseExtension(name, os.path.join(STATIC_DIR, template), weight=weight)


SpiderExtension = _construct_extension("gmscaffold.extensions.spider.Spider", "spider")
WebExtension = _construct_extension("gmscaffold.extensions.web.Web", "web")
GmAppExtension = _construct_extension("gmscaffold.extensions.GmApp.GmApp", "GmApp")

EXTENSIONS = [
    SpiderExtension,
    GmAppExtension,
    WebExtension,
]

EXTENSIONS_BASE = {}

EXTENSIONS_TEMPLATE_PATH = {}

for ext_info in EXTENSIONS:
    """ """
    EXTENSIONS_BASE[ext_info["module"]] = ext_info["weight"]
    EXTENSIONS_TEMPLATE_PATH[ext_info["module"]] = ext_info["template"]
