package inbound

import (
	"fmt"
	"io/fs"
	"net/http"

	httpserver "github.com/loopforge-ai/utils/html"
)

// RegisterRoutes configures all routes on a new ServeMux and returns it.
func RegisterRoutes(
	healthHandler *HealthHandler,
	indexHandler *IndexHandler,
	staticFS fs.FS,
) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	// Health check.
	mux.Handle("GET /health", httpserver.SecurityHeaders(httpserver.Log(httpserver.Recover(httpserver.ContentType(healthHandler)))))

	// Static assets from embed.FS.
	staticSub, err := fs.Sub(staticFS, "static")
	if err != nil {
		return nil, fmt.Errorf("sub static fs: %w", err)
	}
	mux.Handle("GET /static/{path...}",
		httpserver.CacheControl(http.StripPrefix("/static/", http.FileServerFS(staticSub))))

	// Pages.
	mux.Handle("GET /{$}", httpserver.SecurityHeaders(httpserver.Log(httpserver.Recover(httpserver.ContentType(indexHandler)))))

	return mux, nil
}
