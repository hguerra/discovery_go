package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMap(t *testing.T) {
	cfg, err := loadMap("../../test/testdata/config.test.json")
	assert.Nil(t, err)
	assert.NotNil(t, cfg)

	strValue, found := searchMap(cfg, []string{"name"})
	assert.True(t, found)
	assert.Equal(t, "MyApp", strValue)

	floatValue, found := searchMap(cfg, []string{"version"})
	assert.True(t, found)
	assert.Equal(t, float64(1), floatValue)

	nested1, found := searchMap(cfg, []string{"a", "b"})
	assert.True(t, found)
	assert.Equal(t, "c", nested1)

	nested2, found := searchMap(cfg, []string{"a", "d", "e"})
	assert.True(t, found)
	assert.Equal(t, float64(2), nested2)

	nested3, found := searchMap(cfg, []string{"endpoints"})
	assert.True(t, found)
	assert.Equal(t, 1, len(nested3.([]interface{})))

	nested4, found := searchMap(cfg, []string{"endpoints", "0", "endpoint"})
	assert.True(t, found)
	assert.Equal(t, "https://api.test.com", nested4)

	nested5, found := searchMap(cfg, []string{"fakeArray"})
	assert.True(t, found)
	assert.Equal(t, 0, len(nested5.([]interface{})))
}
