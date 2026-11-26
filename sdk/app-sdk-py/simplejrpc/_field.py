# -*- encoding: utf-8 -*-

from abc import ABCMeta, abstractmethod
from typing import Any, Dict, List, Tuple

from simplejrpc import exceptions
from simplejrpc.field import BaseField, Field


# +--------------------------------------------------
# Field function enhancement
# Beware of unfamiliar advice
# +--------------------------------------------------
# Decorator for applying type checking
def Typed(expected_type, cls=None):
    """ Performance is better than pure inheritance

    # --------------------------------
    Usage example:

    class IntegerField(Field[int]):
        """ """
        expected_type = int

    # Custom Unsigned :Reference Typed
    # Decorator for unsigned values
    def Unsigned(cls):
        super_set = cls.__set__

        def __set__(self, instance, value):
            if value < 0:
                raise ValueError('Expected >= 0')
            super_set(self, instance, value)

        cls.__set__ = __set__
        return cls

    @Unsigned
    class UnsignedInteger(IntegerField):
        pass

    # Use
    class Personal:
        """ """
        age = UnsignedInteger()
    # --------------------------------
    """
    if cls is None:
        return lambda cls: Typed(expected_type, cls)
    super_set = cls.__set__

    def __set__(self, instance, value):
        """ """
        if not isinstance(value, expected_type):
            raise exceptions.TypeError(
                f"Field {cls.name}, expect type {expected_type.__name__}"
            )
        super_set(self, instance, value)

    cls.__set__ = __set__
    return cls


# +--------------------------------------------------
# Form of field
# +--------------------------------------------------
class FieldBase:
    """ """

    def __init__(self, name: str, field: BaseField, value: Any = None) -> None:
        """ """
        self._name = name
        self._field = field
        self._value = value

    def get_name(self):
        """ """
        return self._name

    def get_value(self):
        """ """
        return self._value

    def get_field(self):
        """ """
        return self._field

    def get_field_validator(self):
        """ """
        return self._field.valid

    def get_field_label(self):
        return self._field.label

    @abstractmethod
    def accept(self, visitor: "FieldRuler"): ...


class SimpleFieldRule(FieldBase):
    """ """

    def accept(self, visitor: "FieldRuler"):
        return visitor.check(self)


# visitor
class FieldRuler:
    """ """

    RULE: str

    def __init__(self, controller: "RuleControllerBase") -> None:
        """ """
        self._controller = controller

    @abstractmethod
    def check(self, rule: FieldBase):
        """ """

    def get_kwds(self):
        """ """
        return self._controller.get_kwds()

    def express_eval(self, field_info_list: List):
        """ """
        kwd = {}
        for i in range(0, len(field_info_list), 2):
            first_value = field_info_list[i]
            second_value = (
                field_info_list[i + 1] if i + 1 < len(field_info_list) else None
            )
            if not second_value:
                continue
            kwd[first_value] = second_value
        return kwd

    def make_condition(self, rule: FieldBase) -> Dict[str, Any] | List[Any]:
        """ """
        validator = rule.get_field_validator()
        _, field_infos = validator.split(":")
        field_info_list = field_infos.split(",")
        kwd = self.express_eval(field_info_list)
        return kwd


class DefaultRuler(FieldRuler):
    """ """

    def check(self, rule: FieldBase): ...


class RequiredRuler(FieldRuler):
    """ """

    def check(self, rule: FieldBase):
        return


class RequiredIfRuler(FieldRuler):
    """
    Rule: required if: field, value
    Description: Required parameter (when any given field value is equal to the given value, i.e. when the value of the field field is value, the current validation field is a required parameter). Multiple fields are separated by asterisks.
    Example: When the Gender field is 1, the WifeName field must not be empty, and when the Gender field is 2, the HusbandName field must not be empty.
    """

    RULE: str = "required-if"

    # required-if:gender,1
    def check(self, rule: FieldBase):
        """ """
        name = rule.get_name()
        label = rule.get_field_label()
        kwd_cond = self.make_condition(rule)
        kwds = self.get_kwds()

        flag = False
        for _k, _v in kwd_cond.items():
            """ """
            if _k not in kwds:
                continue
            if str(kwds[_k]) != str(_v):
                continue
            flag = True

        if flag and not kwds.get(name):
            """ """
            raise exceptions.TypeError(f"Please enter a valid {label or name} value")


