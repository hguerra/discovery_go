package home

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/web/res"
)

func RegisterHomeRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res.HTML(w, http.StatusOK, "home", "word")
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		res.JSON2(w, http.StatusOK, res.NewResponse(res.M{"message": "pong!"}))
	})

	return r
}
