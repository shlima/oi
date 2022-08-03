package db

import "github.com/Masterminds/squirrel"

type (
	// Sqlizer is the interface that wraps the ToSql method
	Sqlizer  = squirrel.Sqlizer
	Eq       = squirrel.Eq
	NotEq    = squirrel.NotEq
	Like     = squirrel.Like
	ILike    = squirrel.ILike
	NotLike  = squirrel.NotLike
	NotILike = squirrel.NotILike
	And      = squirrel.And
	Or       = squirrel.Or
)

// Query вернет squirrel SQL Builder объект
func Query() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
}

// Expr builds an expression from a SQL fragment and arguments
// Передавать аргументы в запрос необходимо через долларовую нотацию,
// например: Expr("SELECT $1", 55)
func Expr(sql string, args ...interface{}) Sqlizer {
	return squirrel.Expr(sql, args...)
}
