package services

import (
	"database/sql"
	"fmt"
	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/constants"
	"github.com/dimuska139/urlshortener/internal/logging"
	"github.com/dimuska139/urlshortener/internal/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type Migrator struct {
	cfg    *config.Config
	logger logging.Loggerer
}

func NewMigrator(conf *config.Config, logger logging.Loggerer) *Migrator {
	return &Migrator{
		cfg:    conf,
		logger: logger,
	}
}

func (m *Migrator) Up() error {
	db, err := sql.Open("postgres", m.cfg.Db.Dsn)
	if err != nil {
		return fmt.Errorf("open database connection error: %v", err)
	}
	defer db.Close()

	goose.SetTableName(constants.MigrationsTableName)
	goose.SetBaseFS(migrations.Migrations)
	goose.SetLogger(m.logger)

	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("migrate error: %w", err)
	}

	return nil
}
