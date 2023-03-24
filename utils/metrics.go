package utils

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
)

var (
	TotalResponsesWithStatusCode = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "http",
			Name:      "total_responses_with_status_code",
			Help:      "Total http responses with code served by the server",
		},
		[]string{
			// HTTP status code.
			"code",
		},
	)
)

// RegisterMetrics registers all metrics for the server.
func RegisterMetrics() {
	prometheus.MustRegister(TotalResponsesWithStatusCode)
	log.Print("Metrics registered successfully")
}
