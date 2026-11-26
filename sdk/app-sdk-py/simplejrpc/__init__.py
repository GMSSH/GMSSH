# -*- encoding: utf-8 -*-
from simplejrpc import exceptions  # type: ignore
from simplejrpc._field import *  # type: ignore
from simplejrpc._mapping import DefaultMapping  # type: ignore
from simplejrpc._text import ErrorTextMessage as TextMessage  # type: ignore
from simplejrpc.app import ServerApplication  # type: ignore
from simplejrpc.client import Request  # type: ignore
from simplejrpc.config import Settings  # type: ignore
from simplejrpc.field import *  # type: ignore
from simplejrpc.i18n import T as i18n  # type: ignore
from simplejrpc.interfaces import (  # type: ignore
    BaseServer,
    BaseValidator,
    ClientTransport,
    RPCMiddleware,
)
from simplejrpc.validate import *  # type: ignore

try:
    # For python 3.8 and later
    import importlib.metadata as importlib_metadata
except ImportError:
    # For everyone else
    import importlib_metadata
try:
    __version__ = importlib_metadata.version("simplejrpc")
except importlib_metadata.PackageNotFoundError:
    # package is not installed
    __version__ = "0.0.0"
