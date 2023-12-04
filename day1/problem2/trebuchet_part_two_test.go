package problem2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTrebuchet(t *testing.T) {
	t.Run("return 0", func(t *testing.T) {
		resp := Trebuchet()
		require.Equal(t, 54100, resp)
	})
}
