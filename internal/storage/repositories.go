package storage

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/dimuska139/urlshortener/internal/repository"
	"github.com/jackc/pgtype/pgxtype"
)

type RepositoriesFactoryInterface interface {
	GetLinkRepository() repository.LinkRepository
	GetStatisticsRepository() repository.StatisticsRepository
}

type RepositoriesFactory struct {
	querier      pgxtype.Querier
	queryBuilder sq.StatementBuilderType
}

func NewRepositoriesFactory(q pgxtype.Querier) *RepositoriesFactory {
	return &RepositoriesFactory{
		querier:      q,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (factory *RepositoriesFactory) GetLinkRepository() repository.LinkRepository {
	return repository.NewLinkPostgresqlRepository(factory.querier, factory.queryBuilder)
}

func (factory *RepositoriesFactory) GetStatisticsRepository() repository.StatisticsRepository {
	return repository.NewStatisticsPostgresqlRepository(factory.querier, factory.queryBuilder)
}
