---
stepsCompleted: ['step-01-init', 'step-02-discovery', 'step-02b-vision', 'step-02c-executive-summary', 'step-03-success', 'step-04-journeys', 'step-05-domain-skipped', 'step-06-innovation-skipped', 'step-07-project-type', 'step-08-scoping', 'step-09-functional', 'step-10-nonfunctional', 'step-11-polish', 'step-12-complete']
status: 'complete'
completedAt: '2026-03-14'
inputDocuments: ['_bmad-output/project-context.md', 'docs/index.md', 'docs/project-overview.md', 'docs/architecture.md', 'docs/api-contracts.md', 'docs/source-tree-analysis.md', 'docs/development-guide.md', 'docs/deployment-guide.md', 'forge/codegen/skills/bootstrap.md']
documentCounts:
  briefs: 0
  research: 0
  brainstorming: 0
  projectDocs: 7
  projectContext: 1
  forgeSkills: 1
workflowType: 'prd'
classification:
  projectType: 'developer_tool_scaffold'
  domain: 'general'
  complexity: 'medium'
  projectContext: 'brownfield'
  architecture: 'modulith'
  consumers: ['forge_bootstrap_skill', 'developers']
---

# Product Requirements Document - template

**Author:** Andy
**Date:** 2026-03-14

## Executive Summary

**Template** is a minimal-but-complete Go reference architecture consumed by LoopForge Forge's bootstrap skill to scaffold production-grade modulith projects. It targets Go engineers and AI code generation agents who need a deterministic starting point embodying 2026 best practices: hexagonal architecture, domain-driven design, and a controlled dependency surface limited to stdlib and a first-party shared library.

The modulith architecture is a deliberate choice: single-binary deployment simplicity with the internal structure of bounded contexts that enforce clean boundaries through hexagonal layering. Each context could be extracted into a separate service if needed, but you don't pay the distributed systems tax upfront.

Three entry points ship by default — HTTP server, CLI, and MCP server — making every bootstrapped project Forge-connected from day one, ready for AI-assisted code generation without additional setup.

This is a brownfield refinement of an existing scaffold. The current phase focuses on making the template irreducibly minimal: refining the `dashboard` bounded context and improving the placeholder structure for bootstrap consumption. Future phases will add outbound layer examples, additional bounded contexts, and expanded patterns.

### What Makes This Special

- **Controlled dependency surface**: Only stdlib and `loopforge-ai/utils` (first-party, organizationally controlled). No supply chain risk, no version conflicts, no abandoned maintainer scenarios.
- **Zero decision fatigue**: No choosing a router. No configuring a logger. No debating test frameworks. No setting up Docker. No writing middleware. Every infrastructure decision is pre-made so the first commit is productive code.
- **Deterministic codebase**: Every bootstrapped project has identical structure, patterns, and conventions. Zero ramp-up between projects. Forge targets one canonical structure for code generation.
- **Proof of minimality**: The `dashboard` context demonstrates every hexagonal pattern with nothing redundant. Each file maps to exactly one pattern. Remove any file and a pattern is lost. Add any file and you've introduced bloat.
- **Self-documenting code**: Each file teaches its pattern through clarity, not comments. The code IS the tutorial.
- **Forge-connected by default**: The MCP binary integrates every bootstrapped project into the LoopForge ecosystem, enabling AI-assisted skill generation and code generation from day one.
- **Quality contract**: The template must pass all linters (all enabled), all tests green, and Docker build successful — not just in the template repo, but after every bootstrap operation. This is a testing contract between template and Forge.

### Project Classification

- **Type**: Developer Tool — Project Scaffold / Reference Architecture
- **Domain**: General (developer tooling / code generation infrastructure)
- **Complexity**: Medium (low domain complexity, high implementation quality bar — every pattern propagates to all bootstrapped projects)
- **Context**: Brownfield (refining existing template)
- **Architecture**: Modulith (modular monolith — single-binary deployment, bounded-context internal structure)
- **Consumers**: Forge bootstrap skill (automated string replacement) and developers (human, reading exemplary code)
- **Ecosystem**: Part of LoopForge AI (`loopforge-ai/template`, `loopforge-ai/forge`, `loopforge-ai/utils`)

## Success Criteria

### User Success

