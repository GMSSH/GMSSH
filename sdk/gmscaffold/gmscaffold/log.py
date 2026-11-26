# -*- encoding: utf-8 -*-
"""
@文件        :command.py
@说明        :
@时间        :2024/11/26 14:57:46
@作者        :Zack
@版本        :1.0
"""

import logging
from functools import partial

_LOG_FORMAT = "%(asctime)s - %(name)s - %(levelname)s - %(message)s"


class LoggerWriter:
    """ """

    def __init__(self, logger, level):
        """ """
        self.logger = logger
        self.level = level

    def write(self, message):
        if message != "\n":  # avoid printing bare newlines, if you like
            self.logger.log(self.level, message)

    def flush(self):
        # doesn't actually do anything, but might be expected of a file-like
        # object - so optional depending on your situation
        pass

    def close(self):
        # doesn't actually do anything, but might be expected of a file-like
        # object - so optional depending on your situation. You might want
        # to set a flag so that later calls to write raise an exception
        pass


class Logger:
    """ """

    def __init__(self, level=logging.DEBUG, **kwargs):
        """
        Do basic configuration for the logging system.

        This function does nothing if the root logger already has handlers
        configured, unless the keyword argument *force* is set to ``True``.
        It is a convenience method intended for use by simple scripts
        to do one-shot configuration of the logging package.

        The default behaviour is to create a StreamHandler which writes to
        sys.stderr, set a formatter using the BASIC_FORMAT format string, and
        add the handler to the root logger.

        A number of optional keyword arguments may be specified, which can alter
        the default behaviour.

        filename  Specifies that a FileHandler be created, using the specified
                filename, rather than a StreamHandler.
        filemode  Specifies the mode to open the file, if filename is specified
                (if filemode is unspecified, it defaults to 'a').
        format    Use the specified format string for the handler.
        datefmt   Use the specified date/time format.
        style     If a format string is specified, use this to specify the
                type of format string (possible values '%', '{', '$', for
                %-formatting, :meth:`str.format` and :class:`string.Template`
                - defaults to '%').
        level     Set the root logger level to the specified level.
        stream    Use the specified stream to initialize the StreamHandler. Note
                that this argument is incompatible with 'filename' - if both
                are present, 'stream' is ignored.
        handlers  If specified, this should be an iterable of already created
                handlers, which will be added to the root handler. Any handler
                in the list which does not have a formatter assigned will be
                assigned the formatter created in this function.
        force     If this keyword  is specified as true, any existing handlers
                attached to the root logger are removed and closed, before
                carrying out the configuration as specified by the other
                arguments.
        encoding  If specified together with a filename, this encoding is passed to
                the created FileHandler, causing it to be used when the file is
                opened.
        errors    If specified together with a filename, this value is passed to the
                created FileHandler, causing it to be used when the file is
                opened in text mode. If not specified, the default value is
                `backslashreplace`.

        Note that you could specify a stream created using open(filename, mode)
        rather than passing the filename and mode in. However, it should be
        remembered that StreamHandler does not close its stream (since it may be
        using sys.stdout or sys.stderr), whereas FileHandler closes its stream
        when the handler is closed.

        .. versionchanged:: 3.2
        Added the ``style`` parameter.

        .. versionchanged:: 3.3
        Added the ``handlers`` parameter. A ``ValueError`` is now thrown for
        incompatible arguments (e.g. ``handlers`` specified together with
        ``filename``/``filemode``, or ``filename``/``filemode`` specified
        together with ``stream``, or ``handlers`` specified together with
        ``stream``.

        .. versionchanged:: 3.8
        Added the ``force`` parameter.

        .. versionchanged:: 3.9
        Added the ``encoding`` and ``errors`` parameters.
        """
        logging.basicConfig(level=level)
        self.logger = logging.getLogger("GmScaffold")

    @classmethod
    def from_log_file(cls, log_file: str, level=logging.DEBUG, formatter=None):
        """ """
        fh = logging.FileHandler(log_file)
        formatter = logging.Formatter(formatter or _LOG_FORMAT)
        fh.setFormatter(formatter)
        fh.setLevel(level)
        logger = cls(level)
        logger.add_handle(fh)
        return logger

    def add_handle(self, fh):
        """ """
        self.logger.addHandler(fh)

    def io_logger(self, logger, mode) -> LoggerWriter:
        """ """
        return LoggerWriter(logger, mode)

    def io_partial(self, mode):
        """ """
        return partial(self.io_logger, logger=self.logger, mode=mode)

    def debug(self, message: str):
        """ """
        print(message, file=self.io_logger(logger=self.logger, mode=logging.DEBUG))

    def info(self, message: str):
        """ """
        print(message, file=self.io_logger(logger=self.logger, mode=logging.INFO))

    def warn(self, message: str):
        """ """
        print(message, file=self.io_logger(logger=self.logger, mode=logging.WARN))

    def error(self, message: str):
        """ """
        print(message, file=self.io_logger(logger=self.logger, mode=logging.ERROR))

    def critical(self, message: str):
        """ """
        print(message, file=self.io_logger(logger=self.logger, mode=logging.CRITICAL))


logger = Logger()
