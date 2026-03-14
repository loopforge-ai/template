---
project_name: 'template'
user_name: 'Andy'
date: '2026-03-14'
sections_completed: ['technology_stack', 'language_rules', 'framework_rules', 'testing_rules', 'code_quality', 'workflow_rules', 'critical_rules']
status: 'complete'
rule_count: 67
optimized_for_llm: true
---

# Project Context for AI Agents

_This file contains critical rules and patterns that AI agents must follow when implementing code in this project. Focus on unobvious details that agents might otherwise miss._

---

## Technology Stack & Versions

| Technology | Version | Constraint |
|-----------|---------|-----------|
| Go | 1.25.6 | Exact version in go.mod; use 1.22+ ServeMux features freely |
| loopforge-ai/utils | v0.1.0 | Canonical implementations: assert, env, fs, html, llm, mcp, yaml |
| golangci-lint | v2 config | All linters on, selective disables in .golangci.yml |
| Docker base (build) | golang:1.25-alpine | CGO_ENABLED=0 always — no CGO dependencies allowed |
| Docker base (runtime) | alpine:3.21 | Non-root `app` user |
| just | latest | Task runner; build chain: lint → test → build |

**Hard Constraints:**
- **No external dependencies** beyond stdlib and `github.com/loopforge-ai` modules. Never introduce third-party modules.
- **CGO_ENABLED=0** for all builds. Do not introduce packages requiring CGO.
- **Utils are canonical, not optional**: `loopforge-ai/utils` packages (assert, env, fs, html, llm, mcp, yaml) are the exclusive implementations. Do not reinvent, wrap, or replace them with alternatives like testify, gorilla/mux, or manual os.Getenv fallback.
- **Single assertion function**: `assert.That(t, description, got, expected)` — four arguments, every test. No testify, no manual `if/t.Fatalf`.
- **embed.FS for web assets**: Static files and templates go in `web/` via `//go:embed`. No external asset pipeline or bundler.
- **Use utils/html middleware**: SecurityHeaders, Log, Recover, ContentType, CacheControl. No hand-rolled middleware or third-party routers.
- **One module, three binaries**: `cmd/cli`, `cmd/mcp`, `cmd/server` share a single `go.mod`. Do not split into separate modules or create `go.work` workspaces.
- **Module path from go.mod**: Always read the actual module path from `go.mod` for internal imports. Do not hardcode `loopforge-ai/template` — it changes when the project is cloned and renamed.
- **Build pipeline order**: `just build` enforces lint → test → build. Do not skip gates with direct `go build`.

## Critical Implementation Rules

### Language-Specific Rules (Go)

- **Error handling**: Wrap with context: `fmt.Errorf("operation: %w", err)`. Check immediately — never defer error checking.
- **Constructors**: `NewTypeName(deps) *TypeName` — always return pointer. First function after its type definition.
- **File ordering**: Alphabetical within each file: `const` → `type` → `var` blocks, then functions/methods. `NewTypeName` comes first after its type (sole exception).
- **Switch/case ordering**: Alphabetical `case` clauses in `switch` statements.
- **Import grouping**: stdlib first, blank line, then internal/`loopforge-ai` packages. Alias when the package name doesn't convey its role at the call site — follow existing patterns (e.g., `httpserver`, `dashboard`).
- **Logging**: Use `log/slog` for all production logging (`slog.Info`, `slog.Error` with key-value pairs). `log.Fatalf` is allowed only in `main()` for fatal startup errors. Never use `log.Println` or `fmt.Println` for operational output.
- **Context propagation**: Create `context.Context` at the entry point (`signal.NotifyContext`), pass it down. Never create `context.Background()` inside a function that already receives a context. Use `context.WithTimeout`/`context.WithCancel` for scoped operations only.
- **Module-level `var`**: Only for build-time injection (`appVersion`) and package-level config objects (`RendererConfig`). Never for mutable shared state.
- **Comments**: GoDoc comments on exported symbols only (`// TypeName does X`). No inline comments unless logic is genuinely non-obvious. Do not over-comment.
- **Interface satisfaction**: Interfaces defined in `domain/`, implementations in `inbound/` or `outbound/`. Implicit satisfaction — no `var _ Interface = (*Type)(nil)` checks.
- **No dead code**: Eliminate unused functions, types, variables, imports on every change. No commented-out code.

### Framework-Specific Rules (Hexagonal Architecture + stdlib HTTP)

- **Layer boundaries are inviolable**:
  - `domain/` must never import from `inbound/`, `outbound/`, or `web/`.
  - `inbound/` depends on `domain/` (and may reference `web/` for embed.FS). Never on `outbound/`.
  - `outbound/` depends on `domain/`. Never on `inbound/` or `web/`.
  - `web/` is a shared resource package — only `cmd/server/main.go` and `inbound/` may import it.
  - `cmd/` packages are composition roots, not shared libraries. Never import one `cmd/` package from another.