- **Two commands to running dashboard**: Bootstrap with Forge, `go run ./cmd/server`, browser open at `localhost:8080` — dashboard renders with project name. Zero manual configuration between bootstrap and working server.
- **Instant pattern recognition**: A Go engineer reads any file and immediately understands the pattern it demonstrates without external documentation.
- **First handler by replication**: Developer adds a custom handler by copying and modifying an existing handler file. The patterns are obvious enough that no guide is needed.

### Business Success

- **Forge ecosystem standard**: Every new LoopForge AI project starts from this template. Zero projects bypass the scaffold.
- **Convention compliance**: All bootstrapped projects pass lint (all linters), tests, and Docker build on first run.
- **Pattern propagation**: Template conventions are consistently replicated in Forge-generated code.

### Technical Success

- **Proof of minimality verified**: Every file in the `dashboard` context maps to exactly one demonstrated pattern (see pattern map). No file is redundant. No required pattern is missing.
- **Bootstrap contract holds**: After Forge's string replacement, no occurrence of `template`/`Template` remains in source files except legitimate Go identifiers (e.g., `html/template` import). The output project builds, tests, and lints cleanly for any valid Go project name.
- **Zero lint issues**: `golangci-lint run ./...` with `default: all`. No inline suppressions, no `//nolint`.
- **All tests green**: Happy path + error path coverage for every handler. Route-level integration tests for every registered route.
- **Docker builds cleanly**: Multi-stage build produces working binaries for all three entry points.

### Pattern Map (Minimality Verification)

| File | Pattern Demonstrated |
|------|---------------------|
| `domain/renderer.go` | Domain types + package-level config object |
| `inbound/handler_health.go` | Stateless handler (no dependencies, pure JSON) |
| `inbound/handler_index.go` | Stateful handler (dependency injection, template rendering) |
| `inbound/routes.go` | Route registration + middleware wiring + static asset serving |
| `inbound/handler_health_test.go` | Unit test for stateless handler |
| `inbound/handler_index_test.go` | Unit test with fixtures (happy path + error path) |
| `inbound/routes_test.go` | Route-level integration tests through mux |
| `inbound/testhelper_test.go` | Shared test fixtures pattern |
| `web/embed.go` | Embedded filesystem declaration |
| `web/templates/layouts/base.html` | Template layout composition |
| `web/templates/pages/index.html` | Page template with named blocks |
| `web/templates/partials/header.html` | Partial without data binding (static) |
| `web/templates/partials/footer.html` | Partial with data binding (`{{.Version}}`) |
| `web/static/css/style.css` | Static asset serving + theming |
| `cmd/server/main.go` | Composition root + graceful shutdown + env config |
| `cmd/mcp/main.go` | MCP entry point (Forge integration) |
| `cmd/cli/main.go` | CLI entry point (signal-aware placeholder) |

### Measurable Outcomes

| Metric | Target |
|--------|--------|
| Commands from bootstrap to browser | 2 (bootstrap + `go run ./cmd/server`) |
| Manual configuration steps | 0 |
| Lint issues after bootstrap | 0 |
| Test pass rate after bootstrap | 100% |
| Residual `template`/`Template` in source after bootstrap | 0 (except legitimate Go identifiers) |
| Files in dashboard context | 17 (each maps to exactly one pattern) |
| External dependencies | 0 (stdlib + `loopforge-ai/utils` only) |

## User Journeys

### Journey 1: Forge Bootstrap Skill — Automated Project Scaffolding

**Actor:** Forge Bootstrap Skill (automated process triggered by developer or CI)

**Opening Scene:** A developer invokes `skill_bootstrap` through Forge's MCP interface, providing three inputs: `module_path`, `project_dir`, and `project_name`. The bootstrap skill validates that `project_name` is a valid Go package name before proceeding.

**Rising Action:** The skill clones `loopforge-ai/template` into the target directory, removes `.git`, then walks every text file replacing `github.com/loopforge-ai/template` with the new module path. It targets specific files for project name replacement: `.justfile`, `.mcp.json`, `CLAUDE.md`, `Dockerfile`, `README.md`, `cmd/cli/main.go`, `cmd/mcp/main.go`. HTML templates get the capitalized project name. Finally, `git init` and `go mod tidy` finalize the project.

