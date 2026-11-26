# -*- encoding: utf-8 -*-

import enum
import functools
import os
from abc import ABCMeta, abstractmethod
from pathlib import Path
from typing import Any, Final, List, Union, cast

from simplejrpc.exceptions import ValueError

__all__ = [
    "Ten",
    "Tja",
    "Tru",
    "TzhCN",
    "TzhTW",
    "TranslateBase",
    "DefaultTranslate",
    "T",
    "GI18n",
]


class Language(str, enum.Enum):
    """Language codes"""

    EN = "en"
    JA = "ja"
    RU = "ru"
    ZH_CN = "zh-CN"
    ZH_TW = "zh-TW"

    @classmethod
    def values(cls):
        """ """
        return [x.value for x in cls.__members__.values()]


# +--------------------------------------------------
# i18n
# +--------------------------------------------------
class _T:
    """Internal class for translation-related properties and methods"""

    # Default language code
    LANG = Language.ZH_CN
    # Directory of i18n
    I18N_ROOT: Path

    @classmethod
    def exists_t(cls, file):
        """Check if the specified i18n file exists, raise a PyGmValueError if not"""
        if not os.path.exists(file):
            raise ValueError(f"File does not exist: {file}")
        ini_utils = CustomIniConfig()
        return ini_utils.read(file)

    @classmethod
    def translate(cls, key):
        """Get translation for the specified key"""

        i18n_file = os.path.join(cls.I18N_ROOT, f"{cls.LANG.value}.ini")
        t = cls.exists_t(i18n_file)
        if value := t.get(key):
            return value
        raise ValueError("Not found language key")

    @classmethod
    def translate_ctx(cls, key, *value):
        """Get formatted translation for the specified key and value"""
        fmt: str = cls.translate(key)
        if not fmt:
            return None
        return fmt.format(*value)


class Ten(_T):
    """English i18n class"""

    I18N_ROOT: Path = None
    LANG: Final = Language.EN


class Tja(_T):
    """Japanese i18n class"""

    I18N_ROOT: Path = None
    LANG: Final = Language.JA


class Tru(_T):
    """Russian i18n class"""

    I18N_ROOT: Path = None
    LANG: Final = Language.RU


class TzhCN(_T):
    """Chinese i18n class"""

    I18N_ROOT: Path = None
    LANG: Final = Language.ZH_CN


class TzhTW(_T):
    """Chinese i18n class"""

    I18N_ROOT: Path = None
    LANG: Final = Language.ZH_TW


# +--------------------------------------------------
# Load i18n config from file
# +--------------------------------------------------
# Define a class named CustomIniConfig, inheriting from the dictionary type
class CustomIniConfig(dict):
    """ """

    # Initialize method, accepts a file path parameter
    def __init__(self, file_path=None):
        """ """
        # If the file path parameter is not empty, assign it to the file_path attribute
        if file_path is not None:
            self.file_path = file_path

    # Method to read the file and create a dictionary
    def read(self, file_path):
        """ """
        # Assign the file path parameter to the file_path attribute
        self.file_path = file_path
        # Return the created dictionary
        return self.__create_dict()

    # Method to create the dictionary from the file
    def __create_dict(self):
        """ """
        # Open the file in read mode
        with open(self.file_path, "r", encoding="utf-8") as f:
            # Initialize an empty dictionary
            variables = {}
            # Iterate over each line in the file
            for line in f.readlines():
                # Update the dictionary with the parsed line
                variables.update(self.__parse_line(line))
            # Return the created dictionary
            return variables

    # Method to parse a line from the file
    def __parse_line(self, line):
        """ """
        # If the line starts with a comment symbol, discard it and return an empty dictionary
        if line.lstrip().startswith("#"):
            return {}
        # If the line is not empty
        if line.lstrip():
            # Find the second occurrence of a quote mark
            quote_delimit = max(
                line.find("'", line.find("'") + 1), line.find('"', line.rfind('"')) + 1
            )
            # Find the first comment mark after the second quote mark
            comment_delimit = line.find("#", quote_delimit)
            # Trim the line and split it into key and value
            key, value = map(
                lambda x: x.strip().strip("'").strip('"'), line.split("=", 1)
            )
            # Return a dictionary with the key and value
            return {key: value}
        else:
            # Return an empty dictionary
            return {}

    # Method to persist the changes to the file
    def __persist(self):
        """ """
        # Open the file in write mode
        with open(self.file_path, "w") as f:
            # Iterate over each key-value pair in the dictionary
            for key, value in self.items():
                # Write the key and value to the file, separated by an equal sign
                f.write("%s=%s\n" % (key, value))

    # Method to set a value for a key in the dictionary
    def __setitem__(self, key, value):
        """ """
        # Call the parent class's __setitem__ method to set the value for the key
        super(CustomIniConfig, self).__setitem__(key, value)
        # Persist the changes to the file
        self.__persist()

    # Method to delete a key from the dictionary
    def __delitem__(self, key):
        """ """
        # Call the parent class's __delitem__ method to delete the key
        super(CustomIniConfig, self).__delitem__(key)
        # Persist the changes to the file
        self.__persist()


