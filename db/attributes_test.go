package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAttributes_Add(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got := make(Attributes).Add("foo", "bar")
		require.EqualValues(t, "bar", got["foo"])
	})
}

func TestAttributes_Columns(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got := Attributes{"foo": "bar"}
		require.Equal(t, []string{"foo"}, got.Columns())
	})
}

func TestAttributes_PickValues(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		got := Attributes{"a": "1", "b": "2"}
		require.Equal(t, []any{"1", "2"}, got.PickValues([]string{"a", "b"}))
		require.Equal(t, []any{"2", "1"}, got.PickValues([]string{"b", "a"}))
		require.Equal(t, []any{nil}, got.PickValues([]string{""}))
	})
}
