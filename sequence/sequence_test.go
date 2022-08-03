package sequence

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSequence_NextInt64(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		seq := MustNewSequence(t)
		require.Equal(t, int64(1), seq.NextInt64())
		require.Equal(t, int64(2), seq.NextInt64())
	})
}

func TestSequence_NextInt32(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		seq := MustNewSequence(t)
		require.Equal(t, int32(1), seq.NextInt32())
		require.Equal(t, int32(2), seq.NextInt32())
	})
}

func TestSequence_NextString(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		seq := MustNewSequence(t)
		require.Equal(t, "1", seq.NextString())
		require.Equal(t, "2", seq.NextString())
	})
}

func TestSequence_RandomInt64(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		seq := MustNewSequence(t)
		require.NotEqual(t, seq.RandomInt64(), seq.RandomInt64())
	})
}

func TestSequence_RandomInt32(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		seq := MustNewSequence(t)
		require.NotEqual(t, seq.RandomInt32(), seq.RandomInt32())
	})
}

func TestSequence_RandomTime(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		seq := MustNewSequence(t)
		require.NotEqual(t, seq.RandomTime(), seq.RandomTime())
	})
}

func MustNewSequence(t *testing.T) *Sequence {
	return New(0)
}
