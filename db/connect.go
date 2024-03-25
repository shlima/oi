package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect to to database, example DSN:
// postgres://jack:secret@pg.example.com:5432/db?sslmode=verify-ca&pool_max_conns=10
// all query params:
// - pool_max_conns
// - pool_min_conns
// - pool_max_conn_lifetime
// - pool_max_conn_idle_time
// - pool_health_check_period
func Connect(ctx context.Context, dsn string) (*Pool, error) {
	return pgxpool.New(ctx, dsn)
}