- **Bounded context pattern**: Each context lives at `internal/<context-name>/` with `domain/`, `inbound/`, `outbound/` subdirectories. Contexts are independent — no cross-context imports within `internal/`.
- **Domain layer purity**: `domain/` defines data types, configuration values, and interfaces. No implementations, no I/O, no HTTP, no filesystem. All side effects live in adapter layers.
- **Struct-based handlers**: Implement `http.Handler` via struct with `ServeHTTP` method. Inject dependencies through constructor fields. Do not use `http.HandlerFunc` closures.
- **Route registration per context**: Each bounded context has its own `RegisterRoutes()` in `inbound/routes.go`. Takes handler dependencies as parameters — does not construct them. `cmd/server/main.go` calls each context's `RegisterRoutes` and mounts them.
- **Middleware chain**: Page/API handlers: `SecurityHeaders(Log(Recover(ContentType(handler))))`. Static assets: `CacheControl` only.
- **Template system**: Use `httpserver.NewRenderer(fs, config)`. Define `RendererConfig` in domain with `CommonFiles` (layouts + partials) and `Pages` (page identifiers). Call `httpserver.RenderPage(w, renderer, pageName, data)`. Never call `template.ParseFiles` directly.
- **Adding a new page**: Add name to `Pages` slice in domain `RendererConfig` → create `web/templates/pages/<name>.html` with `{{define "title"}}` and `{{define "content"}}` blocks → create handler struct → register route.
- **Adding a new route checklist**: Handler struct → constructor → `ServeHTTP` → add to `RegisterRoutes` → write handler test → add route test proving expected status code.
- **Embed directive**: `//go:embed all:static all:templates` — `all:` prefix required. New assets in `web/static/`, new templates in `web/templates/`.
- **Static assets**: `http.FileServerFS` on `/static/{path...}` via `fs.Sub(staticFS, "static")` with `CacheControl` middleware.
- **Server wiring order**: Domain services → inbound handlers → register routes → `http.Server` with utils/html timeout constants → graceful shutdown via `signal.NotifyContext` + `srv.Shutdown`.

### Testing Rules

- **Naming**: `Test_<Unit>_With_<Condition>_Should_<Outcome>` — every test follows this exact pattern.
- **Structure**: Strict Arrange/Act/Assert with explicit `// Arrange`, `// Act`, `// Assert` comments. No exceptions.
- **Parallelism**: Every test starts with `t.Parallel()`.
- **Assertions**: `assert.That(t, "description", got, expected)` — four arguments, always. No other assertion library, no `if/t.Fatalf` for assertions.
- **External test package**: Tests use `package <name>_test`, accessing only exported symbols.
- **Test helpers**: Shared fixtures in `testhelper_test.go`. Lowercase unexported functions named `new<Purpose><Type>` (e.g., `newTestRenderer`, `newBrokenRenderer`). First line: `t.Helper()`. Use `t.Fatalf` for setup failures only.
- **Happy path + error path**: Every handler requires at minimum one happy path test and one error path test. Use broken fixtures (e.g., `newBrokenRenderer`) to test graceful degradation.
- **No mocking frameworks**: Use stdlib test doubles: `httptest.NewRequest`, `httptest.NewRecorder`, `testing/fstest.MapFS`. Only create custom test doubles to simulate failure conditions.
- **Route-level integration tests**: Every route in `RegisterRoutes` gets a test that constructs real handlers with real test fixtures → builds the mux → sends a request → asserts status code and response. Not just isolated handler tests.
- **Test boundary mirrors architecture**: `inbound/` tests verify HTTP behavior (request → response). `domain/` tests verify pure logic (no I/O). `outbound/` tests verify adapter behavior against real or faked externals.
- **Integration tests**: `//go:build integration` build tag on first line. Run with `go test ./... -tags=integration`. Separate from unit tests.

### Code Quality & Style Rules

- **golangci-lint must pass with zero issues**: Run `golangci-lint run ./...` after every change. Never modify `.golangci.yml` to suppress findings — fix the code instead.
- **All linters enabled by default**: Config uses `default: all` with documented disables. Disabled linters reflect architectural decisions (e.g., `ireturn` off for hexagonal interface returns, `exhaustruct` off for constructor-based init, `forbidigo` off for CLI output). Do not work around disabled linters.
- **No inline `//nolint` directives**: If a linter needs exemption, it belongs in `.golangci.yml` with a justification comment. Never silence individual instances in code.
- **gofmt rewrites enforced**: `interface{}` → `any`, `a[b:len(a)]` → `a[b:]`. Write the modern form directly.
- **File naming**: Lowercase with underscores. Handlers: `handler_<name>.go`. Tests: `handler_<name>_test.go`. Routes: `routes.go`. Test helpers: `testhelper_test.go`. Domain types: named by purpose (e.g., `renderer.go`), not generic (`types.go`, `models.go`).
- **One file per handler**: Each HTTP handler gets its own file. Do not combine multiple handlers.
- **Template structure**: `web/templates/layouts/base.html` (skeleton with `{{template}}` calls), `web/templates/pages/<name>.html` (defines `title` + `content` blocks), `web/templates/partials/<name>.html` (reusable components like `header`, `footer`). Follow this exact structure.
- **CSS conventions**: Vanilla CSS only — no preprocessors, no Tailwind, no CSS frameworks. Extend existing custom properties (`--color-*`, `--font-*`, `--radius`). Dark theme default, responsive at 768px breakpoint, respect `prefers-reduced-motion`.

