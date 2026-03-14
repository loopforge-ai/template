package inbound_test

import (
	"testing"
	"testing/fstest"

	dashboard "github.com/loopforge-ai/template/internal/dashboard/domain"
	httpserver "github.com/loopforge-ai/utils/html"
)

// newBrokenRenderer creates a Renderer whose templates fail at execution time.
// The "base" layout uses {{call .Title}} which errors because Title is a string, not a function.
func newBrokenRenderer(t *testing.T) *httpserver.Renderer {
	t.Helper()
	fsys := fstest.MapFS{
		"templates/layouts/base.html": &fstest.MapFile{
			Data: []byte(`{{define "base"}}{{call .Title}}{{end}}`),
		},
		"templates/partials/header.html": &fstest.MapFile{
			Data: []byte(`{{define "header"}}h{{end}}`),
		},
		"templates/partials/footer.html": &fstest.MapFile{
			Data: []byte(`{{define "footer"}}f{{end}}`),
		},
		"templates/pages/index.html": &fstest.MapFile{
			Data: []byte(`{{define "title"}}t{{end}}{{define "content"}}c{{end}}`),
		},
	}
	r, err := httpserver.NewRenderer(fsys, dashboard.RendererConfig)
	if err != nil {
		t.Fatalf("create broken renderer: %v", err)
	}
	return r
}

func newTestRenderer(t *testing.T) *httpserver.Renderer {
	t.Helper()
	fsys := fstest.MapFS{
		"templates/layouts/base.html": &fstest.MapFile{
			Data: []byte(`{{define "base"}}<!DOCTYPE html><html><head><title>{{template "title" .}}</title></head><body>{{template "header" .}}{{template "content" .}}{{template "footer" .}}</body></html>{{end}}`),
		},
		"templates/partials/header.html": &fstest.MapFile{
			Data: []byte(`{{define "header"}}<header>nav</header>{{end}}`),
		},
		"templates/partials/footer.html": &fstest.MapFile{
			Data: []byte(`{{define "footer"}}<footer>{{.Version}}</footer>{{end}}`),
		},
		"templates/pages/index.html": &fstest.MapFile{
			Data: []byte(`{{define "title"}}Dashboard{{end}}{{define "content"}}<h1>Dashboard</h1><p>version={{.Version}}</p>{{end}}`),
		},
	}
	r, err := httpserver.NewRenderer(fsys, dashboard.RendererConfig)
	if err != nil {
		t.Fatalf("create test renderer: %v", err)
	}
	return r
}