**Climax:** The replacement completes. Zero occurrences of `template`/`Template` remain in source files except legitimate Go identifiers like `html/template`. The module path is consistent across `go.mod` and all import statements. `go mod tidy` resolves cleanly.

**Resolution:** A fresh, valid Go project exists in the target directory. Every file compiles, every test passes, every lint rule holds. The project is indistinguishable from one hand-crafted by a Go expert — because it was.

**Failure Scenarios:**
- Target directory already exists → skill fails fast with clear error
- Network unavailable for clone → skill fails with clone error
- Project name is not a valid Go package name → skill validates and rejects before cloning
- String replacement hits unintended matches → residual `template` references break imports

**Requirements Revealed:** Clean replacement points, no ambiguous string matches, legitimate `template` references (Go imports) must survive replacement, all `<!-- TODO -->` placeholders must be meaningful and discoverable, bootstrap input validation for valid Go package names.

### Journey 2: Developer — Bootstrap to Browser to Forge (The "Aha" Moment)

**Actor:** Kai, a senior Go engineer starting a new microservice for their company's internal tooling platform.

**Opening Scene:** Kai has been told to build a new service. They've used Go for years but are tired of spending the first day of every project arguing about project structure, choosing a router, configuring linters, and writing Dockerfiles. They've heard about LoopForge and decide to try it.

**Rising Action:** Kai runs the Forge bootstrap skill: `module_path: github.com/acme/inventory`, `project_name: inventory`. In seconds, a new directory appears. Kai opens it in their editor and sees clean, organized code — `cmd/`, `internal/`, `web/` — instantly recognizable as hexagonal architecture. They run `go run ./cmd/server`.

**Climax:** Browser opens to `localhost:8080`. A dark-themed dashboard renders with "Inventory" in the header and footer, a version stat card, and a clean navigation bar. It's not a "Hello World" — it's a production-ready starting point. Kai thinks: "I didn't configure anything and this just works."

**Extended Arc:** Kai opens Claude Code. Forge connects via the MCP binary (`cmd/mcp`) — already configured and ready. Kai asks Forge to generate a `GET /products` handler. Forge reads the project's canonical structure, generates `handler_products.go`, `handler_products_test.go`, and updates `routes.go` — all matching the exact patterns from the template. The three binaries form a complete workflow: `server` for running, `mcp` for AI-assisted development, `cli` for future automation.

**Resolution:** Kai spends zero time on infrastructure. Within minutes they have a bootstrapped project running in the browser and Forge generating handlers that match the template's conventions perfectly. The first commit is business logic, not boilerplate.

**Failure Scenarios:**
- `go run ./cmd/server` fails to compile → broken template, trust destroyed immediately
- Dashboard renders but shows "Template" instead of "Inventory" → bootstrap replacement failed
- Lint fails on the bootstrapped project → quality contract violated
- Developer can't figure out where to add their first handler → patterns aren't self-documenting enough
- MCP binary doesn't connect to Forge → AI-assisted workflow broken

**Requirements Revealed:** Working server out of the box, correct name replacement in all visible UI, obvious patterns for handler creation, lint/test/build passing immediately, MCP entry point functional for Forge integration.

### Journey 3: Developer — Reading the Template as a Tutorial

**Actor:** Priya, a mid-level Go developer who has been assigned to a project already bootstrapped from this template. She needs to understand the architecture before contributing.

**Opening Scene:** Priya opens the repository and sees `cmd/`, `internal/`, `web/`. She's heard of hexagonal architecture but never worked in a codebase that implements it. She needs to understand the patterns fast — her first PR is due tomorrow.

**Rising Action:** Priya starts with `cmd/server/main.go` — it reads top to bottom like a recipe: create renderer, create handlers, register routes, start server, handle shutdown. Each step is one or two lines. She follows the imports into `internal/dashboard/domain/` and finds `PageData` and `RendererConfig` — simple types, no magic. Then into `inbound/` where `handler_health.go` shows the simplest possible handler (12 lines, no dependencies) and `handler_index.go` shows one with injected dependencies. `routes.go` shows how they're wired together with middleware.

