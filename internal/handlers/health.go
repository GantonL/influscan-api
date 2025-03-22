package handlers

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

// HealthResponse represents the health check response structure
type HealthResponse struct {
	Status      string    `json:"status"`
	Timestamp   time.Time `json:"timestamp"`
	Environment string    `json:"environment"`
	Version     string    `json:"version"`
	GoVersion   string    `json:"go_version"`
	Uptime      string    `json:"uptime"`
}

var startTime = time.Now()

// HealthCheck handles the health check endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// Create health response
	response := HealthResponse{
		Status:      "healthy",
		Timestamp:   time.Now().UTC(),
		Environment: "development", // This should be configured via environment variables
		Version:     "1.0.0",
		GoVersion:   runtime.Version(),
		Uptime:      time.Since(startTime).String(),
	}

	// Set HTTP status code
	w.WriteHeader(http.StatusOK)

	// Encode and send response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
