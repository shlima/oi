package db

import (
	"context"
	"fmt"
	"strings"
)

func TruncateExcept(ctx context.Context, db IDB, except []string) error {
	query := Query().Select("table_name").
		From("information_schema.tables").
		Where("table_schema = 'public'").
		Where(NotEq{"table_name": except}).
		Where("table_type = 'BASE TABLE'")

	tables := make([]string, 0)
	err := Select(ctx, db, &tables, query)
	if err != nil {
		return fmt.Errorf("failed to select tables: %w", err)
	}

	truncation := Expr(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", strings.Join(tables, ",")))
	return Exec(ctx, db, truncation)
}
