"""
@文件        :__init__.py
@说明        :This is an example
@时间        :2025/06/30 09:17:23
@作者        :xxx
@邮箱        :
@版本        :1.0.0
"""

from simplejrpc import BaseForm,StringField,RequireValidator

from simplejrpc import TextMessage as _

class ExampleForm(BaseForm):
    """ """

    name = StringField(validators=[RequireValidator(_("REQUIRE_VALIDATION_TM"))])