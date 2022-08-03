package db

// ErrRow implements pgx.Row
type ErrRow struct {
	err error
}

func NewErrRow(err error) ErrRow {
	return ErrRow{err: err}
}

// Scan implements sql.Row
func (e ErrRow) Scan(dest ...interface{}) error {
	return e.err
}
