package codec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProto2Hex(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		proto, hex := MustBuildProtoHex(t)
		got, err := Proto2Hex(proto)
		require.NoError(t, err)
		require.Equal(t, hex, got)
	})
}

func TestMustProto2Hex(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		proto, hex := MustBuildProtoHex(t)
		got := MustProto2Hex(proto)
		require.Equal(t, hex, got)
	})
}

func TestHex2Proto(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		proto, hex := MustBuildProtoHex(t)
		got, err := Hex2Proto[*TestProto](hex, new(TestProto))
		require.NoError(t, err)
		require.Equal(t, proto.Title, got.Title)
	})

	t.Run("it errors", func(t *testing.T) {
		_, err := Hex2Proto[*TestProto]("foo", new(TestProto))
		require.Error(t, err)
	})
}

func TestMustHex2Proto(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		proto, hex := MustBuildProtoHex(t)
		got := MustHex2Proto[*TestProto](hex, new(TestProto))
		require.Equal(t, proto.Title, got.Title)
	})

	t.Run("it panics", func(t *testing.T) {
		require.Panics(t, func() {
			_ = MustHex2Proto[*TestProto]("foo", new(TestProto))
		})
	})
}

func TestMarshalProto(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		proto, _ := MustBuildProtoHex(t)
		bytea, err := MarshalProto(proto)
		require.NoError(t, err)

		got, err := UnmarshalProto(bytea, new(TestProto))
		require.NoError(t, err)
		require.Equal(t, proto.Title, got.Title)
	})
}

func TestUnmarshalProto(t *testing.T) {
	t.Run("it errors", func(t *testing.T) {
		_, err := UnmarshalProto([]byte{1, 1, 1, 1, 1}, new(TestProto))
		require.Error(t, err)
	})
}

func MustBuildProtoHex(t *testing.T) (*TestProto, string) {
	return &TestProto{Title: "foo"}, "1203666f6f"
}
