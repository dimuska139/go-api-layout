// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

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
)

// Injectors from wire.go:

func InitConfig(configPath string) (*config.Config, error) {
	configConfig, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}
	return configConfig, nil
}

func InitLogger(config2 *config.Config) logging.Loggerer {
	loggerer := logging.NewLogger(config2)
	return loggerer
}

func InitRestAPI(config2 *config.Config, logger logging.Loggerer) (*api.RestAPI, error) {
	middlewareFactory, err := middleware.NewMiddlewareFactory(logger)
	if err != nil {
		return nil, err
	}
	pool, err := postgresql.NewPostgresPool(config2, logger)
	if err != nil {
		return nil, err
	}
	client := http.NewHttpClient()
	storage := services.NewDatabase(config2, logger, pool, client)
	shrinkService := services.NewShrinkService(config2, storage)
	statisticsService := services.NewStatisticsService(config2, storage)
	mapper := handlers.NewResponseMapper(config2)
	shrinkHandler := handlers.NewShrinkHandler(logger, shrinkService, statisticsService, mapper)
	restAPI, err := api.NewRestAPI(config2, logger, middlewareFactory, shrinkHandler)
	if err != nil {
		return nil, err
	}
	return restAPI, nil
}

func InitMigrator(config2 *config.Config, logger logging.Loggerer) (*services.Migrator, error) {
	migrator := services.NewMigrator(config2, logger)
	return migrator, nil
}
