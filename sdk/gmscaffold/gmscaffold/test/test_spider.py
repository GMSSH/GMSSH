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
import unittest

sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.dirname(__file__))))

from gmscaffold.cmdline import GmScaffoldMain
from gmscaffold.commands.SpiderCommand import SpiderCommand
from gmscaffold.log import Logger
from gmscaffold.settings.default_settings import STATIC_DIR
from gmscaffold.utils.path import *
from gmscaffold.utils.template import *


class TestSpider(unittest.TestCase):
    """ """

    @unittest.skip
    def test_get_dirs(self):
        """ """
        directory = STATIC_DIR
        dirs = get_files_in_directory(directory)
        print("[*] dirs >>> ", dirs)

    @unittest.skip
    def test_spider_template(self):
        """ """
        test_dir = os.getcwd()
        tmp = Template(test_dir)
        print("[*] test_spider_template >> ", tmp.to_cwd("logs"))

    @unittest.skip
    def test_spider(self):
        """ """
        sc = SpiderCommand()
        sc.create_spider()

    @unittest.skip
    def test_loader_web(self):
        """ """
        gm = GmScaffoldMain()
        gm.web("start_project")

    # @unittest.skip
    def test_loader(self):
        """ """
        gm = GmScaffoldMain()
        gm.spider("create_spider")
        # build_component_list()
        # pl = ProcessLoader()
        # pl.load_extensions()

    @unittest.skip
    def test_logger(self):
        """ """
        file = os.path.join(os.path.dirname(os.path.dirname(__file__)), "logs", "test.log")
        logger = Logger.from_log_file(file)
        logger.info("ddddddddddddd")


if __name__ == "__main__":
    """ """
    unittest.main()
