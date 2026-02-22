package inbound_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/loopforge-ai/template/internal/dashboard/inbound"
	"github.com/loopforge-ai/utils/assert"
)

func Test_HealthHandler_With_Request_Should_ReturnOK(t *testing.T) {
	t.Parallel()

	// Arrange
	handler := inbound.NewHealthHandler()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	// Act
	handler.ServeHTTP(rec, req)

	// Assert
	assert.That(t, "status code", rec.Code, http.StatusOK)
	assert.That(t, "content type", rec.Header().Get("Content-Type"), "application/json")
	assert.That(t, "body", rec.Body.String(), `{"status":"ok"}`)
}
