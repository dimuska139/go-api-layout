package statistics

import (
	"context"
	"fmt"
)

type StatisticsService struct {
	statisticsRepository StatisticsRepository
}

func NewStatisticsService(rep StatisticsRepository) *StatisticsService {
	return &StatisticsService{
		statisticsRepository: rep,
	}
}

func (s *StatisticsService) SaveRedirectEvent(ctx context.Context, code string, userAgent string) error {
	if err := s.statisticsRepository.SaveRedirectEvent(ctx, code, userAgent); err != nil {
		return fmt.Errorf("save redirect event: %w", err)
	}

	return nil
}
