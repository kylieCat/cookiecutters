import falcon
import json


class HealthCheckResource(object):
    """
    A basic health check endpoint.
    """
    @staticmethod
    def on_get(request, response):
        """
        Simple health check that returns an OK response if the service is healthy.

        Args:
            request (falcon.Request):
            response (falcon.Response):
        """
        response.status = falcon.HTTP_200
        response.data = json.dumps({"status": "OK"}).encode()
