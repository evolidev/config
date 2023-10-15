package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	SetDirectory("./configs")

	t.Run("config should get config value", func(t *testing.T) {
		conf := NewConfig("storage")

		result := conf.Get("local.path")

		assert.Equal(t, "storage", result.Value())
	})

	t.Run("config should get config value directly", func(t *testing.T) {
		conf := NewConfig("storage.local.path")

		result := conf.Value()

		assert.Equal(t, "storage", result)
	})

	t.Run("config should return sub config if key points to sub config", func(t *testing.T) {
		conf := NewConfig("storage")

		result := conf.Get("local")

		if _, ok := result.Value().(*Config); !ok {
			t.Errorf("Not a config")
			t.Fail()
		}

		result = result.Value().(*Config).Get("path")
		assert.Equal(t, "storage", result.Value())
	})

	t.Run("config should be overridden by environment variable", func(t *testing.T) {
		os.Setenv("STORAGE_LOCAL_PATH", "test")
		conf := NewConfig("storage.local.path")

		result := conf.Value()

		assert.Equal(t, "test", result)
	})
}

func TestSetConfig(t *testing.T) {
	SetDirectory("./configs")

	t.Run("config should set config value", func(t *testing.T) {
		conf := NewConfig("storage")
		conf.Set("local.path", "test")

		result := conf.Get("local.path")

		assert.Equal(t, "test", result.Value())
	})
}
