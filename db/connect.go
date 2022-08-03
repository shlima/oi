package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
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
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DSN: %w", err)
	}

	return pgxpool.ConnectConfig(ctx, config)
}
