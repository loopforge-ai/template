# Development Guide

## Prerequisites

| Tool | Version | Purpose |
|------|---------|---------|
| Go | 1.25.6+ | Language runtime |
| golangci-lint | v2+ | Linting (config in `.golangci.yml`) |
| just | latest | Task runner (`.justfile`) |
| Docker | latest | Container builds (optional) |

## Getting Started

### Clone and Setup

```bash
git clone https://github.com/loopforge-ai/template.git
cd template
```

### Install Dependencies

```bash
go mod download
```

No external dependencies to install beyond the Go toolchain. The only dependency (`loopforge-ai/utils`) is fetched automatically.

## Common Tasks

### Run the HTTP Server

```bash
go run ./cmd/server
# or
just server
```

Server starts on the default address (configurable via `SERVER_ADDR` environment variable).

### Build All Binaries

```bash
just build
# Produces: bin/cli, bin/mcp, bin/server
```

Individual builds:
```bash
just build-cli
just build-mcp
just build-server
```

### Install Binaries

```bash
just install
```

### Run Tests

```bash
go test ./...          # Unit tests only
go test -cover ./...   # With coverage
just test              # Via task runner (includes coverage)
```

### Run Integration Tests

```bash
go test ./... -tags=integration
```

### Run Linter

```bash
golangci-lint run ./...
just lint
```

Must pass with **zero issues**. Do not modify `.golangci.yml` to suppress findings.

## Docker

### Build Image

```bash
just docker-build                    # Default version "dev"
just docker-build version="1.0.0"   # Specific version
```

### Run Container

```bash
just docker-run                              # Default port 8080
just docker-run version="1.0.0" port="9090"  # Custom
```

### Stop Containers

```bash
just docker-stop
```

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_ADDR` | `html.DefaultAddr` | HTTP server listen address |

## Project Conventions

### File Structure

New bounded contexts follow this pattern:
```
internal/<context-name>/
├── domain/      # Business logic, types, interfaces
├── inbound/     # HTTP handlers, adapters for incoming requests
└── outbound/    # Repository implementations, external service clients
```

### Naming Conventions

| Element | Convention | Example |
|---------|-----------|---------|
| Test functions | `Test_<Unit>_With_<Condition>_Should_<Outcome>` | `Test_HealthHandler_With_Request_Should_ReturnOK` |
| Constructors | `New<TypeName>(deps) *<TypeName>` | `NewHealthHandler() *HealthHandler` |
| Error wrapping | `fmt.Errorf("<operation>: %w", err)` | `fmt.Errorf("sub static fs: %w", err)` |
| Test structure | Arrange/Act/Assert with explicit comments | `// Arrange`, `// Act`, `// Assert` |

### Adding a New Handler

1. Create domain types in `internal/<ctx>/domain/` if needed
2. Create handler in `internal/<ctx>/inbound/handler_<name>.go`
3. Add constructor `NewXxxHandler(deps) *XxxHandler`
4. Implement `http.Handler` interface (`ServeHTTP`)
5. Register route in `routes.go` `RegisterRoutes()` function
6. Write tests in `handler_<name>_test.go`
7. Run lint-fix loop: `golangci-lint run ./...`

### Adding a New Entry Point

1. Create `cmd/<name>/main.go`
2. Add signal handling (`signal.NotifyContext`)
3. Add `appVersion` var with `-ldflags` support
4. Add build targets to `.justfile`
5. Add build stage to `Dockerfile`

## Forge Integration

The project integrates with LoopForge Forge via MCP (`.mcp.json`). Before writing code manually, check forge skills:

```
search_skill → generate_skill (match) or define_skill (no match) → score_skill → refine_skill
```

## Workflow for Every Change

1. **Skill-First Loop** — Check forge for existing skills
2. **Red-Green Loop** — Write failing test → implement → pass
3. **Lint-Fix Loop** — `golangci-lint run ./...` → zero issues
4. **Verify-Implement-Verify Loop** — End-to-end verification
5. On failure → **Self-Healing Loop** (max 5 iterations)
