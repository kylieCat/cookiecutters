from datetime import datetime
import logging
import sys

_logger = None


def get_logger():
    """
    Checks the global scope for an existing logger and creates one if it's ``None``.

    Creates a logger suitable for use with Docker/Kubernetes that logs to stdout
    and not to a file. Logs in the following format::

        2018-04-07 01:32:12,918,918 INFO     [middleware.py:8:p8] ...

    Returns:
        logging.Logger
    """
    global _logger
    if _logger is None:
        ch = logging.StreamHandler(sys.stdout)
        ch.setLevel(logging.INFO)
        log_format = "%(asctime)s,%(msecs)d %(levelname)-8s [%(filename)s:%(lineno)d:p%(process)s] %(message)s"
        logging.basicConfig(format=log_format, level=logging.INFO, handlers=[ch])
        _logger = logging.getLogger(__name__)
    return _logger


def get_log_message_from_grpc_metadata(md):
    """
    Parses a gRPC metadata object and returns a string suitable for logging.

    Args:
        md: A metadata object form a gRPC context.

    Returns:
        str
    """
    data = {k: v for k, v in md}
    endpoint = data["endpoint"]
    user_agent = data["user-agent"]
    caller = data["caller"]
    template = "endpoint: %s user-agent: %s response_code: [200] caller: %s"
    return template % (endpoint, user_agent, caller)


def get_timestamp(dt=None):
    """
    Get a Unix timestamp from the given ``dt``. If no argument is provided a
    Unix UTC timestamp for the current time will be returned.

    Args:
        dt (datetime.datetime):

    Returns:
        float
    """
    if dt is None:
        dt = datetime.utcnow()
    return dt.timestamp()
