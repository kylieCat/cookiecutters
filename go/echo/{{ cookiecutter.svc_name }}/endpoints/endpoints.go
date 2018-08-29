package endpoints

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/config"
)

type Endpoint struct {
	Url string
	Method string
	Handler echo.HandlerFunc
}

type Response struct {
	Message string
	Status  int
}

func Init() []Endpoint {
	conf := config.GetConfig()
	endpoints := []Endpoint {
		{"/hc", http.MethodGet, HC},
		{"/hello", http.MethodGet, HelloWorld},
		{conf.GetString("metrics.endpoint"), http.MethodGet, echo.WrapHandler(promhttp.Handler())},
	}
	return endpoints
}
