package res

import (
	"fmt"
	"text/template"

	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/config"
)

func assetsUrl() string {
	return config.GetString("server.assets.url")
}

func javascriptSrc(file string) string {
	return fmt.Sprintf("%s/assets/javascripts/%s.js", assetsUrl(), file)
}

func javascriptStylesheetSrc(file string) string {
	return fmt.Sprintf("%s/assets/javascripts/%s.css", assetsUrl(), file)
}

func javascriptImageSrc(file string) string {
	return fmt.Sprintf("%s/%s", assetsUrl(), file)
}

func stylesheetSrc(file string) string {
	return fmt.Sprintf("%s/assets/stylesheets/%s", assetsUrl(), file)
}

func registerHelpers() template.FuncMap {
	return template.FuncMap{
		"isDev":                   config.IsDev,
		"assetsUrl":               assetsUrl,
		"javascriptSrc":           javascriptSrc,
		"javascriptStylesheetSrc": javascriptStylesheetSrc,
		"javascriptImageSrc":      javascriptImageSrc,
		"stylesheetSrc":           stylesheetSrc,
	}
}
