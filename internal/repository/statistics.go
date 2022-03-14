package repository

import (
	"context"
)

var (
	//ErrConflict = errors.New("user with same email already exists")
)

type StatisticsRepository interface {
	SaveRedirectEvent(ctx context.Context, code string, userAgent string) error
}
