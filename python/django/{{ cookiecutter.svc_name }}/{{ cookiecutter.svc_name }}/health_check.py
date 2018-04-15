from http import HTTPStatus
import json

from django.http import HttpResponse
from django.views import View


class HealthCheckView(View):
    def get(self, request):
        return HttpResponse(
            status=HTTPStatus.OK,
            content=json.dumps({"message": "OK!", "status": HTTPStatus.OK})
        )
