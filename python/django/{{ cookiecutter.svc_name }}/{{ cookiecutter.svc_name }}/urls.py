from django.contrib import admin
from django.urls import path, url

from .health_check import HealthCheckView

urlpatterns = [
    path('admin/', admin.site.urls),
    url("hc/", HealthCheckView.as_view(), name="health_check")
]
