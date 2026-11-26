# -*- encoding: utf-8 -*-
"""
@文件        :test_cmdline.py
@说明        :
@时间        :2024/11/25 13:40:03
@作者        :Zack
@版本        :1.0
"""
import os
import sys

sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.dirname(__file__))))

import unittest

from gmscaffold.utils.conf import get_config


class TestConfig(unittest.TestCase):
    """ """

    def test_get_config(self):
        """ """
        cfg = get_config()
        print("[*] cfg ", cfg)


if __name__ == "__main__":
    """ """
    unittest.main()
