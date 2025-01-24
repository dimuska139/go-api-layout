package shrink

import (
	"context"

	"github.com/dimuska139/urlshortener/internal/model"
)

//go:generate mockgen -source=dependency.go -destination=./dependency_mock.go -package=shrink

type (
	StatisticsService interface {
		SaveRedirectEvent(ctx context.Context, code string, userAgent string) error
	}

	ShrinkService interface {
		CreateShortCode(ctx context.Context, longUrl string) (model.Link, error)
		GetLongUrlByCode(ctx context.Context, shortCode string) (string, error)
	}
)
