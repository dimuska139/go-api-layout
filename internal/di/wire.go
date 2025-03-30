//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/service/application"
	"github.com/dimuska139/urlshortener/internal/service/server/rest"
	"github.com/dimuska139/urlshortener/internal/service/server/rest/handler/shrink"
	"github.com/dimuska139/urlshortener/internal/service/server/rest/middleware"
	"github.com/dimuska139/urlshortener/internal/service/server/rest/middleware/cors"
	shrinkService "github.com/dimuska139/urlshortener/internal/service/shrink"
	statisticsService "github.com/dimuska139/urlshortener/internal/service/statistics"
	"github.com/dimuska139/urlshortener/internal/storage/postgresql/link"
	"github.com/dimuska139/urlshortener/internal/storage/postgresql/migrator"
	"github.com/dimuska139/urlshortener/internal/storage/postgresql/statistics"
	"github.com/dimuska139/urlshortener/pkg/postgresql"
	"github.com/dimuska139/urlshortener/pkg/postgresql/tx"
)

func InitConfig(configPath string, version config.VersionParam) (*config.Config, error) {
	panic(wire.Build(config.NewConfig))
}

func InitMigrator(_ *config.Config) (*migrator.Migrator, func(), error) {
	panic(wire.Build(
		wire.FieldsOf(new(*config.Config), "Migrator"),
		migrator.NewMigrator,
	))
}

func initPgxPool(_ *config.Config) (*pgxpool.Pool, func(), error) {
	panic(wire.Build(
		wire.FieldsOf(new(*config.Config), "PostgreSQL"),
		postgresql.NewPgxPool,
	))
}

func initPostresPool(_ *pgxpool.Pool) (*postgresql.PostgresPool, func(), error) {
	panic(wire.Build(
		postgresql.NewPostgresPool,
	))
}

func initTransactionManager(_ *pgxpool.Pool) (*tx.Manager, func(), error) {
	panic(wire.Build(
		tx.NewManager,
	))
}

func initLinkRepository(_ *postgresql.PostgresPool) (*link.Repository, func(), error) {
	panic(wire.Build(
		link.NewRepository,
	))
}

func initShrinkService(
	_ *config.Config,
	_ *postgresql.PostgresPool,
	_ *tx.Manager) (*shrinkService.ShrinkService, func(), error) {
	panic(wire.Build(
		initLinkRepository,
		wire.Bind(new(shrinkService.LinkRepository), new(*link.Repository)),
		wire.Bind(new(shrinkService.TransactionManager), new(*tx.Manager)),

		wire.FieldsOf(new(*config.Config), "shrink"),
		shrinkService.NewShrinkService,
	))
}

func initStatisticsRepository(_ *postgresql.PostgresPool) (*statistics.Repository, func(), error) {
	panic(wire.Build(
		statistics.NewRepository,
	))
}

func initStatisticsService(_ *postgresql.PostgresPool) (*statisticsService.StatisticsService, func(), error) {
	panic(wire.Build(
		initStatisticsRepository,
		wire.Bind(new(statisticsService.StatisticsRepository), new(*statistics.Repository)),
		statisticsService.NewStatisticsService,
	))
}

func initCorsMiddleware(_ *middleware.Config) (*cors.Middleware, error) {
	panic(wire.Build(
		wire.FieldsOf(new(*middleware.Config), "cors"),
		cors.NewMiddleware,
	))
}

func initMiddlewareFactory(_ *rest.Config) (*middleware.Factory, error) {
	panic(wire.Build(
		wire.FieldsOf(new(*rest.Config), "middleware"),
		initCorsMiddleware,
		wire.Struct(new(middleware.Factory), "*"),
	))
}

func initShrinkHandler(
	_ *config.Config,
	_ *shrinkService.ShrinkService,
	_ *statisticsService.StatisticsService,
) (*shrink.Handler, func(), error) {
	panic(wire.Build(
		wire.Bind(new(shrink.StatisticsService), new(*statisticsService.StatisticsService)),
		wire.Bind(new(shrink.ShrinkService), new(*shrinkService.ShrinkService)),
		shrink.NewHandler,
	))
}

func initRestApiServer(
	_ *config.Config,
	_ *shrinkService.ShrinkService,
	_ *statisticsService.StatisticsService,
) (*rest.Server, func(), error) {
	panic(wire.Build(
		initMiddlewareFactory,
		initShrinkHandler,

		wire.FieldsOf(new(*config.Config), "HttpServer"),
		rest.NewServer,
	))
}

func InitApplication(_ *config.Config) (*application.Application, func(), error) {
	panic(wire.Build(
		initPgxPool,
		initPostresPool,
		initTransactionManager,

		initShrinkService,
		initStatisticsService,

		initRestApiServer,

		application.NewApplication,
	))
}
