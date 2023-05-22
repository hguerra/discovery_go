package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/config"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/logging"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/web/handlers/home"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/web/handlers/user"
)

const (
	requestTimeout    = 60 * time.Second
	readHeaderTimeout = 3 * time.Second
)

func NewServer() {
	logger := logging.GetLogger()
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Timeout(requestTimeout))

	r.Use(middleware.Heartbeat("/health"))

	// https://go-chi.io/#/pages/routing
	// https://thedevelopercafe.com/articles/restful-routing-with-chi-in-go-d05a2f952b3d#path-parameters
	r.Mount("/", home.RegisterHomeRoutes())

	r.Route("/api/v1.0", func(r chi.Router) {
		r.Mount("/users", user.RegisterUserRoutes())
	})

	address := fmt.Sprintf(":%s", config.GetString("server.port"))
	s := &http.Server{
		Addr:              address,
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	logger.Infof("Active profile: %s", config.GetActiveProfile())
	logger.Infof("Listening and serving HTTP on %s", address)
	err := s.ListenAndServe()
	if err != nil {
		logger.Fatalf("Error to starting server in %s: %v", address, err)
	}
}
