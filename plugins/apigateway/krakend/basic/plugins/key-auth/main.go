package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// pluginName is the plugin name
var pluginName = "krakend-key-auth"

// HandlerRegisterer is the symbol the plugin loader will try to load. It must implement the Registerer interface
var HandlerRegisterer = registerer(pluginName)

type registerer string

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

func (r registerer) registerHandlers(_ context.Context, extra map[string]interface{}, h http.Handler) (http.Handler, error) {
	// If the plugin requires some configuration, it should be under the name of the plugin. E.g.:
	/*
			   "extra_config":{
			       "plugin/http-server":{
			           "name":["krakend-key-auth"],
			           "krakend-key-auth":{
			               "path": ["/api"],
		                 "consumer": "krakend",
		                 "key": "1234",
		                 "key_name": "auth_token"
			           }
			       }
			   }
	*/
	// The config variable contains all the keys you have defined in the configuration
	// if the key doesn't exists or is not a map the plugin returns an error and the default handler
	config, ok := extra[pluginName].(map[string]interface{})
	if !ok {
		return h, errors.New("configuration not found")
	}

	rawPath, ok := config["path"].([]interface{})
	if !ok || len(rawPath) == 0 {
		return h, errors.New("path not found in configuration")
	}

	path := make([]string, len(rawPath))
	for i := range rawPath {
		path[i] = rawPath[i].(string)
	}

	if len(path) == 0 || strings.TrimSpace(path[0]) == "" {
		return h, errors.New("path not found in configuration")
	}
	logger.Debug(fmt.Sprintf("Add key authentication to path %v", path))

	consumer, ok := config["consumer"].(string)
	if !ok || strings.TrimSpace(consumer) == "" {
		return h, errors.New("consumer (custom identifier) not found in configuration")
	}
	logger.Debug(fmt.Sprintf("Add key authentication for consumer %s", consumer))

	key, ok := config["key"].(string)
	if !ok || strings.TrimSpace(key) == "" {
		return h, errors.New("key not found in configuration")
	}

	keyName, ok := config["key_name"].(string)
	if !ok || strings.TrimSpace(key) == "" {
		keyName = "apikey"
	}

	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http handler
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if containsBasePath(path, req.URL.Path) {
			v := req.URL.Query().Get(keyName)
			if v == "" {
				v = req.Header.Get(keyName)
			}

			if key != v {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		// If the requested path is not what we defined, continue.
		h.ServeHTTP(w, req)
	}), nil
}

func main() {}

// This logger is replaced by the RegisterLogger method to load the one from KrakenD
var logger Logger = noopLogger{}

func (registerer) RegisterLogger(v interface{}) {
	l, ok := v.(Logger)
	if !ok {
		return
	}
	logger = l
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Logger loaded", HandlerRegisterer))
}

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
}

// Empty logger implementation
type noopLogger struct{}

func (n noopLogger) Debug(_ ...interface{})    {}
func (n noopLogger) Info(_ ...interface{})     {}
func (n noopLogger) Warning(_ ...interface{})  {}
func (n noopLogger) Error(_ ...interface{})    {}
func (n noopLogger) Critical(_ ...interface{}) {}
func (n noopLogger) Fatal(_ ...interface{})    {}

func containsBasePath(s []string, str string) bool {
	for _, v := range s {
		if strings.HasPrefix(str, v) {
			return true
		}
	}
	return false
}
