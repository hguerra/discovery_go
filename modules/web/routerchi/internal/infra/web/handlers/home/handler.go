package home

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/config"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/web/res"
)

func RegisterHomeRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res.HTML(w, http.StatusOK, "home", "word")
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		dto := res.M{
			"message": "pong!",
			"name":    config.GetString("app.name"),
			"secret":  config.GetString("mysecret"),
		}
		res.JSON(w, http.StatusOK, res.NewResponse(dto))
	})

	return r
}
