package migrator

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/dimuska139/urlshortener/pkg/logging"
)

const (
	defaultMigrationsTableName = "db_version"
)

type Migrator struct {
	connection *sql.DB
}

func NewMigrator(config Config) (*Migrator, func(), error) {
	if config.MigrationsTableName == "" {
		config.MigrationsTableName = defaultMigrationsTableName
	}

	connection, err := sql.Open("postgres", config.DSN)
	if err != nil {
		return nil, nil, fmt.Errorf("open database connection: %w", err)
	}

	migrationsTableName := defaultMigrationsTableName
	if config.MigrationsTableName != "" {
		migrationsTableName = config.MigrationsTableName
	}

	goose.SetTableName(migrationsTableName)
	goose.SetBaseFS(migrations)
	goose.SetLogger(NewLoggerAdaptor())

	return &Migrator{
			connection: connection,
		}, func() {
			if err := connection.Close(); err != nil {
				logging.Error(context.Background(), "Can't close database connection",
					"err", err.Error())
			}
		}, nil
}

func (m *Migrator) Up() error {
	if err := goose.Up(m.connection, ".", goose.WithAllowMissing()); err != nil {
		return fmt.Errorf("apply migrations: %w", err)
	}

	return nil
}
