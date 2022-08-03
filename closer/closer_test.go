package closer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCloser_close(t *testing.T) {
	t.Run("it calls the closer", func(t *testing.T) {
		a := 0
		closer := New(os.Kill)
		closer.Add(func() { a = 1 })
		closer.close()
		require.Equal(t, 1, a)
	})
}
