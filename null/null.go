package null

import (
	"database/sql"
	"time"
)

type (
	String = sql.NullString
	Int64  = sql.NullInt64
	Int32  = sql.NullInt32
	Time   = sql.NullTime
	Date   = sql.NullTime
	Bool   = sql.NullBool
)

func NewValidString(input string) String {
	return String{String: input, Valid: true}
}

func NewValidInt64(input int64) Int64 {
	return Int64{Int64: input, Valid: true}
}

func NewAutoInt64(input int64) Int64 {
	return Int64{Int64: input, Valid: input != 0}
}

func NewValidInt32(input int32) Int32 {
	return Int32{Int32: input, Valid: true}
}

func NewAutoInt32(input int32) Int32 {
	return Int32{Int32: input, Valid: input != 0}
}

func NewAutoTime(input time.Time) Time {
	return Time{Time: input, Valid: !input.IsZero()}
}

func NewAutoString(input string) String {
	return String{String: input, Valid: input != ""}
}

func NewValidBool(input bool) Bool {
	return Bool{Bool: input, Valid: true}
}

func NewAutoDate(input time.Time) Date {
	return Date{Time: ToDate(input), Valid: !input.IsZero()}
}

func NewValidDate(input time.Time) Date {
	return Date{Time: ToDate(input), Valid: true}
}

func NewValidTime(input time.Time) Time {
	return Time{Time: input, Valid: true}
}

func ToDate(input time.Time) time.Time {
	return time.Date(input.Year(), input.Month(), input.Day(), 0, 0, 0, 0, input.Location())
}
