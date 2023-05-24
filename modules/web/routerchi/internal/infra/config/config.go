package config

import (
	"github.com/spf13/viper"
)

type configCache struct {
	strings map[string]string
	ints    map[string]int
}

var config = &configCache{
	strings: make(map[string]string),
	ints:    make(map[string]int),
}

func GetString(key string) string {
	v, ok := config.strings[key]
	if ok {
		return v
	}
	v = viper.GetString(key)
	config.strings[key] = v
	return v
}

func GetInt(key string) int {
	v, ok := config.ints[key]
	if ok {
		return v
	}
	v = viper.GetInt(key)
	config.ints[key] = v
	return v
}
