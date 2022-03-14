package repository

import (
	"context"
	"github.com/dimuska139/urlshortener/internal/models"
)

type LinkRepository interface {
	Create(ctx context.Context, longUrl string) (models.Link, error)
	GetLongUrlByCode(ctx context.Context, shortCode string) (string, error)
	SetShortcode(ctx context.Context, id int, code string) error
}
