package home

import (
	"net/http"

	"github.com/go-chi/chi"
)

func RegisterHomeRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		/* #nosec G104 */
		w.Write([]byte("handlers home")) //nolint:errcheck
	})

	return r
}
