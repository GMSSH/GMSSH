# -*- encoding: utf-8 -*-
"""
@文件        :SpiderCommands.py
@说明        :
@时间        :2024/11/26 18:51:23
@作者        :Zack
@版本        :1.0
"""

from collections import defaultdict
from typing import Any


class _Parameter:
    """Used to represent default parameter values."""

    def __repr__(self):
        return self.__class__.__name__


class _Anonymous(_Parameter):
    """Singleton used to signal "Anonymous Sender"

    The Anonymous object is used to signal that the sender
    of a message is not specified (as distinct from being
    "any sender").  Registering callbacks for Anonymous
    will only receive messages sent without senders.  Sending
    with anonymous will only send messages to those receivers
    registered for Any or Anonymous.

    Note:
        The default sender for connect is Any, while the
        default sender for send is Anonymous.  This has
        the effect that if you do not specify any senders
        in either function then all messages are routed
        as though there was a single sender (Anonymous)
        being used everywhere.
    """


Anonymous = _Anonymous()


class dispatcher(object):
    """ """

    Anonymous = _Anonymous()
    handlers = defaultdict(list)

    @classmethod
    def connect(cls, receiver: Any, signal: Any, **kwargs: Any):
        """ """
        cls.handlers[signal].append(receiver)

    @classmethod
    def send(cls, signal=Any, *arguments, **kwargs):
        """ """
        for handler in cls.handlers[signal]:
            handler(*arguments, **kwargs)

    @classmethod
    def disconnect(cls, receiver, signal: Any):
        """ """
        raise NotImplementedError

    @classmethod
    def disconnect_all(cls, signal: Any, **kwargs: Any):
        """ """
        raise NotImplementedError

    @classmethod
    def liveReceivers(cls, signal: Any):
        """ """
        raise NotImplementedError

    @classmethod
    def getAllReceivers(cls, sender=Any, signal=Any):
        """ """
        raise NotImplementedError
