package constants

import "time"

const (
	DbPoolIdleConns         = 1
	DbPoolMaxConns          = 2
	DbPoolHealthcheckPeriod = 5 * time.Second
	DbMaxConnIdleTime       = 10 * time.Second
	DbMaxConnLifetime       = 50 * time.Second
	DbTimeout               = 1 * time.Second
	MigrationsTableName     = "migrations"
)
