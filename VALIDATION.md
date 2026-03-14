# Validation Procedure

Repeatable validation procedure for verifying template correctness, convention coverage, and bootstrap safety before every release.

## 1. Template Quality Gates

Run each command and confirm the expected output.

```bash
# Lint — must report zero issues
golangci-lint run ./...
# Expected: 0 issues.

# Unit tests — all must pass
go test ./...
# Expected: ok for all packages with tests

# Docker build — must succeed
docker build -t template:validate .
# Expected: build completes without errors
```

## 2. Pattern Map Verification

Every source file maps to exactly one demonstrated pattern. No file is removable without losing a demonstrated pattern.

| # | File | Demonstrated Pattern | FR |
|---|------|---------------------|----|
| 1 | `cmd/cli/main.go` | CLI entry point with signal handling | FR4 |
| 2 | `cmd/mcp/main.go` | MCP entry point (stdio protocol server) | FR4 |
| 3 | `cmd/server/main.go` | Composition root (dependency wiring) + graceful shutdown | FR19, FR20 |
| 4 | `internal/dashboard/domain/renderer.go` | Domain types + package-level config (`RendererConfig`) | FR18 |
| 5 | `internal/dashboard/inbound/handler_health.go` | Stateless HTTP handler (no dependencies) | FR11 |
| 6 | `internal/dashboard/inbound/handler_index.go` | Stateful HTTP handler (dependency injection) | FR12 |
| 7 | `internal/dashboard/inbound/routes.go` | Route registration + middleware chaining | FR13 |
| 8 | `internal/dashboard/inbound/handler_health_test.go` | Unit test for stateless handler (Arrange/Act/Assert) | FR21 |
| 9 | `internal/dashboard/inbound/handler_index_test.go` | Unit test for stateful handler (happy + error path) | FR22 |
| 10 | `internal/dashboard/inbound/routes_test.go` | Route-level integration tests (through actual mux) | FR23 |
| 11 | `internal/dashboard/inbound/testhelper_test.go` | Shared test fixture creation | FR24 |
| 12 | `web/embed.go` | Embedded filesystem (`embed.FS`) | FR14 |
| 13 | `web/static/css/style.css` | Static asset (served via embedded FS) | FR14 |
| 14 | `web/templates/layouts/base.html` | Template composition (layout + blocks) | FR15 |
| 15 | `web/templates/pages/index.html` | Page template with data binding | FR17 |
| 16 | `web/templates/partials/header.html` | Partial template without data binding (static) | FR16 |
| 17 | `web/templates/partials/footer.html` | Partial template with data binding (dynamic) | FR17 |

**Verification**: Each row maps to exactly one pattern. Removing any row loses a demonstrated pattern.

## 3. Placeholder Audit Procedure

All TODO markers must be categorized as `TODO(bootstrap)` or `TODO(developer)`.

```bash
# Must return zero results (excluding _bmad planning artifacts)
grep -rn 'TODO' . --include='*.md' \
  | grep -v 'TODO(bootstrap)' \
  | grep -v 'TODO(developer)' \
  | grep -v '_bmad' \
  | grep -v 'docs/' \
  | grep -v 'VALIDATION.md'
# Expected: no output (all TODOs categorized)
```

## 4. Bootstrap Replacement Safety

Verify no Go identifiers survive bootstrap string replacement with confusing names.

```bash
# No TemplateConfig references in Go source
grep -rn 'TemplateConfig' --include='*.go' .
# Expected: no output

# Remaining template/Template references are module paths or html/template strings only
grep -rn 'template\|Template' --include='*.go' --exclude-dir=_bmad . \
  | grep -v 'html/template' \
  | grep -v 'github.com/loopforge-ai/template' \
  | grep -v 'templates/' \
  | grep -v 'template-cli' \
  | grep -v '"template"' \
  | grep -v 'every template'
# Expected: no output (all references are module paths or template-system strings)
```

## 5. Convention-to-File Mapping

Every CLAUDE.md convention is demonstrated in at least one source file. The code IS the spec (NFR6).

