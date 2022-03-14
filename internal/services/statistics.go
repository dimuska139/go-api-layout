package services

import (
	"context"
	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/storage"
)

//go:generate mockgen -source=statistics.go -destination=./statistics_mock.go -package=services

type StatisticsServiceInterface interface {
	SaveRedirectEvent(ctx context.Context, code string, userAgent string) error
}

type StatisticsService struct {
	config *config.Config
	db     storage.Database
}

func NewStatisticsService(cfg *config.Config, db storage.Database) *StatisticsService {
	return &StatisticsService{
		config: cfg,
		db:     db,
	}
}

func (s *StatisticsService) SaveRedirectEvent(ctx context.Context, code string, userAgent string) error {
	statisticsRepository := s.db.Repositories.GetStatisticsRepository()
	return statisticsRepository.SaveRedirectEvent(ctx, code, userAgent)
}
