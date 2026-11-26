# -*- encoding: utf-8 -*-
"""
@文件        :conf.py
@说明        :
@时间        :2024/11/25 13:33:13
@作者        :Zack
@版本        :1.0
"""
import os
from configparser import ConfigParser
from pathlib import Path
from typing import Iterator, List

from gmscaffold.settings import Settings


def build_component_list(Gm):
    """ """
    from gmscaffold.loader.process import ProcessLoader

    Gm._from_settings(Settings())
    return ProcessLoader(Gm)


def get_config(config_path: str = None) -> ConfigParser:
    """Get Scrapy config file as a ConfigParser"""
    sources = get_sources(config_path)
    cfg = ConfigParser()
    cfg.read(sources)
    return cfg


def get_sources(config_path: str = None) -> List[str]:
    xdg_config_home = os.environ.get("GM_CONFIG_HOME") or Path("~/.config").expanduser()
    sources = [
        str(Path(xdg_config_home) / "gm.cfg"),
    ]
    if config_path:
        sources.append(config_path)
    return sources


def get_extensions_path():
    """ """
    extensions_path = os.path.join(os.path.dirname(os.path.dirname(__file__)), "extensions")
    return extensions_path


def gen_list_dir(base_dir) -> Iterator:
    """ """
    yield from os.listdir(base_dir)


def get_module_name(module: object):
    """ """
    ext_name = "{0.__class__.__module__}.{0.__class__.__name__}".format(module)
    return ext_name


def get_module_basename_with_path(path: str):
    """ """
    return os.path.basename(path)