**Climax:** Priya reads the test files. `handler_health_test.go` is 15 lines of obvious Arrange/Act/Assert. `handler_index_test.go` shows her how to test with fixtures — including an error path she wouldn't have thought to test. She opens `testhelper_test.go` and discovers the shared fixture pattern: `newTestRenderer` and `newBrokenRenderer` — unexported functions, `t.Helper()` first line, `t.Fatalf` for setup failures. `routes_test.go` proves every route through the actual mux. She thinks: "I can write my feature now. I know exactly where every piece goes."

**Resolution:** Priya's mental model is complete. Domain types go in `domain/`. Handlers go in `inbound/` as struct-based `http.Handler` implementations. Routes are registered in one function. Test helpers live in `testhelper_test.go`. Tests mirror the source structure. She writes her first handler confidently, following the patterns she just read. Her PR passes lint and tests on the first push.

**Failure Scenarios:**
- Code requires comments to understand → not self-documenting enough
- Handler pattern isn't obvious from reading one file → too many indirections
- Test structure isn't obvious → assertion pattern unclear, AAA structure not visible
- Test helper pattern missed → Priya duplicates fixture code in every test file
- She has to look at external docs to understand the architecture → code failed as tutorial

**Requirements Revealed:** Every file must be readable in isolation. Patterns must be learnable by reading, not by being told. Minimal indirection. No "magic" — every dependency is explicit. Test fixture pattern must be discoverable.

### Journey 4: Forge Code Generation Skills — Targeting Canonical Structure

**Actor:** A Forge skill generating a new handler for a project bootstrapped from the template.

**Opening Scene:** A developer asks Forge to generate a handler: "Create a GET /products endpoint that returns a list of products as JSON." Forge's skill system activates, targeting the project's canonical structure.

**Rising Action:** The skill reads the project structure and identifies the existing patterns: bounded context at `internal/<ctx>/`, handlers as struct-based `http.Handler` in `inbound/handler_<name>.go`, constructors as `NewXxxHandler(deps) *XxxHandler`, tests in `handler_<name>_test.go` with `package <name>_test`. It extracts conventions from the code itself — file ordering (alphabetical declarations), import grouping (stdlib then internal), test naming (`Test_<Unit>_With_<Condition>_Should_<Outcome>`), assertion pattern (`assert.That` with four arguments). The code is the spec — not CLAUDE.md, not external documentation.

**Climax:** The generated code is indistinguishable from hand-written code in the template. Same file naming, same struct pattern, same test naming convention, same middleware wiring, same declaration ordering. `golangci-lint` passes. All tests pass. A developer reading the codebase cannot tell which handlers were hand-written and which were generated.

**Resolution:** The deterministic structure made generation trivial. Every project bootstrapped from this template speaks the same language, so Forge skills work identically across all of them.

**Failure Scenarios:**
- Template patterns are inconsistent → skill generates inconsistent code
- Template uses patterns that are hard to replicate programmatically → skill can't target the structure
- Template has implicit conventions only documented in CLAUDE.md but not visible in code → skill misses conventions
- Convention exists in one file but not another → skill doesn't know which version to follow

**Requirements Revealed:** Patterns must be consistent and mechanical. Every convention must be visible in the code — if it's not extractable by reading source files, Forge can't replicate it. The code is the spec.

### Journey 5: Template Maintainer — Refining and Validating Minimality

**Actor:** Andy, the template maintainer, working to refine the template after a round of real-world usage.

**Opening Scene:** Andy bootstrapped a project last week and noticed a `<!-- TODO: add license -->` marker in the README that the bootstrap skill didn't clean up. The bootstrapped project still had raw TODO markers visible. This triggers a refinement cycle: are all placeholders intentional? Are they replaced by bootstrap or left for the developer?

**Rising Action:** Andy reviews every `<!-- TODO -->` marker in the template and categorizes each as: (a) replaced by bootstrap skill, (b) left for developer to fill in, or (c) should be removed. He then walks the pattern map file by file. For each file, he asks: "If I delete this, which pattern is lost?" He verifies each is load-bearing. He asks the inverse: "Is there a pattern that should be demonstrated but isn't?" He checks against hexagonal architecture requirements: domain types, inbound handlers, route registration, middleware, template rendering, static assets, tests. All covered.

**Climax:** Andy runs the quality contract: `golangci-lint run ./...` (zero issues), `go test ./...` (all green), `docker build` (success). Then he runs the automated bootstrap validation — an integration test in the template repo that executes the bootstrap skill with a test project name and verifies: no residual `template`/`Template`, all tests pass, lint clean, dashboard renders with the test name.

