package storage

import (
	sq "github.com/Masterminds/squirrel"
)

func NewQueryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
