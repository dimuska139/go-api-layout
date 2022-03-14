package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type TransactionManager struct {
	Repositories *RepositoriesFactory
	tr           pgx.Tx
}

func NewTransactionManager(tr pgx.Tx) TransactionManager {
	return TransactionManager{
		Repositories: NewRepositoriesFactory(tr),
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
