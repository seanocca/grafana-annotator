package cmd

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var grafanaURL string

func init() {
	grafanaURL = constructURL()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/annotate", annotateHandler).Methods("POST")
	r.HandleFunc("/metrics", metricsHandler).Methods("GET")
	r.HandleFunc("/readiness", readinessHandler).Methods("GET")
	r.HandleFunc("/liveness", livenessHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
