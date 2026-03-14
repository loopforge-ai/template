package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	dashboard "github.com/loopforge-ai/template/internal/dashboard/domain"
	"github.com/loopforge-ai/template/internal/dashboard/inbound"
	"github.com/loopforge-ai/template/web"
	"github.com/loopforge-ai/utils/env"
	httpserver "github.com/loopforge-ai/utils/html"
)

// appVersion is set at build time via -ldflags "-X main.appVersion=x.y.z".
var appVersion = "dev"

func main() {
	addr := env.Get("SERVER_ADDR", httpserver.DefaultAddr)

	// Domain services.
	renderer, err := httpserver.NewRenderer(web.FS, dashboard.RendererConfig)
	if err != nil {
		log.Fatalf("create renderer: %v", err)
	}

	// Inbound handlers.
	healthHandler := inbound.NewHealthHandler()
	indexHandler := inbound.NewIndexHandler(renderer, appVersion)

	// Routes.
	mux, err := inbound.RegisterRoutes(healthHandler, indexHandler, web.FS)
	if err != nil {
		log.Fatalf("register routes: %v", err)
	}

	// HTTP server.
	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		IdleTimeout:       httpserver.IdleTimeout,
		MaxHeaderBytes:    httpserver.MaxHeaderBytes,
		ReadHeaderTimeout: httpserver.ReadHeaderTimeout,
		ReadTimeout:       httpserver.ReadTimeout,
		WriteTimeout:      httpserver.WriteTimeout,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		slog.Info("server starting", "addr", addr, "version", appVersion)
		if listenErr := srv.ListenAndServe(); listenErr != nil && !errors.Is(listenErr, http.ErrServerClosed) {
			slog.Error("listen", "error", listenErr)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), httpserver.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("shutdown", "error", err)
	}
}
