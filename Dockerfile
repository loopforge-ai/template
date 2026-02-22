# Stage 1: Build
FROM golang:1.25-alpine AS build

ARG VERSION=dev

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -trimpath \
    -ldflags "-s -w -X main.appVersion=${VERSION}" \
    -o /bin/cli ./cmd/cli

RUN CGO_ENABLED=0 go build -trimpath \
    -ldflags "-s -w -X main.appVersion=${VERSION}" \
    -o /bin/mcp ./cmd/mcp

RUN CGO_ENABLED=0 go build -trimpath \
    -ldflags "-s -w -X main.appVersion=${VERSION}" \
    -o /bin/server ./cmd/server

# Stage 2: Runtime
FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata wget \
    && addgroup -S app && adduser -S app -G app

COPY --from=build /bin/cli /usr/local/bin/cli
COPY --from=build /bin/mcp /usr/local/bin/mcp
COPY --from=build /bin/server /usr/local/bin/server

USER app

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD ["wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:8080/health"]

ENTRYPOINT ["server"]
