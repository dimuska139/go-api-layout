package storage

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	QueryBuilder sq.StatementBuilderType
	Repositories *RepositoriesFactory
	pool         *pgxpool.Pool
}

func NewDatabase(dbpool *pgxpool.Pool) Database {
	return Database{
		QueryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		pool:         dbpool,
		Repositories: NewRepositoriesFactory(dbpool),
	}
}

func (db *Database) BeginTx(ctx context.Context) (*TransactionManager, error) {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to open transaction: %v", err)
	}
	transaction := NewTransactionManager(tx)
	return &transaction, nil
}
