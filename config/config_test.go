package config_test

import (
	"github.com/dionofrizal88/go-allocator/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigGetConfiguration(t *testing.T) {
	conf := config.GetConfig()

	t.Run("positive case to test config, expected no error", func(t *testing.T) {
		t.Run("positive case while use func get configuration, expected no error", func(t *testing.T) {

			assert.NotEmpty(t, conf.AppName)
			assert.NotEmpty(t, conf.QiscusSecretKey)
		})
	})

}