class RequiredUnlessRuler(FieldRuler):
    """
    Rule: required uncles: field, value
    Description: Required parameter (when the given field value is not equal to the given value, i.e. when the value of the field field is not value, the current validation field is a required parameter). Multiple fields are separated by asterisks.
    Example: When Gender is not equal to 0 and Gender is not equal to 2, WifeName must not be empty; When Id is not equal to 0 and Gender is not equal to 2, HusbandName must not be empty.
    """

    RULE: str = "required-unless"

    def check(self, rule: FieldBase):

        name = rule.get_name()
        label = rule.get_field_label()
        kwd_cond = self.make_condition(rule)
        kwds = self.get_kwds()

        flag = False
        for _k, _v in kwd_cond.items():
            """ """
            if _k not in kwds:
                continue
            if str(kwds[_k]) == str(_v):
                continue
            flag = True

        if flag and not kwds.get(name):
            """ """
            raise exceptions.TypeError(f"Please enter a valid {label or name} value")


class RequiredWithRuler(FieldRuler):
    """
    Rule: required with: field1, field2
    Description: Required parameter (when one of the given field values is not empty).
    Example: When WifeName is not empty, HusbandName must not be empty.
    """

    RULE: str = "required-with"

    # required-with:WifeName

    def make_condition(self, rule):
        """ """
        validator = rule.get_field_validator()
        _, field_infos = validator.split(":")
        field_info_list = field_infos.split(",")
        return field_info_list

    def check(self, rule: FieldBase):
        """ """

        name = rule.get_name()
        label = rule.get_field_label()
        kwd_cond = self.make_condition(rule)
        kwds = self.get_kwds()

        flag = False
        for _v in kwd_cond:
            """ """
            if _v in kwds and kwds.get(_v) in [None, ""]:
                continue
            flag = True

        if flag and not kwds.get(name):
            """ """
            raise exceptions.TypeError(f"Please enter a valid {label or name} value")


class RequiredWithoutRuler(FieldRuler):
    """
    Rule: required without: field1, field2,...
    Description: Required parameter (when one of the given field values is empty).
    Example: When Id or WifeName is empty, HusbandName must not be empty.
    """

    # required-without:WifeName
    RULE: str = "required-without"

    def make_condition(self, rule):
        """ """
        validator = rule.get_field_validator()
        _, field_infos = validator.split(":")
        field_info_list = field_infos.split(",")
        return field_info_list

    def check(self, rule: FieldBase):

        name = rule.get_name()
        label = rule.get_field_label()
        kwd_cond = self.make_condition(rule)
        kwds = self.get_kwds()

        flag = False
        for _v in kwd_cond:
            """ """
            if _v in kwds and kwds.get(_v) not in [None, ""]:
                continue
            flag = True

        if flag and not kwds.get(name):
            """ """
            raise exceptions.TypeError(f"Please enter a valid {label or name} value")


class BaseRuleBuilder:
    """ """


class DefaultRuleBuilder(BaseRuleBuilder):
    """ """

    def __init__(self, rule: FieldBase) -> None:
        """ """
        self.rules = []
        self._rule = rule

    def then(self, ruler: FieldRuler):
        """ """
        self.rules.append(ruler)

    def when(self, rule: FieldBase):
        """ """
        return DefaultRuleBuilder(rule)

    def combine(self, rule_builder: BaseRuleBuilder):
        """ """
        self.rules.extend(rule_builder.rules)

    def evaluate(self):
        """ """
        for rule in self.rules:
            rule.check(self._rule)


class RuleControllerBase(metaclass=ABCMeta):
    """ """

    @abstractmethod
    def get_kwds(self) -> Dict[str, Any]: ...

    @abstractmethod
    def get_fields(self) -> List[Tuple[str, Field]]: ...

    @abstractmethod
    def evaluate(self): ...


class SimpleRuleController(RuleControllerBase):
    """ """

    RULE_OP = {
        "required": RequiredRuler,
        "required_if": RequiredIfRuler,
        "required_unless": RequiredUnlessRuler,
        "required_without": RequiredWithoutRuler,
        "required_with": RequiredWithRuler,
        "default": DefaultRuler,
    }

    def __init__(self, kwds: Dict[str, Any], fields: List[Tuple[str, Field]]) -> None:
        """ """
        self._kwds = kwds
        self._fields = fields

    def get_kwds(self):
        """ """
        return self._kwds

    def get_fields(self):
        """ """
        return self._fields

    def attach_op(self, express: Dict[str, Any]):
        """ """
        self.RULE_OP.update(express)

    def build(self):
        """ """
        for name, field in self._fields:
            value = self._kwds.get(name)
            rule_builder = DefaultRuleBuilder(SimpleFieldRule(name, field, value))
            validator_prefix = field.valid.split(":", maxsplit=1)[0]
            rule_operator = self.RULE_OP.get(
                validator_prefix.replace("-", "_"), DefaultRuler
            )
            rule_builder.then(rule_operator(self))
            yield rule_builder

    def evaluate(self):
        """ """
        for builder in self.build():
            builder.evaluate()
