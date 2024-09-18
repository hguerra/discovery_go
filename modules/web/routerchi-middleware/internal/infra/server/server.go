package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"routerchi-middleware/internal/infra/config"
	"routerchi-middleware/internal/infra/logger"
)

func NewServer(cfg config.Configuration) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      NewMux(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.NewHandler(cfg), slog.LevelError),
	}
}
