package handlers

//go:generate mockgen -source=mapper.go -destination=./mapper_mock.go -package=handlers

import (
	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/gen/models"
	"github.com/dimuska139/urlshortener/internal/gen/restapi/operations"
	models2 "github.com/dimuska139/urlshortener/internal/models"
)

type Mapper interface {
	PostShrinkOKBody(link models2.Link) *operations.PostShrinkOKBody
}

type ResponseMapper struct {
	config *config.Config
}

func NewResponseMapper(config *config.Config) Mapper {
	return &ResponseMapper{config: config}
}

func (mapper *ResponseMapper) PostShrinkOKBody(link models2.Link) *operations.PostShrinkOKBody {
	return &operations.PostShrinkOKBody{
		Data: &models.ProcessedLink{
			ID:       link.Code,
			LongURL:  link.LongURL,
			ShortURL: link.ShortURL,
		},
		Success: true,
	}
}
