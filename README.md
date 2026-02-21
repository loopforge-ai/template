# <!-- TODO: project-name -->

<!-- TODO: one-line project description -->

## Quick Start

```bash
git clone https://github.com/loopforge-ai/template.git <project-name>
cd <project-name> && rm -rf .git && git init
go mod init github.com/<org>/<project-name>
mkdir -p cmd/<app> internal/<context>/domain internal/<context>/inbound internal/<context>/outbound
```

1. Update `.mcp.json` with your local paths.
2. Fill in all `<!-- TODO -->` placeholders in `CLAUDE.md` and `README.md`.
3. Run the first build and test cycle: `go test ./... && golangci-lint run ./...`

## What's Included

| File | Purpose |
|------|---------|
| `CLAUDE.md` | AI agent guidelines: conventions, agentic loops, forge integration |
| `.mcp.json` | Forge MCP server configuration |
| `.claude/settings.local.json` | Claude Code tool permissions |
| `README.md` | This file (update after cloning) |

## Architecture

This project follows **hexagonal architecture** (ports & adapters):

- **Domain** (`internal/<context>/domain/`) — Business logic and interfaces. No imports from inbound or outbound.
- **Inbound** (`internal/<context>/inbound/`) — Adapts external requests to domain services. Depends on domain only.
- **Outbound** (`internal/<context>/outbound/`) — Implements domain interfaces. Depends on domain only.

Shared utilities come from [loopforge-ai/utils](https://github.com/loopforge-ai/utils). No external dependencies beyond the standard library and `github.com/loopforge-ai` modules.

<!-- TODO: describe project-specific bounded contexts and their dependency directions -->

## Build & Test

```bash
go test ./...                          # unit tests only
go test ./... -tags=integration        # include integration tests
golangci-lint run ./...                # lint (must pass with 0 issues)
# <!-- TODO: add go run ./cmd/... for project binary -->
```

## Technology Stack

| Layer | Technology |
|-------|------------|
| Architecture | Hexagonal (ports & adapters) |
| Language | Go 1.23+ |
| Linter | golangci-lint |
| Testing | stdlib + `assert.That` |
| Utilities | [loopforge-ai/utils](https://github.com/loopforge-ai/utils) |
| Code Generation | [loopforge-ai/forge](https://github.com/loopforge-ai/forge) |

## Project Structure

<!-- TODO: update after scaffolding -->
```
<project>/
├── cmd/<app>/          # Entry point
├── internal/<ctx>/
│   ├── domain/         # Business logic and interfaces
│   ├── inbound/        # Adapters (MCP handlers, HTTP)
│   └── outbound/       # Adapters (persistence, writers)
└── ...
```

## Related Repositories

| Repository | Description |
|------------|-------------|
| [loopforge-ai/forge](https://github.com/loopforge-ai/forge) | Skill forging MCP server |
| [loopforge-ai/utils](https://github.com/loopforge-ai/utils) | Shared Go utility packages |
| [loopforge-ai/docs](https://github.com/loopforge-ai/docs) | Skills, workflows, agent personas |

## License

<!-- TODO: add license -->