**Resolution:** The template is verified minimal and correct. The placeholder issue is fixed — each TODO marker is now categorized and documented. Andy tags a release. Every future bootstrap produces a project at this verified quality level.

**Failure Scenarios:**
- A file is redundant but Andy doesn't catch it → template has bloat
- A pattern is missing but Andy doesn't notice → bootstrapped projects lack a needed example
- Bootstrap replacement breaks something the integration test doesn't cover → silent failure
- A lint rule catches something after a Go version upgrade → quality contract needs maintenance

**Requirements Revealed:** Pattern map must be authoritative and maintained. Quality contract must be automated as integration tests. Placeholder structure must be categorized (bootstrap-replaced vs developer-fills-in). Refinement is trigger-driven by real bootstrap friction.

### Journey Requirements Summary

| Capability Area | Revealed By Journeys | Description |
|----------------|---------------------|-------------|
| Clean replacement points | 1, 2 | All `template`/`Template` occurrences intentional and replaceable, no ambiguous matches |
| Bootstrap input validation | 1 | `project_name` validated as valid Go package name before cloning |
| Self-documenting code | 2, 3, 4 | Every file readable in isolation, patterns learnable by reading |
| Three-binary developer workflow | 2 | server (run) + mcp (AI-assist) + cli (automate) form complete workflow |
| Consistent mechanical patterns | 3, 4 | Patterns replicable by humans (copy-modify) and Forge skills (generation) |
| Test fixture pattern | 3 | `testhelper_test.go` clearly demonstrates shared fixture creation |
| Code-as-spec for conventions | 4 | All conventions extractable by reading code, not external docs |
| Quality contract validation | 1, 2, 5 | Lint + test + Docker build pass in template AND after bootstrap (see Measurable Outcomes) |
| Automated bootstrap validation | 5 | Integration test that bootstraps and validates output project |
| Pattern map maintenance | 5 | Authoritative list of files → patterns, verified on every change |
| Placeholder categorization | 1, 5 | Each `<!-- TODO -->` classified as bootstrap-replaced or developer-fills-in |
| Trigger-driven refinement | 5 | Template refinement driven by real bootstrap friction |
| Minimal + complete tension | 3, 5 | Every file load-bearing, no redundancy, no gaps |

## Developer Tool (Project Scaffold) Specific Requirements

### Project-Type Overview

This is a **project scaffold** consumed by an automated code generation tool (Forge bootstrap skill). Unlike traditional developer tools (SDKs, libraries, CLIs), the template is never used directly — it's cloned, transformed, and discarded. The output project is what matters. Every requirement must be evaluated from the perspective of the *output project*, not the template repository itself.

### Technical Architecture Considerations

**Bootstrap Consumption Contract:**
- All files must be valid after string replacement (`template` → `projectName`, `Template` → `ProjectName`, module path replacement)
- `README.md` is fully rewritten by the bootstrap skill — template README exists only as the bootstrap input. Post-bootstrap README must explicitly identify the `dashboard` context as example code to replace.
- `CLAUDE.md` is modified by bootstrap to replace project-specific references
- `.mcp.json`, `.justfile`, `Dockerfile` are modified by bootstrap for project name references
- `cmd/cli/main.go` and `cmd/mcp/main.go` are modified by bootstrap for binary naming

**Naming for Bootstrap Survival:**
- Domain variables must use generic, context-appropriate names that make sense in any project — not template-specific names
- `TemplateConfig` must be renamed to `RendererConfig` — it's a renderer configuration object, and `TemplateConfig` would be confusing after bootstrap in a project called e.g. "inventory"
- `PageData` is acceptable — it describes what it is (page data) without template-specific connotation

**Dashboard as Example Code:**
- The `dashboard` bounded context must be obviously identifiable as example/starter code
- Naming convention signals example status: `dashboard` is a generic name that developers will naturally replace with their domain context
- No business logic in the dashboard — only infrastructure patterns demonstrated
- The dashboard should feel like scaffolding (useful to learn from, natural to replace) not like a feature (precious, hard to remove)
- Post-bootstrap README explicitly lists the dashboard as example code to replace

**Template File Categories:**

