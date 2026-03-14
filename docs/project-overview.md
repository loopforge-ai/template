# Project Overview — Template

## Executive Summary

**Template** is a Go-based project scaffold and starter kit by LoopForge AI. It provides a hexagonal architecture foundation with three entry-point binaries (CLI, MCP server, HTTP server), a single bounded context (`dashboard`), and embedded web assets for server-side rendered HTML pages. The project enforces strict coding conventions via `golangci-lint` and uses only standard library plus first-party `loopforge-ai/utils` modules — no external dependencies.

## Technology Stack

| Category | Technology | Version | Notes |
|----------|-----------|---------|-------|
| Language | Go | 1.25.6 | Set in `go.mod` |
| Architecture | Hexagonal (Ports & Adapters) | — | Domain/Inbound/Outbound layers |
| Linter | golangci-lint | v2 config | All linters enabled by default, selective disables |
| Testing | stdlib `testing` + `assert.That` | — | From `loopforge-ai/utils/assert` |
| HTTP Server | stdlib `net/http` | — | `http.ServeMux` with middleware chain |
| HTML Rendering | Go `html/template` + `embed.FS` | — | Via `loopforge-ai/utils/html` |
| MCP Server | `loopforge-ai/utils/mcp` | — | Model Context Protocol integration |
| Environment Config | `loopforge-ai/utils/env` | — | `env.Get()` with defaults |
| Task Runner | just (justfile) | — | Build, test, lint, docker tasks |
| Containerization | Docker (multi-stage) | Alpine 3.21 | `golang:1.25-alpine` build stage |
| Code Generation | LoopForge Forge | — | MCP-based skill forging |

## Architecture Type

**Hexagonal Architecture** (Ports & Adapters) with clear layer separation:

- **Domain** — Business logic, interfaces, data types. Zero infrastructure imports.
- **Inbound** — HTTP handlers adapting external requests to domain. Depends on domain only.
- **Outbound** — (Not yet implemented) Will implement domain interfaces for persistence/external services.

## Repository Structure

- **Type:** Monolith
- **Parts:** 1 (single Go module)
- **Bounded Contexts:** 1 (`dashboard`)
- **Entry Points:** 3 (`cmd/cli`, `cmd/mcp`, `cmd/server`)

## Key Dependencies

| Module | Purpose |
|--------|---------|
| `github.com/loopforge-ai/utils` | Shared utilities: `assert`, `env`, `html`, `mcp` |

No external third-party dependencies. All imports are standard library or `loopforge-ai` organization modules.

## Links to Detailed Documentation

- [Architecture](./architecture.md)
- [Source Tree Analysis](./source-tree-analysis.md)
- [API Contracts](./api-contracts.md)
- [Development Guide](./development-guide.md)
- [Deployment Guide](./deployment-guide.md)
