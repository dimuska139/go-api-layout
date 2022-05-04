package services

import (
	"context"
	"fmt"
	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/logging"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

type TransactionManager struct {
	Repositories *TransactionalRepositoryFactory
	tr           pgx.Tx
}

func NewTransactionManager(config *config.Config, logger logging.Loggerer, tr pgx.Tx) TransactionManager {
	return TransactionManager{
		Repositories: NewTransactionalRepositoryFactory(config, logger, tr),
		tr:           tr,
	}
}

func (t *TransactionManager) Commit(ctx context.Context) error {
	if err := t.tr.Commit(ctx); err != nil {
		return fmt.Errorf("unable to commit transaction: %v", err)
	}

	return nil
}

func (t *TransactionManager) Rollback(ctx context.Context) error {
	if err := t.tr.Rollback(ctx); err != nil {
		return fmt.Errorf("unable to rollback transaction: %v", err)
	}

	return nil
}

type Storage struct {
	config         *config.Config
	logger         logging.Loggerer
	httpClient     *http.Client
	Repositories   *RepositoryFactory
	postgresqlPool *pgxpool.Pool
}

func NewDatabase(config *config.Config, logger logging.Loggerer, postgresqlPool *pgxpool.Pool, httpClient *http.Client) Storage {
	return Storage{
		config:         config,
		logger:         logger,
		postgresqlPool: postgresqlPool,
		httpClient:     httpClient,
		Repositories:   NewRepositoryFactory(config, logger, postgresqlPool, httpClient),
	}
}

func (storage *Storage) BeginTx(ctx context.Context) (*TransactionManager, error) {
	tx, err := storage.postgresqlPool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to open transaction: %v", err)
	}
	transaction := NewTransactionManager(storage.config, storage.logger, tx)
	return &transaction, nil
}
