# API Contracts

## Overview

The HTTP server exposes 3 routes via `net/http.ServeMux`, configured in `internal/dashboard/inbound/routes.go`. All routes use a middleware chain from `loopforge-ai/utils/html`: `SecurityHeaders` → `Log` → `Recover` → `ContentType`.

## Endpoints

### GET /health

**Purpose:** Health check endpoint for container orchestration and monitoring.

| Property | Value |
|----------|-------|
| Method | `GET` |
| Path | `/health` |
| Handler | `HealthHandler` (`handler_health.go`) |
| Middleware | SecurityHeaders → Log → Recover → ContentType |
| Authentication | None |

**Response:**
```json
{
  "status": "ok"
}
```

| Status Code | Content-Type | Description |
|-------------|-------------|-------------|
| 200 | `application/json` | Server is healthy |

**Docker HEALTHCHECK** uses this endpoint:
```
wget --no-verbose --tries=1 --spider http://localhost:8080/health
```

---

### GET /

**Purpose:** Serve the dashboard HTML page (server-side rendered).

| Property | Value |
|----------|-------|
| Method | `GET` |
| Path | `/{$}` (exact root match) |
| Handler | `IndexHandler` (`handler_index.go`) |
| Middleware | SecurityHeaders → Log → Recover → ContentType |
| Authentication | None |

**Template Data:**
```go
type PageData struct {
    Title   string  // "Dashboard"
    Version string  // Build version (e.g., "dev", "1.0.0")
}
```

| Status Code | Content-Type | Description |
|-------------|-------------|-------------|
| 200 | `text/html` | Rendered dashboard page |
| 500 | `text/plain` | Template rendering error |

---

### GET /static/{path...}

**Purpose:** Serve embedded static assets (CSS, vendor files).

| Property | Value |
|----------|-------|
| Method | `GET` |
| Path | `/static/{path...}` |
| Handler | `http.FileServerFS` (stdlib) |
| Middleware | CacheControl |
| Authentication | None |
| Source FS | `web.FS` sub-directory `static/` |

**Available Assets:**
- `/static/css/style.css` — Main stylesheet (dark theme, responsive)

| Status Code | Headers | Description |
|-------------|---------|-------------|
| 200 | `Cache-Control: public, ...` | Asset served with cache headers |
| 404 | — | Asset not found |

## Middleware Chain

All page/API handlers use this middleware pipeline (applied in `RegisterRoutes`):

1. **SecurityHeaders** — Sets security-related HTTP headers
2. **Log** — Request logging via `slog`
3. **Recover** — Panic recovery
4. **ContentType** — Content-Type negotiation

Static assets use only:
1. **CacheControl** — Sets public cache headers

## Configuration

| Environment Variable | Default | Description |
|---------------------|---------|-------------|
| `SERVER_ADDR` | (from `html.DefaultAddr`) | HTTP listen address |

## Authentication

No authentication is currently implemented. All endpoints are publicly accessible.
