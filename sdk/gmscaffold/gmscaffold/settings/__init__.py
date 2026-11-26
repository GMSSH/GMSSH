# -*- encoding: utf-8 -*-
"""
@文件        :test_cmdline.py
@说明        :
@时间        :2024/11/26 15:40:03
@作者        :Zack
@版本        :1.0
"""

from pathlib import Path

from gmscaffold.settings import default_settings


class Settings(object):
    """ """

    def __init__(self):
        """ """
        self._settings = default_settings

    def upsert_settings(self):
        """ """
        self.merge_extensions()
        self.reversed_settings()

    def reversed_settings(self):
        """ """

    def merge_extensions(self):
        """ """

        if not self._settings.EXTENSIONS_BASE or not isinstance(self._settings.EXTENSIONS_BASE, dict):
            return

        self._settings.EXTENSIONS = {} if not self._settings.EXTENSIONS else self._settings.EXTENSIONS
        self._settings.EXTENSIONS_BASE |= self._settings.EXTENSIONS

    def __getattr__(self, __name: str):
        """ """
        return getattr(self._settings, __name)

    def extensions(self):
        """Get all extensions info, exclude incomplete data"""
        return dict(
            filter(
                lambda x: bool(x[0]) and bool(x[1]),
                self._settings.EXTENSIONS_TEMPLATE_PATH.items(),
            )
        )

    def extension_template(self, ext_name: str) -> Path | None:
        """Get extension template path

        :param ext_name: Extension name
        :return: Extension template path or None
        """
        ext_dict = self.extensions()
        return ext_dict.get(ext_name)
