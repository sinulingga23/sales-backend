package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	M *Metric
)

type Metric struct {
	TotalRequestEndpoint    *prometheus.CounterVec
	DurationRequestEndpoint *prometheus.HistogramVec
}

func NewMetric(registry prometheus.Registerer) *Metric {
	m := &Metric{
		TotalRequestEndpoint: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "sales_backend",
			Name:      "total_request_endpoint",
			Help:      "It's show total request for each endpoint",
		}, []string{"service_name", "http_method", "http_status"}),
		DurationRequestEndpoint: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "sales_backend",
			Name:      "duration_request_endpoint",
			Help:      "it'show duration request for each endpoint",
		}, []string{"service_name", "http_method", "http_status"}),
	}

	registry.MustRegister(m.TotalRequestEndpoint, m.DurationRequestEndpoint)
	return m
}
