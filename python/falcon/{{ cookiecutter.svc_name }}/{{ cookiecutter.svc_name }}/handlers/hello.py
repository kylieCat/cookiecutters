import falcon
import json


class HelloResource(object):
    @staticmethod
    def on_get(request, response):
        name = request.get_param("name") or "World"
        response.status = falcon.HTTP_200
        response.data = json.dumps({"status": 200, "message": f"Hello, {name}"}).encode()
