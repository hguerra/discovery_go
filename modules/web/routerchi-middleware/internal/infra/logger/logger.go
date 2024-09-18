package logger

import (
	"log/slog"
	"os"
	"strings"

	"routerchi-middleware/internal/infra/config"
)

func parseLevel(levelStr string) slog.Level {
	switch strings.ToUpper(levelStr) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func NewHandlerOptions(cfg config.Configuration) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: cfg.Log.AddSource,
		Level:     parseLevel(cfg.Log.Level),
	}
}

func NewHandler(cfg config.Configuration) slog.Handler {
	opts := NewHandlerOptions(cfg)

	var handler slog.Handler = slog.NewJSONHandler(os.Stdout, opts)
	if cfg.IsDevelopment() {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return handler
}

func NewLogger(cfg config.Configuration) *slog.Logger {
	return slog.New(NewHandler(cfg)).
		With("service", cfg.Name).
		With("version", cfg.Version).
		With("env", cfg.Env)
}
