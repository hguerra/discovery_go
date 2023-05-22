package config_test

import (
	"testing"

	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/config"
	"github.com/stretchr/testify/assert"
)

func TestGetActiveProfile(t *testing.T) {
	assert.Equal(t, config.GetActiveProfile(), "test", "default profile should be test")
}

func TestIsProd(t *testing.T) {
	assert.False(t, config.IsProd())
}

func TestIsDev(t *testing.T) {
	assert.False(t, config.IsDev())
}

func TestIsTest(t *testing.T) {
	assert.True(t, config.IsTest())
}
