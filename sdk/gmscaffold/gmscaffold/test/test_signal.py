# -*- encoding: utf-8 -*-
"""
@文件        :test_signal.py
@说明        :
@时间        :2024/11/25 11:45:30
@作者        :Zack
@版本        :1.0
"""
import unittest

from pydispatch import dispatcher


def handle_print1(sender, **kwargs):
    """测试函数"""
    print("[*] handle_print1:", sender, kwargs)


def handle_print2(sender, **kwargs):
    """测试函数"""
    print("[*] handle_print2:", sender, kwargs)


signal_print1 = object()
signal_print2 = object()


class TestSignal(unittest.TestCase):
    """ """

    def test_signal(self):
        """ """
        dispatcher.connect(handle_print1, signal_print1)
        dispatcher.connect(handle_print2, signal_print2)
        dispatcher.send(signal=signal_print1, sender=dispatcher.Anonymous, item="hello print")
        dispatcher.send(signal=signal_print1, sender=dispatcher.Anonymous, item="hello print")
        dispatcher.send(signal=signal_print2, sender=dispatcher.Anonymous, item="hello print")


if __name__ == "__main__":
    unittest.main()
