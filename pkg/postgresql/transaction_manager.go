package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionManager struct {
	pool *pgxpool.Pool
}

func NewTransactionManager(pool *pgxpool.Pool) *TransactionManager {
	return &TransactionManager{
		pool: pool,
	}
}

func (t *TransactionManager) createCtx(ctx context.Context, opts pgx.TxOptions) (context.Context, pgx.Tx, error) {
	if ctx.Value(transactionContextKey) != nil {
		tx, ok := ctx.Value(transactionContextKey).(pgx.Tx)
		if !ok {
			return ctx, nil, fmt.Errorf("cast %s to pgx.Tx", transactionContextKey)
		}

		return ctx, tx, nil
	}

	tx, err := t.pool.BeginTx(ctx, opts)
	if err != nil {
		return ctx, nil, fmt.Errorf("begin transaction: %w", err)
	}

	return context.WithValue(ctx, transactionContextKey, tx), tx, nil
}

func handleTransaction(ctx context.Context, tx pgx.Tx, functionError error) error {
	if functionError != nil {
		if err := tx.Rollback(ctx); err != nil {
			return fmt.Errorf("rollback transaction: %w", err)
		}

		return fmt.Errorf("function: %w", functionError)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (t *TransactionManager) WithTx(ctx context.Context, fn func(ctx context.Context) error, opts pgx.TxOptions) error {
	ctx, tx, err := t.createCtx(ctx, opts)
	if err != nil {
		return fmt.Errorf("create ctx: %w", err)
	}

	if err := handleTransaction(ctx, tx, fn(ctx)); err != nil {
		return fmt.Errorf("handle transaction: %w", err)
	}

	return nil
}
