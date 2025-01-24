package link

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"

	"github.com/dimuska139/urlshortener/internal/model"
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

// Create сохраняет ссылку в базу данных и возвращает структуру ссылки
func (r *Repository) Create(ctx context.Context, longUrl string) (model.Link, error) {
	sql, args, err := r.qb.Insert("link").
		Columns("full_url", "created_at").
		Suffix("RETURNING id, created_at").
		Values(longUrl, "now()").
		ToSql()
	if err != nil {
		return model.Link{}, fmt.Errorf("build query: %w", err)
	}

	var link model.Link
	if err := r.pgPool.QueryRow(ctx, sql, args...).Scan(&link.ID, &link.CreatedAt); err != nil {
		return model.Link{}, db.NewErrQueryExecutionFailed(sql, args, err)
	}

	link.LongURL = longUrl

	return link, nil
}

// SetShortcode записывает код для ссылки
func (r *Repository) SetShortcode(ctx context.Context, id int, shortcode string) error {
	sql, args, err := r.qb.Update("link").
		Set("code", shortcode).
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	_, err = r.pgPool.Exec(ctx, sql, args...)
	if err != nil {
		return db.NewErrQueryExecutionFailed(sql, args, err)
	}

	return nil
}

// GetLongUrlByCode returns full url by short code
func (r *Repository) GetLongUrlByCode(ctx context.Context, shortCode string) (string, error) {
	sql, args, err := r.qb.Select("full_url").
		From("link").
		Where("code = ?", shortCode).
		Limit(1).
		ToSql()
	if err != nil {
		return "", fmt.Errorf("unable to build query: %w", err)
	}

	var fullUrl string

	if err := r.pgPool.QueryRow(ctx, sql, args...).Scan(&fullUrl); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", nil
		}

		return "", db.NewErrQueryExecutionFailed(sql, args, err)
	}

	return fullUrl, nil
}
