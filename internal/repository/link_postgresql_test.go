package repository

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/bxcodec/faker/v3"
	"github.com/dimuska139/urlshortener/internal/logging"
	"github.com/dimuska139/urlshortener/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestLinkPostgresqlRepository_Create(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	link := models.Link{
		ID:      10,
		LongURL: faker.URL(),
	}

	expectedSql := "INSERT INTO links (full_url,created_at) VALUES ($1,$2) RETURNING id, created_at"

	type fields struct {
		querier pgxtype.Querier
		qb      squirrel.StatementBuilderType
	}
	type args struct {
		ctx     context.Context
		longUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(*MockQuerier)
		want    models.Link
		wantErr bool
	}{
		{
			name: "creation",
			fields: fields{
				qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
			},
			args: args{
				ctx:     ctx,
				longUrl: link.LongURL,
			},
			prepare: func(querierMock *MockQuerier) {
				rowMock := NewMockRow(ctrl)
				rowMock.EXPECT().
					Scan(gomock.Any(), gomock.Any()).
					Return(nil).
					Times(1)

				querierMock.EXPECT().
					QueryRow(ctx, expectedSql, link.LongURL, "now()").
					Return(rowMock).
					Times(1)
			},
			want: link,
		},
		{
			name: "scan error",
			fields: fields{
				qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
			},
			args: args{
				ctx:     ctx,
				longUrl: link.LongURL,
			},
			prepare: func(querierMock *MockQuerier) {
				rowMock := NewMockRow(ctrl)
				rowMock.EXPECT().
					Scan(gomock.Any(), gomock.Any()).
					Return(errors.New("error")).
					Times(1)

				querierMock.EXPECT().
					QueryRow(ctx, expectedSql, link.LongURL, "now()").
					Return(rowMock).
					Times(1)
			},
			want:    models.Link{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			querierMock := NewMockQuerier(ctrl)

			r := &LinkPostgresqlRepository{
				querier: querierMock,
				qb:      tt.fields.qb,
			}
			if tt.prepare != nil {
				tt.prepare(querierMock)
			}
			got, err := r.Create(tt.args.ctx, tt.args.longUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want.LongURL, got.LongURL)
		})
	}
}

func TestLinkPostgresqlRepository_GetLongUrlByCode(t *testing.T) {
	type fields struct {
		querier pgxtype.Querier
		qb      squirrel.StatementBuilderType
	}
	type args struct {
		ctx       context.Context
		shortCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &LinkPostgresqlRepository{
				querier: tt.fields.querier,
				qb:      tt.fields.qb,
			}
			got, err := r.GetLongUrlByCode(tt.args.ctx, tt.args.shortCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLongUrlByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLongUrlByCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkPostgresqlRepository_SetShortcode(t *testing.T) {
	type fields struct {
		querier pgxtype.Querier
		qb      squirrel.StatementBuilderType
	}
	type args struct {
		ctx       context.Context
		id        int
		shortcode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &LinkPostgresqlRepository{
				querier: tt.fields.querier,
				qb:      tt.fields.qb,
			}
			if err := r.SetShortcode(tt.args.ctx, tt.args.id, tt.args.shortcode); (err != nil) != tt.wantErr {
				t.Errorf("SetShortcode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewLinkPostgresqlRepository(t *testing.T) {
	type args struct {
		logger  logging.Loggerer
		querier pgxtype.Querier
	}
	tests := []struct {
		name string
		args args
		want *LinkPostgresqlRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkPostgresqlRepository(tt.args.logger, tt.args.querier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkPostgresqlRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
