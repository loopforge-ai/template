---
stepsCompleted: [1, 2, 3, 4, 5, 6, 7, 8]
lastStep: 8
status: 'complete'
completedAt: '2026-03-14'
inputDocuments: ['_bmad-output/planning-artifacts/prd.md', '_bmad-output/project-context.md', 'docs/index.md', 'docs/project-overview.md', 'docs/architecture.md', 'docs/api-contracts.md', 'docs/source-tree-analysis.md', 'docs/development-guide.md', 'docs/deployment-guide.md']
workflowType: 'architecture'
project_name: 'template'
user_name: 'Andy'
date: '2026-03-14'
---

# Architecture Decision Document

_This document builds collaboratively through step-by-step discovery. Sections are appended as we work through each architectural decision together._

## Project Context Analysis

### Requirements Overview

**Functional Requirements (38 FRs):**
- FR1-FR5: Template structure — hexagonal patterns through single bounded context, three entry points, minimality verification
- FR6-FR10: Bootstrap consumption — replacement safety, directory exclusions, identifier naming, post-bootstrap README
- FR11-FR20: Pattern demonstration — two handler types (stateless/stateful), route registration, middleware, templates, composition root, graceful shutdown
- FR21-FR26: Testing patterns — AAA structure, happy/error paths, route-level integration, shared fixtures, external test packages, assertion/parallelism conventions
- FR27-FR30: Quality assurance — validation checklist, post-bootstrap verification, replacement audit, placeholder categorization
- FR31-FR35: Developer experience — bootstrap-to-browser, self-documenting code, handler replication, page addition, bounded context creation
- FR36-FR38: Forge integration — MCP connectivity, code-as-spec, convention extractability

**Non-Functional Requirements (18 NFRs):**
- Code quality: Zero lint issues with all linters, no inline suppressions, single-purpose files, consistent imports (NFR1-4)
- Maintainability: Pattern map validation in <5 min, conventions in code not docs, predictable file naming (NFR5-7)
- Portability: Cross-platform Go builds, Docker compatibility, justfile as optional convenience (NFR8-10)
- Security patterns: SecurityHeaders on all handlers, non-root Docker, no hardcoded secrets, documented gosec exclusions (NFR11-14)
- Accessibility: Semantic HTML, aria-labels, lang attribute, reduced-motion support (NFR15-18)

**FR Categorization for Brownfield Refinement:**

| Category | FRs | Architectural Implication |
|----------|-----|--------------------------|
| Already satisfied | FR1, FR4-5, FR11-FR20, FR21-FR26, FR31-FR34, FR36 | No work — existing architecture handles these |
| Drives MVP work | FR8-FR9, FR27-FR30, FR37-FR38 | Active architectural decisions needed |
| Satisfied after MVP | FR2-3, FR6-7, FR10, FR35 | Verified once MVP work completes |

**Scale & Complexity:**
- Primary domain: Go backend scaffold with server-side rendered web dashboard
- Complexity level: Low volume (17 files), high quality requirements (each file exemplary, propagates to all projects)
- Architectural components: 5 (domain layer, inbound layer, web assets layer, entry points layer, configuration layer)

### Technical Constraints & Dependencies

| Constraint | Source | Architectural Impact |
|-----------|--------|---------------------|
| stdlib + loopforge-ai/utils only | PRD, Project Context | No third-party packages for any functionality — use utils or implement from scratch |
| CGO_ENABLED=0 | Project Context | Static binaries only, no C library dependencies |
| Single go.mod, three binaries | Project Context, FR4 | Modulith structure — shared internal packages, independent entry points |
| Bootstrap string replacement | FR6-FR9, Journey 1 | Constrains identifier naming — no `template`/`Template` in Go source except `html/template` |
| Code-as-spec | FR37, NFR6, Journey 4 | Every convention must be visible in source code, not just documentation |
| Pattern minimality | FR2-FR3, Journey 5 | Every file must map to exactly one pattern — no redundancy, no gaps |
| Go 1.22+ ServeMux | Project Context | Method-specific routing, exact match patterns, automatic HEAD handling |
| Existing test suite | 7 tests, 4 files | Any change must preserve or provide equivalent test coverage |

### Cross-Cutting Concerns

**Constraining Concerns (limit what we can do):**

