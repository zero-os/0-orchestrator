"""
Auto-generated class for HTTPProxy
"""
from .HTTPType import HTTPType

from . import client_support


class HTTPProxy(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(destinations, host, types):
        """
        :type destinations: list[str]
        :type host: str
        :type types: list[HTTPType]
        :rtype: HTTPProxy
        """

        return HTTPProxy(
            destinations=destinations,
            host=host,
            types=types,
        )

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'HTTPProxy'
        create_error = '{cls}: unable to create {prop} from value: {val}: {err}'
        required_error = '{cls}: missing required property {prop}'

        data = json or kwargs

        property_name = 'destinations'
        val = data.get(property_name)
        if val is not None:
            datatypes = [str]
            try:
                self.destinations = client_support.list_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'host'
        val = data.get(property_name)
        if val is not None:
            datatypes = [str]
            try:
                self.host = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'types'
        val = data.get(property_name)
        if val is not None:
            datatypes = [HTTPType]
            try:
                self.types = client_support.list_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
