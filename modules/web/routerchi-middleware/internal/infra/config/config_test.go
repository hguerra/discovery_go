package config_test

import (
	"fmt"
	"testing"

	"routerchi-middleware/internal/infra/config"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	assert.Equal(t, "development", config.GetAppEnv())

	env := "test"
	t.Setenv(config.KEY_APP_ENV, env)
	assert.Equal(t, env, config.GetAppEnv())

	p := fmt.Sprintf("../../../configs/config.%s.json", env)
	cfg, err := config.NewConfig(p)
	assert.Nil(t, err)
	assert.Equal(t, "test", cfg.Env)
}
