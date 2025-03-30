package tx

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type contextKey string

const (
	TransactionContextKey contextKey = "postgresql_transaction"
)

type Manager struct {
	pool Pool
}

func NewManager(pool *pgxpool.Pool) *Manager {
	return &Manager{
		pool: pool,
	}
}

func (t *Manager) createCtx(ctx context.Context, opts pgx.TxOptions) (context.Context, pgx.Tx, error) {
	/*
		Обработка случая открытия транзакции внутри другой транзакции
		Если в контексте уже есть транзакция, то используем ее
	*/
	txFromContext := ctx.Value(TransactionContextKey)
	if txFromContext != nil {
		tx, ok := txFromContext.(pgx.Tx)
		if !ok {
			return ctx, nil, fmt.Errorf("cast %s to pgx.Tx", TransactionContextKey)
		}

		return ctx, tx, nil
	}

	tx, err := t.pool.BeginTx(ctx, opts)
	if err != nil {
		return ctx, nil, fmt.Errorf("begin transaction: %w", err)
	}

	return context.WithValue(ctx, TransactionContextKey, tx), tx, nil
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

func (t *Manager) WithTx(ctx context.Context, fn func(ctx context.Context) error, opts pgx.TxOptions) error {
	isNested := ctx.Value(TransactionContextKey) != nil

	ctx, tx, err := t.createCtx(ctx, opts)
	if err != nil {
		return fmt.Errorf("create ctx: %w", err)
	}

	if isNested {
		if err := fn(ctx); err != nil {
			return fmt.Errorf("call fn: %w", err)
		}
	} else {
		if err := handleTransaction(ctx, tx, fn(ctx)); err != nil {
			return fmt.Errorf("handle transaction: %w", err)
		}
	}

	return nil
}
