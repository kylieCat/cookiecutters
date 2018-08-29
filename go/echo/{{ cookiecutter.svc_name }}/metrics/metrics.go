package metrics

import (
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus"
	"github.com.com/{{ cookiecutter.github_username }}/{{ cookiecutter.svc_name }}/config"
)

// MetricsMiddleware creates and registers prometheus metrics for HTTP endpoints then returns a middleware function
// to record them.
func MetricsMiddleware() echo.MiddlewareFunc {
	ApiCalls := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "sre",
			Subsystem: "{{ cookiecutter.svc_name }}",
			Name:      "api_calls",
			Help:      "A count of total calls by api endpoint",
			Buckets:   []float64{10, 20, 40, 80, 160, 320, 640},
		},
		[]string{"method", "endpoint", "response_code"},
	)
	prometheus.MustRegister(ApiCalls)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}

			metrics := []string{c.Request().Method, c.Path(), strconv.Itoa(c.Response().Status)}
			ApiCalls.WithLabelValues(metrics...).Observe(time.Since(start).Seconds())
			return nil
		}
	}
}

// Init can be used to register any additional prometheus metrics
func Init() {
	conf := config.GetConfig()
	if conf.GetBool("metrics.enabled") {
		// Add additional metrics here to be registered.
		prometheus.MustRegister()
	}
}