| Concern | Impact | Example |
|---------|--------|---------|
| Bootstrap replacement safety | Limits identifier naming and file content | `RendererConfig` not `TemplateConfig` — can't contain `template` substring |
| Convention consistency | All files must follow identical patterns | Alphabetical ordering, import grouping must be universal — no exceptions |
| Existing test suite preservation | Any change must maintain test coverage | RendererConfig rename requires updating 3 files: `renderer.go`, `main.go`, `testhelper_test.go` |

**Validation Concerns (verify what we did):**

| Concern | Verification Method |
|---------|-------------------|
| Pattern minimality invariant | Walk pattern map — each file maps to one pattern, no file removable |
| Quality contract propagation | Run lint/test/build in template AND post-bootstrap |
| Self-documenting readability | Read each file in isolation — purpose clear within 10 lines |

## Starter Template Evaluation

### Primary Technology Domain

Go backend modulith with server-side rendered web — all technology decisions pre-established.

### Technology Stack (Pre-Established)

This project IS a starter template. The technology stack is not a choice — it's the product. All decisions are locked as hard constraints in the project context (67 rules).

| Layer | Technology | Version | Rationale |
|-------|-----------|---------|-----------|
| Language | Go | 1.25.6 | Exact version; leverage 1.22+ ServeMux features |
| HTTP | stdlib `net/http` | — | Zero-dependency HTTP server; `ServeMux` with method routing |
| Templates | stdlib `html/template` | — | Server-side rendering via `embed.FS` |
| Shared utils | `loopforge-ai/utils` | v0.1.0 | Organizational trust boundary; canonical implementations |
| Linter | golangci-lint | v2 | All linters enabled; architectural decisions encoded in disables |
| Testing | stdlib `testing` + `assert.That` | — | Single assertion function; no mocking frameworks |
| Container | Docker | Alpine 3.21 | Multi-stage build; non-root runtime; CGO_ENABLED=0 |
| Task runner | just | latest | Convenience layer; not required for core validation |

### Rejected Alternatives

| Category | Rejected | Why |
|----------|----------|-----|
| Router | chi, gorilla/mux, echo | stdlib `ServeMux` (Go 1.22+) covers method routing and path params — third-party adds dependency for unused features |
| Logger | zap, zerolog | stdlib `log/slog` (Go 1.21+) provides structured logging — it's the standard answer now |
| Assertions | testify | `assert.That` is one function, four arguments. Testify adds 50+ functions, most unused — violates controlled dependency surface |
| Config | viper, envconfig | `env.Get(key, default)` handles the only pattern needed. Viper adds file-based config we explicitly don't want |
| Middleware | negroni, alice | `loopforge-ai/utils/html` provides the exact middleware chain needed — no reason to add a dependency for function composition |

### Why No Starter Evaluation Is Needed

- The template already exists as a working codebase (17 source files, 7 tests passing)
- Technology decisions are hard constraints, not choices — changing any would break the controlled dependency surface promise
- The "initialization command" is Forge's `skill_bootstrap` — not a framework CLI
- This architecture document formalizes existing decisions rather than making new ones
- Formalizing these decisions IS the architectural work — downstream agents reference this document for implementation constraints

