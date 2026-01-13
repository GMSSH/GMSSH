"""
@文件        :__init__.py
@说明        :This is an example
@时间        :2025/06/30 09:17:23
@作者        :xxx
@邮箱        :
@版本        :1.0.0
"""

from simplejrpc.interfaces import RPCMiddleware

class ExampleMiddleware(RPCMiddleware):
    """ """

    def process_request(self, request, context):
        # print("[middleware-request] ", request, context)
        return request

    def process_response(self, response, context):
        # print("[middleware-response] ", response, context)
        return response