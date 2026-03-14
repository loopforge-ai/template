# Source Tree Analysis

## Directory Structure

```
template/
├── cmd/                            # Entry points (3 binaries)
│   ├── cli/
│   │   └── main.go                 # CLI binary — placeholder, signal-aware
│   ├── mcp/
│   │   └── main.go                 # MCP server binary — Model Context Protocol
│   └── server/
│       └── main.go                 # HTTP server binary — main application entry
├── internal/                       # Private application code
│   └── dashboard/                  # Bounded context: Dashboard
│       ├── domain/
│       │   └── renderer.go         # PageData type + RendererConfig
│       └── inbound/
│           ├── handler_health.go   # GET /health — JSON health check
│           ├── handler_index.go    # GET / — Dashboard HTML page
│           ├── routes.go           # RegisterRoutes() — mux setup + middleware
│           ├── handler_health_test.go
│           ├── handler_index_test.go
│           ├── routes_test.go
│           └── testhelper_test.go  # Shared test renderer fixtures
├── web/                            # Embedded web assets
│   ├── embed.go                    # embed.FS declaration (static + templates)
│   ├── static/
│   │   ├── css/
│   │   │   └── style.css           # Dark theme, responsive, CSS custom properties
│   │   └── vendor/
│   │       └── .gitkeep
│   └── templates/
│       ├── layouts/
│       │   └── base.html           # Base layout template (head, body structure)
│       ├── pages/
│       │   └── index.html          # Dashboard page (version stat card)
│       └── partials/
│           ├── header.html         # Navigation header
│           └── footer.html         # Footer with version
├── docs/                           # Project documentation (this folder)
├── _bmad/                          # BMAD Method configuration and workflows
├── CLAUDE.md                       # AI agent guidelines and conventions
├── README.md                       # Project readme (template with TODOs)
├── Dockerfile                      # Multi-stage Docker build (3 binaries)
├── .golangci.yml                   # Linter config (v2, all linters, selective disables)
├── .justfile                       # Task runner (build, test, lint, docker)
├── .mcp.json                       # Forge MCP server configuration
├── .gitignore                      # Standard Go + IDE ignores
└── go.mod                          # Go module: github.com/loopforge-ai/template
```

## Critical Folders

| Folder | Purpose | Layer |
|--------|---------|-------|
| `cmd/server/` | HTTP server entry point — wires domain, inbound, and web layers | Entry Point |
| `cmd/mcp/` | MCP server entry point — Model Context Protocol | Entry Point |
| `cmd/cli/` | CLI entry point — placeholder for future CLI logic | Entry Point |
| `internal/dashboard/domain/` | Business types and template configuration | Domain |
| `internal/dashboard/inbound/` | HTTP handlers and route registration | Inbound |
| `web/` | Embedded static assets and HTML templates | Assets |
| `web/templates/` | Go HTML templates (layouts, pages, partials) | Templates |
| `web/static/` | CSS and vendor assets served at `/static/` | Static |

## Entry Points

| Binary | Path | Description |
|--------|------|-------------|
| `server` | `cmd/server/main.go` | HTTP server on `:8080` with graceful shutdown |
| `mcp` | `cmd/mcp/main.go` | MCP protocol server via stdio |
| `cli` | `cmd/cli/main.go` | CLI tool — placeholder awaiting logic |

## File Counts

| Category | Count |
|----------|-------|
| Go source files | 8 |
| Go test files | 4 |
| HTML templates | 4 |
| CSS files | 1 |
| Config files | 5 (.golangci.yml, .justfile, .mcp.json, go.mod, Dockerfile) |
| **Total source** | **22** |
