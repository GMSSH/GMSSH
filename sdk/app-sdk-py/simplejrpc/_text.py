# -*- encoding: utf-8 -*-
import json

from loguru import logger

from simplejrpc.i18n import T as i18n


# Define the internationalized error message object
class ErrorTextMessage:
    """ """

    def __init__(self, message: str, *args, translate=None):
        """ """
        self.message = message
        self.args = args
        self.translate = translate

    def __str__(self):
        """ """
        try:
            if self.message:
                if self.translate:
                    return self.translate(self.message)
                if not self.args:
                    return i18n.translate(self.message)
                return i18n.translate_ctx(self.message, *self.args)
        except Exception as e:
            logger.error(f"ErrorTextMessage: {e}")
            return self.message
        return self.message

    def __repr__(self):
        """ """
        return f"<ErrorTextMessage: {self.message}>"

    def __eq__(self, value):
        return super().__eq__(value)

    def concat(self, value):
        """ """
        if isinstance(value, str):
            return value + str(self)
        elif isinstance(value, ErrorTextMessage):
            return str(value) + str(self)
        return self

    def __iadd__(self, value):
        """ """
        return self.concat(value)

    def __radd__(self, value):
        """ """
        return self.concat(value)

    def __add__(self, value):
        """ """
        return self.concat(value)


class TextMessageDecoder(json.JSONEncoder):
    """ """

    def default(self, o):
        if isinstance(o, ErrorTextMessage):
            return str(o)
        return o


def json_dumps(obj, **kwargs):
    """ """
    return json.dumps(obj, cls=TextMessageDecoder, **kwargs)
