package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// appVersion is set at build time via -ldflags "-X main.appVersion=x.y.z".
var appVersion = "dev"

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	fmt.Printf("template-cli %s\n", appVersion)

	// Placeholder: add CLI logic here.
	<-ctx.Done()
}
