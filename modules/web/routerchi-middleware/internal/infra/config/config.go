package config

import (
	"encoding/json"
	"os"
)

const (
	KEY_APP_ENV     = "APP_ENV"
	DEFAULT_APP_ENV = "development"
)

type Configuration struct {
	Version  int    `json:"version"`
	Name     string `json:"name"`
	Env      string `json:"env"`
	LogLevel string `json:"logLevel"`
	Server   struct {
		Port int `json:"port"`
	} `json:"server"`
}

func Getenv(key, fallback string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return fallback
}

func GetAppEnv() string {
	return Getenv(KEY_APP_ENV, DEFAULT_APP_ENV)
}

func NewConfig(file string) (*Configuration, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var cfg Configuration
	err = json.Unmarshal(f, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Configuration) IsDevelopment() bool {
	return c.Env == DEFAULT_APP_ENV
}
