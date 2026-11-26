from collections.abc import Mapping


class MappingObject(dict):
    """ """

    def __init__(self, *args, **kwargs):  # pylint: disable=super-init-not-called
        self.update(*args, **kwargs)

    def __getattr__(self, k):
        """ """
        try:
            return object.__getattribute__(self, k)
        except AttributeError:
            try:
                return self[k]
            except KeyError:
                raise AttributeError(k)

    def __setattr__(self, k, v):
        """ """
        try:
            # Throws exception if not in prototype chain
            object.__getattribute__(self, k)
        except AttributeError:
            try:
                self[k] = v
            except:
                raise AttributeError(k)
        else:
            object.__setattr__(self, k, v)

    def __delattr__(self, k):
        """ """
        try:
            # Throws exception if not in prototype chain
            object.__getattribute__(self, k)
        except AttributeError:
            try:
                del self[k]
            except KeyError:
                raise AttributeError(k)
        else:
            object.__delattr__(self, k)

    def toDict(self):
        """ """
        return unmunchify(self)

    @property
    def __dict__(self):
        return self.toDict()

    def __repr__(self):
        """Invertible* string-form of a MappingObject.

        >>> b = MappingObject(foo=MappingObject(lol=True), hello=42, ponies='are pretty!')
        >>> print (repr(b))
        MappingObject({'ponies': 'are pretty!', 'foo': MappingObject({'lol': True}), 'hello': 42})
        >>> eval(repr(b))
        MappingObject({'ponies': 'are pretty!', 'foo': MappingObject({'lol': True}), 'hello': 42})

        >>> with_spaces = MappingObject({1: 2, 'a b': 9, 'c': MappingObject({'simple': 5})})
        >>> print (repr(with_spaces))
        MappingObject({'a b': 9, 1: 2, 'c': MappingObject({'simple': 5})})
        >>> eval(repr(with_spaces))
        MappingObject({'a b': 9, 1: 2, 'c': MappingObject({'simple': 5})})

        (*) Invertible so long as collection contents are each repr-invertible.
        """
        return f"{self.__class__.__name__}({dict.__repr__(self)})"

    def __dir__(self):
        return list(self.keys())

    def __getstate__(self):
        """Implement a serializable interface used for pickling.

        See https://docs.python.org/3.6/library/pickle.html.
        """
        return {k: v for k, v in self.items()}

    def __setstate__(self, state):
        """Implement a serializable interface used for pickling.

        See https://docs.python.org/3.6/library/pickle.html.
        """
        self.clear()
        self.update(state)

    __members__ = __dir__  # for python2.x compatibility

    @classmethod
    def fromDict(cls, d):
        """Recursively transforms a dictionary into a MappingObject via copy.

        >>> b = MappingObject.fromDict({'urmom': {'sez': {'what': 'what'}}})
        >>> b.urmom.sez.what
        'what'

        See munchify for more info.
        """
        return munchify(d, cls)

    def copy(self):
        return type(self).fromDict(self)

    def update(self, *args, **kwargs):
        """
        Override built-in method to call custom __setitem__ method that may
        be defined in subclasses.
        """
        for k, v in dict(*args, **kwargs).items():
            self[k] = v

    def get(self, k, d=None):
        """
        D.get(k[,d]) -> D[k] if k in D, else d.  d defaults to None.
        """
        if k not in self:
            return d
        return self[k]

    def setdefault(self, k, d=None):
        """
        D.setdefault(k[,d]) -> D.get(k,d), also set D[k]=d if k not in D
        """
        if k not in self:
            self[k] = d
        return self[k]


def unmunchify(x):
    """Recursively converts a MappingObject into a dictionary.

    >>> b = MappingObject(foo=MappingObject(lol=True), hello=42, ponies='are pretty!')
    >>> sorted(unmunchify(b).items())
    [('foo', {'lol': True}), ('hello', 42), ('ponies', 'are pretty!')]

    unmunchify will handle intermediary dicts, lists and tuples (as well as
    their subclasses), but ymmv on custom datatypes.

    >>> b = MappingObject(foo=['bar', MappingObject(lol=True)], hello=42,
    ...         ponies=('are pretty!', MappingObject(lies='are trouble!')))
    >>> sorted(unmunchify(b).items()) #doctest: +NORMALIZE_WHITESPACE
    [('foo', ['bar', {'lol': True}]), ('hello', 42), ('ponies', ('are pretty!', {'lies': 'are trouble!'}))]

    nb. As dicts are not hashable, they cannot be nested in sets/frozensets.
    """

    # Munchify x, using `seen` to track object cycles
    seen = dict()

    def unmunchify_cycles(obj):
        # If we've already begun unmunchifying obj, just return the already-created unmunchified obj
        try:
            return seen[id(obj)]
        except KeyError:
            pass

        # Otherwise, first partly unmunchify obj (but without descending into any lists or dicts) and save that
        seen[id(obj)] = partial = pre_unmunchify(obj)
        # Then finish unmunchifying lists and dicts inside obj (reusing unmunchified obj if cycles are encountered)
        return post_unmunchify(partial, obj)

    def pre_unmunchify(obj):
        # Here we return a skeleton of unmunchified obj, which is enough to save for later (in case
        # we need to break cycles) but it needs to filled out in post_unmunchify
        if isinstance(obj, Mapping):
            return dict()
        elif isinstance(obj, list):
            return type(obj)()
        elif isinstance(obj, tuple):
            type_factory = getattr(obj, "_make", type(obj))
            return type_factory(unmunchify_cycles(item) for item in obj)
        else:
            return obj

    def post_unmunchify(partial, obj):
        # Here we finish unmunchifying the parts of obj that were deferred by pre_unmunchify because they
        # might be involved in a cycle
        if isinstance(obj, Mapping):
            partial.update((k, unmunchify_cycles(obj[k])) for k in obj.keys())
        elif isinstance(obj, list):
            partial.extend(unmunchify_cycles(v) for v in obj)
        elif isinstance(obj, tuple):
            for value_partial, value in zip(partial, obj):
                post_unmunchify(value_partial, value)

        return partial

    return unmunchify_cycles(x)