# +--------------------------------------------------
# i18n translate
# +--------------------------------------------------
class TranslateBase(metaclass=ABCMeta):
    """ """

    def __init__(self, lang: Language, i18n_root: str = None) -> None:
        """
        :param lang: 国际化语言
        :param i18n_path: 国际化的目录位置
        """
        self._lang = lang
        self._i18n_path = i18n_root

    @abstractmethod
    def translate(self, *args, **kwargs): ...

    @abstractmethod
    def translate_ctx(self, *args, **kwargs): ...

    @abstractmethod
    def translate_load(self, *args, **kwargs): ...


# Plugin
class DefaultTranslate(TranslateBase):
    """ """

    def translate(self, key: Language):
        """ """
        match self._lang:
            case Language.EN:
                translate_cls = Ten
            case Language.JA:
                translate_cls = Tja
            case Language.RU:
                translate_cls = Tru
            case Language.ZH_CN:
                translate_cls = TzhCN
            case Language.ZH_TW:
                translate_cls = TzhTW
            case _:
                raise ValueError("Unknown language code.")

        if self._i18n_path is not None:
            translate_cls.I18N_ROOT = self._i18n_path
        return translate_cls.translate(key)

    def translate_ctx(self, key, *value):
        """Get formatted translation for the specified language and value"""
        match self._lang:
            case Language.EN:
                translate_cls = Ten
            case Language.JA:
                translate_cls = Tja
            case Language.RU:
                translate_cls = Tru
            case Language.ZH_CN:
                translate_cls = TzhCN
            case Language.ZH_TW:
                translate_cls = TzhTW
            case _:
                raise ValueError("Unkonw language type")

        if self._i18n_path is not None:
            translate_cls.I18N_ROOT = self._i18n_path
        return translate_cls.translate_ctx(key, *value)

    def translate_load(self, file, ignore) -> Path:
        """ """
        path = Path(file)
        file_dir = os.path.join(os.path.dirname(file), self._lang)
        if not os.path.exists(file_dir):
            os.makedirs(file_dir)
        file_path = os.path.join(file_dir, path.name)
        if not os.path.exists(file_path) and not ignore:
            raise ValueError(f"File does not exist: {file_path}")
        return file_path


# +--------------------------------------------------
# Translation
# +--------------------------------------------------


