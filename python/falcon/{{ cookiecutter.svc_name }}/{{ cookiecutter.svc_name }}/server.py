from wsgiref import simple_server

import falcon

from . import handlers
from .middleware import ResponseLoggerMiddleware


app = falcon.API(middleware=[ResponseLoggerMiddleware()])

"""
Alert URL's
"""
app.add_route("/hello", handlers.HelloResource())

"""
Health Check
"""
app.add_route("/hc", handlers.HealthCheckResource())


if __name__ == "__main__":
    httpd = simple_server.make_server("0.0.0.0", {{ cookiecutter.port }}, app)
    httpd.serve_forever()
