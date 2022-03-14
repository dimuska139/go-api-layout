package services

import (
	"context"
	"fmt"
	"github.com/dimuska139/urlshortener/internal/config"
	helpers2 "github.com/dimuska139/urlshortener/internal/helpers"
	"github.com/dimuska139/urlshortener/internal/models"
	"github.com/dimuska139/urlshortener/internal/storage"
)

//go:generate mockgen -source=shrink.go -destination=./shrink_mock.go -package=services

type ShrinkServiceInterface interface {
	CreateShortCode(ctx context.Context, longUrl string) (models.Link, error)
	GetLongUrlByCode(ctx context.Context, shortCode string) (string, error)
}

type ShrinkService struct {
	config *config.Config
	db     storage.Database
}

func NewShrinkService(cfg *config.Config, db storage.Database) *ShrinkService {
	return &ShrinkService{
		config: cfg,
		db:     db,
	}
}

func (s *ShrinkService) CreateShortCode(ctx context.Context, longUrl string) (models.Link, error) {
	tx, err := s.db.BeginTx(ctx)
	if err != nil {
		return models.Link{}, fmt.Errorf("unable to upen transaction: %w", err)
	}

	linkRepository := tx.Repositories.GetLinkRepository()
	link, err := linkRepository.Create(ctx, longUrl)
	if err != nil {
		tx.Rollback(ctx)
		return models.Link{}, fmt.Errorf("unable to save link: %w", err)
	}

	link.Code = helpers2.GenerateShortcode(link.ID)
	if err := linkRepository.SetShortcode(ctx, link.ID, link.Code); err != nil {
		tx.Rollback(ctx)

		return models.Link{}, fmt.Errorf("unable to update link code: %w", err)
	}
	tx.Commit(ctx)

	return link, nil
}

func (s *ShrinkService) GetLongUrlByCode(ctx context.Context, shortCode string) (string, error) {
	linkRepository := s.db.Repositories.GetLinkRepository()
	return linkRepository.GetLongUrlByCode(ctx, shortCode)
}