class T:
    """
    # This class provides methods for translating text using a specified language and source, with caching
    # functionality and the ability to set a default language.
    """

    I18n_ROOT: Path
    LANG: Language = Language.EN
    Translate = DefaultTranslate

    @classmethod
    def _translate(cls, lang, key):
        """ """
        return cls.Translate(lang, i18n_root=cls.I18n_ROOT).translate(key)

    @classmethod
    def translate(cls, key: str):
        """Get translation for the specified language"""
        return cls._translate(cls.LANG, key)

    def _translate_partial(cls, lang, key):
        """ """
        return cls.Translate(lang, i18n_root=cls.I18n_ROOT).translate(key)

    @classmethod
    def translate_partial(cls, key: str):
        """ """
        return functools.partial(cls._translate_partial, cls=cls, key=key)

    @classmethod
    def _translate_ctx(cls, lang, key, *value):
        """ """
        return cls.Translate(lang, i18n_root=cls.I18n_ROOT).translate_ctx(key, *value)

    @classmethod
    def translate_ctx(cls, key, *value):
        """Get formatted translation for the specified language and value"""
        """Get translation for the specified language"""
        return cls._translate_ctx(cls.LANG, key, *value)

    @classmethod
    def _translate_ctx_partial(cls, lang, key, value):
        """ """
        return cls.Translate(lang, i18n_root=cls.I18n_ROOT).translate_ctx(key, *value)

    @classmethod
    def translate_ctx_partial(cls, key: str, value: List[Any]):
        """Get formatted translation for the specified language"""
        return functools.partial(
            cls._translate_ctx_partial, cls=cls, key=key, value=value
        )

    @classmethod
    def _translate_load(cls, lang, file, ignore=False):
        """Return a translate from a file .
        Returns:
            [type]: [description]
        """
        return cls.Translate(lang, i18n_root=cls.I18n_ROOT).translate_load(file, ignore)

    @classmethod
    def translate_load(cls, file: Path, ignore=False):
        """ """
        return cls._translate_load(cls.LANG, file, ignore)

    @classmethod
    def _translate_load_partial(cls, lang, file, ignore=False):
        """ """
        return cls.Translate(lang, i18n_root=cls.I18n_ROOT).translate_load(file, ignore)

    @classmethod
    def translate_load_partial(cls, file: Path, ignore=False):
        """ """
        return functools.partial(
            cls._translate_load_partial,
            cls=cls,
            file=file,
            ignore=ignore,
        )

    @classmethod
    def translate_app_info(cls, file_root: Path, filename: str):
        """
        :param : file_root 国际化目标目录
        :param : filename : 需要国际化的文件
        """
        file = os.path.join(file_root, filename)
        return cls._translate_load(cls.LANG, file, ignore=True)

    @classmethod
    def set_lang(cls, lang: Language = Language.ZH_CN):
        """Set the default language"""
        cls.LANG = lang

    @classmethod
    def get_lang(cls):
        """ """
        return cls.LANG

    @classmethod
    def set_path(cls, path: str):
        """ """
        cls.I18n_ROOT = path

    @property
    def lang(self):
        """ """
        return self.LANG


GI18nTranslate = DefaultTranslate


class GI18n:
    """ """

    def __init__(self, path: str, lang: str, adapter: TranslateBase = None) -> None:
        """
        :param path: 国际化配置文件位置
        :param lang: 国际化语言配置
        :param adapter: 国际化翻译器(如果传入该参数则默认使用该适配器)
        """
        GI18nMan = T
        GI18nMan = cast(T, GI18nMan)
        GI18nMan.set_lang(lang)
        GI18nMan.set_path(path)
        if adapter is not None:
            GI18nMan.Translate = adapter
            self._gm = GI18nMan
        else:
            self._gm = GI18nMan

    def translate(self, key) -> str:
        """ """
        return self._gm.translate(key)

    def get_lang(self) -> str:
        """ """
        return self._gm.get_lang()

    def set_path(self, path: str):
        """ """
        self._gm.set_path(path)

    def set_lang(self, lang: str):
        """ """
        self._gm.set_lang(lang)

    def translate_ctx(self, key, *value) -> str:
        """ """
        return self._gm.translate_ctx(key, *value)

    def translate_load(self, file: Union[str, Path], ignore=False) -> str:
        """ """
        return self._gm.translate_load(file, ignore)

    def translate_load_partial(self, file: Path, ignore=False) -> str:
        """ """
        return self._gm.translate_load_partial(file, ignore)

    def translate_app_info(self, file_root: Path, filename: str) -> str:
        """ """
        return self._gm.translate_app_info(file_root, filename)