| Category | Files | Bootstrap Behavior |
|----------|-------|-------------------|
| Rewritten by bootstrap | `README.md` | Fully replaced with project-specific content |
| Modified by bootstrap | `go.mod`, all `.go` files, `.justfile`, `.mcp.json`, `CLAUDE.md`, `Dockerfile`, HTML templates | String replacement for module path and project name |
| Carried through unchanged | `.golangci.yml`, `.gitignore`, `web/static/css/style.css` | Used as-is in bootstrapped project |
| Removed by bootstrap | `.git/` | Fresh `git init` after removal |
| Template-only (excluded from bootstrap) | `_bmad/`, `_bmad-output/`, `.claude/`, `docs/` | Not included in bootstrapped projects — these are template repo infrastructure |

### Implementation Considerations

**Replacement Safety:**
- The word `template` must not appear in Go source code except in `html/template` imports
- The word `Template` must not appear in Go source code — `TemplateConfig` renamed to `RendererConfig`
- All replacement targets in HTML, markdown, and config files must use exact case: `template` (lowercase) or `Template` (capitalized)

**Example Code Clarity:**
- The `dashboard` package name signals "this is the example" — developers replace `internal/dashboard/` with `internal/<their-context>/`
- `PageData` and `RendererConfig` in domain are dashboard-specific types that get replaced with the developer's domain types
- Handler names (`HealthHandler`, `IndexHandler`) are concrete examples, not abstractions to extend
- Post-bootstrap README explicitly guides developers to replace the dashboard context

**Quality Gates — Template Repository:**
- All targets per Measurable Outcomes table
- Pattern map verification: every file maps to exactly one demonstrated pattern
- Placeholder audit: all `<!-- TODO -->` markers categorized as bootstrap-replaced or developer-fills-in

**Quality Gates — Post-Bootstrap Output:**
- All targets per Measurable Outcomes table
- No stale template-specific content (TODOs, placeholder descriptions that reference the template)
- Dashboard renders with correct project name in browser
- Module path consistent across `go.mod` and all imports

## Project Scoping & Phased Development

### MVP Strategy & Philosophy

**MVP Approach:** Refinement MVP — the product already exists and functions. The MVP is proving that the existing template is irreducibly minimal, bootstrap-safe, and convention-complete. No new patterns; only verification, tightening, and correction.

**Why this approach:** Adding features to a scaffold before validating the foundation creates compounding debt. Every new pattern added to an unverified template propagates unverified patterns to all bootstrapped projects. Verify first, expand second.

**Resource Requirements:** Single developer (template maintainer) with access to Forge bootstrap skill for validation.

### MVP Feature Set (Phase 1 — Refinement)

**Core Journeys Supported:** All five journeys at current scope.

**Must-Have Capabilities:**

| # | Capability | Rationale | Acceptance Criterion |
|---|-----------|-----------|---------------------|
| 1 | Rename `TemplateConfig` → `RendererConfig` | Survives bootstrap without confusion; contains no `template` substring | No `Template`-prefixed identifiers in Go source except `html/template` imports |
| 2 | Verify pattern map completeness | Every file maps to one pattern, no redundancy | Pattern map table validated — each file justified, no file removable without losing a pattern |
| 3 | Standardize placeholder structure | Clean `<!-- TODO: description -->` format, each categorized | Every TODO marker documented as bootstrap-replaced or developer-fills-in, documented in a manifest section of README |
| 4 | Validate replacement safety | No ambiguous `template`/`Template` matches | Post-bootstrap grep returns zero matches (except `html/template`) |
| 5 | Validate quality contract end-to-end | Lint + test + build in template AND after bootstrap | Documented validation procedure: scripted command sequence the maintainer executes before each release |
| 6 | Document template-only directories | `_bmad/`, `_bmad-output/`, `.claude/`, `docs/` are template-only | Template declares exclusions; Forge team updates bootstrap skill separately (known issue: currently cloned) |
| 7 | Verify bootstrap-to-browser flow | Two commands, dashboard renders with project name | Scripted validation confirms end-to-end flow |

**Explicitly Out of MVP Scope:**
- No new patterns introduced (file count may change if minimality verification reveals redundancy)
- No outbound layer
- No new bounded contexts
- No automated integration test (scripted validation instead; automation is Phase 2)
- No bootstrap skill changes (documented as Forge team responsibility)

