package shrink

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/dimuska139/urlshortener/internal/service/server/rest/gen/models"
	"github.com/dimuska139/urlshortener/internal/service/server/rest/gen/restapi/operations"
	"github.com/dimuska139/urlshortener/pkg/logging"
)

type Handler struct {
	statisticsService StatisticsService
	shrinkService     ShrinkService
}

func NewHandler(statisticsService StatisticsService, shrinkService ShrinkService) *Handler {
	return &Handler{
		statisticsService: statisticsService,
		shrinkService:     shrinkService,
	}
}

func (h *Handler) Shrink(params operations.PostShrinkParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	internalError := operations.NewPostShrinkInternalServerError().
		WithPayload(&models.InternalError{
			Common: "Something went wrong",
		})

	link, err := h.shrinkService.CreateShortCode(ctx, *params.Body.LongURL)
	if err != nil {
		logging.Error(ctx, "Can't create short code",
			"err", err.Error())
		return internalError
	}

	return operations.NewPostShrinkOK().
		WithPayload(mapLinkToResponse(link))
}

func (h *Handler) Redirect(params operations.GetShortCodeParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	internalError := operations.NewGetShortCodeInternalServerError().
		WithPayload(&models.InternalError{
			Common: "Something went wrong",
		})

	longUrl, err := h.shrinkService.GetLongUrlByCode(params.HTTPRequest.Context(), params.ShortCode)
	if err != nil {
		logging.Error(ctx, "Can't get long url by code",
			"short_code", params.ShortCode,
			"err", err.Error(),
		)

		return internalError
	}

	if longUrl == "" {
		return operations.NewGetShortCodeNotFound().WithPayload(&models.NotFoundError{
			Common: "Not found",
		})
	}

	var userAgent string
	if params.UserAgent != nil {
		userAgent = *params.UserAgent
	}

	if err := h.statisticsService.SaveRedirectEvent(ctx, params.ShortCode, userAgent); err != nil {
		logging.Error(ctx, "Can't save redirect event",
			"short_code", params.ShortCode,
			"err", err.Error(),
		)

		return operations.NewGetShortCodeInternalServerError()
	}

	return operations.NewGetShortCodeFound().
		WithPayload(&models.RedirectURL{
			LongURL: longUrl,
		}).
		WithLocation(longUrl)
}
