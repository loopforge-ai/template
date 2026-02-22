package inbound

import (
	"net/http"

	dashboard "github.com/loopforge-ai/template/internal/dashboard/domain"
	httpserver "github.com/loopforge-ai/utils/html"
)

// IndexHandler serves the home page.
type IndexHandler struct {
	renderer *httpserver.Renderer
	version  string
}

// NewIndexHandler creates a new index handler.
func NewIndexHandler(renderer *httpserver.Renderer, version string) *IndexHandler {
	return &IndexHandler{
		renderer: renderer,
		version:  version,
	}
}

// ServeHTTP renders the dashboard page.
func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	data := dashboard.PageData{
		Title:   "Dashboard",
		Version: h.version,
	}
	httpserver.RenderPage(w, h.renderer, "index", data)
}
