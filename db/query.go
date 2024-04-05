package db

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

func WithTx(ctx context.Context, db IDB, fn func(tx Tx) error) error {
	return pgx.BeginFunc(ctx, db, func(tx pgx.Tx) error {
		return fn(tx)
	})
}

// Select вернет много строк
func Select(ctx context.Context, db IDB, dest interface{}, query Sqlizer) error {
	stmt, args, err := query.ToSql()
	if err != nil {
		return err
	}

	rows, err := db.Query(ctx, stmt, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

// Get returns exactly one record or throws ErrNotFound
func Get(ctx context.Context, db IDB, dest interface{}, query Sqlizer) error {
	stmt, args, err := query.ToSql()
	if err != nil {
		return err
	}

	rows, err := db.Query(ctx, stmt, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, rows)
}

// Exec will execute SQL statement
func Exec(ctx context.Context, db IDB, query Sqlizer) error {
	_, err := ExecCmd(ctx, db, query)
	return err
}

// Exec will execute SQL statement
func ExecCmd(ctx context.Context, db IDB, query Sqlizer) (CommandTag, error) {
	stmt, args, err := query.ToSql()
	if err != nil {
		return CommandTag{}, fmt.Errorf("SQL builder failed: %w", err)
	}

	return db.Exec(ctx, stmt, args...)
}

// QueryRow returns exactly one Row or throws ErrNotFound
func QueryRow(ctx context.Context, db IDB, query Sqlizer) Row {
	stmt, args, err := query.ToSql()
	if err != nil {
		return NewErrRow(err)
	}

	return db.QueryRow(ctx, stmt, args...)
}