### Development Workflow Rules

- **CLAUDE.md is authoritative for agentic loops**: Skill-First, Red-Green, Lint-Fix, Verify-Implement-Verify, Self-Healing loops are defined there. Do not duplicate — follow them.
- **No CI safety net**: There is no CI/CD pipeline. The agent IS the CI. Skipping lint or test gates means broken code ships.
- **Work on `main`**: No branching strategy established yet. Commit directly to `main`.
- **Run `go mod tidy`** after any change that touches imports. If `go.mod` changes unexpectedly (new `require` lines), an unauthorized dependency was introduced — revert.
- **Adding a new handler** (within existing context): Search forge → create handler file → write failing test → implement `ServeHTTP` → pass test → add to `RegisterRoutes` → write route-level test → run full suite → lint.
- **Adding a new page** (with UI): All of the above, plus: add page name to domain `RendererConfig.Pages` → create `web/templates/pages/<name>.html` with `title` + `content` blocks → extend `style.css` if needed.
- **Adding a new bounded context** (architectural change): Create `internal/<ctx>/domain/`, `inbound/`, `outbound/` directories → define domain types/interfaces → implement inbound handlers → create context's `RegisterRoutes` → wire into `cmd/server/main.go` → full test coverage → lint. Recognize this is fundamentally different from adding a handler.
- **Adding a new binary**: Create `cmd/<name>/main.go` with signal handling + `appVersion` → add build targets to `.justfile` → add build stage to `Dockerfile` → test build.
- **Post-implementation**: Update `README.md` and `CLAUDE.md` if behavior changed. Identify refactoring opportunities.

### Critical Don't-Miss Rules

**Anti-Patterns to Avoid:**
- Never introduce third-party modules (testify, chi, gorilla, viper, cobra, etc.). The answer is always stdlib + `loopforge-ai/utils`.
- Never use `http.HandleFunc` or `http.HandlerFunc` closures — struct-based handlers with dependency injection only.
- Never call `template.ParseFiles` directly — use `httpserver.NewRenderer` + `httpserver.RenderPage`.
- Never add inline `//nolint` comments or modify `.golangci.yml`.
- Never create `context.Background()` inside a function that receives a context parameter.
- Never put business logic in `inbound/` or I/O in `domain/`.
- Never create `shared/` or `common/` packages for cross-context types. Each bounded context owns its types. If two contexts need the same shape, they each define their own.

**Go 1.22+ ServeMux Gotchas:**
- Method patterns (`GET /health`) are method-specific — POST returns 405 automatically. Do not add manual method checking in handlers.
- HEAD requests are handled automatically for GET routes. Do not register HEAD separately.
- `/{$}` is exact root match. Without `{$}`, `/` is a catch-all that matches every path.
- `http.StripPrefix` is used for static routes — understand the prefix being stripped or paths will break.

**Template Gotchas:**
- `RenderPage` returns 500 if template execution fails — but missing template fields render as empty silently (no error). Manually verify template `{{.Field}}` references match the data struct.
- The broken renderer test pattern (`newBrokenRenderer`) catches execution errors but not missing fields. Always visually verify new template fields.

**HTTP Response Gotchas:**
- `httptest.NewRecorder` captures the first `WriteHeader` call. Error paths must write error status before any success path writes. `RenderPage` handles this correctly — follow its pattern.

**Security:**
- `SecurityHeaders` middleware must be the outermost wrapper on all page/API handlers — it sets CSP, X-Frame-Options, etc. Never register a handler without it.
- Middleware order matters: `SecurityHeaders(Log(Recover(ContentType(handler))))` — security headers first.
- Docker runtime uses non-root `app` user — never switch to root in Dockerfile.
- `gosec` exclusions (G301, G703, G704, G706) have documented justifications — do not re-enable without understanding the context.

---

## Usage Guidelines

**For AI Agents:**
- Read this file before implementing any code
- Follow ALL rules exactly as documented
- When in doubt, prefer the more restrictive option
- Update this file if new patterns emerge

**For Humans:**
- Keep this file lean and focused on agent needs
- Update when technology stack changes
- Review quarterly for outdated rules
- Remove rules that become obvious over time

Last Updated: 2026-03-14
