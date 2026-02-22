package inbound_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/loopforge-ai/template/internal/dashboard/inbound"
	"github.com/loopforge-ai/utils/assert"
)

func Test_IndexHandler_With_BrokenRenderer_Should_Return500(t *testing.T) {
	t.Parallel()

	// Arrange
	handler := inbound.NewIndexHandler(newBrokenRenderer(t), "1.0.0")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Act
	handler.ServeHTTP(rec, req)

	// Assert
	assert.That(t, "status code", rec.Code, http.StatusInternalServerError)
	assert.That(t, "body should contain error", strings.Contains(rec.Body.String(), "Internal Server Error"), true)
}

func Test_IndexHandler_With_ValidRequest_Should_RenderDashboard(t *testing.T) {
	t.Parallel()

	// Arrange
	handler := inbound.NewIndexHandler(newTestRenderer(t), "1.0.0")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Act
	handler.ServeHTTP(rec, req)

	// Assert
	assert.That(t, "status code", rec.Code, http.StatusOK)
	assert.That(t, "should contain version", strings.Contains(rec.Body.String(), "1.0.0"), true)
}
