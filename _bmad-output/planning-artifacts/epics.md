---
stepsCompleted: ['step-01-validate-prerequisites', 'step-02-design-epics', 'step-03-create-stories', 'step-04-final-validation']
status: 'complete'
completedAt: '2026-03-14'
inputDocuments: ['_bmad-output/planning-artifacts/prd.md', '_bmad-output/planning-artifacts/architecture.md']
---

# template - Epic Breakdown

## Overview

This document provides the complete epic and story breakdown for template, decomposing the requirements from the PRD and Architecture into implementable stories.

## Requirements Inventory

### Functional Requirements

- FR1: The template can demonstrate every hexagonal architecture pattern through a single `dashboard` bounded context
- FR2: The template maintainer can verify that each source file maps to exactly one demonstrated pattern using the pattern map
- FR3: The template maintainer can verify that no source file is removable without losing a demonstrated pattern
- FR4: The template can provide three independent entry points (HTTP server, MCP server, CLI) from a single Go module
- FR5: The `dashboard` context can be identified as example code through naming convention alone
- FR6: The Forge bootstrap skill can clone the template and produce a valid Go project by replacing module path and project name
- FR7: The bootstrap skill can replace all occurrences of `template`/`Template` without affecting legitimate Go identifiers
- FR8: The template can declare which directories are template-only and should be excluded from bootstrapped projects
- FR9: All Go identifiers in the template can survive bootstrap string replacement without producing confusing names
- FR10: The template can provide a post-bootstrap README that identifies the dashboard as example code to replace
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
- FR21: The template can demonstrate unit tests for stateless handlers using Arrange/Act/Assert
- FR22: The template can demonstrate unit tests for stateful handlers with happy path and error path coverage
- FR23: The template can demonstrate route-level integration tests through the actual mux
- FR24: The template can demonstrate shared test fixture creation via `testhelper_test.go`
- FR25: The template can demonstrate external test package usage (`package <name>_test`)
- FR26: The template can demonstrate the standard assertion pattern (`assert.That`) and test parallelism (`t.Parallel()`) consistently across all test files
- FR27: The template can provide a documented validation checklist of commands for pre-release verification
- FR28: The template maintainer can execute post-bootstrap validation including server start and HTTP response body inspection
- FR29: The template maintainer can verify that no residual `template`/`Template` exists in post-bootstrap source files
- FR30: The template maintainer can audit all `<!-- TODO -->` placeholders and classify each as bootstrap-replaced or developer-fills-in
- FR31: A developer can start the HTTP server and see the dashboard in a browser with two commands after bootstrap
- FR32: A developer can understand any template file's purpose by reading the file in isolation
- FR33: A developer can add a new handler by replicating the existing handler pattern
- FR34: A developer can add a new HTML page by adding to the `RendererConfig` pages list and creating a template file
- FR35: A developer can add a new bounded context by following the single-context pattern
- FR36: The MCP binary can connect to Forge for AI-assisted code generation in every bootstrapped project
- FR37: The template's coding conventions can be fully extracted by reading source files without referencing external documentation
- FR38: The template's patterns are consistent and mechanical enough that generated code following them is indistinguishable from hand-written code

### NonFunctional Requirements

- NFR1: Zero inline `//nolint` directives; global config with documented rationale
- NFR2: Happy path + error path test coverage for every handler
- NFR3: Every file has single clear purpose apparent within first 10 lines
- NFR4: Consistent import grouping across all Go files
- NFR5: Single-file changes validated against pattern map in <5 minutes
- NFR6: Convention set fully expressed in source code, not only documentation
- NFR7: File naming follows predictable pattern
- NFR8: Cross-platform builds (macOS, Linux, Windows) with CGO_ENABLED=0
- NFR9: Docker builds on any compatible host
- NFR10: Justfile is convenience layer, not required
- NFR11: SecurityHeaders middleware on all handlers
- NFR12: Non-root Docker runtime user
- NFR13: No hardcoded secrets or environment-specific values
- NFR14: gosec with documented exclusions only
- NFR15: Semantic HTML elements
- NFR16: aria-label attributes on navigation
- NFR17: lang="en" on html element
- NFR18: prefers-reduced-motion CSS support

### Additional Requirements

