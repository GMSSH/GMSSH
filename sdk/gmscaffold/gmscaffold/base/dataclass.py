# -*- encoding: utf-8 -*-
"""
@文件        :spider.py
@说明        :
@时间        :2024/11/25 20:01:51
@作者        :Zack
@版本        :1.0
"""
import datetime
from dataclasses import dataclass, field
from typing import Any, Dict, List, Optional


@dataclass
class BaseData:
    """ """

    #: 包名
    package: Optional[str] = field(default="")
    # 项目名称
    name: Optional[str] = field(default="")
    # 作者
    author: Optional[str] = field(default="")
    # 邮箱
    email: Optional[str] = field(default="")
    # 版本
    version: Optional[str] = field(default="")
    # 描述信息
    description: Optional[str] = field(default="")
    # 创建时间
    date: Optional[str] = field(default=datetime.datetime.now().strftime("%Y/%m/%d %H:%M:%S"))


@dataclass
class TemplateData:
    """ """

    files: List[str]
    render: Dict[str, Any]
    template_path: str
    template_path_root: str
    target_dir: Optional[str] = field(default=None)
    app: str = False
