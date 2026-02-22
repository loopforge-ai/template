// Package web provides embedded static assets and HTML templates.
package web

import "embed"

//go:embed all:static all:templates
var FS embed.FS