**Note:** No project initialization story is needed — the project is already initialized. The first implementation story is the `RendererConfig` rename (MVP capability #1).

## Core Architectural Decisions

### Decision Priority Analysis

**Critical Decisions (Block MVP Implementation):**

| # | Decision | Choice | Rationale |
|---|----------|--------|-----------|
| ADR-1 | Template-only directory declaration | `.bootstrapignore` file | Machine-readable, self-documenting, familiar `.gitignore` pattern |
| ADR-2 | Placeholder categorization format | `TODO(bootstrap)` / `TODO(developer)` markers | Single pattern, in-place self-documenting, grep-friendly |
| ADR-3 | Validation procedure location | `just validate` + `VALIDATION.md` | Executable + documented; VALIDATION.md is template-only |
| ADR-4 | Code-as-spec verification | Convention-to-file mapping table in `VALIDATION.md` | Explicit contract; if convention isn't in table, it's invisible to Forge |

**Deferred Decisions (Post-MVP):**

| # | Decision | Phase | Why Deferred |
|---|----------|-------|-------------|
| ADR-5 | Outbound layer pattern | 2a | No outbound in MVP — pattern decided when implementing |
| ADR-6 | Multi-context wiring pattern | 2b | Single context in MVP — composition root pattern decided when adding second context |
| ADR-7 | Integration test pattern | 2a | No integration tests in MVP — pattern decided when implementing |
| ADR-8 | CLI expansion pattern | 3 | CLI is placeholder in MVP — pattern decided when expanding |
| ADR-9 | Automated bootstrap validation test | 2a | Scripted validation in MVP — automation decided in Phase 2a |

### ADR-1: Template-Only Directory Declaration

**Decision:** Add `.bootstrapignore` file to the template repository root.

**Format:**
```
_bmad/
_bmad-output/
.claude/
docs/
.bootstrapignore
VALIDATION.md
```

**Phasing:** `.bootstrapignore` is a declaration of intent. The current bootstrap skill does not implement `.bootstrapignore` support. Until Forge adds this capability, the file serves as documentation. The bootstrap skill team implements support as a separate Forge task.

**Implications:**
- Bootstrap skill reads `.bootstrapignore` and skips listed paths during clone/copy (once implemented)
- `.bootstrapignore` itself is listed (it's template-only infrastructure)
- `VALIDATION.md` is listed (maintainer-only, not relevant to bootstrapped projects)
- Until Forge implements support, this file documents which directories are template-only

**Affects:** FR8, Journey 1, Journey 5

### ADR-2: Placeholder Categorization Format

**Decision:** Use categorized TODO markers: `TODO(bootstrap)` and `TODO(developer)`.

**Scope clarification:** README.md is fully rewritten by bootstrap — its TODOs don't need categorization. The audit focuses on files that survive bootstrap (primarily CLAUDE.md). Post-bootstrap output must contain zero uncategorized `TODO` markers.

**Current markers to migrate:**
- CLAUDE.md: `<!-- TODO: add go run ./cmd/... for project binary -->` → `<!-- TODO(developer): add go run ./cmd/... for project binary -->`
- README.md TODOs: Moot — file is fully rewritten by bootstrap. No categorization needed.

**Verification command:**
```bash
grep -rn 'TODO' . | grep -v 'TODO(bootstrap)' | grep -v 'TODO(developer)' | grep -v '_bmad' | grep -v 'node_modules'
# Should return zero results — catches all uncategorized TODOs
```

**Implications:**
- Existing `<!-- TODO -->` markers must be audited and recategorized (only in files surviving bootstrap)
- Uncategorized `TODO` (without parenthesized category) in post-bootstrap output is a validation failure
- Bootstrap skill may need updating to recognize `TODO(bootstrap)` as replacement triggers

**Affects:** FR30, Journey 5, MVP capability #3

### ADR-3: Validation Procedure

**Decision:** Dual approach — executable `just validate` recipe + documented `VALIDATION.md`.

**`just validate` recipe:**
```
validate: lint test
    docker build -t template:validate .
    @echo "Template quality contract: PASS"
```

**`VALIDATION.md` contains:**
1. Template quality gates (lint, test, build, Docker)
2. Pattern map verification checklist
3. Convention-to-file mapping (ADR-4)
4. Placeholder audit procedure (grep command from ADR-2)
5. Post-bootstrap validation procedure (scripted commands)
6. Bootstrap-to-browser verification steps

**`VALIDATION.md` is template-only** — listed in `.bootstrapignore`, not carried to bootstrapped projects.

**Implications:**
- `just validate` is the quick-check for maintainers (NFR10: just is convenience, not required)
- `VALIDATION.md` is the comprehensive reference when just isn't available
- Post-bootstrap validation is scripted commands, not automated tests (automation deferred to Phase 2a)

**Affects:** FR27-FR28, Journey 5, MVP capability #5

### ADR-4: Code-as-Spec Verification

**Decision:** Convention-to-file mapping table in `VALIDATION.md`, referencing identifiers (not line numbers) for stability across edits.

**Convention-to-File Mapping:**

| Convention (from CLAUDE.md) | Demonstrated In | Identifier |
|----------------------------|----------------|-----------|
| Alphabetical file ordering | All .go files | Declarations ordered alphabetically in every file |
| Constructor pattern | handler_health.go | `NewHealthHandler` — first function after type |
| Constructor with deps | handler_index.go | `NewIndexHandler(renderer, version)` |
| Error wrapping | routes.go | `fmt.Errorf("sub static fs: %w", err)` in `RegisterRoutes` |
| Struct-based handlers | handler_health.go, handler_index.go | `HealthHandler.ServeHTTP`, `IndexHandler.ServeHTTP` |
| Import grouping | All .go files | stdlib, blank line, internal — consistent everywhere |
| slog logging | cmd/server/main.go | `slog.Info("server starting", ...)` in `main` |
| Context propagation | cmd/server/main.go | `signal.NotifyContext` → passed to `srv.Shutdown` |
| Graceful shutdown | cmd/server/main.go | `signal.NotifyContext` + `srv.Shutdown` in `main` |
| Middleware chain | routes.go | `SecurityHeaders(Log(Recover(ContentType(...))))` in `RegisterRoutes` |
| embed.FS | web/embed.go | `var FS embed.FS` with `//go:embed all:static all:templates` |
| Test naming | All _test.go files | `Test_Unit_With_Condition_Should_Outcome` pattern |
| AAA structure | All _test.go files | `// Arrange`, `// Act`, `// Assert` comments |
| t.Parallel() | All _test.go files | First line of every test function |
| assert.That | All _test.go files | Four arguments: `t, description, got, expected` |
| External test package | All _test.go files | `package inbound_test` (not `package inbound`) |
| Shared fixtures | testhelper_test.go | `newTestRenderer`, `newBrokenRenderer` with `t.Helper()` |
| env.Get for config | cmd/server/main.go | `env.Get("SERVER_ADDR", httpserver.DefaultAddr)` |

**Contract rule:** If a convention is in CLAUDE.md but not in this table, it's either (a) invisible to Forge and must be added to the template, or (b) not demonstrable in the current file set and should be removed from CLAUDE.md.

**Implications:**
- This table is the authoritative link between CLAUDE.md conventions and template code
- Maintained in `VALIDATION.md` alongside the pattern map
- Reviewed during every template release (Journey 5)

**Affects:** FR37-FR38, NFR6, Journey 4

### Decision Impact Analysis

**Implementation Sequence:**
1. ADR-2 first (placeholder recategorization) — affects no code, only markers in CLAUDE.md
2. ADR-1 second (`.bootstrapignore`) — new file, no code changes
3. ADR-3 third (`VALIDATION.md` + `just validate`) — new file + justfile update
4. ADR-4 fourth (convention mapping) — content within `VALIDATION.md`
5. Then: `RendererConfig` rename (MVP capability #1) — code change across 3 files, validated by ADR-3/4

**Cross-Decision Dependencies:**
- ADR-3 depends on ADR-1 (`VALIDATION.md` listed in `.bootstrapignore`)
- ADR-4 lives inside ADR-3's `VALIDATION.md`
- ADR-2's verification command is part of ADR-3's validation procedure
- All ADRs feed into ADR-3 as the single validation reference

## Implementation Patterns & Consistency Rules

### Pattern Source of Truth

The canonical implementation patterns for this project are documented in two places:
- **`project-context.md`** (67 rules) — the comprehensive ruleset for AI agents
- **CLAUDE.md** — the authoritative reference loaded into every agent context

This architecture document does NOT duplicate those rules. Instead, it addresses **conflict points specific to working on a project scaffold** — patterns that agents wouldn't know from a normal Go project.

### Bootstrap-Aware Coding Patterns

These patterns apply specifically because this is a template consumed by string replacement:

**Identifier Naming Rule:**
- Before creating any new Go identifier, check: does it contain `template` or `Template` as a substring?
- If yes: choose an alternative name. `RendererConfig` not `TemplateConfig`. `PageData` not `TemplatePageData`.
- Verification: `grep -rn 'template\|Template' --include='*.go' --exclude-dir=_bmad . | grep -v 'html/template'` must return zero results.

**String Literal Rule:**
- Avoid the word `template` in string literals, log messages, or comments in Go files
- Exception: `html/template` import path (legitimate Go identifier)
- Exception: `web/templates/` path in embed directive (directory name, not a replacement target — bootstrap replaces the module path, not directory names)

**Import Alias Convention:**
- Alias when the package name doesn't convey its role at the call site
- Follow existing style: short, role-descriptive, lowercase (`httpserver` for `utils/html`, `dashboard` for domain package)
- New aliases must be consistent with this pattern — no `capitalCase` or `SCREAMING_CASE` aliases

**New File Pattern:**
- New Go files follow existing naming: `<purpose>.go` for source, `<purpose>_test.go` for tests
- New non-Go files (`.bootstrapignore`, `VALIDATION.md`) are template-only and listed in `.bootstrapignore`
- No new file should introduce a pattern that isn't already demonstrated — MVP adds no new patterns

### Cross-Repository Dependency Pattern

When a needed capability doesn't exist in `loopforge-ai/utils`:
1. Identify the need in the template
2. Open an issue on `loopforge-ai/utils` describing the required package/function
3. Implement and release a new utils version
4. Update the template's `go.mod` to reference the new utils version
5. Never add the capability directly to the template — utils is the canonical home

### Configuration File Patterns

**`.bootstrapignore` format:**
- One path per line, trailing `/` for directories
- Comments with `#`
- No glob patterns — explicit paths only (simplicity over flexibility)

**`VALIDATION.md` structure:**
- Section per validation category (quality gates, pattern map, convention map, placeholder audit, post-bootstrap)
- Each section has: purpose, commands to run, expected output, pass/fail criteria
- Machine-friendly command blocks (copy-pasteable)

**`.justfile` additions:**
- New recipes follow existing pattern: verb-noun naming, dependencies declared inline
- `validate` recipe depends on `lint` and `test` (existing recipes)

### Pattern Map Update Rule

Any code change that affects the pattern map (adding, removing, or modifying a file's demonstrated pattern) requires updating the pattern map in the same commit. The pattern map must never drift from reality. This applies to:
- The pattern map table in `VALIDATION.md`
- The convention-to-file mapping in `VALIDATION.md`
- The pattern map in the PRD (if referenced)

### Anti-Patterns for Template Development

| Anti-Pattern | Why It's Wrong | Correct Pattern |
|-------------|---------------|-----------------|
| Adding `//nolint` to pass linter | Hides issues that propagate to all projects | Fix the code; update `.golangci.yml` with documented rationale if needed |
| Using `template` in new identifier names | Breaks bootstrap string replacement | Use context-appropriate names (`Renderer`, `Page`, `Dashboard`) |
| Adding a file that demonstrates the same pattern as an existing file | Violates minimality invariant | Verify against pattern map before creating any file |
| Adding an external dependency | Breaks controlled dependency surface | Use stdlib or `loopforge-ai/utils`; if utils doesn't have it, follow cross-repo dependency pattern |
| Writing a convention in CLAUDE.md without demonstrating it in code | Convention invisible to Forge (code-as-spec violation) | Add code that demonstrates the convention, then document it |
| Adding comments explaining what pattern a file demonstrates | Violates self-documenting code principle — telling instead of showing | The code demonstrates the pattern by being clear, not by saying what it demonstrates |
| Letting the pattern map drift from code | Downstream agents reference stale map | Update pattern map in the same commit as the code change |

### Enforcement

**Pre-commit verification:**
All MVP changes must pass before commit:
1. `golangci-lint run ./...` — zero issues
2. `go test ./...` — all pass
3. `grep -rn 'template\|Template' --include='*.go' --exclude-dir=_bmad . | grep -v 'html/template'` — zero results (bootstrap safety)
4. Pattern map check — every file still maps to exactly one pattern

**Pre-release verification:**
`just validate` or the full `VALIDATION.md` procedure before tagging any release.

## Project Structure & Boundaries

### Complete Project Directory Structure

```
template/
├── cmd/                                    # Entry points (3 binaries)
│   ├── cli/
│   │   └── main.go                         # CLI binary — signal-aware placeholder
│   ├── mcp/
│   │   └── main.go                         # MCP server — Forge integration
│   └── server/
│       └── main.go                         # HTTP server — composition root
├── internal/                               # Private application code
│   └── dashboard/                          # Bounded context: Dashboard (example)
│       ├── domain/
│       │   └── renderer.go                 # PageData + RendererConfig (renamed from TemplateConfig)
│       └── inbound/
│           ├── handler_health.go           # GET /health — stateless handler
│           ├── handler_health_test.go      # Unit test — stateless handler
│           ├── handler_index.go            # GET / — stateful handler with DI
│           ├── handler_index_test.go       # Unit test — happy + error path
│           ├── routes.go                   # RegisterRoutes — middleware wiring
│           ├── routes_test.go              # Route-level integration tests
│           └── testhelper_test.go          # Shared test fixtures
├── web/                                    # Embedded web assets
│   ├── embed.go                            # embed.FS declaration
│   ├── static/
│   │   ├── css/
│   │   │   └── style.css                   # Dark theme, responsive, CSS custom props
│   │   └── vendor/
│   │       └── .gitkeep
│   └── templates/
│       ├── layouts/
│       │   └── base.html                   # Layout composition
│       ├── pages/
│       │   └── index.html                  # Page with named blocks
│       └── partials/
│           ├── header.html                 # Partial — static (no data binding)
│           └── footer.html                 # Partial — dynamic ({{.Version}})
├── .bootstrapignore                        # [NEW - MVP] Template-only directory declaration
├── .gitignore                              # Standard Go + IDE ignores
├── .golangci.yml                           # Linter config (v2, all linters)
├── .justfile                               # Task runner (build, test, lint, docker, validate)
├── .mcp.json                               # Forge MCP server configuration
├── CLAUDE.md                               # AI agent guidelines and conventions
├── Dockerfile                              # Multi-stage Docker build (3 binaries)
├── README.md                               # Project readme (bootstrap-rewritten)
├── VALIDATION.md                           # [NEW - MVP] Validation procedure + convention map
├── go.mod                                  # Go module declaration
│
├── _bmad/                                  # [TEMPLATE-ONLY] BMAD workflow config
├── _bmad-output/                           # [TEMPLATE-ONLY] BMAD planning artifacts
├── .claude/                                # [TEMPLATE-ONLY] Claude Code settings
└── docs/                                   # [TEMPLATE-ONLY] Template documentation
```

**Note:** `go.sum` is not listed because it's a generated file — it appears after `go mod tidy` runs (during bootstrap or first build). It IS committed to VCS as a dependency verification lockfile.

### MVP Change Footprint

MVP touches only **6 files** (3 modified + 3 new) out of 17+ source files:

| Action | Files |
|--------|-------|
| **Modified** | `domain/renderer.go`, `cmd/server/main.go`, `inbound/testhelper_test.go` (RendererConfig rename) |
| | `CLAUDE.md` (placeholder recategorization) |
| | `.justfile` (add validate recipe) |
| **New** | `.bootstrapignore`, `VALIDATION.md` |
| **Unchanged** | `handler_health.go`, `handler_index.go`, `handler_health_test.go`, `handler_index_test.go`, `routes.go`, `routes_test.go`, `embed.go`, all HTML templates, `style.css`, `cmd/mcp/main.go`, `cmd/cli/main.go`, `.golangci.yml`, `.gitignore`, `Dockerfile`, `.mcp.json`, `go.mod` |

### Architectural Boundaries

**Hexagonal Layer Boundaries:**

```
cmd/server/main.go (composition root)
        │
        ├── imports → internal/dashboard/domain/     (types, config)
        ├── imports → internal/dashboard/inbound/     (handlers, routes)
        └── imports → web/                            (embed.FS)

internal/dashboard/inbound/  ──depends on──→  internal/dashboard/domain/
                             ──may reference──→  web/ (for embed.FS type)
                             ──never imports──→  outbound/ (doesn't exist yet)

internal/dashboard/domain/   ──imports only──→  stdlib
                             ──never imports──→  inbound/, outbound/, web/
```

**Bootstrap Boundary:**

| Category | Files | Survives Bootstrap | Notes |
|----------|-------|-------------------|-------|
| Module file | `go.mod` | Yes (module path replaced) | Critical — wrong path means nothing compiles |
| Project source | All `.go` files | Yes (module path in imports replaced) | Internal package paths updated |
| Web assets | `.html`, `.css` files | Yes (project name in HTML replaced) | Capitalized name in templates |
| Build config | `.justfile`, `.mcp.json`, `Dockerfile` | Yes (project name replaced) | Binary names updated |
| Agent config | `CLAUDE.md` | Yes (project references replaced) | Conventions carried through |
| Linter config | `.golangci.yml`, `.gitignore` | Yes (unchanged) | Carried through as-is |
| Documentation | `README.md` | Rewritten entirely | Bootstrap generates project-specific README |
| Validation | `VALIDATION.md`, `.bootstrapignore` | No (template-only) | Listed in `.bootstrapignore` |
| BMAD infra | `_bmad/`, `_bmad-output/`, `.claude/`, `docs/` | No (template-only) | Listed in `.bootstrapignore` |

### Data Flow

**HTTP Server Flow:**
```
Developer: go run ./cmd/server
    │
    ▼
main.go: env.Get("SERVER_ADDR") → NewRenderer(web.FS, RendererConfig)
    → NewHealthHandler() → NewIndexHandler(renderer, version)
    → RegisterRoutes(health, index, web.FS) → http.Server{} → ListenAndServe
    │
    ▼
Browser: GET / → ServeMux → SecurityHeaders → Log → Recover → ContentType
    → IndexHandler.ServeHTTP → RenderPage("index", PageData{}) → HTML
```

**MCP Server Flow:**
```
Developer: Claude Code opens project with .mcp.json
    │
    ▼
cmd/mcp/main.go: mcp.NewServer(projectName, version) → server.Serve(ctx)
    │
    ▼
Forge connects via stdio → Developer asks for code generation
    → Forge reads project structure → generates code matching template conventions
```

**Bootstrap Flow:**
```
Developer: invoke skill_bootstrap(module_path, project_dir, project_name)
    │
    ▼
git clone template → remove .git → replace module path in all text files
    → replace project name in specific files → replace Title in HTML
    → git init → go mod tidy
    │
    ▼
Output: Valid Go project → go run ./cmd/server → browser shows dashboard
```

### Requirements to Structure Mapping

**FR Categories → Directory Mapping:**

| FR Category | Primary Location | Supporting Files | Phase |
|-------------|-----------------|-----------------|-------|
| Template Structure (FR1-5) | `internal/dashboard/` | Pattern map in `VALIDATION.md` | MVP |
| Bootstrap Consumption (FR6-10) | `.bootstrapignore`, `README.md` | Bootstrap skill in `forge/codegen/skills/` | MVP |
| Pattern Demonstration (FR11-20) | `internal/dashboard/`, `web/`, `cmd/` | Every source file | Already satisfied |
| Testing Patterns (FR21-26) | `internal/dashboard/inbound/*_test.go` | `testhelper_test.go` | Already satisfied |
| Quality Assurance (FR27-30) | `VALIDATION.md`, `.justfile` | `.bootstrapignore` | MVP |
| Developer Experience (FR31-35) | All source files (code-as-tutorial) | `CLAUDE.md` | Already satisfied |
| Forge Integration (FR36-38) | `cmd/mcp/main.go`, all source files | `.mcp.json` | MVP (verify) |

## Architecture Validation Results

### Coherence Validation

**Decision Compatibility:**
- Technology stack is locked and internally consistent — Go 1.25.6, stdlib, utils v0.1.0. No version conflicts possible (zero external dependencies).
- ADR-1 through ADR-4 are sequentially dependent and non-contradictory: ADR-2 feeds ADR-3, ADR-3 contains ADR-4, ADR-1 gates ADR-3.
- Implementation patterns align with ADRs: bootstrap-aware coding patterns directly support ADR-1/ADR-2 decisions.

**Pattern Consistency:**
- All naming patterns consistent: lowercase with underscores for files, PascalCase for types, camelCase for local variables — matching existing code.
- Convention-to-file mapping (ADR-4) verifies all 18 conventions map to specific identifiers. No orphan conventions.

**Structure Alignment:**
- MVP change footprint (6 modified/new files) aligns with 7 MVP capabilities. No structural changes unaccounted for.
- RendererConfig rename footprint verified: 3 Go files (`renderer.go`, `main.go`, `testhelper_test.go`) + 3 documentation files (`project-context.md`, `docs/architecture.md`, `docs/source-tree-analysis.md`) = 6 files for the rename alone.

**Verdict:** No coherence issues found.

### Requirements Coverage Validation

**Functional Requirements (38 FRs):**

| Category | FRs | Coverage Status |
|----------|-----|----------------|
| Template Structure (FR1-5) | Already satisfied + pattern map in VALIDATION.md | Fully covered |
| Bootstrap Consumption (FR6-10) | ADR-1 + ADR-2 + RendererConfig rename | Fully covered |
| Pattern Demonstration (FR11-20) | Already satisfied by existing source files | Fully covered |
| Testing Patterns (FR21-26) | Already satisfied by existing test files | Fully covered |
| Quality Assurance (FR27-30) | ADR-3 (VALIDATION.md) + ADR-4 (convention map) | Fully covered |
| Developer Experience (FR31-35) | Already satisfied + code-as-spec verification | Fully covered |
| Forge Integration (FR36-38) | MCP binary exists + code-as-spec verified by ADR-4 | Fully covered |

**Non-Functional Requirements (18 NFRs):** All covered by existing architecture + enforcement rules.

**Verdict:** All 38 FRs and 18 NFRs have architectural support. Zero gaps.

### Implementation Readiness Validation

**Decision Completeness:**
- 4 critical ADRs with choice, rationale, implications, affected FRs
- 5 deferred ADRs with phase and deferral reason
- Implementation sequence with cross-dependencies
- Concrete examples (file formats, commands, tables)

**Structure Completeness:**
- Complete project tree with all files including 3 new MVP files
- Bootstrap boundary table with `go.mod` explicitly listed
- MVP change footprint with exact file list
- FR-to-directory mapping with phase column

**Pattern Completeness:**
- 7 anti-patterns documented
- Pre-commit and pre-release enforcement commands
- Cross-repo dependency workflow
- Pattern map update rule

**Verdict:** AI agents have sufficient guidance for consistent MVP implementation.

### Gap Analysis Results

**Critical Gaps:** None.

**Important Gaps (documented, not blocking):**

| Gap | Impact | Mitigation |
|-----|--------|-----------|
| Bootstrap skill lacks `.bootstrapignore` support | Template-only dirs cloned into bootstrapped projects | Documented as Forge team responsibility; file serves as declaration of intent |
| No automated post-bootstrap validation test | Validation is scripted, not CI-enforced | Scripted procedure in VALIDATION.md; automation in Phase 2a |
| `project-context.md` line 50 references `TemplateConfig` | Stale reference after rename | Update as part of RendererConfig rename (step 6 of implementation) |
| `docs/architecture.md` references `TemplateConfig` | Stale reference in component tables | Update as part of RendererConfig rename |
| `docs/source-tree-analysis.md` references `TemplateConfig` | Stale reference in file descriptions | Update as part of RendererConfig rename |

### Architecture Completeness Checklist

**Requirements Analysis**
- [x] Project context analyzed (67 rules, 38 FRs, 18 NFRs)
- [x] Scale and complexity assessed (low volume, high quality bar)
- [x] Technical constraints identified (7 constraints + 1 existing test suite)
- [x] Cross-cutting concerns mapped (3 constraining, 3 validation)
- [x] FR categorization for brownfield (already satisfied vs drives MVP work)

**Architectural Decisions**
- [x] Critical decisions documented with rationale (ADR-1 through ADR-4)
- [x] Technology stack fully specified and locked
- [x] Rejected alternatives documented (5 categories)
- [x] Deferred decisions catalogued with phase (ADR-5 through ADR-9)

**Implementation Patterns**
- [x] Bootstrap-aware coding patterns established
- [x] Anti-patterns documented (7 entries)
- [x] Cross-repo dependency workflow defined
- [x] Pattern map update rule established
- [x] Import alias convention documented
- [x] Enforcement commands provided (pre-commit + pre-release)

**Project Structure**
- [x] Complete directory structure with all files
- [x] MVP change footprint identified (6 modified/new files)
- [x] Architectural boundaries defined (hexagonal + bootstrap)
- [x] Requirements-to-structure mapping with phase column
- [x] Data flow diagrams (HTTP, MCP, Bootstrap)

### Architecture Readiness Assessment

**Overall Status:** READY FOR IMPLEMENTATION

**Confidence Level:** High — brownfield refinement with small change surface (6 files), well-understood constraints (67 rules), and comprehensive validation artifacts.

**Key Strengths:**
- Architecture already exists and works — formalizing, not inventing
- Minimal MVP footprint — high confidence in scope containment
- Code-as-spec principle ensures architecture stays aligned with implementation
- Convention-to-file mapping provides verifiable traceability

**Areas for Future Enhancement:**
- Outbound layer pattern (Phase 2a)
- Multi-context composition pattern (Phase 2b)
- Automated bootstrap validation (Phase 2a)
- CI/CD pipeline integration (when needed)

### Implementation Handoff

**Document Authority Priority (highest to lowest):**
1. **CLAUDE.md** — coding convention authority, loaded first by every agent. Wins on conflict.
2. **project-context.md** — extended ruleset with 67 rules for unobvious details.
3. **This architecture document** — decision record, structure mapping, validation artifacts.

**Progressive Validation:**
Validation is progressively available during implementation:
- Steps 1-2 (placeholders, .bootstrapignore): Use existing checks — `golangci-lint`, `go test`
- Step 3 (VALIDATION.md created): Full validation procedure now available
- Step 4 (convention map populated): Code-as-spec verification now possible
- Steps 5-7 (rename + doc updates + final validation): Run complete VALIDATION.md procedure

**Implementation Priority (MVP):**
1. ADR-2: Recategorize TODO markers in CLAUDE.md
2. ADR-1: Create `.bootstrapignore`
3. ADR-3: Create `VALIDATION.md` + add `just validate` recipe
4. ADR-4: Populate convention-to-file mapping in VALIDATION.md
5. Rename `TemplateConfig` → `RendererConfig` (3 Go files: `renderer.go`, `main.go`, `testhelper_test.go`)
6. Update documentation references (3 doc files: `project-context.md`, `docs/architecture.md`, `docs/source-tree-analysis.md`)
7. Run full VALIDATION.md procedure to verify everything
