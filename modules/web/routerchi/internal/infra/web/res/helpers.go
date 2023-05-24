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
	assetsURL     = ""
	assetsURLOnce = sync.Once{}
	manifest      = make(map[string]string)
	manifestOnce  = sync.Once{}
)

func getAssetsURL() string {
	assetsURLOnce.Do(func() {
		assetsURL = config.GetString("server.assets.url")
	})
	return assetsURL
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
			_, keyPath, keyFound := strings.Cut(k, assetsBasePath)
			_, valuePath, valueFound := strings.Cut(v, assetsBasePath)
			if keyFound && valueFound {
				manifest[keyPath] = fmt.Sprintf("%s%s%s", getAssetsURL(), assetsBasePath, valuePath)
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
	return fmt.Sprintf("%s%s/stylesheets/%s.css", getAssetsURL(), assetsBasePath, file)
}

func imageSrc(file string) string {
	return fmt.Sprintf("%s%s/images/%s", getAssetsURL(), assetsBasePath, file)
}

func publicSrc(file string) string {
	return fmt.Sprintf("%s/public/%s", getAssetsURL(), file)
}

func registerHelpers() template.FuncMap {
	return template.FuncMap{
		"isDev":                   config.IsDev,
		"isProd":                  config.IsProd,
		"assetsURL":               getAssetsURL,
		"javascriptSrc":           javascriptSrc,
		"javascriptStylesheetSrc": javascriptStylesheetSrc,
		"stylesheetSrc":           stylesheetSrc,
		"imageSrc":                imageSrc,
		"publicSrc":               publicSrc,
	}
}
