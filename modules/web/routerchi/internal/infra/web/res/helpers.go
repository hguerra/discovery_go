package res

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"text/template"

	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/config"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/logging"
)

const assetsBasePath = "/assets/"

var (
	assetsUrl     = ""
	assetsUrlOnce = sync.Once{}
	manifest      = make(map[string]string)
	manifestOnce  = sync.Once{}
)

func getAssetsUrl() string {
	assetsUrlOnce.Do(func() {
		assetsUrl = config.GetString("server.assets.url")
	})
	return assetsUrl
}

func getManifest(file string) string {
	manifestOnce.Do(func() {
		jsonFile, err := os.Open(config.GetString("server.assets.manifest"))
		logging.Catch(err)
		defer jsonFile.Close()

		jsonBytes, err := io.ReadAll(jsonFile)
		logging.Catch(err)

		var rawManifest map[string]string
		logging.Catch(json.Unmarshal(jsonBytes, &rawManifest))

		for k, v := range rawManifest {
			_, kPath, kFound := strings.Cut(k, assetsBasePath)
			_, vPath, vFound := strings.Cut(v, assetsBasePath)
			if kFound && vFound {
				manifest[kPath] = fmt.Sprintf("%s%s%s", getAssetsUrl(), assetsBasePath, vPath)
			}
		}
	})
	return manifest[file]
}

func javascriptSrc(file string) string {
	return getManifest(fmt.Sprintf("javascripts/%s.ts", file))
}

func javascriptStylesheetSrc(file string) string {
	return getManifest(fmt.Sprintf("javascripts/%s.css", file))
}

func stylesheetSrc(file string) string {
	return fmt.Sprintf("%s%s/stylesheets/%s.css", getAssetsUrl(), assetsBasePath, file)
}

func imageSrc(file string) string {
	return fmt.Sprintf("%s%s/images/%s", getAssetsUrl(), assetsBasePath, file)
}

func publicSrc(file string) string {
	return fmt.Sprintf("%s/public/%s", getAssetsUrl(), file)
}

func registerHelpers() template.FuncMap {
	return template.FuncMap{
		"isDev":                   config.IsDev,
		"isProd":                  config.IsProd,
		"assetsUrl":               getAssetsUrl,
		"javascriptSrc":           javascriptSrc,
		"javascriptStylesheetSrc": javascriptStylesheetSrc,
		"stylesheetSrc":           stylesheetSrc,
		"imageSrc":                imageSrc,
		"publicSrc":               publicSrc,
	}
}
