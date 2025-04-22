package routes_test

import (
	"github.com/dionofrizal88/go-allocator/pkg/tests"
	"github.com/dionofrizal88/go-allocator/routes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoutesInitiate(t *testing.T) {
	testSuite := tests.InitTestSuite()

	// Init Router
	e := routes.
		NewRouter(
			routes.WithConfig(testSuite.Config),
			routes.WithRedisDB(testSuite.RedisClient),
		).
		Init()

	t.Run("positive case to test routes, expected no error", func(t *testing.T) {
		t.Run("positive case while use func routes initiate, expected no error", func(t *testing.T) {

			require.NotNil(t, e)
		})
	})

}
