package postgresql

import (
	"context"
	"fmt"
	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/constants"
	"github.com/dimuska139/urlshortener/internal/logging"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresPool(conf *config.Config, logger logging.Loggerer) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(conf.Db.Dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database dsn: %v", err)
	}
	poolConfig.HealthCheckPeriod = constants.DbPoolHealthcheckPeriod
	poolConfig.MaxConnIdleTime = constants.DbMaxConnIdleTime
	poolConfig.MaxConnLifetime = constants.DbMaxConnLifetime
	poolConfig.MaxConns = constants.DbPoolIdleConns
	poolConfig.MinConns = constants.DbPoolMaxConns

	if conf.Loglevel == constants.LogLevelDebug {
		poolConfig.ConnConfig.Logger = logger.NewPgxLogger()
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("postgresql connection failed: %v", err)
	}

	return pool, nil
}
