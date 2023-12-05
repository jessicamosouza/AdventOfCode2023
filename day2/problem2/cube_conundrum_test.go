package problem1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCubeConundrumTwo(t *testing.T) {
	t.Run("get result right", func(t *testing.T) {
		resp := CubeConundrumTwo()
		require.Equal(t, 62811, resp)
	})
}
