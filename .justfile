# Build all binaries
build: lint test
    go build -trimpath -ldflags "-s -w" -o bin/cli ./cmd/cli
    go build -trimpath -ldflags "-s -w" -o bin/mcp ./cmd/mcp
    go build -trimpath -ldflags "-s -w" -o bin/server ./cmd/server

# Build the CLI binary
build-cli: lint test
    go build -trimpath -ldflags "-s -w" -o bin/cli ./cmd/cli

# Build the MCP binary
build-mcp: lint test
    go build -trimpath -ldflags "-s -w" -o bin/mcp ./cmd/mcp

# Build the server binary
build-server: lint test
    go build -trimpath -ldflags "-s -w" -o bin/server ./cmd/server

# Build the Docker image
docker-build version="dev":
    docker build --build-arg VERSION={{ version }} -t template:{{ version }} .

# Run the Docker container (server on port 8080)
docker-run version="dev" port="8080":
    docker run --rm -p {{ port }}:8080 template:{{ version }}

# Stop all running template containers
docker-stop:
    docker ps -q --filter ancestor=template | xargs -r docker stop

# Install all binaries
install: lint test
    go install -trimpath -ldflags "-s -w" ./cmd/cli
    go install -trimpath -ldflags "-s -w" ./cmd/mcp
    go install -trimpath -ldflags "-s -w" ./cmd/server

# Run linter
lint:
    golangci-lint run ./...

# Start the HTTP server locally
server:
    go run ./cmd/server

# Run tests with coverage
test:
    go test -cover ./...

# Validate template quality contract (lint + test + Docker build)
validate: lint test
    docker build -t template:validate .
    @echo "Template quality contract: PASS"
