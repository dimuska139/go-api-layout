package services

import (
	"context"
	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/logging"
	"github.com/dimuska139/urlshortener/internal/models"
	"github.com/dimuska139/urlshortener/internal/repository"
	"github.com/jackc/pgtype/pgxtype"
	"net/http"
)

type LinkRepository interface {
	Create(ctx context.Context, longUrl string) (models.Link, error)
	GetLongUrlByCode(ctx context.Context, shortCode string) (string, error)
	SetShortcode(ctx context.Context, id int, code string) error
}

type StatisticsRepository interface {
	SaveRedirectEvent(ctx context.Context, code string, userAgent string) error
}

type RepositoryFactoryInterface interface {
	GetLinkRepository() LinkRepository
	GetStatisticsRepository() StatisticsRepository
}

// TransactionalRepositoryFactory фабрика репозиториев, которые могут работать в транзакции
type TransactionalRepositoryFactory struct {
	config  *config.Config
	logger  logging.Loggerer
	querier pgxtype.Querier
}

func NewTransactionalRepositoryFactory(config *config.Config, logger logging.Loggerer, q pgxtype.Querier) *TransactionalRepositoryFactory {
	return &TransactionalRepositoryFactory{
		config:  config,
		logger:  logger,
		querier: q,
	}
}

// RepositoryFactory фабрика репозиториев, которые не могут работать в транзакции
type RepositoryFactory struct {
	TransactionalRepositoryFactory
	httpClient *http.Client
}

func NewRepositoryFactory(config *config.Config, logger logging.Loggerer, q pgxtype.Querier, httpClient *http.Client) *RepositoryFactory {
	return &RepositoryFactory{
		TransactionalRepositoryFactory: TransactionalRepositoryFactory{
			config:  config,
			logger:  logger,
			querier: q,
		},
		httpClient: httpClient,
	}
}

func (factory *TransactionalRepositoryFactory) LinkRepository() LinkRepository {
	return repository.NewLinkPostgresqlRepository(factory.logger, factory.querier)
}

func (factory *TransactionalRepositoryFactory) StatisticsRepository() StatisticsRepository {
	return repository.NewStatisticsPostgresqlRepository(factory.logger, factory.querier)
}