def munchify(x, factory=MappingObject):
    """Recursively transforms a dictionary into a Munch via copy.

    >>> b = munchify({'urmom': {'sez': {'what': 'what'}}})
    >>> b.urmom.sez.what
    'what'

    munchify can handle intermediary dicts, lists and tuples (as well as
    their subclasses), but ymmv on custom datatypes.

    >>> b = munchify({ 'lol': ('cats', {'hah':'i win again'}),
    ...         'hello': [{'french':'salut', 'german':'hallo'}] })
    >>> b.hello[0].french
    'salut'
    >>> b.lol[1].hah
    'i win again'

    nb. As dicts are not hashable, they cannot be nested in sets/frozensets.
    """
    # Munchify x, using `seen` to track object cycles
    seen = dict()

    def munchify_cycles(obj):
        partial, already_seen = pre_munchify_cycles(obj)
        if already_seen:
            return partial
        return post_munchify(partial, obj)

    def pre_munchify_cycles(obj):
        # If we've already begun munchifying obj, just return the already-created munchified obj
        try:
            return seen[id(obj)], True
        except KeyError:
            pass

        # Otherwise, first partly munchify obj (but without descending into any lists or dicts) and save that
        seen[id(obj)] = partial = pre_munchify(obj)
        return partial, False

    def pre_munchify(obj):
        # Here we return a skeleton of munchified obj, which is enough to save for later (in case
        # we need to break cycles) but it needs to filled out in post_munchify
        if isinstance(obj, Mapping):
            return factory({})
        elif isinstance(obj, list):
            return type(obj)()
        elif isinstance(obj, tuple):
            type_factory = getattr(obj, "_make", type(obj))
            return type_factory(pre_munchify_cycles(item)[0] for item in obj)
        else:
            return obj

    def post_munchify(partial, obj):
        # Here we finish munchifying the parts of obj that were deferred by pre_munchify because they
        # might be involved in a cycle
        if isinstance(obj, Mapping):
            partial.update((k, munchify_cycles(obj[k])) for k in obj.keys())
        elif isinstance(obj, list):
            partial.extend(munchify_cycles(item) for item in obj)
        elif isinstance(obj, tuple):
            for item_partial, item in zip(partial, obj):
                post_munchify(item_partial, item)

        return partial

    return munchify_cycles(x)


class DefaultMapping(MappingObject):
    """
    A MappingObject that returns a user-specified value for missing keys.
    """

    def __init__(self, *args, **kwargs):
        """Construct a new DefaultMunch. Like collections.defaultdict, the
        first argument is the default value; subsequent arguments are the
        same as those for dict.
        """
        # Mimic collections.defaultdict constructor
        if args:
            default = args[0]
            args = args[1:]
        else:
            default = None
        super().__init__(*args, **kwargs)
        self.__default__ = default

    def __getattr__(self, k):
        """Gets key if it exists, otherwise returns the default value."""
        try:
            return super().__getattr__(k)
        except AttributeError:
            return self.__default__

    def __setattr__(self, k, v):
        if k == "__default__":
            object.__setattr__(self, k, v)
        else:
            super().__setattr__(k, v)

    def __getitem__(self, k):
        """Gets key if it exists, otherwise returns the default value."""
        try:
            return super().__getitem__(k)
        except KeyError:
            return self.__default__

    def __getstate__(self):
        """Implement a serializable interface used for pickling.

        See https://docs.python.org/3.6/library/pickle.html.
        """
        return (self.__default__, {k: v for k, v in self.items()})

    def __setstate__(self, state):
        """Implement a serializable interface used for pickling.

        See https://docs.python.org/3.6/library/pickle.html.
        """
        self.clear()
        default, state_dict = state
        self.update(state_dict)
        self.__default__ = default

    @classmethod
    def fromDict(cls, d, default=None):
        # pylint: disable=arguments-differ
        return munchify(d, factory=lambda d_: cls(default, d_))

    def copy(self):
        return type(self).fromDict(self, default=self.__default__)

    def __repr__(self):
        return f"{type(self).__name__}({self.__undefined__!r}, {dict.__repr__(self)})"
