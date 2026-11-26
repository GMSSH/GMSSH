# -*- encoding: utf-8 -*-

from abc import abstractmethod
from typing import Any, Dict, List, NoReturn


class _Validator:
    """ """

    @abstractmethod
    def update_attr(self, value: Any) -> None:
        """
        The function `update_attr` updates the `value` attribute of the instance with the provided value.
        :param value: The value to update the attribute with.
        :type value: Any
        """

    @abstractmethod
    def raise_except(self, e=None) -> NoReturn:
        """
        The function `raise_except` raises an exception with the provided error message.
        :param e: The error message to be raised as an exception.
        :type e: Exception
        :return: None
        """

    @abstractmethod
    def clean(self, instance: Any) -> NoReturn:
        """
        The function `clean` attempts to clean the data of the instance and raises an exception if an
        error occurs.
        :param instance: The instance to be cleaned.
        :type instance: Any
        :return: None
        """

    @abstractmethod
    def clean_data(self, instance: Any) -> Any:
        """
        The function `clean_data` cleans the data of the instance and returns the cleaned data.
        :param instance: The instance to be cleaned.
        :type instance: Any
        :return: The cleaned data.
        """


# Getting the file creation time This class is used to verify that the data is valid
class _NoValue: ...


class _BaseForm:

    @abstractmethod
    def is_valid(self) -> bool:
        """
        The function `is_valid` checks if the form is valid by calling the `validate` method and
        returning the result.
        """

    @abstractmethod
    def get_errors(self) -> List[Any]:
        """
        The function `get_errors` returns the errors attribute of the object.
        :return: The method is returning the value of the attribute "errors".
        """

    @abstractmethod
    def raise_valid(self) -> None:
        """
        The function raises a FormValidateError with the error message for each error in the self.errors
        list.
        """

    @abstractmethod
    def raise_all_errors(self) -> None:
        """
        The function `raise_all_errors` raises a `FormValidateError` exception if the form is not valid,
        with the error messages joined together.
        """

    @abstractmethod
    def form_data(self) -> Dict[str, Any]:
        """Retrieve the valid form data.

        Returns:
            Any: The valid form data.
        """

    @abstractmethod
    def get_attrs(self) -> Dict[str, Any]:
        """Retrieve the valid form data.

        Returns:
            Any: The valid form data.
        """

    def raise_valid_error(self):
        """ """
        raise NotImplementedError