- ADR-1: Create `.bootstrapignore` file declaring template-only directories
- ADR-2: Recategorize TODO markers to `TODO(bootstrap)` / `TODO(developer)` format
- ADR-3: Create `VALIDATION.md` with quality gates + convention map; add `just validate` recipe
- ADR-4: Populate convention-to-file mapping table in VALIDATION.md (18 conventions)
- ARCH-1: Rename `TemplateConfig` → `RendererConfig` across 3 Go files + 3 doc files
- ARCH-2: Bootstrap-aware identifier naming (no `template`/`Template` in Go identifiers)
- ARCH-3: Pattern map update rule — update map in same commit as code changes
- ARCH-4: Implementation sequence: ADR-2 → ADR-1 → ADR-3 → ADR-4 → rename → doc updates → validation

### UX Design Requirements

N/A — not applicable for this project type (server-side rendered scaffold).

### FR Coverage Map

- FR1: Already satisfied — existing dashboard demonstrates all hexagonal patterns
- FR2: Epic 2 — pattern map verification in VALIDATION.md
- FR3: Epic 2 — minimality verification via pattern map
- FR4: Already satisfied — three entry points exist
- FR5: Already satisfied — `dashboard` naming convention signals example code
- FR6: Epic 2 — verified by post-bootstrap validation
- FR7: Epic 2 — verified by replacement safety check
- FR8: Epic 1 — .bootstrapignore declares template-only directories
- FR9: Epic 1 — RendererConfig rename removes template-substring identifiers
- FR10: Forge team — post-bootstrap README is bootstrap skill responsibility
- FR11-FR20: Already satisfied — all pattern demonstration files exist
- FR21-FR26: Already satisfied — all testing patterns exist
- FR27: Epic 2 — VALIDATION.md + just validate recipe
- FR28: Epic 2 — post-bootstrap server start + HTTP response inspection
- FR29: Epic 2 — grep verification for residual template/Template
- FR30: Epic 1 — TODO markers recategorized to TODO(bootstrap)/TODO(developer)
- FR31: Already satisfied — two commands to browser (verified in Epic 2)
- FR32-FR34: Already satisfied — self-documenting code, handler replication, page addition
- FR35: Already satisfied — bounded context pattern demonstrated by convention
- FR36: Already satisfied — MCP binary exists
- FR37: Epic 2 — convention-to-file mapping in VALIDATION.md
- FR38: Epic 2 — convention consistency verified via mapping table
- NFR1-18: Enforced by existing architecture; verified by Epic 2 validation procedure

## Epic List

### Epic 1: Bootstrap-Safe Template Refinement

**Goal:** Make the template safe and clean for Forge bootstrap consumption — no confusing identifiers, clear directory exclusions, and categorized placeholders.

After this epic: The bootstrap skill produces projects with no `TemplateConfig` confusion, template-only directories are declared, and all TODO markers are categorized. All existing tests pass. Lint is clean.

**FRs covered:** FR8, FR9, FR30
**ADRs implemented:** ADR-1 (.bootstrapignore), ADR-2 (placeholder categorization), ARCH-1 (RendererConfig rename)

**Stories:**
- Story 1.1: Template infrastructure files (.bootstrapignore + placeholder recategorization)
- Story 1.2: RendererConfig rename (3 Go files)

### Epic 2: Validation Infrastructure & Verification

**Goal:** Create a repeatable validation procedure that verifies template correctness, convention coverage, and bootstrap safety — then execute it to confirm the template is release-ready.

After this epic: The maintainer has a documented, executable validation procedure. All documentation references are current. The template is verified minimal, correct, and bootstrap-safe. Ready to tag a release.

**FRs covered:** FR2, FR3, FR6, FR7, FR27, FR28, FR29, FR37, FR38
**ADRs implemented:** ADR-3 (VALIDATION.md + just validate), ADR-4 (convention-to-file mapping)

**Stories:**
- Story 2.1: Create VALIDATION.md with quality gates and pattern map
- Story 2.2: Convention-to-file mapping and just validate recipe
- Story 2.3: Update stale documentation references (3 doc files)
- Story 2.4: Execute full validation and confirm release readiness

### Forge Team Handoff (Not Stories — External Dependencies)

These items are the Forge team's responsibility, documented here for coordination:
- Implement `.bootstrapignore` support in the bootstrap skill
- Update post-bootstrap README to identify dashboard as example code (FR10)
- Validate `project_name` as valid Go package name before cloning (Journey 1 requirement)

## Epic 1: Bootstrap-Safe Template Refinement

**Goal:** Make the template safe and clean for Forge bootstrap consumption — no confusing identifiers, clear directory exclusions, and categorized placeholders.

