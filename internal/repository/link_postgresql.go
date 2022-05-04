package repository

//go:generate mockgen -source=mapper.go -destination=./mapper_mock.go -package=handlers

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/dimuska139/urlshortener/internal/logging"
	"github.com/dimuska139/urlshortener/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4"
)

type LinkPostgresqlRepository struct {
	logger  logging.Loggerer
	querier pgxtype.Querier
	qb      sq.StatementBuilderType
}

func NewLinkPostgresqlRepository(logger logging.Loggerer, querier pgxtype.Querier) *LinkPostgresqlRepository {
	return &LinkPostgresqlRepository{
		logger:  logger,
		querier: querier,
		qb:      sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

// Create сохраняет ссылку в базу данных и возвращает структуру ссылки
func (r *LinkPostgresqlRepository) Create(ctx context.Context, longUrl string) (models.Link, error) {
	sql, args, err := r.qb.Insert("links").
		Columns("full_url", "created_at").
		Suffix("RETURNING id, created_at").
		Values(longUrl, "now()").
		ToSql()
	if err != nil {
		return models.Link{}, fmt.Errorf("unable to build query: %w", err)
	}

	link := models.Link{}
	if err := r.querier.QueryRow(ctx, sql, args...).Scan(&link.ID, &link.CreatedAt); err != nil {
		return models.Link{}, fmt.Errorf("unable to insert record to links table: %w", err)
	}

	link.LongURL = longUrl

	return link, nil
}

// LinkPostgresqlRepository записывает код для ссылки
func (r *LinkPostgresqlRepository) SetShortcode(ctx context.Context, id int, shortcode string) error {
	sql, args, err := r.qb.Update("links").
		Set("code", shortcode).
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("unable to build query: %w", err)
	}
	_, err = r.querier.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("unable to update record (#%d) in users table: %w", id, err)
	}
	return err
}

// GetLongUrlByCode возвращает полный url по коду
func (r *LinkPostgresqlRepository) GetLongUrlByCode(ctx context.Context, shortCode string) (string, error) {
	sql, args, err := r.qb.Select("full_url").
		From("links").
		Where("code = ?", shortCode).
		Limit(1).
		ToSql()
	if err != nil {
		return "", fmt.Errorf("unable to build query: %w", err)
	}
	var fullUrl string
	if err := r.querier.QueryRow(ctx, sql, args...).Scan(&fullUrl); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", nil
		}
		return "", fmt.Errorf("unable to get full url by code (%s): %w", shortCode, err)
	}

	return fullUrl, nil
}
