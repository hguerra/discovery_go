package main

import (
	"fmt"
	"log"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

const PluginName = "header-validation"
const Version = "0.1.0"
const Priority = 1000

const FailedResponse = `{"error": "%s is required"}`

type Config struct {
	HeaderKey string `json:"header_key"`
}

func main() {
	err := server.StartServer(New, Version, Priority)
	if err != nil {
		log.Fatalf("Failed start %s plugin", PluginName)
	}
}

func New() interface{} {
	return &Config{}
}

func (conf *Config) Access(kong *pdk.PDK) {
	headerKey, err := kong.Request.GetHeader(conf.HeaderKey)
	if err != nil {
		kong.Log.Err(err.Error())
	}

	headerResponse := make(map[string][]string)
	headerResponse["Content-Type"] = []string{"application/json"}

	if headerKey == "" {
		kong.Response.Exit(401, fmt.Sprintf(FailedResponse, conf.HeaderKey), headerResponse)
	}
}
