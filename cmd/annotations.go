package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// CustomAnnotation is the struct that will be sent to Grafana as an annotation
type CustomAnnotation struct {
	Time           int64    `json:"time"`
	Title          string   `json:"title"`
	Text           string   `json:"text"`
	Tags           []string `json:"tags"`
	AnnotationName string   `json:"annotationName"`
}

// AnnotateMetric sends the data to the Grafana server
func annotateMetric(url *string, data *CustomAnnotation) error {
	// Annotate the data to the Grafana server

	jsonData, err := json.Marshal(data)
	if err != nil {
		requestErrorTotal.Inc()
		return err
	}

	for _, tag := range data.Tags {
		uniqueTag(tag)
	}

	req, err := http.NewRequest("POST", *url, bytes.NewBuffer(jsonData))
	if err != nil {
		requestErrorTotal.Inc()
		return err
	}

	requestsTotal.Inc()
	fmt.Print(req)

	return nil
}

func constructURL() string {
	// Get the Grafana URL from the ENV variables
	grafanaURL := os.Getenv("GRAFANA_URL")
	// Create the Grafana annotations URL
	url := fmt.Sprintf("%s/api/annotations", grafanaURL)

	return url
}
