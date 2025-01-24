package statistics

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"

	db "github.com/dimuska139/urlshortener/internal/storage/postgresql"
	"github.com/dimuska139/urlshortener/pkg/postgresql"
)

type Repository struct {
	pgPool *postgresql.PostgresPool
	qb     sq.StatementBuilderType
}

func NewRepository(pgPool *postgresql.PostgresPool) *Repository {
	return &Repository{
		pgPool: pgPool,
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *Repository) SaveRedirectEvent(ctx context.Context, code string, userAgent string) error {
	sql, args, err := r.qb.Insert("redirect_statistics").
		Columns("link_id", "user_agent", "created_at").
		Values(sq.Expr("(SELECT id FROM link WHERE code=?)", code), userAgent, "now()").
		ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	if _, err := r.pgPool.Exec(ctx, sql, args...); err != nil {
		return db.NewErrQueryExecutionFailed(sql, args, err)
	}

	return nil
}
