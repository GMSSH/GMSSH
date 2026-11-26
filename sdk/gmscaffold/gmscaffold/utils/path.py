# -*- encoding: utf-8 -*-
"""
@文件        :path.py
@说明        :
@时间        :2024/11/27 15:14:10
@作者        :Zack
@版本        :1.0
"""
import os
from pathlib import Path


def get_files_in_directory(directory: str, ignore_folders=[]):
    """
    递归获取目录中的所有文件。
    :param directory: 要检查的目录。
    :return: 包含文件路径的列表。
    """
    files_list = []
    for entry in os.scandir(directory):
        if entry.is_file():
            files_list.append(entry.path)
        elif entry.is_dir():
            if entry.path in ignore_folders:
                continue
            files_list.extend(get_files_in_directory(entry.path, ignore_folders=ignore_folders))
    return files_list


def replace_path(source: str, old: str, new: str):
    """ """
    return source.replace(old, new)


def replace_suffix(source: str, old: str, new: str) -> str:
    """ """
    s_path = Path(source)
    return os.path.join(s_path.parent, s_path.name.replace(old, new))