### Story 1.1: Template Infrastructure Files [XS]

As a **template maintainer**,
I want to declare template-only directories and categorize all TODO placeholders,
So that the bootstrap skill knows what to exclude and post-bootstrap projects have no uncategorized markers.

**Acceptance Criteria:**

**Given** the template repository has no `.bootstrapignore` file
**When** the maintainer creates `.bootstrapignore` at the repository root
**Then** the file lists `_bmad/`, `_bmad-output/`, `.claude/`, `docs/`, `.bootstrapignore`, and `VALIDATION.md` — one path per line
**And** the file uses `#` for comments explaining each exclusion

**Given** `CLAUDE.md` contains `<!-- TODO: add go run ./cmd/... for project binary -->`
**When** the maintainer recategorizes the TODO marker
**Then** it becomes `<!-- TODO(developer): add go run ./cmd/... for project binary -->`
**And** `grep -rn 'TODO' . | grep -v 'TODO(bootstrap)' | grep -v 'TODO(developer)' | grep -v '_bmad'` returns zero results for files that survive bootstrap

**Given** both changes are complete
**When** `golangci-lint run ./...` is executed
**Then** zero issues are reported
**And** `go test ./...` passes with all tests green

### Story 1.2: Rename TemplateConfig to RendererConfig [S]

As a **developer bootstrapping a project**,
I want all Go identifiers to survive bootstrap string replacement without confusion,
So that a project called "inventory" doesn't contain a misleading `TemplateConfig` variable.

**Acceptance Criteria:**

**Given** `internal/dashboard/domain/renderer.go` declares `var TemplateConfig`
**When** the maintainer renames it to `var RendererConfig`
**Then** the variable name is `RendererConfig` in `renderer.go`
**And** `cmd/server/main.go` references `dashboard.RendererConfig` instead of `dashboard.TemplateConfig`
**And** `inbound/testhelper_test.go` references `dashboard.RendererConfig` in both `newTestRenderer` and `newBrokenRenderer`

**Given** the rename is complete
**When** `grep -rn 'TemplateConfig' --include='*.go' .` is executed
**Then** zero results are returned

**Given** the rename is complete
**When** `grep -rn 'template\|Template' --include='*.go' --exclude-dir=_bmad . | grep -v 'html/template'` is executed
**Then** zero results are returned (bootstrap replacement safety verified)

**Given** all three files are updated
**When** `golangci-lint run ./...` is executed
**Then** zero issues are reported
**And** `go test ./...` passes with all 7 tests green
**And** `go run ./cmd/server` starts without compilation errors (full server verification in Story 2.4)

## Epic 2: Validation Infrastructure & Verification

**Goal:** Create a repeatable validation procedure that verifies template correctness, convention coverage, and bootstrap safety — then execute it to confirm the template is release-ready.

### Story 2.1: Create VALIDATION.md with Quality Gates and Pattern Map [M]

As a **template maintainer**,
I want a documented validation procedure with quality gates and a pattern map,
So that I can verify template correctness with a repeatable, scripted process before every release.

**Note:** Pattern map must use `RendererConfig` (post-rename name from Story 1.2), not `TemplateConfig`.

**Acceptance Criteria:**

**Given** no `VALIDATION.md` exists in the repository
**When** the maintainer creates `VALIDATION.md` at the repository root
**Then** the file contains these sections:
1. Template Quality Gates (lint, test, build, Docker commands with expected output)
2. Pattern Map Verification (table mapping each of the 17 source files to its demonstrated pattern)
3. Placeholder Audit Procedure (grep command from ADR-2 with pass/fail criteria)
4. Post-Bootstrap Validation Procedure (scripted commands: bootstrap test project, lint, test, build, grep for residual template/Template, start server, check HTTP response)
5. Bootstrap-to-Browser Verification Steps

**Given** `VALIDATION.md` is created
**When** the `.bootstrapignore` file is checked
**Then** `VALIDATION.md` is already listed (added in Story 1.1)

**Given** the Pattern Map section is populated
**When** the maintainer reviews each row
**Then** every source file in `internal/dashboard/`, `web/`, and `cmd/` maps to exactly one pattern
**And** no file appears in multiple rows
**And** removing any row would leave a pattern undemonstrated (FR2, FR3)

### Story 2.2: Convention-to-File Mapping and Just Validate Recipe [M]

