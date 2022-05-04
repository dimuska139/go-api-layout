//+build wireinject

package wire

import (
	"github.com/dimuska139/urlshortener/internal/api"
	"github.com/dimuska139/urlshortener/internal/api/middleware"
	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/handlers"
	"github.com/dimuska139/urlshortener/internal/http"
	"github.com/dimuska139/urlshortener/internal/logging"
	"github.com/dimuska139/urlshortener/internal/postgresql"
	"github.com/dimuska139/urlshortener/internal/services"
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
		http.NewHttpClient,
		postgresql.NewPostgresPool,
		services.NewDatabase,
		services.NewShrinkService,
		wire.Bind(new(handlers.ShrinkServiceInterface), new(*services.ShrinkService)),
		services.NewStatisticsService,
		wire.Bind(new(handlers.StatisticsServiceInterface), new(*services.StatisticsService)),
		handlers.NewResponseMapper,
		handlers.NewShrinkHandler,
		middleware.NewMiddlewareFactory,
		api.NewRestAPI,
	)
	return &api.RestAPI{}, nil
}

func InitMigrator(config *config.Config, logger logging.Loggerer) (*services.Migrator, error) {
	wire.Build(
		services.NewMigrator,
	)
	return &services.Migrator{}, nil
}
