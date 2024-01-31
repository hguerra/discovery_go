package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	assert.Equal(t, "development", GetEnv())

	env := "test"
	os.Setenv(KEY_APP_ENV, env)
	assert.Equal(t, env, GetEnv())

	cfg, err := Load(fmt.Sprintf("../../test/testdata/config.%s.json", env))
	assert.Nil(t, err)
	assert.Equal(t, "test", cfg.GetEnv())
}

func TestLoad(t *testing.T) {
	_, err := Load("abc.json")
	assert.NotNil(t, err)

	cfg, err := Load("../../test/testdata/config.test.json")
	assert.Nil(t, err)

	strCfg, found := cfg.GetString("name")
	assert.True(t, found)
	assert.Equal(t, "MyApp", strCfg)

	strCache, found := cfg.GetString("name")
	assert.True(t, found)
	assert.Equal(t, "MyApp", strCache)

	strPath, found := cfg.GetString("endpoints.0.endpoint")
	assert.True(t, found)
	assert.Equal(t, "https://api.test.com", strPath)

	strNotFoundCfg, found := cfg.GetString("notFound GetString")
	assert.False(t, found)
	assert.Empty(t, strNotFoundCfg)

	strNotFoundCache, found := cfg.GetString("notFound GetString")
	assert.False(t, found)
	assert.Empty(t, strNotFoundCache)

	os.Setenv("APP_ABC_CDE", "test-var")
	strEnvCfg, found := cfg.GetString("app.abc.cde")
	assert.True(t, found)
	assert.Equal(t, "test-var", strEnvCfg)

	strEnvCache, found := cfg.GetString("app.abc.cde")
	assert.True(t, found)
	assert.Equal(t, "test-var", strEnvCache)

	floatCfg, found := cfg.GetFloat64("version")
	assert.True(t, found)
	assert.Equal(t, float64(1), floatCfg)

	floatCache, found := cfg.GetFloat64("version")
	assert.True(t, found)
	assert.Equal(t, float64(1), floatCache)

	floatNotFoundCfg, found := cfg.GetFloat64("notFound GetFloat64")
	assert.False(t, found)
	assert.Empty(t, floatNotFoundCfg)

	floatNotFoundCache, found := cfg.GetFloat64("notFound GetFloat64")
	assert.False(t, found)
	assert.Empty(t, floatNotFoundCache)

	os.Setenv("SERVER_PORT", "3000")
	floatEnvCfg, found := cfg.GetFloat64("server.port")
	assert.True(t, found)
	assert.Equal(t, float64(3000), floatEnvCfg)

	floatEnvCache, found := cfg.GetFloat64("server.port")
	assert.True(t, found)
	assert.Equal(t, float64(3000), floatEnvCache)

	int8EnvCfg, found := cfg.GetInt8("a.d.e")
	assert.True(t, found)
	assert.Equal(t, int8(2), int8EnvCfg)

	int16EnvCfg, found := cfg.GetInt16("server.port")
	assert.True(t, found)
	assert.Equal(t, int16(3000), int16EnvCfg)

	int32EnvCfg, found := cfg.GetInt32("server.port")
	assert.True(t, found)
	assert.Equal(t, int32(3000), int32EnvCfg)

	int64EnvCfg, found := cfg.GetInt64("server.port")
	assert.True(t, found)
	assert.Equal(t, int64(3000), int64EnvCfg)
}
