"""
@文件        :__init__.py
@说明        :This is an example
@时间        :2025/06/30 09:17:23
@作者        :xxx
@邮箱        :
@版本        :1.0.0
"""
import os

from simplejrpc.app import ServerApplication
from simplejrpc.response import jsonify
from simplejrpc.i18n import T as i18n

from app.consts import settings
from app.services.example import Example
from app.schemas.example import ExampleForm
from app.middlewares.example import ExampleMiddleware

app = ServerApplication(socket_path=settings.APP_SOCKET_FILE_PATH, i18n_dir=settings.APP_I18N_DIR_PATH,
                        config_path=settings.APP_CONFIG_FILE_PATH)
app.middleware(ExampleMiddleware())


@app.route(name="hello", form=ExampleForm)
async def hello(**kwargs):
    """ """
    example = Example()
    data = await example.hello(kwargs)
    return jsonify(data=data, msg=i18n.translate("STATUS_OK"))


# 状态检查接口
@app.route(name="ping")
async def ping(**kwargs):
    """ """
    return jsonify(data="pong", msg=i18n.translate("STATUS_OK"))
