"""
@文件        :__init__.py
@说明        :示例
@时间        :2025/05/13 14:06:36
@作者        :xxx
@邮箱        :xxxxxxxxxxx.163.com
@版本        :0.0.1
"""

from app.services.example import Example


class ExampleService:
    """ """
    def __init__(self):
        self.example = Example()

    def hello(self):
        """ """
        return self.example.hello()
