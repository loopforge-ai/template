package dashboard

import (
	httpserver "github.com/loopforge-ai/utils/html"
)

// PageData is the top-level data passed to every template.
type PageData struct {
	Title   string
	Version string
}

// RendererConfig is the configuration for the dashboard server templates.
var RendererConfig = httpserver.RendererConfig{
	CommonFiles: []string{
		"templates/layouts/base.html",
		"templates/partials/footer.html",
		"templates/partials/header.html",
	},
	Pages: []string{"index"},
}
