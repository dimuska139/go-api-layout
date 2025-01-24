package statistics

import "context"

type (
	StatisticsRepository interface {
		SaveRedirectEvent(ctx context.Context, code string, userAgent string) error
	}
)
