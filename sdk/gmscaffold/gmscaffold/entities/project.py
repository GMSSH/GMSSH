# -*- encoding: utf-8 -*-
"""
@文件        :cmdline.py
@说明        :
@时间        :2024/11/25 10:54:59
@作者        :Zack
@版本        :1.0
"""
import enum


class ProjectEnum(enum.Enum):
    """Project type enumeration class,
        used to specify different styles of creating projects.

    The following are the available project types:
    - SPIDER: represents the style of spider script projects, suitable for [script projects of specific applicable script types].
    - SCRAPY: represents the style of spider Scrapy projects, used for [specifically applicable Scrapy script projects].
    - GM_APP: represents the style of GM (App) projects, often used for [specifically applicable GMB-APP application projects].
    - WEB_SERVER: represents the style of traditional Web projects, often used for [specifically using traditional Web projects].
    """

    SPIDER = enum.auto()
    SCRAPY = enum.auto()
    GM_APP = enum.auto()
    WEB_SERVER = enum.auto()