### Post-MVP Features

**Phase 2a — Hexagonal Completion (outbound layer):**

| # | Feature | Rationale |
|---|---------|-----------|
| 1 | Add `outbound/` layer to dashboard context | Complete hexagonal triad — domain interface + outbound adapter |
| 2 | Add integration test example | Demonstrate `//go:build integration` pattern |
| 3 | Automated bootstrap integration test | Go test in template repo that bootstraps and validates output |

**Phase 2b — Multi-Context (second bounded context):**

| # | Feature | Rationale |
|---|---------|-----------|
| 1 | Add second bounded context | Demonstrate cross-context isolation and independent `RegisterRoutes` |
| 2 | Update `cmd/server/main.go` to wire both contexts | Demonstrate multi-context composition root |
| 3 | Update pattern map for expanded file set | Maintain minimality invariant with additional patterns |

**Phase 3 — Expansion (advanced patterns):**

| # | Feature | Rationale |
|---|---------|-----------|
| 1 | Domain events between contexts | Demonstrate event-driven communication |
| 2 | Expand CLI entry point | Move beyond placeholder to demonstrate CLI patterns |
| 3 | Template variants | API-only, CLI-only options selectable during bootstrap |
| 4 | Forge skills for template patterns | Generate handlers, contexts, adapters targeting canonical structure |

### Risk Mitigation Strategy

**Technical Risks:**

| Risk | Impact | Mitigation |
|------|--------|-----------|
| `RendererConfig` rename breaks downstream | Requires `cmd/server/main.go` update | Two-file change; no bootstrap conflict (no `template` substring) |
| Placeholder changes break bootstrap | Bootstrap skill may reference specific markers | Coordinate with Forge team; document placeholder format |
| Bootstrap clones template-only directories | Bootstrapped projects contain `_bmad/`, `.claude/`, `docs/` | Known issue — template documents exclusions, Forge team updates bootstrap skill |

**Process Risks:**

| Risk | Impact | Mitigation |
|------|--------|-----------|
| No real bootstrap usage to validate against | Refined in theory, not practice | Bootstrap 2-3 test projects with different names before tagging release |
| Pattern map becomes stale after Phase 2 | Minimality invariant violated silently | Pattern map verification as part of release checklist; automated in Phase 2a |
| Convention drift between template and CLAUDE.md | CLAUDE.md describes conventions code doesn't demonstrate | Code-as-spec: every CLAUDE.md convention must be visible in template code |

## Functional Requirements

### Template Structure & Minimality

- FR1: The template can demonstrate every hexagonal architecture pattern through a single `dashboard` bounded context
- FR2: The template maintainer can verify that each source file maps to exactly one demonstrated pattern using the pattern map
- FR3: The template maintainer can verify that no source file is removable without losing a demonstrated pattern
- FR4: The template can provide three independent entry points (HTTP server, MCP server, CLI) from a single Go module
- FR5: The `dashboard` context can be identified as example code through naming convention alone (no explicit markers in source code)

### Bootstrap Consumption

- FR6: The Forge bootstrap skill can clone the template and produce a valid Go project by replacing module path and project name
- FR7: The bootstrap skill can replace all occurrences of `template`/`Template` without affecting legitimate Go identifiers (e.g., `html/template` imports)
- FR8: The template can declare which directories are template-only and should be excluded from bootstrapped projects
- FR9: All Go identifiers in the template can survive bootstrap string replacement without producing confusing or misleading names
- FR10: The template can provide a post-bootstrap README (via bootstrap skill) that identifies the dashboard as example code to replace

### Pattern Demonstration

- FR11: The template can demonstrate a stateless HTTP handler with no dependencies (health check)
- FR12: The template can demonstrate a stateful HTTP handler with dependency injection (renderer-backed page)
- FR13: The template can demonstrate route registration with middleware chaining
- FR14: The template can demonstrate static asset serving via embedded filesystem
- FR15: The template can demonstrate Go HTML template composition (layout + pages + partials)
- FR16: The template can demonstrate a partial template without data binding (static content)
- FR17: The template can demonstrate a partial template with data binding (dynamic content)
- FR18: The template can demonstrate domain types and package-level configuration objects
- FR19: The template can demonstrate the composition root pattern (dependency wiring in `main.go`)
- FR20: The template can demonstrate graceful shutdown with signal handling

