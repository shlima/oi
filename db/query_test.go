package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	t.Run("when exists", func(t *testing.T) {
		ctx := context.Background()
		pool := MustConnect(t)
		t.Cleanup(pool.Close)

		out := make([]string, 0)
		query := Query().Select("*").From("(VALUES('1'), ('2')) t")
		err := Select(ctx, pool, &out, query)
		require.NoError(t, err)
		require.ElementsMatch(t, out, []string{"1", "2"})
	})

	t.Run("when empty", func(t *testing.T) {
		ctx := context.Background()
		pool := MustConnect(t)
		t.Cleanup(pool.Close)

		out := make([]string, 0)
		query := Query().Select("1").Where("1=0")
		err := Select(ctx, pool, &out, query)
		require.NoError(t, err)
		require.Len(t, out, 0)
	})
}

func TestGet(t *testing.T) {
	t.Run("when exists", func(t *testing.T) {
		ctx := context.Background()
		pool := MustConnect(t)
		t.Cleanup(pool.Close)

		out := ""
		query := Query().Select("*").From("(VALUES('1')) t")
		err := Get(ctx, pool, &out, query)
		require.NoError(t, err)
		require.Equal(t, out, "1")
	})

	t.Run("when empty", func(t *testing.T) {
		ctx := context.Background()
		pool := MustConnect(t)
		t.Cleanup(pool.Close)

		out := ""
		query := Query().Select("1").Where("1=0")
		err := Get(ctx, pool, &out, query)
		require.ErrorIs(t, err, ErrNotFound)
	})
}

func TestQueryRow(t *testing.T) {
	t.Run("when exists", func(t *testing.T) {
		ctx := context.Background()
		pool := MustConnect(t)
		t.Cleanup(pool.Close)

		out := ""
		query := Query().Select("*").From("(VALUES('1')) t")
		err := QueryRow(ctx, pool, query).Scan(&out)
		require.NoError(t, err)
		require.Equal(t, out, "1")
	})

	t.Run("when empty", func(t *testing.T) {
		ctx := context.Background()
		pool := MustConnect(t)
		t.Cleanup(pool.Close)

		out := ""
		query := Query().Select("1").Where("1=0")
		err := QueryRow(ctx, pool, query).Scan(&out)
		require.ErrorIs(t, err, ErrNotFound)
	})
}

func TestWithTx(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		ctx := context.Background()
		pool := MustConnect(t)
		t.Cleanup(pool.Close)

		id1, id2 := 0, 0
		err1 := Get(ctx, pool, &id1, Expr("SELECT txid_current()"))
		err2 := WithTx(ctx, pool, func(tx Tx) error {
			return Get(ctx, pool, &id2, Expr("SELECT txid_current()"))
		})
		require.NoError(t, err1)
		require.NoError(t, err2)
		require.NotEqual(t, id1, id2)
	})
}
