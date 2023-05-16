package res

import (
	"bytes"
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/config"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/logging"
	"github.com/unrolled/render"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger
var r *render.Render

func init() {
	logger = logging.GetLogger()
	r = render.New(render.Options{
		IsDevelopment: config.IsDev(),
		StreamingJSON: true,
		Directory:     "web/templates",
		Layout:        "layouts/application",
		Funcs:         []template.FuncMap{registerHelpers()},
	})
}

func HTML(w http.ResponseWriter, status int, name string, v any) {
	if err := r.HTML(w, status, name, v); err != nil {
		logger.Panic(err)
	}
}

func JSON(w http.ResponseWriter, status int, v any) {
	if err := r.JSON(w, status, v); err != nil {
		logger.Panic(err)
	}
}

// JSON marshals 'v' to JSON, automatically escaping HTML and setting the
// Content-Type as application/json.
// Based on:
// https://github.com/go-chi/render/blob/master/responder.go#L93
// https://github.com/gmhafiz/go8/blob/master/internal/utility/respond/json.go
func JSON2(w http.ResponseWriter, status int, v any) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	/* #nosec G104 */
	w.Write(buf.Bytes()) //nolint:errcheck
}
