package inbound

import (
	"log/slog"
	"net/http"
)

// HealthHandler serves the health check endpoint.
type HealthHandler struct{}

// NewHealthHandler creates a new health check handler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// ServeHTTP returns a JSON health status.
func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(`{"status":"ok"}`)); err != nil {
		slog.Error("write health response", "error", err)
	}
}
