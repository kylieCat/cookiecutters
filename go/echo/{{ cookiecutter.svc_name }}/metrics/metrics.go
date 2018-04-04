package metrics

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/config"
)

var (
	ApiCalls = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "sre",
			Subsystem: "{{ cookiecutter.svc_name }}",
			Name:      "api_calls",
			Help:      "A count of total calls by api endpoint",
			Buckets:   []float64{10, 20, 40, 80, 160, 320, 640},
		},
		[]string{"endpoint", "protocol", "response_code"},
	)
)

func Init() {
	prometheus.MustRegister(ApiCalls)
	if config.GetConfig().GetBool("metrics.enabled") == true {
		http.Handle(config.GetConfig().GetString("metrics.endpoint"), prometheus.Handler())
	}
}
