# Deployment Guide

## Overview

The Template project produces 3 binaries and a Docker image. The primary deployment target is the `server` binary, which serves the HTTP dashboard.

## Build Artifacts

| Binary | Source | Purpose |
|--------|--------|---------|
| `server` | `cmd/server/main.go` | HTTP server (primary deployment target) |
| `mcp` | `cmd/mcp/main.go` | MCP protocol server |
| `cli` | `cmd/cli/main.go` | CLI tool |

## Docker Deployment

### Image Build

```bash
# Build with version tag
docker build --build-arg VERSION=1.0.0 -t template:1.0.0 .

# Or via justfile
just docker-build version="1.0.0"
```

### Image Details

| Property | Value |
|----------|-------|
| Base (build) | `golang:1.25-alpine` |
| Base (runtime) | `alpine:3.21` |
| Runtime user | `app` (non-root) |
| Exposed port | `8080` |
| Entrypoint | `server` |
| Healthcheck | `wget --spider http://localhost:8080/health` (30s interval) |
| Installed packages | `ca-certificates`, `tzdata`, `wget` |

### Build Optimizations

- **Multi-stage build** — build tools not included in runtime image
- **CGO disabled** — `CGO_ENABLED=0` for static binaries
- **Trimpath** — removes local file paths from binary
- **Strip flags** — `-s -w` removes debug symbols and DWARF info

### Run Container

```bash
docker run --rm -p 8080:8080 template:1.0.0

# Or via justfile
just docker-run version="1.0.0" port="8080"
```

## Environment Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_ADDR` | `html.DefaultAddr` | HTTP listen address (e.g., `:8080`) |

## Health Monitoring

The `/health` endpoint returns JSON:

```json
{"status": "ok"}
```

Docker HEALTHCHECK configuration:
- **Interval:** 30 seconds
- **Timeout:** 10 seconds
- **Start period:** 5 seconds
- **Retries:** 3

## Binary Deployment (without Docker)

### Build

```bash
CGO_ENABLED=0 go build -trimpath \
    -ldflags "-s -w -X main.appVersion=1.0.0" \
    -o server ./cmd/server
```

### Run

```bash
SERVER_ADDR=:8080 ./server
```

### Signals

The server handles:
- `SIGINT` (Ctrl+C) — graceful shutdown
- `SIGTERM` — graceful shutdown (container orchestration)

Shutdown drains in-flight requests within the configured timeout before exiting.

## CI/CD

No CI/CD pipeline is currently configured. Recommended setup:

1. **Lint:** `golangci-lint run ./...`
2. **Test:** `go test -cover ./...`
3. **Build:** `go build -trimpath -ldflags "-s -w -X main.appVersion=$VERSION" -o server ./cmd/server`
4. **Docker:** `docker build --build-arg VERSION=$VERSION -t template:$VERSION .`
5. **Push:** Push to container registry
