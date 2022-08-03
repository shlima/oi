package db

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
)

func WithTx(ctx context.Context, db IDB, fn func(tx Tx) error) error {
	return db.BeginFunc(ctx, fn)
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
	stmt, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("SQL builder failed: %w", err)
	}

	_, err = db.Exec(ctx, stmt, args...)
	return err
}

// QueryRow returns exactly one Row or throws ErrNotFound
func QueryRow(ctx context.Context, db IDB, query Sqlizer) Row {
	stmt, args, err := query.ToSql()
	if err != nil {
		return NewErrRow(err)
	}

	return db.QueryRow(ctx, stmt, args...)
}
