package server

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type M map[string]interface{}

const READ_HEADER_TIMEOUT = 3 * time.Second

var appEnv = os.Getenv("APP_ENV")

func LoadTemplatesV1(logger *slog.Logger) *template.Template {
	tpl := template.Must(
		template.Must(template.ParseGlob("web/templates/*.html")).ParseGlob("web/templates/layouts/*.html"),
	)

	for _, t := range tpl.Templates() {
		logger.Info("Loading template...", slog.String("name", t.Name()))
	}

	return tpl
}

func LoadTemplatesV2(logger *slog.Logger) *template.Template {
	tpl := template.Must(
		template.Must(template.ParseGlob("web/templates/*.html")).ParseGlob("web/templates/*/*.html"),
	)

	for _, t := range tpl.Templates() {
		logger.Info("Loading template...", slog.String("name", t.Name()))
	}

	return tpl
}

func NewServer() {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)
	if appEnv == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	address := fmt.Sprintf(":%s", "8080")
	mux := http.NewServeMux()

	// tpl := LoadTemplatesV1(logger)
	tpl := LoadTemplatesV2(logger)

	mux.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		data := M{"name": "Batman"}
		err := tpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		data := M{"name": "Batman"}
		err := tpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	s := &http.Server{
		Addr:              address,
		Handler:           mux,
		ReadHeaderTimeout: READ_HEADER_TIMEOUT,
		ErrorLog:          slog.NewLogLogger(handler, slog.LevelError),
	}

	logger.Info("Listening and serving HTTP", slog.String("address", address))
	err := s.ListenAndServe()
	if err != nil {
		logger.Error("Error to starting server", slog.String("address", address), slog.Any("error", err))
	}
}
