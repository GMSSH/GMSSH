# -*- encoding: utf-8 -*-
import os
from typing import Any, Dict

from simplejrpc._mapping import DefaultMapping

PROJECT_ROOT = os.path.dirname(__file__)
PROJECT_I18n_PATH = os.path.join(PROJECT_ROOT, "i18n")
DEFAULT_GA_SOCKET = os.path.join("/.__gmssh/tmp", "rpc.sock")
DEFAULT_LOGGING_CONFIG = os.path.join(PROJECT_ROOT, "config.yaml")


class Settings:
    """ """

    def __init__(self, config_content: Dict[str, Any] = None):
        """ """
        config_content = config_content or {}
        self.config = DefaultMapping.fromDict(config_content)

    def get_section(self, key: str, default=None) -> Dict[str, Any]:
        """ """
        return default or self.config.__dict__[key]

    def get_config_object(self):
        """ """
        return self.config

    def get_option(self, key: str) -> DefaultMapping:
        """ """
        return getattr(self.config, key)

    def __getattr__(self, key: str) -> DefaultMapping:
        """ """
        return self.get_option(key)
