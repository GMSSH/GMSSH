# -*- encoding: utf-8 -*-
import asyncio
import atexit
import inspect
import os
from functools import partial, wraps
from pathlib import Path
from typing import Any, Callable, Dict, NoReturn, Optional, Tuple, Union, cast, overload

from jsonrpcserver import method
from loguru import logger
from wtforms import Form as _Form
from wtforms import ValidationError

from simplejrpc import exceptions
from simplejrpc._sockets import JsonRpcServer
from simplejrpc.config import Settings
from simplejrpc.daemon.daemon import DaemonContext
from simplejrpc.i18n import T as i18n
from simplejrpc.interfaces import RPCMiddleware
from simplejrpc.parse import IniConfigParser, JsonConfigParser, YamlConfigParser
from simplejrpc.response import raise_exception
from simplejrpc.schemas import BaseForm as _BaseForm
from simplejrpc.validate import BaseForm, Form


class ServerApplication:
    """ """

    def __init__(
        self,
        socket_path: str,
        config: Optional[object] = Settings(),
        config_path: Optional[str] = os.path.join(os.getcwd(), "config.yaml"),
        i18n_dir: Optional[str] = os.path.join(os.getcwd(), "app", "i18n"),
    ):
        self.server = JsonRpcServer(socket_path)
        self.config_path = config_path
        self.config = config
        i18n.set_path(i18n_dir)
        if self.config_path is not None and os.path.exists(self.config_path):
            self.from_config(config_path=self.config_path)

    def from_config(
        self,
        config_content: Optional[Dict[str, Any]] = None,
        config_path: Optional[str] = None,
    ) -> Settings:
        """ """
        if config_content:
            self.config = Settings(config_content)
        if config_path:
            """ """
            config_content = self.load_config(config_path)
        return self.config

    @overload
    def route(
        self,
        name: Optional[str] = None,
        form: Optional[_Form] = _BaseForm,
        fn: Callable[[Union[Tuple, Dict[str, Any]]], Any] = None,
    ) -> Callable[[Union[Tuple, Dict[str, Any]]], Any]: ...

    @overload
    def route(
        self,
        name: Optional[str] = None,
        form: Optional[Form] = BaseForm,
        fn: Callable[[Union[Tuple, Dict[str, Any]]], Any] = None,
    ) -> Callable[[Union[Tuple, Dict[str, Any]]], Any]: ...

    def route(
        self,
        name: Optional[str] = None,
        form: Optional[Union[Form, _Form]] = BaseForm,
        fn=None,
    ) -> Callable[[Union[Tuple, Dict[str, Any]]], Any]:
        """路由装饰器"""
        if fn is None:
            return partial(self.route, name, form)

        @wraps(fn)
        def wrapper(*args, **kwargs):
            """ """
            if form:
                params = dict(zip(inspect.getfullargspec(fn).args, args))
                params.update(kwargs)
                form_validate = form(**params)
                if isinstance(form_validate, _Form):
                    form_validate = cast(_Form, form_validate)
                    if not form_validate.validate():
                        for _, errors in form_validate.errors.items():
                            for error in errors:
                                raise_exception(ValidationError, msg=error)
                else:
                    form_validate = cast(Form, form_validate)
                    form_validate.raise_all_errors()

            return fn(*args, **kwargs)

        method(wrapper, name=name or fn.__name__)
        return wrapper

    def load_config(self, config_path: str) -> Dict[str, Any]:
        """ """

        if not os.path.exists(config_path):
            """ """
            raise exceptions.FileNotFoundError(f"Not found path {config_path}")

        path = Path(config_path)
        base_file = path.name
        _, filetype = base_file.split(".")

        match filetype:
            case "yml" | "yaml":
                parser = YamlConfigParser(config_path)
            case "ini":
                parser = IniConfigParser(config_path)
            case "json":
                parser = JsonConfigParser(config_path)
            case _:
                raise exceptions.ValueError("Unable to parse the configuration file")
        config_content: Dict[str, Any] = parser.read()
        self.config = Settings(config_content)
        self.setup_logger(config_content)
        return config_content

    def setup_logger(self, config_content: Dict[str, Any]):
        """ """
        # NOTE:: logger必须携带且sink必须携带
        logger_config_items = config_content.get("logger", {})
        if "sink" not in logger_config_items:
            return

        sink = self.config.logger.sink
        os.makedirs(Path(sink).parent, exist_ok=True)
        logger.add(**logger_config_items)

    def clear_socket(self):
        """ """
        self.server.clear_socket()

    def middleware(self, middleware_instance: RPCMiddleware):
        """中间件配置"""
        return self.server.middleware(middleware_instance)

    def run_daemon(self, fpidfile: Path) -> NoReturn:
        """Start service in daemon mode"""
        with DaemonContext(fpidfile=fpidfile):
            asyncio.run(self.server.run())

    async def run(self, daemon: bool = False, fpidfile: Path = None) -> NoReturn:
        """
        :param daemon: Whether to run as a daemon process
        :param fpidfile: Guardian process PID file
        :return:
        """
        atexit.register(self.clear_socket)
        if daemon:
            self.run_daemon(fpidfile)
        else:
            await self.server.run()
