package bug

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStdErr(t *testing.T) {
	t.Run("it uses STDERR as output", func(t *testing.T) {
		got := NewStdErr("foo").Logger.Writer()
		file, ok := got.(*os.File)

		require.True(t, ok)
		require.Equal(t, "/dev/stderr", file.Name())
	})
}

func TestLogger_Ensure(t *testing.T) {
	t.Run("it writes if error", func(t *testing.T) {
		logger, buffer := MustNewLogger(t, "foo")
		logger.Ensure(errors.New("bar"))
		got := buffer.String()
		require.Contains(t, got, "foo", "contains prefix")
		require.Contains(t, got, "logger.go", "contains filename")
		require.Contains(t, got, "bar", "contains error")
	})

	t.Run("it do nothing if no", func(t *testing.T) {
		logger, buffer := MustNewLogger(t, "foo")
		logger.Ensure(nil)
		require.Equal(t, "", buffer.String())
	})
}

func MustNewLogger(t *testing.T, prefix string) (*Logger, *bytes.Buffer) {
	buffer := bytes.NewBufferString("")
	logger := NewLogger(buffer, prefix)
	return logger, buffer
}
