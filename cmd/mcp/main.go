package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/loopforge-ai/utils/mcp"
)

// appVersion is set at build time via -ldflags "-X main.appVersion=x.y.z".
var appVersion = "dev"

func main() {
	server := mcp.NewServer("template", appVersion)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	err := server.Serve(ctx)
	stop()

	if err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("serve: %v", err)
	}
}
