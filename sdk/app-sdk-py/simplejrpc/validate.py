# -*- encoding: utf-8 -*-
import re
from typing import Generic, NoReturn, Type, TypeVar, Union

from simplejrpc import exceptions
from simplejrpc._types import _BaseForm
from simplejrpc._types import _NoValue as NoValue
from simplejrpc.field import StringField
from simplejrpc.form import Form
from simplejrpc.i18n import Language
from simplejrpc.i18n import T as i18n

T = Type[Union[str, int, float, bool]]
F = TypeVar("F", bound=_BaseForm)


# +--------------------------------------------------
# Validator
# +--------------------------------------------------
class ValidatorBase:
    """ """

    require = False

    def __init__(self, err_msg: str = "", code: int | None = None) -> None:
        """ """
        self.name: str = NoValue
        self.label: str = ""
        self.value: Type[T] = NoValue
        self.err_msg: str = err_msg
        self.instance: Type[F] = NoValue
        self.code = code

    def raise_except(self, e=None) -> NoReturn:
        """
        The function raises an exception based on the value of the err_msg attribute.
        """
        if self.err_msg:
            if self.code is not None:
                raise exceptions.ValidationError(self.err_msg, code=self.code)
            raise exceptions.ValidationError(self.err_msg)
        elif e is not None:
            raise e
        else:
            msg = f"Please enter a valid {self.label or self.name} value"
            if self.code is not None:
                raise exceptions.AttributeError(msg, code=self.code)
            raise exceptions.AttributeError(msg)

    def clean(self, instance: Type[F]) -> NoReturn:
        """
        The function "clean" attempts to clean data, raises a FormValidateError if validation fails, and
        raises any other exceptions.
        """
        try:
            self.clean_data(instance)
        except exceptions.ValidationError as e:
            raise e
        except Exception as e:
            self.raise_except(e)

    # @abstractmethod
    def clean_data(self, instance: Type[F]):
        """
        The function clean_data is not implemented yet.

        If the parameter does not meet the condition,
        please use exception thrown,
        will automatically catch processing
        """
        raise NotImplementedError


class Validator(ValidatorBase):
    """ """

    """
    The function updates an attribute of an instance with a new value.

    :param value: The parameter "value" in the above code is of type "Type[T]". This means that it can
    accept any type of value, but it must be a subclass of type "T"
    :type value: Type[T]
    """

    def update_attr(self, value: Type[T]):
        """
        Updates the attribute value of the instance.

        Args:
            value: The new value to be assigned to the attribute.

        Returns:
            None

        Examples:
            # Update the attribute value of the instance
            >>> update_attr(10)
        """
        self.instance.__dict__[self.name] = value
        self.value = value


class BooleanValidator(Validator):
    """
    The BooleanValidator class is a subclass of ValidatorBase that validates a boolean value and raises
    an exception if the value is not true.
    """

    def clean(self, instance: Type[F]) -> NoReturn:
        """
        The function "clean" attempts to clean data, raises an exception if an error occurs, and raises an
        exception if the cleaned data is empty.
        """
        try:
            value = self.clean_data(instance)
        except Exception as e:
            self.raise_except(e)
        if not value:
            self.raise_except()


class RequireValidator(Validator):
    """
    The `RequireValidator` class is a subclass of `ValidatorBase` that checks if a value is empty or
    None and raises an exception if it is.
    """

    NULL_VALUE = ["", None, [], {}]
    require = True

    def clean_data(self, instance: Type[F]) -> NoReturn:
        """
        The function `clean_data` checks if a value is empty or None and raises an exception if it is.
        """
        if self.value in self.NULL_VALUE:
            if self.err_msg:
                if self.code is not None:
                    raise exceptions.ValidationError(
                        f"{self.err_msg} : {self.label or self.name}", code=self.code
                    )
                raise exceptions.ValidationError(
                    f"{self.err_msg} : {self.label or self.name}"
                )

            msg = f"Please enter a valid {self.label or self.name} value"
            if self.code is not None:
                raise exceptions.ValidationError(
                    message=msg,
                    code=self.code,
                )
            raise exceptions.ValidationError(msg)


class StrictPasswordValidator(Validator):
    """ """

    PASSWD_VALIDATE_LENGTH = 8

    def __init__(self, err_msg="", length=8) -> None:
        """ """
        self.name = NoValue
        self.value = NoValue
        self.err_msg = err_msg
        self.instance = NoValue
        self.length = self.PASSWD_VALIDATE_LENGTH or length

    def clean_data(self, instance):
        """
        Clean the data by validating the password input.

        Args:
            instance: The instance of the data to be cleaned.

        Raises:
            FormValidateError: If the password input is invalid, such as empty, too short, not starting with a letter,
            not containing at least three character types, or not meeting the specified criteria.
        """

        _err_message = f"Check whether the password is valid: it must start with a letter and contain at least {self.length} digits and uppercase letters."
        if self.value == "" or self.value is None:
            raise exceptions.ValidationError(_err_message)
        if len(self.value) < self.length:
            raise exceptions.ValidationError(_err_message)
        if not self.value[0].isalpha():
            raise exceptions.ValidationError(_err_message)
        char_types = 0
        if any(c.isupper() for c in self.value):
            char_types += 1
        if any(c.islower() for c in self.value):
            char_types += 1
        if any(c.isdigit() for c in self.value):
            char_types += 1
        if any(not c.isalnum() for c in self.value):
            char_types += 1
        if char_types < 3:  # 满足四分之三原则
            raise exceptions.ValidationError(_err_message)


class NameValidator(Validator):
    """ """

    def clean_data(self, instance):
        if not bool(re.fullmatch(r"[a-zA-Z_][a-zA-Z0-9_]*", self.value)):
            tmp = f"lease enter valid names, including digits and underscores, starting with a letter :{self.label or self.name}"
            raise exceptions.ValidationError(tmp)


class StringLangValidator(Validator):
    """ """

    def clean_data(self, instance):
        values = Language.values()
        if self.value not in values:
            tmp = f"Please enter a valid language code, such as {values} :{self.label or self.name}"
            raise exceptions.ValidationError(tmp)
        i18n.set_lang(self.value)


# The `BaseForm` class is a subclass of `MetaBase`.
class BaseForm(Form, Generic[F]):
    """ """

    lang = StringField(validators=[StringLangValidator()])
