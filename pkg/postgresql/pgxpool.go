package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"

	"github.com/dimuska139/urlshortener/pkg/logging"
)

const (
	DbPoolIdleConns         = 2
	DbPoolMaxConns          = 10
	DbPoolHealthcheckPeriod = 3 * time.Second
	DbMaxConnIdleTime       = 10 * time.Second
	DbMaxConnLifetime       = 1 * time.Minute
)

func NewPgxPool(conf Config) (*pgxpool.Pool, func(), error) {
	poolConfig, err := pgxpool.ParseConfig(conf.DSN)
	if err != nil {
		return nil, nil, fmt.Errorf("parse database dsn: %w", err)
	}

	poolConfig.HealthCheckPeriod = DbPoolHealthcheckPeriod
	poolConfig.MaxConnIdleTime = DbMaxConnIdleTime
	poolConfig.MaxConnLifetime = DbMaxConnLifetime
	poolConfig.MaxConns = DbPoolMaxConns
	poolConfig.MinConns = DbPoolIdleConns

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("connect to PostgreSQL: %w", err)
	}

	return pool, func() {
		logging.Info(context.Background(), "Closing PostgreSQL pool...")
		pool.Close()
	}, nil
}
