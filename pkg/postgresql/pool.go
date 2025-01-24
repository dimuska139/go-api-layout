package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPool struct {
	pool *pgxpool.Pool
}

func newQueryRowErr(err error) QueryRowErr {
	return QueryRowErr{
		err: err,
	}
}

type QueryRowErr struct {
	err error
}

func (e QueryRowErr) Scan(dest ...any) error {
	return nil
}

type contextKey string

const (
	transactionContextKey contextKey = "postgresql_transaction"
)

func NewPostgresPool(pool *pgxpool.Pool) (*PostgresPool, error) {
	return &PostgresPool{
		pool: pool,
	}, nil
}

func (p *PostgresPool) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	if ctx.Value(transactionContextKey) == nil {
		cmd, err := p.pool.Exec(ctx, sql, args...)
		if err != nil {
			return pgconn.CommandTag{}, fmt.Errorf("execute: %w", err)
		}

		return cmd, nil
	}

	tx, ok := ctx.Value(transactionContextKey).(pgx.Tx)
	if !ok {
		return pgconn.CommandTag{}, fmt.Errorf("cast %s to pgx.Tx", transactionContextKey)
	}

	cmd, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return pgconn.CommandTag{}, fmt.Errorf("execute: %w", err)
	}

	return cmd, nil
}

func (p *PostgresPool) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if ctx.Value(transactionContextKey) == nil {
		return p.pool.QueryRow(ctx, sql, args...)
	}

	tx, ok := ctx.Value(transactionContextKey).(pgx.Tx)
	if !ok {
		return newQueryRowErr(fmt.Errorf("cast %s to pgx.Tx", transactionContextKey))
	}

	return tx.QueryRow(ctx, sql, args...)
}

func (p *PostgresPool) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if ctx.Value(transactionContextKey) == nil {
		rows, err := p.pool.Query(ctx, sql, args...)
		if err != nil {
			return nil, fmt.Errorf("query: %w", err)
		}

		return rows, nil
	}

	tx, ok := ctx.Value(transactionContextKey).(pgx.Tx)
	if !ok {
		return nil, fmt.Errorf("cast %s to pgx.Tx", transactionContextKey)
	}

	rows, err := tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return rows, nil
}
