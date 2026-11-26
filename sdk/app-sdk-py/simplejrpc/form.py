# -*- encoding: utf-8 -*-
import inspect
from collections import ChainMap
from copy import deepcopy
from typing import Any, Dict, List, NoReturn, Type

from simplejrpc import exceptions
from simplejrpc._field import RuleControllerBase, SimpleRuleController
from simplejrpc.field import BaseField
from simplejrpc.func import make_signature, order_dict

_KT = Type[str]
_VT = Type[Any]


# +--------------------------------------------------
# Form
# +--------------------------------------------------
class from_meta(type):
    """ """

    def __new__(cls, classname, bases, methods: Dict[str, Any]):
        """
        The function dynamically creates a new class with a custom signature based on the fields defined in
        the class methods.

        :param cls: The `cls` parameter refers to the class object that is being created. It is a reference
        to the metaclass itself
        :param classname: The `classname` parameter is the name of the class being created
        :param bases: The `bases` parameter is a tuple of the base classes that the new class inherits from
        :param methods: The `methods` parameter is a dictionary that contains the methods and attributes of
        the class being created. Each key-value pair in the dictionary represents a method or attribute,
        where the key is the name of the method or attribute, and the value is the actual method or
        attribute object
        :type methods: Dict[str, Any]
        :return: The `__new__` method is returning a new instance of the class.
        """

        parent_attributes = {}
        for base in bases:
            parent_attributes |= base.__dict__

        for attr in parent_attributes:
            """Get the attribute"""
            if attr in methods:
                continue
            methods[attr] = parent_attributes[attr]

        _fields, fields_, _kw_default = [], [], {}
        for key, value in methods.items():
            if isinstance(value, BaseField):
                fields_.append((key, value))
                _fields.append((key, value.required))
                value.name = key
                if value.default is None:
                    continue
                _kw_default[key] = value.default
        methods["__kw_default__"] = _kw_default
        methods["__fields__"] = fields_
        methods["__signature__"] = make_signature(_fields)
        return super().__new__(cls, classname, bases, methods)

    def __call__(cls, *args: Any, **kwds: Any) -> Any:
        """New verification rule extension"""

        obj = object.__new__(cls)
        cls.__init__(obj, *args, **kwds)
        rule_cls = SimpleRuleController
        if hasattr(obj, "rule_adapter"):
            """ """
            rule_cls = (
                obj.rule_adapter
                if obj.rule_adapter is not None
                else SimpleRuleController
            )
        src = rule_cls(kwds=kwds, fields=obj.__fields__)
        src.evaluate()
        return obj


class Form(metaclass=from_meta):
    """
    The above code defines a class with methods for initializing an object, representing the object as a
    string, checking if the object is valid, and raising exceptions if the object is not valid.

    Use the parameters from form data as input parameters
    """

    rule_adapter: "RuleControllerBase" = None

    SUPER_NUM = 2

    def __init__(self, *args, **kw) -> None:
        """The function initializes an object by binding arguments to its attributes."""
        self.errors = []
        self.code = None
        kw_cp = deepcopy(kw)

        if is_super := self.__is_bases():
            __sig_args = self.__get_bases_sig_args()
        else:
            __sig_args = self.__get_sig_args()
        kwargs = {k: kw_cp[k] for k in kw_cp if k in __sig_args}

        self._valid_attr = {}
        try:
            bound_values = self.__signature__.bind(
                *args, **ChainMap(order_dict(kwargs), self.__kw_default__)
            )
        except TypeError as e:
            raise exceptions.ValidationError(str(e)) from e
        for name, value in bound_values.arguments.items():
            setattr(self, name, value)
            self._valid_attr[name] = self.__dict__.get("files") or value

    def __get_sig_args(self) -> Type[set]:
        """
        The function __get_sig_args returns a set of all the arguments and keyword-only arguments of the
        current class.
        :return: a set of all the arguments and keyword-only arguments of the class method.
        """

        __class = type(self)
        prepare_bind_args = inspect.getfullargspec(__class)
        return set(prepare_bind_args.args.__add__(prepare_bind_args.kwonlyargs))

    def __get_bases_sig_args(self) -> Type[set]:
        """
        The function `__get_bases_sig_args` returns a set of all the argument names from the base classes of
        the current class.
        :return: a set of signature arguments (__sig_args) that are obtained by iterating over the base
        classes of the current class (__bases) and adding the arguments from each base class to the set.
        """
        __bases = inspect.getmro(self.__class__)
        __sig_args = set()
        for __cls in __bases:
            """ """
            if __cls.__name__ in [Form.__name__, object.__name__]:
                continue
            __bind_args = inspect.getfullargspec(__cls)
            __sig_args = __sig_args.union(
                set(__bind_args.args.__add__(__bind_args.kwonlyargs))
            )

        return __sig_args

    def __is_bases(self) -> bool:
        """
        The function checks if the number of base classes of an object is greater than a specified number.
        :return: a boolean value indicating whether the number of classes in the method resolution order
        (MRO) of the current class is greater than the value of `self.SUPER_NUM`.
        """

        return len(inspect.getmro(self.__class__)) > self.SUPER_NUM

    def __repr__(self) -> str:
        """ """
        f_str = "".join(
            f"{__name}={self.__dict__.get(__name)},"
            for __name in self.__signature__.parameters.keys()
        )
        return f"<(class={self.__class__.__name__}, args:{f_str})>"

    def is_valid(self) -> bool:
        """
        The function `is_valid` returns True if there are no errors, and the function `get_errors` returns
        the value of the attribute "errors".
        :return: The method `get_errors` is returning the value of the attribute "errors".
        """
        return not self.errors

    def get_errors(self) -> List[Any]:
        """
        The function `get_errors` returns the errors attribute of the object.
        :return: The method is returning the value of the attribute "errors".
        """
        return self.errors

    def raise_valid(self) -> NoReturn:
        """
        The function raises a FormValidateError with the error message for each error in the self.errors
        list.
        """
        for e in self.errors:
            if self.code is not None:
                raise exceptions.ValidationError(str(e), code=self.code)
            raise exceptions.ValidationError(str(e))

    def raise_all_errors(self) -> NoReturn:
        """
        The function `raise_all_errors` raises a `FormValidateError` exception if the form is not valid,
        with the error messages joined together.
        """
        err_msg = ""
        if not self.is_valid():
            for error in self.errors:
                if isinstance(error, Dict):
                    """ """
                    # err_msg += ",".join(reduce(lambda x, y: x + y, [x.get("msg", "") for x in error.values()]))
                    err_msg1 = error.get("msg", "")

                    if err_msg == "":
                        err_msg = err_msg1
                    elif err_msg1 not in err_msg:
                        err_msg += f"#{err_msg1}"
                else:
                    if error not in err_msg:
                        err_msg += f"#{error}"
                if err_msg:
                    break
        if bool(err_msg):
            if self.code is not None:
                raise exceptions.ValidationError(err_msg, code=self.code, data=None)
            raise exceptions.ValidationError(err_msg, data=None)

    def form_data(self) -> Dict[_KT, _VT]:
        """Retrieve the valid form data.

        Returns:
            Any: The valid form data.

        """
        return self._valid_attr

    def get_attrs(self) -> Dict[_KT, _VT]:
        """Retrieve the valid form data.

        Returns:
            Any: The valid form data.

        """
        return self._valid_attr

    def raise_valid_error(self):
        """ """
        raise NotImplementedError
