package res

import (
	"net/http"
	"sync"
	"text/template"

	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/config"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/logging"
	"github.com/unrolled/render"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger = logging.GetLogger()
	r      *render.Render
	rOnce  = sync.Once{}
)

func getRender() *render.Render {
	rOnce.Do(func() {
		r = render.New(render.Options{
			IsDevelopment: config.IsDev(),
			StreamingJSON: true,
			Directory:     "web/templates",
			Layout:        "layouts/application",
			Funcs:         []template.FuncMap{registerHelpers()},
		})
	})
	return r
}

func HTML(w http.ResponseWriter, status int, name string, v any) {
	if err := getRender().HTML(w, status, name, v); err != nil {
		logger.Panic(err)
	}
}

func JSON(w http.ResponseWriter, status int, v any) {
	if err := getRender().JSON(w, status, v); err != nil {
		logger.Panic(err)
	}
}
