# -*- encoding: utf-8 -*-

from simplejrpc.exceptions import RPCException


class DaemonError(RPCException):
    """ """


class DaemonOSEnvironmentError(DaemonError):
    """ """


class DaemonProcessDetachError(DaemonError):
    """ """
