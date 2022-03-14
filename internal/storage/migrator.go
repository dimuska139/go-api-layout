package storage

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/constants"
	"github.com/dimuska139/urlshortener/internal/logging"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*
var embedMigrations embed.FS

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

func (migrator *Migrator) Up() error {
	db, err := sql.Open("postgres", migrator.cfg.Db.Dsn)
	if err != nil {
		return fmt.Errorf("open database connection error: %v", err)
	}
	defer db.Close()

	goose.SetTableName(constants.MigrationsTableName)
	goose.SetBaseFS(embedMigrations)
	goose.SetLogger(migrator.logger)

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("migrate error: %w", err)
	}

	return nil
}
