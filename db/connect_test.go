package db

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func MustConnect(t *testing.T) *Pool {
	pool, err := Connect(context.Background(), os.Getenv("PG_DSN"))
	require.NoError(t, err)
	return pool
}
