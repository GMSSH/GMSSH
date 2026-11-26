# -*- encoding: utf-8 -*-
import asyncio
import json
from typing import Any, Dict, Optional, Tuple, Union

from jsonrpcclient import request
from jsonrpcclient.sentinels import NOID

from simplejrpc.config import DEFAULT_GA_SOCKET
from simplejrpc.interfaces import ClientTransport
from simplejrpc.response import Response


class UnixSocketTransport(ClientTransport):
    """ """

    def __init__(self, socket_path: str):
        self.reader = None
        self.writer = None
        self._socket_path = socket_path

    async def connect(self):
        self.reader, self.writer = await asyncio.open_unix_connection(
            path=self._socket_path
        )

    async def make_header(self, message: str):
        """ """
        message = f"Content-Length: {len(message)}\r\n\r\n{message}"
        self.writer.write(message.encode("utf-8"))

    async def make_body(self, body: str):
        """ """
        self.writer.write(body.encode("utf-8"))

    async def send_message(self, message):
        if isinstance(message, dict):
            message = json.dumps(message)

        await self.make_header(message)
        await self.make_body(message)
        await self.writer.drain()  # Ensure the data is sent

        # Read response
        return await self._read_response()

    # Read Content-Length header
    async def read_from_header(self):
        """ """
        return await self.reader.readuntil(b"\r\n\r\n")

    # Parse Content-Length
    async def read_content_length_from_header(self, header: str):
        """ """
        return int(header.split(b":")[1].strip())

    async def _read_response(self):
        """ """
        header = await self.read_from_header()
        content_length = await self.read_content_length_from_header(header)

        # Read response body
        response_body = await self.reader.readexactly(content_length)
        return Response(json.loads(response_body))

    def close(self):
        if self.writer:
            self.writer.close()


class Request:
    """ """

    def __init__(self, socket_path: str, adapter: Optional[ClientTransport] = None):
        """ """
        self._adapter = adapter or UnixSocketTransport(socket_path)

    async def make_session(self):
        """ """
        await self._adapter.connect()

        # send-message

    async def async_send_request(
        self,
        method: str,
        params: Union[Dict[str, Any], Tuple[Any, ...], None] = None,
        id: Any = NOID,
    ):
        """ """
        await self.make_session()
        try:
            request_body = request(method, params, id=id)
            response = await self._adapter.send_message(request_body)
            return response
        finally:
            self._adapter.close()

    def send_request(
        self,
        method: str,
        params: Union[Dict[str, Any], Tuple[Any, ...], None] = None,
        id: Any = NOID,
    ):

        return asyncio.run(self.async_send_request(method, params=params, id=id))


class GmRequest(Request):
    """ """

    def __init__(self, socket_path=DEFAULT_GA_SOCKET, adapter=None):
        """ """
        super().__init__(socket_path, adapter)
