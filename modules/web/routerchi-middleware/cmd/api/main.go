package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"routerchi-middleware/internal/infra/config"
	"routerchi-middleware/internal/infra/logger"
	"routerchi-middleware/internal/infra/server"
	"sync"
	"time"
)

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	p := fmt.Sprintf("./configs/config.%s.json", config.GetAppEnv())
	cfg, err := config.NewConfig(p)
	if err != nil {
		return err
	}

	log := logger.NewLogger(cfg)
	slog.SetDefault(log)

	srv := server.NewServer(cfg)

	go func() {
		log.Info("listening on", slog.String("address", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("error listening and serving", slog.String("address", srv.Addr), slog.Any("error", err))
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()

		log.Info("shutting down http server", slog.String("address", srv.Addr))
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Error("error shutting down http server", slog.String("address", srv.Addr), slog.Any("error", err))
		}
	}()
	wg.Wait()

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
