# -*- encoding: utf-8 -*-
import asyncio
import http
import json
import os
from ast import literal_eval
from typing import Any, Dict, List, Optional, Union

from jsonrpcserver.async_main import dispatch_to_response
from jsonrpcserver.response import (
    Left,
    ResponseType,
    SuccessResponse,
    serialize_success,
    to_serializable,
)
from loguru import logger

import simplejrpc.response as res
from simplejrpc.exceptions import RPCException
from simplejrpc.interfaces import RPCMiddleware


class JsonRpcServer:
    """ """

    def __init__(
        self,
        socket_path: str,
    ):
        self.socket_path = socket_path
        self.middlewares: List[RPCMiddleware] = []

    def _to_deserialized(self, response: Left):
        """ """
        _data = response._error.data
        if isinstance(_data, str):
            """ """
            try:
                return literal_eval(_data)
            except ValueError as _:
                return json.loads(_data)
            except Exception as _:
                return res.res_failure(
                    code=http.HTTPStatus.INTERNAL_SERVER_ERROR.value,
                    data=_data,
                    msg=response._error.message,
                )
        return _data

    def _to_serializable(self, response: Left):
        """ """
        if response._error.code < 0:
            response = SuccessResponse(
                id=response._error.id,
                result=self._to_deserialized(response),
            )
            return serialize_success(response)
        else:
            return serialize_success(
                SuccessResponse(id=response.id, result=response._error.data)
            )

    def to_serializable(self, response: ResponseType):
        """ """
        if isinstance(response, Left):
            return self._to_serializable(response)
        else:
            return to_serializable(response)

    def middleware(self, middleware_instance: RPCMiddleware):
        """ """
        self.middlewares.append(middleware_instance)
        return middleware_instance

    async def _process_request(
        self,
        request_data: Union[str, Any],
        context: Dict[str, Any],
        handler: Optional[str] = "process_request",
        reverse: Optional[int] = 1,
    ):
        """ """
        for middleware in self.middlewares[::reverse]:
            process_handle = getattr(middleware, handler)
            if not process_handle:
                continue
            request_data = process_handle(request_data, context)
        return request_data

    async def dispatch_to_response(
        self, *args, serializer=json.dumps, post_process=None, **kwargs
    ):
        """ """
        request_data = args
        context = {
            "request": request_data,
            "serializer": serializer,
            "post_process": post_process,
            "kwargs": kwargs,
        }
        request_data = await self._process_request(request_data, context)
        response = await dispatch_to_response(
            *args, post_process=post_process or self.to_serializable, **kwargs
        )
        response = await self._process_request(
            response, context, handler="process_response", reverse=-1
        )

        return "" if response is None else serializer(response)

    async def handle_client(self, reader, writer):
        """ """
        try:
            # Read headers
            headers = await self._read_headers(reader)
            content_length = self._get_content_length(headers)

            if content_length is None:
                raise ValueError("Missing Content-Length header")

            body = await reader.read(content_length)
            if not body:
                return

            request = body.decode()
            response = await self.dispatch_to_response(request)

            response_data = self._format_response(response)
            writer.write(response_data)
            await writer.drain()
        except RPCException as e:
            logger.error(f"Client validation error: {e}")
            code, err_msg = e.code, e.message
            error_response = self._format_error_response(code, err_msg)
            writer.write(error_response)
        except Exception as e:
            logger.error(f"Client handling error: {e}")
            code = http.HTTPStatus.INTERNAL_SERVER_ERROR.value
            error_response = self._format_error_response(code, str(e))
            writer.write(error_response)
        finally:
            writer.close()
            await writer.wait_closed()

    async def _read_headers(self, reader) -> dict:
        """ """
        headers = {}
        while True:
            line = await reader.readuntil(b"\r\n")
            if line == b"\r\n":
                break
            key, value = line.decode().strip().split(": ", 1)
            headers[key.lower()] = value
        return headers

    def _get_content_length(self, headers: dict) -> Optional[int]:
        """ """
        try:
            return int(headers.get("content-length", 0))
        except (ValueError, TypeError):
            return None

    def _format_response(self, jsonrpc_response) -> bytes:
        """ """
        response_body = (
            json.dumps(jsonrpc_response)
            if not isinstance(jsonrpc_response, str)
            else jsonrpc_response
        )
        headers = f"Content-Length: {len(response_body)}\r\n" "\r\n"
        return headers.encode() + response_body.encode()

    def _format_error_response(self, code, error_msg: str, id=None) -> bytes:
        """ """
        response = SuccessResponse(
            id=id,
            result=json.dumps(
                res.res_failure(
                    code=code,
                    msg=error_msg,
                )
            ),
        )
        error_body = json.dumps(serialize_success(response))
        headers = f"Content-Length: {len(error_body)}\r\n" "\r\n"
        return headers.encode() + error_body.encode()

    def clear_socket(self):
        """ """
        if os.path.exists(self.socket_path):
            os.remove(self.socket_path)

    async def start_server(self, f_socket):
        """ """
        if not f_socket:
            raise RuntimeError("Socket path not configured")

        self.clear_socket()
        server = await asyncio.start_unix_server(self.handle_client, path=f_socket)

        async with server:
            logger.info(f"Server started on {f_socket}")
            await server.serve_forever()

    async def run(self):
        """ """
        await self.start_server(self.socket_path)
