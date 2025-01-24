package shrink

import (
	"github.com/dimuska139/urlshortener/internal/model"
	"github.com/dimuska139/urlshortener/internal/service/server/rest/gen/models"
)

func mapLinkToResponse(link model.Link) *models.ShortLink {
	return &models.ShortLink{
		ShortURL: link.ShortURL,
		LongURL:  link.LongURL,
	}
}
