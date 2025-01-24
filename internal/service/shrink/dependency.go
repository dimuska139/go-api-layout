package shrink

//go:generate mockgen -source=dependency.go -destination=./dependency_mock.go -package=shrink

import (
	"context"
	"github.com/jackc/pgx/v5"

	"github.com/dimuska139/urlshortener/internal/model"
)

type (
	TransactionManager interface {
		WithTx(ctx context.Context, fn func(ctx context.Context) error, opts pgx.TxOptions) error
	}

	LinkRepository interface {
		Create(ctx context.Context, longUrl string) (model.Link, error)
		SetShortcode(ctx context.Context, id int, shortcode string) error
		GetLongUrlByCode(ctx context.Context, shortCode string) (string, error)
	}
)
