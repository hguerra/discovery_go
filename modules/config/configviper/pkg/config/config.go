package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	strings map[string]string
	ints    map[string]int
}

func (c *Config) GetString(key string) string {
	v, ok := c.strings[key]
	if ok {
		return v
	}
	v = viper.GetString(key)
	c.strings[key] = v
	return v
}

func (c *Config) GetInt(key string) int {
	v, ok := c.ints[key]
	if ok {
		return v
	}
	v = viper.GetInt(key)
	c.ints[key] = v
	return v
}

func NewConfig(cfgFile, path string) (*Config, error) {
	env, ok := os.LookupEnv("APP_ENV")
	if !ok {
		env = "development"
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(path)
		viper.SetConfigType("env")
		viper.SetConfigName(fmt.Sprintf(".env.%s", env))
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	log.Printf("Loading config file: %s", viper.ConfigFileUsed())

	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	err = viper.MergeInConfig()
	if err == nil {
		log.Printf("Override env using config file: %s", viper.ConfigFileUsed())
	}

	return &Config{strings: make(map[string]string), ints: make(map[string]int)}, nil
}
