from typing import Any

from gmscaffold.dispatcher import dispatcher


class SignalManager:
    """ """

    _instance = None

    def __new__(cls, *args, **kwargs):
        """ """
        if cls._instance is None:
            cls._instance = super().__new__(cls, *args, **kwargs)
        return cls._instance

    def __init__(self, sender: Any = dispatcher.Anonymous):
        """ """
        self.sender: Any = sender

    def connect(self, receiver: Any, signal: Any, **kwargs: Any) -> None:
        """
        Connect a receiver function to a signal.

        The signal can be any object, although Scrapy comes with some
        predefined signals that are documented in the :ref:`topics-signals`
        section.

        :param receiver: the function to be connected
        :type receiver: collections.abc.Callable

        :param signal: the signal to connect to
        :type signal: object
        """
        kwargs.setdefault("sender", self.sender)
        dispatcher.connect(receiver, signal, **kwargs)

    def disconnect(self, receiver: Any, signal: Any, **kwargs: Any) -> None:
        """
        Disconnect a receiver function from a signal. This has the
        opposite effect of the :meth:`connect` method, and the arguments
        are the same.
        """
        kwargs.setdefault("sender", self.sender)
        dispatcher.disconnect(receiver, signal, **kwargs)

    def disconnect_all(self, signal: Any, **kwargs: Any) -> None:
        """
        Disconnect all receivers from the given signal.

        :param signal: the signal to disconnect from
        :type signal: object
        """
        kwargs.setdefault("sender", self.sender)
        """Disconnect all signal handlers. Useful for cleaning up after running
        tests
        """
        for receiver in dispatcher.liveReceivers(dispatcher.getAllReceivers(self.sender, signal)):
            dispatcher.disconnect(receiver, signal=signal, sender=self.sender)

    def send(self, signal=Any, *arguments, **kwargs):
        """Send signal from sender to all connected receivers.

        signal -- (hashable) signal value, see connect for details

        sender -- the sender of the signal

            if Any, only receivers registered for Any will receive
            the message.

            if Anonymous, only receivers registered to receive
            messages from Anonymous or Any will receive the message

            Otherwise can be any python object (normally one
            registered with a connect if you actually want
            something to occur).

        arguments -- positional arguments which will be passed to
            *all* receivers. Note that this may raise TypeErrors
            if the receivers do not allow the particular arguments.
            Note also that arguments are applied before named
            arguments, so they should be used with care.

        named -- named arguments which will be filtered according
            to the parameters of the receivers to only provide those
            acceptable to the receiver.

        Return a list of tuple pairs [(receiver, response), ... ]

        if any receiver raises an error, the error propagates back
        through send, terminating the dispatch loop, so it is quite
        possible to not have all receivers called if a raises an
        error.
        """
        kwargs.setdefault("sender", self.sender)
        dispatcher.send(signal=signal, *arguments, **kwargs)
