package format

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDate(t *testing.T) {
	t.Run("when exists", func(t *testing.T) {
		date := time.Date(2000, 1, 2, 11, 12, 13, 15, time.UTC)
		got := Date(date)
		require.Equal(t, "2000-01-02", got)
	})

	t.Run("when empty", func(t *testing.T) {
		got := Date(time.Time{})
		require.Equal(t, "", got)
	})
}

func TestTime(t *testing.T) {
	t.Run("when exists", func(t *testing.T) {
		date := time.Date(2000, 1, 2, 11, 12, 13, 15, time.UTC)
		got := Time(date)
		require.Equal(t, "2000-01-02T11:12:13Z", got)
	})

	t.Run("when empty", func(t *testing.T) {
		got := Time(time.Time{})
		require.Equal(t, "", got)
	})
}
