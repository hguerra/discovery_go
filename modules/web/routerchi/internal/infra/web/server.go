package web

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/web/handlers/home"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/web/handlers/user"
)

func NewServer() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(middleware.Heartbeat("/health"))

	// https://go-chi.io/#/pages/routing
	// https://thedevelopercafe.com/articles/restful-routing-with-chi-in-go-d05a2f952b3d#path-parameters
	r.Mount("/", home.RegisterHomeRoutes())

	r.Route("/api/v1.0", func(r chi.Router) {
		r.Mount("/users", user.RegisterUserRoutes())
	})

	s := &http.Server{
		Addr:              ":3000",
		Handler:           r,
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Error to starting server: %v", err)
	}
}
