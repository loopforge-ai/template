# Project Documentation Index

> Generated: 2026-03-14 | Scan: exhaustive | Mode: initial_scan

## Project Overview

- **Type:** Monolith (single Go module)
- **Primary Language:** Go 1.25.6
- **Architecture:** Hexagonal (Ports & Adapters)
- **Module:** `github.com/loopforge-ai/template`

## Quick Reference

- **Tech Stack:** Go 1.25.6, stdlib `net/http`, `html/template`, `embed.FS`, `loopforge-ai/utils`
- **Entry Points:** `cmd/server` (HTTP), `cmd/mcp` (MCP), `cmd/cli` (CLI)
- **Architecture Pattern:** Hexagonal — domain/inbound/outbound layers per bounded context
- **Bounded Contexts:** `dashboard` (1 context)
- **API Endpoints:** `GET /health` (JSON), `GET /` (HTML), `GET /static/{path...}` (assets)
- **Dependencies:** Zero external — stdlib + `loopforge-ai/utils` only

## Generated Documentation

- [Project Overview](./project-overview.md) — Executive summary, tech stack, and repository structure
- [Architecture](./architecture.md) — Hexagonal architecture, layer rules, component details
- [Source Tree Analysis](./source-tree-analysis.md) — Annotated directory structure and critical folders
- [API Contracts](./api-contracts.md) — HTTP endpoints, middleware chain, configuration
- [Development Guide](./development-guide.md) — Prerequisites, commands, conventions, workflows
- [Deployment Guide](./deployment-guide.md) — Docker build, binary deployment, health monitoring

## Existing Documentation

- [CLAUDE.md](../CLAUDE.md) — AI agent guidelines: coding conventions, agentic loops, forge integration
- [README.md](../README.md) — Project readme (template with TODO placeholders)
- [.golangci.yml](../.golangci.yml) — Linter configuration (v2, all linters enabled)
- [.justfile](../.justfile) — Task runner recipes (build, test, lint, docker)
- [Dockerfile](../Dockerfile) — Multi-stage Docker build producing 3 binaries

## Getting Started

### For Development
```bash
go mod download           # Fetch dependencies
just server               # Run HTTP server locally
just test                 # Run tests with coverage
just lint                 # Run linter
```

### For AI-Assisted Development
1. Read this index for project context
2. Reference [Architecture](./architecture.md) for layer rules
3. Follow conventions in [CLAUDE.md](../CLAUDE.md)
4. Use the Forge skill-first workflow before writing code manually

### For New Features
- New bounded context: `mkdir -p internal/<ctx>/{domain,inbound,outbound}`
- New handler: See [Development Guide](./development-guide.md#adding-a-new-handler)
- New entry point: See [Development Guide](./development-guide.md#adding-a-new-entry-point)
