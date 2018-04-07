from .utils import get_logger

logger = get_logger()


class ResponseLoggerMiddleware(object):
    """
    Middleware class  for request logging
    """
    def process_response(self, req, resp, resource, req_succeeded):
        """
        Request logger in the format of::

            2018-04-07 01:32:12,918,918 INFO     [middleware.py:8:p8] GET /hc 200
        """
        logger.info(f"{req.method} {req.relative_uri} {resp.status[:3]}")
