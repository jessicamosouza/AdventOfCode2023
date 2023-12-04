package problem1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCubeConundrum(t *testing.T) {
	t.Run("get result right", func(t *testing.T) {
		resp := CubeConundrum()
		require.Equal(t, 2551, resp)
	})
}
