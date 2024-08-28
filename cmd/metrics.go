package cmd

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "annotator_requests_total",
			Help: "Total number of annotations received",
		},
	)

	requestErrorTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "annotator_requests_errors_total",
			Help: "Total number of annotations that returned an error",
		},
	)

	responseTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "annotator_response_time",
			Help:    "Response time distribution",
			Buckets: prometheus.LinearBuckets(0, 1000, 10), // Example buckets from 0ms to 1000ms
		},
	)

	uniqueTagCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "annotator_tags",
			Help: "Count of individual unique tags",
		},
		[]string{"tag"},
	)

	tagTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "annotator_tags_total",
			Help: "Total number of tags",
		},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(responseTime)
	prometheus.MustRegister(requestErrorTotal)
	prometheus.MustRegister(uniqueTagCounter)
	prometheus.MustRegister(tagTotal)
}

func uniqueTag(tag string) {
	uniqueTagCounter.WithLabelValues(tag).Inc()
	tagTotal.Inc()
}
