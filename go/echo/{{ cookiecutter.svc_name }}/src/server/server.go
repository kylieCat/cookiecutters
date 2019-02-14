package server

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gitlab.internal.unity3d.com/sre/pipboy/generated/go/schemas"
	"gitlab.internal.unity3d.com/sre/pipeline"
	"gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/endpoints"
	"gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/logger"
	"gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/metrics"
)

// placeholders to ensure these packages don't get removed by dep
var _ = schemas.ObjectType_SERVICE
var _ = pipeline.AlertEventName

type ServerOptFunc func(e *echo.Echo) error

func WithEndpoints(endpoints ...endpoints.Endpoint) ServerOptFunc {
	return func(e *echo.Echo) error {
		for _, ep := range endpoints {
			e.Add(ep.Method, ep.Url, ep.Handler)
		}
		return nil
	}
}

func WithMiddleware(middleware ...echo.MiddlewareFunc) ServerOptFunc {
	return func(e *echo.Echo) error {
			e.Use(middleware...)
			return nil
	}
}

func NewServer(opts ...ServerOptFunc) *echo.Echo {
	e := echo.New()
	for _, opt := range opts {
		err := opt(e)
		if err != nil {
			log.Fatalf("error encountered configuring server: %s", err)
		}
	}
	return e
}

// StartServer ...
func StartServer() {
	server := NewServer(
		WithEndpoints(endpoints.Init()...),
		WithMiddleware(
			middleware.LoggerWithConfig(middleware.LoggerConfig{
				Format: "${method} ${uri} ${status} ${latency_human}\n",
			}),
			metrics.MetricsMiddleware(),
		),
	)
	logger.Info("starting {{ cookiecutter.svc_name }} server")
	log.Fatal(server.Start(":{{ cookiecutter.port }}"))
}
