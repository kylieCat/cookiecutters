import os

from ..common import Common


class Prod(Common):
    SECRET_KEY = os.environ.get("DJANGO_SECRET_KEY")
