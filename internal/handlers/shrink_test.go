package handlers

import (
	"errors"
	"github.com/bxcodec/faker/v3"
	"github.com/dimuska139/urlshortener/internal/gen/models"
	"github.com/dimuska139/urlshortener/internal/gen/restapi/operations"
	"github.com/dimuska139/urlshortener/internal/logging"
	models2 "github.com/dimuska139/urlshortener/internal/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewShrinkHandler(t *testing.T) {
	type args struct {
		logger            logging.Loggerer
		shrinkService     ShrinkServiceInterface
		statisticsService StatisticsServiceInterface
		responseMapper    Mapper
	}
	ctrl := gomock.NewController(t)
	mapper := NewMockMapper(ctrl)
	logger := logging.NewMockLoggerer(ctrl)
	tests := []struct {
		name string
		args args
		want *ShrinkHandler
	}{
		{
			name: "test create new shrink handler",
			args: args{
				logger:            logger,
				shrinkService:     nil,
				statisticsService: nil,
				responseMapper:    mapper,
			},
			want: &ShrinkHandler{
				logger:         logger,
				responseMapper: mapper,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewShrinkHandler(tt.args.logger, tt.args.shrinkService, tt.args.statisticsService, tt.args.responseMapper)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestShrinkHandler_Redirect(t *testing.T) {
	type fields struct {
		logger            *logging.Logger
		shrinkUsecase     ShrinkServiceInterface
		statisticsUsecase StatisticsServiceInterface
	}
	type args struct {
		params operations.GetShortCodeParams
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   middleware.Responder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/*shrinkHandler := &handlers.ShrinkHandler{
				logger:            tt.fields.logger,
				shrinkUsecase:     tt.fields.shrinkUsecase,
				statisticsUsecase: tt.fields.statisticsUsecase,
			}
			if got := shrinkHandler.Redirect(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Redirect() = %v, want %v", got, tt.want)
			}*/
		})
	}
}

func TestShrinkHandler_Shrink(t *testing.T) {
	type fields struct {
		logger            logging.Loggerer
		statisticsUsecase StatisticsServiceInterface
	}

	ctrl := gomock.NewController(t)

	type args struct {
		params operations.PostShrinkParams
	}

	sourceUrl := faker.URL()
	tests := []struct {
		name    string
		fields  fields
		prepare func(shrinkMock *MockShrinkServiceInterface, mapperMock *MockMapper, logger *logging.MockLoggerer)
		args    args
		want    middleware.Responder
	}{
		{
			name: "",
			fields: fields{
				logger:            logging.NewLogger(nil),
				statisticsUsecase: nil,
			},
			prepare: func(shrinkMock *MockShrinkServiceInterface, mapperMock *MockMapper, logger *logging.MockLoggerer) {
				err := errors.New("error")
				shrinkMock.EXPECT().
					CreateShortCode(gomock.Any(), gomock.Any()).
					Return(models2.Link{}, err).
					Times(1)
				logger.EXPECT().
					Error("unable to create short code", err, nil).
					Times(1)
			},
			args: args{
				params: operations.PostShrinkParams{
					HTTPRequest: &http.Request{},
					Body: &models.SourceLink{
						LongURL: &sourceUrl,
					},
				},
			},
			want: operations.NewPostShrinkInternalServerError(),
		},
		{
			name: "",
			fields: fields{
				logger:            logging.NewLogger(nil),
				statisticsUsecase: nil,
			},
			prepare: func(shrinkMock *MockShrinkServiceInterface, mapperMock *MockMapper, logger *logging.MockLoggerer) {
				link := models2.Link{
					Code:     "1b",
					LongURL:  "https://ya.ru",
					ShortURL: "https://b.ru/1b",
				}
				shrinkMock.EXPECT().
					CreateShortCode(gomock.Any(), gomock.Any()).
					Return(link, nil).
					Times(1)
				mapperMock.
					EXPECT().
					PostShrinkOKBody(link).
					Return(&operations.PostShrinkOKBody{}).
					Times(1)
			},
			args: args{
				params: operations.PostShrinkParams{
					HTTPRequest: &http.Request{},
					Body: &models.SourceLink{
						LongURL: &sourceUrl,
					},
				},
			},
			want: &operations.PostShrinkOK{
				Payload: &operations.PostShrinkOKBody{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shrinkMock := NewMockShrinkServiceInterface(ctrl)
			mapper := NewMockMapper(ctrl)
			logger := logging.NewMockLoggerer(ctrl)

			if tt.prepare != nil {
				tt.prepare(shrinkMock, mapper, logger)
			}
			shrinkHandler := NewShrinkHandler(logger, shrinkMock, tt.fields.statisticsUsecase, mapper)

			got := shrinkHandler.Shrink(tt.args.params)
			assert.Equal(t, tt.want, got)
		})
	}
}
