package services

import (
	"context"
	"github.com/dimuska139/urlshortener/internal/config"
)

type StatisticsService struct {
	config *config.Config
	db     Storage
}

func NewStatisticsService(cfg *config.Config, db Storage) *StatisticsService {
	return &StatisticsService{
		config: cfg,
		db:     db,
	}
}

func (s *StatisticsService) SaveRedirectEvent(ctx context.Context, code string, userAgent string) error {
	statisticsRepository := s.db.Repositories.StatisticsRepository()
	return statisticsRepository.SaveRedirectEvent(ctx, code, userAgent)
}
