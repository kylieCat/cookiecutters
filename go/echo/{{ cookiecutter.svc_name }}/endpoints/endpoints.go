package endpoints

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/metrics"
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

func instrumentRequest(ctx context.Context, status int) {
	metrics.ApiCalls.With(prometheus.Labels{
		"endpoint":      ctx.Value("endpoint").(string),
		"protocol":      ctx.Value("protocol").(string),
		"response_code": strconv.Itoa(status),
	}).Observe(float64(time.Since(ctx.Value("start-time").(time.Time)) / time.Millisecond))
}


func Init() []Endpoint {
	endpoints := []Endpoint {
		{"/hc", http.MethodGet, HC},
		{"/hello", http.MethodGet, HelloWorld},
	}
	return endpoints
}
