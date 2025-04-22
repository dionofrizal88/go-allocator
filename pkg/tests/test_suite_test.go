package tests_test

import (
	"github.com/dionofrizal88/go-allocator/pkg/tests"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTestsSuite(t *testing.T) {
	testSuite := tests.InitTestSuite()

	t.Run("positive case to test tests suite, expected no error", func(t *testing.T) {
		t.Run("positive case while use func init test suite, expected no error", func(t *testing.T) {

			require.NotNil(t, testSuite)
			require.NotEmpty(t, testSuite.Config.AppName)
		})
	})

}
