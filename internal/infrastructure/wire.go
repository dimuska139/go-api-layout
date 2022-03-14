//+build wireinject

package infrastructure

import (
	"github.com/dimuska139/urlshortener/internal/api"
	"github.com/dimuska139/urlshortener/internal/api/middleware"
	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/handlers"
	"github.com/dimuska139/urlshortener/internal/logging"
	"github.com/dimuska139/urlshortener/internal/services"
	"github.com/dimuska139/urlshortener/internal/storage"
	"github.com/google/wire"
)

func InitConfig(configPath string) (*config.Config, error) {
	panic(wire.Build(config.NewConfig))
}

func InitLogger(config *config.Config) logging.Loggerer {
	wire.Build(logging.NewLogger)
	return &logging.Logger{}
}

func InitRestAPI(config *config.Config, logger logging.Loggerer) (*api.RestAPI, error) {
	wire.Build(
		storage.NewPostgresPool,
		storage.NewDatabase,
		services.NewShrinkService,
		wire.Bind(new(services.ShrinkServiceInterface), new(*services.ShrinkService)),
		services.NewStatisticsService,
		wire.Bind(new(services.StatisticsServiceInterface), new(*services.StatisticsService)),
		handlers.NewResponseMapper,
		handlers.NewShrinkHandler,
		middleware.NewMiddlewareFactory,
		api.NewRestAPI,
	)
	return &api.RestAPI{}, nil
}

func InitMigrator(config *config.Config, logger logging.Loggerer) (*storage.Migrator, error) {
	wire.Build(
		storage.NewMigrator,
	)
	return &storage.Migrator{}, nil
}