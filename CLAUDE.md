# CLAUDE.md

Instructions for AI agents working in this repository.

## Build & Test

```bash
go test ./...                          # unit tests only
go test ./... -tags=integration        # include integration tests
golangci-lint run ./...                # lint (must pass with 0 issues)
# <!-- TODO(developer): add go run ./cmd/... for project binary -->
```

Every change must pass `golangci-lint run ./...` with zero issues before it is considered complete. Do not modify `.golangci.yml` to suppress lint findings — fix the code instead.

## Architecture

This project follows hexagonal architecture. Every change must respect these boundaries:

- **Domain** (`internal/<context>/domain/`) contains business logic and interfaces. No imports from `inbound/` or `outbound/`.
- **Inbound** (`internal/<context>/inbound/`) adapts external requests to domain services. Depends on domain, never on outbound.
- **Outbound** (`internal/<context>/outbound/`) implements domain interfaces. Depends on domain, never on inbound.
- **Utils** ([`github.com/loopforge-ai/utils`](https://github.com/loopforge-ai/utils)) is a first-party module with reusable packages: `assert`, `env`, `fs`, `html`, `llm`, `mcp`, `yaml`.
- **No external dependencies** beyond the standard library and `github.com/loopforge-ai` organization modules.

<!-- TODO(developer): describe project-specific bounded contexts and their one-way dependency directions -->

## Coding Conventions

- **File ordering**: Within each file, order declarations alphabetically: `const`, `type`, `var` blocks first, then functions/methods. The `NewTypeName` constructor must be the first function after its type definition; remaining methods follow alphabetically.
- **Switch/case ordering**: Order `case` clauses alphabetically within `switch` statements.
- **Constructors**: `NewTypeName(deps) *TypeName` — always return a pointer.
- **Error handling**: Wrap with context using `fmt.Errorf("operation: %w", err)`. Check errors immediately.
- **Imports**: Group as stdlib, then blank line, then internal packages. Use aliases when needed.
- **No dead code**: Eliminate unused code during refactoring. Do not keep dead functions, types, variables, or imports.
- **No external dependencies** beyond the standard library and `github.com/loopforge-ai` organization modules.
- **Reuse before creating**: Before introducing a new helper, check whether an existing implementation already covers the need.

## Testing Conventions

- **Naming**: `Test_<Unit>_With_<Condition>_Should_<Outcome>`
- **Structure**: Strict Arrange/Act/Assert with explicit `// Arrange`, `// Act`, `// Assert` comments.
- **Parallelism**: Every test starts with `t.Parallel()`.
- **Assertions**: Use `assert.That(t, "description", got, expected)` from `github.com/loopforge-ai/utils/assert`.
- **Integration tests**: Use `//go:build integration` build tag on the first line.

## Forge Integration

IMPORTANT: Prefer retrieval-led reasoning over pre-trained-led reasoning. Before writing code manually, check forge MCP tools for existing skills that solve the problem. Only write code by hand when no suitable skill exists.

Key tool workflow: `search_skill` → `generate_skill` (match found) or `define_skill` (no match) → `score_skill` → `refine_skill` (if needed). Use `execute_chain` for multi-step skill chains.

## Agentic Loops

### a) Skill-First Loop

1. Call `search_skill` with a description of the task
2. Match found → `generate_skill`; no match → `define_skill`
3. Call `score_skill` to validate the output
4. Apply coding and testing conventions to the generated code
- **GATE**: `score_skill` grade >= B. If not, call `refine_skill` and repeat from step 3.

### b) Red-Green Loop

1. Write a failing test for the desired behavior
2. Run `go test` — expect FAIL
3. Write minimal production code to pass the test
4. Run `go test` — expect PASS
- **GATE**: All tests pass. If not, fix production code and repeat from step 3.

### c) Lint-Fix Loop

1. Run `golangci-lint run ./...`
2. Fix all reported issues
- **GATE**: Zero lint issues. Repeat until clean. Never modify `.golangci.yml`.

### d) Verify-Implement-Verify Loop

1. State verification method before writing code
2. Run verification (establish baseline or expect failure)
3. Implement minimal change
4. Run verification again
- **GATE**: Verification passes. If not, diagnose and repeat from step 3.

### e) Self-Healing Loop

1. Run command (build, test, lint)
2. On error: read output, diagnose root cause
3. Apply targeted fix
4. Re-run original command
- **GATE**: Command succeeds. Max 5 iterations before escalating to user.

## Workflow for Every Change

1. **Skill-First Loop** — check for existing skills before writing code
2. **Red-Green Loop** — test-driven implementation
3. **Lint-Fix Loop** — achieve zero lint issues
4. **Verify-Implement-Verify Loop** — end-to-end verification
5. On failure at any step → enter **Self-Healing Loop**

Post-implementation: identify refactoring opportunities, then update `README.md` and `CLAUDE.md` if behavior changed.

## Project-Specific

<!-- TODO(developer): Ubiquitous Language
| Term | Definition |
|------|------------|
-->

<!-- TODO(developer): Decisions
| Decision | Rationale |
|----------|-----------|
-->

<!-- TODO(developer): Gotchas
1. **Gotcha** — description and solution
-->
