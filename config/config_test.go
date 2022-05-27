package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPort(t *testing.T) {
	p := GetPort()
	assert.Equal(t, os.Getenv("MOSH_PORT"), p)
}
func TestCacheDir(t *testing.T) {
	p := GetCacheDir()
	assert.Equal(t, os.Getenv("MOSH_CACHE_DIR")+"/", p)
}

func TestGetLogidr(t *testing.T) {
	p := GetLogir()
	assert.Equal(t, os.Getenv("MOSH_LOG_DIR")+"/", p)
}

func TestGetConfig(t *testing.T) {
	conf := GetConfig()
	assert.Equal(t, UNINITIALIZED, conf.Token)
	assert.Equal(t, UNINITIALIZED, conf.Address)
	assert.Equal(t, UNINITIALIZED, conf.Port)
	assert.Equal(t, UNINITIALIZED, conf.Library)
}
