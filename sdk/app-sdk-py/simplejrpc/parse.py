# -*- encoding: utf-8 -*-
import configparser
import json
from abc import abstractmethod
from typing import Any, Dict, NoReturn

import yaml


class ConfigParser:
    """ """

    def __init__(self, config_path: str):
        """ """
        self.config_path = config_path

    @abstractmethod
    def read(self):
        """ """

    @abstractmethod
    def write(self):
        """ """


class JsonConfigParser(ConfigParser):
    """ """

    def read(self) -> Any:
        """ """
        with open(self.config_path, encoding="utf-8", mode="r") as f:
            return json.load(f)

    def write(self, content: Dict[str, Any]) -> NoReturn:
        with open(self.config_path, mode="w") as f:
            json.dump(content, f)


class YamlConfigParser(ConfigParser):
    """ """

    def read(self) -> Any:
        """ """
        with open(self.config_path, "r", encoding="utf-8") as f:
            return yaml.safe_load(f)

    def write(self, content: Dict[str, Any]) -> NoReturn:
        """ """
        with open(self.config_path, "w") as f:
            yaml.dump(content, f)


class IniConfigParser(ConfigParser):
    """ """

    def read(self):
        """ """
        config = configparser.ConfigParser()
        config.read(self.config_path)
        return config._sections

    def write(self, content: str):
        """ """
        with open(self.config_path, "w") as f:
            f.write(content)
