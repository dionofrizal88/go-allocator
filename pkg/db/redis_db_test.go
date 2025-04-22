package db_test

import (
	"github.com/dionofrizal88/go-allocator/config"
	"github.com/dionofrizal88/go-allocator/pkg/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDBRedis(t *testing.T) {
	conf := config.GetConfig()

	t.Run("positive case to test db redis, expected no error", func(t *testing.T) {
		t.Run("positive case while use func to create redis connection, expected no error", func(t *testing.T) {
			dbConnection, errDBConnection := db.NewRedisConnection(conf)

			require.NoError(t, errDBConnection)
			assert.NotNil(t, dbConnection)
		})
	})

}
