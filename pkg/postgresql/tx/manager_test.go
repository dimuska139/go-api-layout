package tx

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestManager_WithTx(t *testing.T) {
	type args struct {
		ctx  context.Context
		fn   func(ctx context.Context) error
		opts pgx.TxOptions
	}
	tests := []struct {
		name          string
		args          args
		getMockedPool func(ctrl *gomock.Controller) *MockPool
		wantErr       bool
	}{
		{
			name: "nested transaction",
			args: args{
				ctx: context.WithValue(context.Background(), TransactionContextKey, &pgxpool.Tx{}),
				fn: func(ctx context.Context) error {
					return nil
				},
				opts: pgx.TxOptions{},
			},
		}, {
			name: "can't create transaction",
			args: args{
				ctx: context.Background(),
				fn: func(ctx context.Context) error {
					return nil
				},
				opts: pgx.TxOptions{
					IsoLevel: pgx.Serializable,
				},
			},
			getMockedPool: func(ctrl *gomock.Controller) *MockPool {
				pool := NewMockPool(ctrl)
				pool.
					EXPECT().
					BeginTx(
						context.Background(),
						pgx.TxOptions{
							IsoLevel: pgx.Serializable,
						},
					).Return(nil, pgx.ErrTxClosed)

				return pool
			},
			wantErr: true,
		}, {
			name: "with rollback",
			args: args{
				ctx: context.Background(),
				fn: func(ctx context.Context) error {
					return assert.AnError
				},
				opts: pgx.TxOptions{
					IsoLevel: pgx.Serializable,
				},
			},
			getMockedPool: func(ctrl *gomock.Controller) *MockPool {
				txMock := NewMockTx(ctrl)
				txMock.
					EXPECT().
					Rollback(gomock.Any()).
					Return(nil)

				pool := NewMockPool(ctrl)
				pool.
					EXPECT().
					BeginTx(
						context.Background(),
						pgx.TxOptions{
							IsoLevel: pgx.Serializable,
						},
					).Return(txMock, nil)

				return pool
			},
			wantErr: true,
		}, {
			name: "success",
			args: args{
				ctx: context.Background(),
				fn: func(ctx context.Context) error {
					return nil
				},
				opts: pgx.TxOptions{
					IsoLevel: pgx.Serializable,
				},
			},
			getMockedPool: func(ctrl *gomock.Controller) *MockPool {
				txMock := NewMockTx(ctrl)
				txMock.
					EXPECT().
					Commit(gomock.Any()).
					Return(nil)

				pool := NewMockPool(ctrl)
				pool.
					EXPECT().
					BeginTx(
						context.Background(),
						pgx.TxOptions{
							IsoLevel: pgx.Serializable,
						},
					).Return(txMock, nil)

				return pool
			},
		}, {
			name: "rollback failed",
			args: args{
				ctx: context.Background(),
				fn: func(ctx context.Context) error {
					return assert.AnError
				},
				opts: pgx.TxOptions{
					IsoLevel: pgx.Serializable,
				},
			},
			getMockedPool: func(ctrl *gomock.Controller) *MockPool {
				txMock := NewMockTx(ctrl)
				txMock.
					EXPECT().
					Rollback(gomock.Any()).
					Return(assert.AnError)

				pool := NewMockPool(ctrl)
				pool.
					EXPECT().
					BeginTx(
						context.Background(),
						pgx.TxOptions{
							IsoLevel: pgx.Serializable,
						},
					).Return(txMock, nil)

				return pool
			},
			wantErr: true,
		}, {
			name: "commit failed",
			args: args{
				ctx: context.Background(),
				fn: func(ctx context.Context) error {
					return nil
				},
				opts: pgx.TxOptions{
					IsoLevel: pgx.Serializable,
				},
			},
			getMockedPool: func(ctrl *gomock.Controller) *MockPool {
				txMock := NewMockTx(ctrl)
				txMock.
					EXPECT().
					Commit(gomock.Any()).
					Return(assert.AnError)

				pool := NewMockPool(ctrl)
				pool.
					EXPECT().
					BeginTx(
						context.Background(),
						pgx.TxOptions{
							IsoLevel: pgx.Serializable,
						},
					).Return(txMock, nil)

				return pool
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txManager := &Manager{}

			ctrl := gomock.NewController(t)
			if tt.getMockedPool != nil {
				txManager.pool = tt.getMockedPool(ctrl)
			}

			err := txManager.WithTx(tt.args.ctx, tt.args.fn, tt.args.opts)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
