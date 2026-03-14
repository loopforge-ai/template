# Architecture — Template

## Executive Summary

Template is a Go monolith following **hexagonal architecture** (ports & adapters). It provides three entry-point binaries sharing one bounded context (`dashboard`), with strict layer separation enforced by convention and linting. The project uses only standard library and first-party `loopforge-ai` modules.

## Architecture Pattern

### Hexagonal Architecture (Ports & Adapters)

```
                    ┌─────────────────────────────────┐
                    │           cmd/server             │
                    │      (composition root)          │
                    └──────────┬──────────────────────┘
                               │ wires
                    ┌──────────▼──────────────────────┐
                    │         inbound/                 │
                    │   HTTP handlers + routes         │
                    │   (adapters for incoming)        │
                    └──────────┬──────────────────────┘
                               │ depends on
                    ┌──────────▼──────────────────────┐
                    │         domain/                  │
                    │   Business types + interfaces    │
                    │   (no infrastructure imports)    │
                    └──────────┬──────────────────────┘
                               │ implemented by
                    ┌──────────▼──────────────────────┐
                    │         outbound/                │
                    │   (not yet implemented)          │
                    │   Persistence, external APIs     │
                    └─────────────────────────────────┘
```

### Layer Rules

| Layer | Directory | May Import | Must Not Import |
|-------|-----------|-----------|-----------------|
| Domain | `internal/<ctx>/domain/` | stdlib only | inbound, outbound |
| Inbound | `internal/<ctx>/inbound/` | domain, stdlib, utils | outbound |
| Outbound | `internal/<ctx>/outbound/` | domain, stdlib, utils | inbound |
| Entry Points | `cmd/*/` | all layers (composition root) | — |

### Bounded Contexts

Currently one context exists:

| Context | Path | Purpose |
|---------|------|---------|
| `dashboard` | `internal/dashboard/` | HTTP dashboard with health check and index page |

New contexts follow the pattern: `internal/<context-name>/domain/`, `internal/<context-name>/inbound/`, `internal/<context-name>/outbound/`.

## Technology Stack

| Category | Technology | Version | Notes |
|----------|-----------|---------|-------|
| Language | Go | 1.25.6 | Single `go.mod` at root |
| HTTP | stdlib `net/http` | — | `ServeMux` + middleware chain |
| Templates | stdlib `html/template` | — | Embedded via `embed.FS` |
| Logging | stdlib `log/slog` | — | Structured logging |
| Utilities | `loopforge-ai/utils` | v0.1.0 | `assert`, `env`, `html`, `mcp` |
| Linter | golangci-lint | v2 config | All linters on, selective disables |
| Container | Docker | Alpine 3.21 | Multi-stage build |
| Task Runner | just | — | `.justfile` |
| Code Gen | Forge MCP | — | Skill-based code generation |

## Component Architecture

### Entry Points

| Binary | Source | Purpose | Key Dependencies |
|--------|--------|---------|-----------------|
| `server` | `cmd/server/main.go` | HTTP server on configurable addr | domain, inbound, web, utils/env, utils/html |
| `mcp` | `cmd/mcp/main.go` | MCP protocol server (stdio) | utils/mcp |
| `cli` | `cmd/cli/main.go` | CLI tool (placeholder) | stdlib only |

### Domain Layer

**`internal/dashboard/domain/`**

| Type | Purpose |
|------|---------|
| `PageData` | Struct — top-level template data (`Title`, `Version`) |
| `RendererConfig` | Var — `RendererConfig` defining common files and page names |

### Inbound Layer

**`internal/dashboard/inbound/`**

| Component | Type | Route | Description |
|-----------|------|-------|-------------|
| `HealthHandler` | `http.Handler` | `GET /health` | JSON health status response |
| `IndexHandler` | `http.Handler` | `GET /{$}` | Server-rendered dashboard page |
| `RegisterRoutes` | Function | — | Wires handlers + middleware onto `ServeMux` |

### Web Assets Layer

**`web/`**

| Component | Path | Description |
|-----------|------|-------------|
| `FS` | `web/embed.go` | `embed.FS` containing static + templates |
| Base layout | `templates/layouts/base.html` | HTML skeleton with head/body |
| Index page | `templates/pages/index.html` | Dashboard content with version stat card |
| Header | `templates/partials/header.html` | Navigation bar |
| Footer | `templates/partials/footer.html` | Version display footer |
| Stylesheet | `static/css/style.css` | Dark theme, responsive, CSS custom properties |

## HTTP Server Architecture

### Request Flow

```
Client Request
    │
    ▼
http.ServeMux (route matching)
    │
    ▼
Middleware Chain: SecurityHeaders → Log → Recover → ContentType
    │
    ▼
Handler (HealthHandler or IndexHandler)
    │
    ▼
Response (JSON or rendered HTML)
```

### Server Configuration

| Setting | Source | Default |
|---------|--------|---------|
| Listen Address | `SERVER_ADDR` env var | `html.DefaultAddr` |
| Idle Timeout | `html.IdleTimeout` | (from utils) |
| Read Timeout | `html.ReadTimeout` | (from utils) |
| Write Timeout | `html.WriteTimeout` | (from utils) |
| Max Header Bytes | `html.MaxHeaderBytes` | (from utils) |
| Shutdown Timeout | `html.ShutdownTimeout` | (from utils) |

### Graceful Shutdown

The server listens for `SIGINT` and `SIGTERM`, then calls `srv.Shutdown()` with a timeout context.

## Testing Strategy

| Aspect | Approach |
|--------|----------|
| Framework | stdlib `testing` + `loopforge-ai/utils/assert` |
| Naming | `Test_<Unit>_With_<Condition>_Should_<Outcome>` |
| Structure | Strict Arrange/Act/Assert with explicit comments |
| Parallelism | `t.Parallel()` in every test |
| HTTP Tests | `httptest.NewRequest` + `httptest.NewRecorder` |
| Template Tests | `testing/fstest.MapFS` with mock templates |
| Test Helpers | `newTestRenderer()` and `newBrokenRenderer()` in `testhelper_test.go` |
| Integration | Build tag `//go:build integration` |

### Test Coverage

| File | Tests | Scenarios |
|------|-------|-----------|
| `handler_health.go` | 1 | Returns 200 + JSON body |
| `handler_index.go` | 2 | Valid render, broken renderer → 500 |
| `routes.go` | 4 | Mux creation, health endpoint, index page, static assets |

## Build & Deployment

### Build Pipeline

```
just lint → just test → just build
```

Produces 3 binaries in `bin/`: `cli`, `mcp`, `server`.

### Docker

Multi-stage build:
1. **Build stage** (`golang:1.25-alpine`) — compiles 3 binaries with `-trimpath -ldflags "-s -w"`
2. **Runtime stage** (`alpine:3.21`) — non-root `app` user, healthcheck on `/health`

### Version Injection

All binaries accept a version at build time:
```
-ldflags "-X main.appVersion=${VERSION}"
```

## Coding Conventions

| Convention | Rule |
|-----------|------|
| File ordering | Alphabetical: const → type → var → functions/methods |
| Constructors | `NewTypeName(deps) *TypeName` — always return pointer |
| Error handling | `fmt.Errorf("operation: %w", err)` |
| Imports | stdlib, blank line, internal packages |
| Switch/case | Alphabetical ordering |
| Dependencies | stdlib + `loopforge-ai` org only |