| # | Convention | Demonstrated In | Reference |
|---|-----------|----------------|-----------|
| 1 | File ordering (const/type/var then functions, alphabetical) | `internal/dashboard/domain/renderer.go` | `PageData` type before `RendererConfig` var |
| 2 | Switch/case alphabetical ordering | `internal/dashboard/inbound/routes.go` | Route registrations in path order |
| 3 | Constructors return pointer (`NewTypeName(deps) *TypeName`) | `internal/dashboard/inbound/handler_health.go` | `NewHealthHandler() *HealthHandler` |
| 4 | Error handling with context wrapping | `internal/dashboard/inbound/routes.go` | `fmt.Errorf("create static sub-filesystem: %w", err)` |
| 5 | Import grouping (stdlib, blank line, internal) | `cmd/server/main.go` | stdlib block then internal block |
| 6 | No dead code | All files | No unused functions/types/variables |
| 7 | No external dependencies | `go.mod` | Only `loopforge-ai/utils` |
| 8 | Test naming (`Test_<Unit>_With_<Condition>_Should_<Outcome>`) | `internal/dashboard/inbound/handler_health_test.go` | `Test_HealthHandler_With_Request_Should_ReturnOK` |
| 9 | Strict Arrange/Act/Assert with comments | `internal/dashboard/inbound/handler_health_test.go` | `// Arrange`, `// Act`, `// Assert` |
| 10 | Test parallelism (`t.Parallel()`) | All test files | First line of every test |
| 11 | Assertions via `assert.That` | `internal/dashboard/inbound/handler_health_test.go` | `assert.That(t, ...)` |
| 12 | External test package (`package <name>_test`) | `internal/dashboard/inbound/handler_health_test.go` | `package inbound_test` |
| 13 | Test helpers in `testhelper_test.go` | `internal/dashboard/inbound/testhelper_test.go` | `newTestRenderer`, `newBrokenRenderer` |
| 14 | Struct-based handlers with `ServeHTTP` | `internal/dashboard/inbound/handler_index.go` | `IndexHandler` struct with `ServeHTTP` method |
| 15 | Route registration per context (`RegisterRoutes`) | `internal/dashboard/inbound/routes.go` | `RegisterRoutes(healthHandler, indexHandler, staticFS)` |
| 16 | Middleware chain order | `internal/dashboard/inbound/routes.go` | `SecurityHeaders(Log(Recover(ContentType(handler))))` |
| 17 | Domain layer purity (no I/O imports) | `internal/dashboard/domain/renderer.go` | Only imports `utils/html` for config type |
| 18 | Embedded filesystem with `all:` prefix | `web/embed.go` | `//go:embed all:static all:templates` |

## 6. Post-Bootstrap Validation Procedure

Execute these steps to verify a bootstrapped project works correctly.

```bash
# 1. Bootstrap a test project (manual — requires forge bootstrap skill)
# forge bootstrap --name testproject --module github.com/example/testproject

# 2. In the bootstrapped project directory:
cd /tmp/testproject

# 3. Lint
golangci-lint run ./...
# Expected: 0 issues

# 4. Test
go test ./...
# Expected: all tests pass

# 5. Build
go build ./cmd/server
# Expected: compiles without errors

# 6. Verify no residual template/Template in Go identifiers
grep -rn 'template\|Template' --include='*.go' . | grep -v 'html/template'
# Expected: no output (module path was replaced, no identifier references remain)

# 7. Start server and verify HTTP response
go run ./cmd/server &
sleep 2
curl -s http://localhost:8080 | grep -i testproject
# Expected: HTML containing "testproject" (not "Template")
kill %1
```

## 7. Bootstrap-to-Browser Verification

After bootstrap, a developer should reach a working browser view with two commands:

```bash
# Command 1: Install dependencies (if needed)
go mod tidy

# Command 2: Start the server
go run ./cmd/server
# Then open http://localhost:8080 in a browser
```

**Expected**: Dashboard page renders with project name and version stat card.

## Release Readiness Checklist

- [ ] Quality Gates pass (lint, test, Docker build)
- [ ] Pattern Map verified (17 files, 17 patterns, no redundancy)
- [ ] Placeholder Audit clean (all TODOs categorized)
- [ ] Bootstrap Replacement Safety verified (no confusing identifiers)
- [ ] Convention-to-File Mapping complete (18 conventions demonstrated)
- [ ] Post-Bootstrap Validation passes (bootstrapped project compiles and runs)
- [ ] Bootstrap-to-Browser works (two commands to running server)
