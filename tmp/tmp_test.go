package tmp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTmp_Dir(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		tmp := New()
		dir1, err1 := tmp.Dir("uno")
		dir2, err2 := tmp.Dir("uno")

		require.NoError(t, err1)
		require.NoError(t, err2)
		require.NotEqual(t, dir1, dir2)

		require.True(t, IsDirExists(t, dir1))
		require.True(t, IsDirExists(t, dir2))

		err := tmp.Close()
		require.NoError(t, err)

		require.False(t, IsDirExists(t, dir1))
		require.False(t, IsDirExists(t, dir2))
	})
}

func TestTmp_Ашду(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		tmp := New()
		file1, err1 := tmp.File("uno")
		file2, err2 := tmp.File("uno")

		require.NoError(t, err1)
		require.NoError(t, err2)
		require.NotEqual(t, file1, file2)

		require.True(t, IsDirExists(t, file1.Name()))
		require.True(t, IsDirExists(t, file2.Name()))

		err := tmp.Close()
		require.NoError(t, err)

		require.False(t, IsDirExists(t, file1.Name()))
		require.False(t, IsDirExists(t, file2.Name()))
	})
}

func IsDirExists(t *testing.T, name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
