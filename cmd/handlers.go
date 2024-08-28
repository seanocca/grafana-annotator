package cmd

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the Prometheus metrics endpoint
	promhttp.Handler().ServeHTTP(w, r)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the readiness probe
	w.Write([]byte("ready"))
	w.WriteHeader(http.StatusOK)
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the liveness probe
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}

func annotateHandler(w http.ResponseWriter, r *http.Request) {

	// Get the data from the request body
	data := &CustomAnnotation{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		requestErrorTotal.Inc()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// annotate the data to the Grafana server
	annotateMetric(&grafanaURL, data)

	w.WriteHeader(http.StatusOK)
}
