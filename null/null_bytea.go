package null

import (
	"database/sql/driver"
	"github.com/jackc/pgtype"
)

type Bytea struct {
	Bytes []byte
	Valid bool
}

// Scan implements the Scanner interface.
func (n *Bytea) Scan(value any) error {
	oid := new(pgtype.Bytea)
	err := oid.Scan(value)
	if err != nil {
		return err
	}

	n.Bytes = oid.Bytes
	n.Valid = oid.Status == pgtype.Present

	return nil
}

// Value implements the driver Valuer interface.
func (n Bytea) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.Bytes, nil
}

func NewValidBytea(input []byte) Bytea {
	return Bytea{Bytes: input, Valid: true}
}

func NewAutoBytea(input []byte) Bytea {
	return Bytea{Bytes: input, Valid: len(input) != 0}
}
