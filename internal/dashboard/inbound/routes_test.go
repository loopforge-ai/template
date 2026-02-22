package inbound_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/loopforge-ai/template/internal/dashboard/inbound"
	"github.com/loopforge-ai/utils/assert"
)

func Test_RegisterRoutes_With_ValidFS_Should_ReturnMux(t *testing.T) {
	t.Parallel()

	// Arrange
	health := inbound.NewHealthHandler()
	idx := inbound.NewIndexHandler(newTestRenderer(t), "1.0.0")
	fsys := fstest.MapFS{
		"static/css/style.css": &fstest.MapFile{Data: []byte("body{}")},
	}

	// Act
	mux, err := inbound.RegisterRoutes(health, idx, fsys)

	// Assert
	assert.That(t, "error should be nil", err, nil)
	assert.That(t, "mux should not be nil", mux != nil, true)
}

func Test_RegisterRoutes_With_ValidFS_Should_ServeHealthEndpoint(t *testing.T) {
	t.Parallel()

	// Arrange
	health := inbound.NewHealthHandler()
	idx := inbound.NewIndexHandler(newTestRenderer(t), "1.0.0")
	fsys := fstest.MapFS{
		"static/css/style.css": &fstest.MapFile{Data: []byte("body{}")},
	}
	mux, _ := inbound.RegisterRoutes(health, idx, fsys)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	// Act
	mux.ServeHTTP(rec, req)

	// Assert
	assert.That(t, "status code", rec.Code, http.StatusOK)
	assert.That(t, "body", rec.Body.String(), `{"status":"ok"}`)
}

func Test_RegisterRoutes_With_ValidFS_Should_ServeIndexPage(t *testing.T) {
	t.Parallel()

	// Arrange
	health := inbound.NewHealthHandler()
	idx := inbound.NewIndexHandler(newTestRenderer(t), "1.0.0")
	fsys := fstest.MapFS{
		"static/css/style.css": &fstest.MapFile{Data: []byte("body{}")},
	}
	mux, _ := inbound.RegisterRoutes(health, idx, fsys)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Act
	mux.ServeHTTP(rec, req)

	// Assert
	assert.That(t, "status code", rec.Code, http.StatusOK)
	assert.That(t, "should contain DOCTYPE", strings.Contains(rec.Body.String(), "<!DOCTYPE html>"), true)
}

func Test_RegisterRoutes_With_ValidFS_Should_ServeStaticAssets(t *testing.T) {
	t.Parallel()

	// Arrange
	health := inbound.NewHealthHandler()
	idx := inbound.NewIndexHandler(newTestRenderer(t), "1.0.0")
	fsys := fstest.MapFS{
		"static/css/style.css": &fstest.MapFile{Data: []byte("body{}")},
	}
	mux, _ := inbound.RegisterRoutes(health, idx, fsys)
	req := httptest.NewRequest(http.MethodGet, "/static/css/style.css", nil)
	rec := httptest.NewRecorder()

	// Act
	mux.ServeHTTP(rec, req)

	// Assert
	assert.That(t, "status code", rec.Code, http.StatusOK)
	assert.That(t, "body should contain CSS", strings.Contains(rec.Body.String(), "body{}"), true)
	assert.That(t, "cache-control should be set", strings.Contains(rec.Header().Get("Cache-Control"), "public"), true)
}
