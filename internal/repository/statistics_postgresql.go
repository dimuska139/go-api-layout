package repository

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/dimuska139/urlshortener/internal/logging"
	"github.com/jackc/pgtype/pgxtype"
)

type StatisticsPostgresqlRepository struct {
	logger  logging.Loggerer
	querier pgxtype.Querier
	qb      sq.StatementBuilderType
}

func NewStatisticsPostgresqlRepository(logger logging.Loggerer, querier pgxtype.Querier) *StatisticsPostgresqlRepository {
	return &StatisticsPostgresqlRepository{
		logger:  logger,
		querier: querier,
		qb:      sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *StatisticsPostgresqlRepository) SaveRedirectEvent(ctx context.Context, code string, userAgent string) error {
	sql, args, err := r.qb.Insert("redirects_statistics").
		Columns("(SELECT id FROM links WHERE code=?)", "user_agent", "created_at").
		Values(code, userAgent, "now()").
		ToSql()
	if err != nil {
		return fmt.Errorf("unable to build query: %w", err)
	}

	if _, err := r.querier.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("unable to insert record to statistics table: %w", err)
	}

	return nil
}
