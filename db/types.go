package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrNotFound = pgx.ErrNoRows

type (
	Pool       = pgxpool.Pool
	Tx         = pgx.Tx
	TxOpts     = pgx.TxOptions
	Rows       = pgx.Rows
	Row        = pgx.Row
	CommandTag = pgconn.CommandTag
)

// IDB works in the same manner with Pool or Tx
type IDB interface {
	// Begin starts a transaction or pseudo nested transaction.
	Begin(ctx context.Context) (Tx, error)
	Query(ctx context.Context, sql string, args ...interface{}) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (CommandTag, error)
	BeginFunc(ctx context.Context, f func(Tx) error) error
}
