package config_test

import (
	"fmt"
	"testing"

	"github.com/hguerra/discovery_go/modules/config/configviper/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	cfgPath := "../../configs"

	t.Run("should load testing config from file", func(t *testing.T) {
		t.Setenv("APP_ENV", "")
		cfg, err := config.NewConfig(fmt.Sprintf("%s/.env.test", cfgPath), cfgPath)
		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		assert.Equal(t, "test", cfg.GetString("APP_ENV"))
		assert.False(t, cfg.IsProduction())
		assert.False(t, cfg.IsDevelopment())
		assert.True(t, cfg.IsTest())
	})

	t.Run("should load testing config from APP_ENV", func(t *testing.T) {
		t.Setenv("APP_ENV", "test")
		cfg, err := config.NewConfig("", cfgPath)
		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		assert.Equal(t, "test", cfg.GetString("APP_ENV"))
		assert.False(t, cfg.IsProduction())
		assert.False(t, cfg.IsDevelopment())
		assert.True(t, cfg.IsTest())
	})
}