### Testing Pattern Demonstration

- FR21: The template can demonstrate unit tests for stateless handlers using Arrange/Act/Assert
- FR22: The template can demonstrate unit tests for stateful handlers with happy path and error path coverage
- FR23: The template can demonstrate route-level integration tests through the actual mux
- FR24: The template can demonstrate shared test fixture creation via `testhelper_test.go`
- FR25: The template can demonstrate external test package usage (`package <name>_test`)
- FR26: The template can demonstrate the standard assertion pattern (`assert.That`) and test parallelism (`t.Parallel()`) consistently across all test files

### Quality Assurance

- FR27: The template can provide a documented validation checklist of commands for pre-release verification (lint, test, build, Docker)
- FR28: The template maintainer can execute post-bootstrap validation including server start and HTTP response body inspection for correct project name
- FR29: The template maintainer can verify that no residual `template`/`Template` exists in post-bootstrap source files (except `html/template`)
- FR30: The template maintainer can audit all `<!-- TODO -->` placeholders and classify each as bootstrap-replaced or developer-fills-in

### Developer Experience

- FR31: A developer can start the HTTP server and see the dashboard in a browser with two commands after bootstrap
- FR32: A developer can understand any template file's purpose by reading the file in isolation
- FR33: A developer can add a new handler by replicating the existing handler pattern (copy-modify)
- FR34: A developer can add a new HTML page by adding to the `RendererConfig` pages list and creating a template file
- FR35: A developer can add a new bounded context by following the single-context pattern (demonstrated by convention, not by explicit multi-context example — explicit example deferred to Phase 2b)

### Forge Integration

- FR36: The MCP binary can connect to Forge for AI-assisted code generation in every bootstrapped project
- FR37: The template's coding conventions can be fully extracted by reading source files without referencing external documentation
- FR38: The template's patterns are consistent and mechanical enough that generated code following them is indistinguishable from hand-written code

## Non-Functional Requirements

### Code Quality

- NFR1: All Go source files pass `golangci-lint` with `default: all` configuration with zero inline `//nolint` directives. Global linter configuration in `.golangci.yml` is permitted with documented rationale for each disable and each `gosec` exclusion.
- NFR2: Test coverage includes both happy path and error path for every handler
- NFR3: Every Go source file and HTML template has a single, clear purpose apparent within the first 10 lines — no file serves multiple unrelated concerns
- NFR4: Import statements use consistent grouping (stdlib, blank line, internal) across all Go files with no exceptions

### Maintainability

- NFR5: Any single-file change to the template (add, modify, remove) can be validated against the pattern map in under 5 minutes
- NFR6: The template's convention set is fully expressed in source code — no convention exists only in documentation
- NFR7: File naming follows a consistent, predictable pattern (`handler_<name>.go`, `handler_<name>_test.go`, `testhelper_test.go`) such that a new file's purpose is clear from its name

### Portability

- NFR8: Go source code builds and tests pass on macOS, Linux, and Windows with Go 1.25+ and `CGO_ENABLED=0` — no platform-specific code or build constraints
- NFR9: Docker build succeeds on any Docker-compatible host without platform-specific configuration
- NFR10: Build tooling (`.justfile`) is a convenience layer — template quality is not dependent on `just` availability. Core validation uses `go test`, `golangci-lint`, and `docker build` directly.

### Security Patterns (Propagated)

- NFR11: All HTTP handlers use the `SecurityHeaders` middleware — no handler is registered without security headers
- NFR12: Docker runtime uses a non-root user (`app`) with minimal installed packages
- NFR13: No secrets, credentials, or environment-specific values are hardcoded in any source file
- NFR14: `gosec` runs as part of the linter suite with documented exclusions only (G301, G703, G704, G706 — each with justification in `.golangci.yml`)

### Accessibility (HTML Templates)

- NFR15: All HTML templates include semantic elements (`<header>`, `<main>`, `<footer>`, `<nav>`)
- NFR16: Navigation elements include `aria-label` attributes for screen reader support
- NFR17: The `<html>` element includes `lang="en"` attribute for screen reader language selection
- NFR18: CSS respects `prefers-reduced-motion` media query for users with motion sensitivity
