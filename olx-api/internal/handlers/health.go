package handlers

import (
	"log/slog"
	"net/http"
)

// A simple health check handler that responds with a JSON indicating the service is healthy. This can be used for monitoring and ensuring the service is running correctly.
func Health(w http.ResponseWriter, _r *http.Request) {
	slog.Info("Health check endpoint hit...")

	// setting the content-type to json
	w.Header().Set("Content-Type", "application/json")

	// methods sends the HTTP response with provided statuscode
	w.WriteHeader(http.StatusOK)

	// final response
	w.Write([]byte(`{"status":"healthy"}`))
}
