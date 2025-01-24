package shrink

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"

	"github.com/dimuska139/urlshortener/internal/model"
)

type ShrinkService struct {
	config         Config
	txManager      TransactionManager
	linkRepository LinkRepository
}

func NewShrinkService(
	config Config,
	txManager TransactionManager,
	linkRepository LinkRepository,
) *ShrinkService {
	return &ShrinkService{
		config:         config,
		txManager:      txManager,
		linkRepository: linkRepository,
	}
}

func generateShortcode(id int) string {
	const shortcodeAlphabet = "0123456789abcdefghijklmnopqrstuvwxyz"

	base := len(shortcodeAlphabet)

	var encoded string
	for id > 0 {
		encoded += string(shortcodeAlphabet[id%base])
		id = id / base
	}

	var reversed string
	for _, v := range encoded {
		reversed = string(v) + reversed
	}

	return reversed
}

func (s *ShrinkService) CreateShortCode(ctx context.Context, longUrl string) (model.Link, error) {
	var shortLink model.Link

	if err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
		link, err := s.linkRepository.Create(ctx, longUrl)
		if err != nil {
			return fmt.Errorf("create link: %w", err)
		}

		link.Code = generateShortcode(link.ID)
		if err := s.linkRepository.SetShortcode(ctx, link.ID, link.Code); err != nil {
			return fmt.Errorf("set shortcode: %w", err)
		}

		link.ShortURL = fmt.Sprintf(s.config.ShortUrlTemplate, link.Code)

		shortLink = link

		return nil
	}, pgx.TxOptions{}); err != nil {
		return model.Link{}, fmt.Errorf("tx: %w", err)
	}

	return shortLink, nil
}

func (s *ShrinkService) GetLongUrlByCode(ctx context.Context, shortCode string) (string, error) {
	longUrl, err := s.linkRepository.GetLongUrlByCode(ctx, shortCode)
	if err != nil {
		return "", fmt.Errorf("get long url by code: %w", err)
	}

	return longUrl, nil
}