As a **template maintainer**,
I want a convention-to-file mapping that links every CLAUDE.md convention to template source code, and a `just validate` recipe for quick checks,
So that I can verify the code-as-spec principle and run quality gates with a single command.

**Acceptance Criteria:**

**Given** `VALIDATION.md` exists from Story 2.1
**When** the maintainer adds the Convention-to-File Mapping section
**Then** the table contains all 18 conventions from ADR-4, each with:
- Convention name (matching CLAUDE.md)
- File where it's demonstrated
- Specific identifier or code reference (not line numbers)

**Given** the convention mapping is complete
**When** the maintainer cross-references every convention in CLAUDE.md against the table
**Then** every convention has a corresponding entry (FR37)
**And** no convention exists only in CLAUDE.md without a code demonstration (NFR6)

**Given** `.justfile` exists with `lint` and `test` recipes
**When** the maintainer adds a `validate` recipe
**Then** the recipe is:
```
validate: lint test
    docker build -t template:validate .
    @echo "Template quality contract: PASS"
```
**And** running `just validate` executes lint, test, and Docker build sequentially
**And** the recipe fails fast if any step fails

### Story 2.3: Update Stale Documentation References [XS]

As a **template maintainer**,
I want all documentation to reflect the RendererConfig rename from Epic 1,
So that no stale `TemplateConfig` references confuse agents or developers reading the docs.

**Acceptance Criteria:**

**Given** Epic 1 renamed `TemplateConfig` to `RendererConfig` in Go source
**When** the maintainer updates `_bmad-output/project-context.md`
**Then** all references to `TemplateConfig` are replaced with `RendererConfig`
**And** the "Module-level var" rule example is updated

**Given** `docs/architecture.md` references `TemplateConfig` in component tables
**When** the maintainer updates the references
**Then** all occurrences of `TemplateConfig` become `RendererConfig`
**And** the Domain Layer table shows `RendererConfig` for the config object

**Given** `docs/source-tree-analysis.md` describes `renderer.go` as "PageData type + TemplateConfig"
**When** the maintainer updates the description
**Then** it reads "PageData type + RendererConfig"

**Given** all three documentation files are updated
**When** `grep -rn 'TemplateConfig' docs/ _bmad-output/` is executed
**Then** zero results are returned

### Story 2.4: Execute Full Validation and Confirm Release Readiness [S]

As a **template maintainer**,
I want to execute the complete VALIDATION.md procedure and confirm the template is release-ready,
So that I can tag a release with confidence that the template is minimal, correct, and bootstrap-safe.

**Acceptance Criteria:**

**Given** all previous stories (1.1, 1.2, 2.1, 2.2, 2.3) are complete
**When** the maintainer executes the Template Quality Gates section of VALIDATION.md
**Then** `golangci-lint run ./...` reports zero issues
**And** `go test ./...` reports all tests passing
**And** `docker build -t template:validate .` succeeds

**Given** quality gates pass
**When** the maintainer walks the Pattern Map Verification
**Then** each of the 17 source files maps to exactly one pattern
**And** no file is removable without losing a demonstrated pattern

**Given** pattern map is verified
**When** the maintainer runs the Placeholder Audit
**Then** `grep -rn 'TODO' . | grep -v 'TODO(bootstrap)' | grep -v 'TODO(developer)' | grep -v '_bmad'` returns zero results

**Given** placeholder audit passes
**When** the maintainer runs the Bootstrap Replacement Safety check
**Then** `grep -rn 'template\|Template' --include='*.go' --exclude-dir=_bmad . | grep -v 'html/template'` returns zero results

**Given** all checks pass
**When** the maintainer executes the Post-Bootstrap Validation (bootstrap a test project, lint, test, build, grep, start server)
**Then** the bootstrapped project passes lint with zero issues
**And** all tests pass in the bootstrapped project
**And** `grep -rn 'template\|Template' --include='*.go' . | grep -v 'html/template'` returns zero in the bootstrapped project
**And** `go run ./cmd/server` starts and serves the dashboard
**And** an HTTP GET to `localhost:8080` returns HTML containing the test project name (not "Template")

**Given** any validation step fails
**When** the maintainer reviews the failure
**Then** the failure is documented, the relevant story is revisited to fix the root cause, and validation re-executes from the failed step

**Given** all validation passes
**When** the maintainer reviews the results
**Then** the template is confirmed release-ready
**And** VALIDATION.md can be committed as the verified validation artifact
