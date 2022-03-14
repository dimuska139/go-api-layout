package handlers

import (
	"github.com/dimuska139/urlshortener/internal/gen/models"
	"github.com/dimuska139/urlshortener/internal/logging"
	"github.com/dimuska139/urlshortener/internal/services"

	//"context"
	//"fmt"
	"github.com/dimuska139/urlshortener/internal/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

type ShrinkHandler struct {
	logger            logging.Loggerer
	shrinkUsecase     services.ShrinkServiceInterface
	statisticsUsecase services.StatisticsServiceInterface
	responseMapper    Mapper
}

func NewShrinkHandler(logger logging.Loggerer, shrinkUsecase services.ShrinkServiceInterface, statisticsUsecase services.StatisticsServiceInterface, responseMapper Mapper) *ShrinkHandler {
	return &ShrinkHandler{logger: logger, shrinkUsecase: shrinkUsecase, statisticsUsecase: statisticsUsecase, responseMapper: responseMapper}
}

func (shrinkHandler *ShrinkHandler) Shrink(params operations.PostShrinkParams) middleware.Responder {
	link, err := shrinkHandler.shrinkUsecase.CreateShortCode(params.HTTPRequest.Context(), *params.Body.LongURL)
	if err != nil {
		shrinkHandler.logger.Error("unable to create short code", err, nil)
		return operations.NewPostShrinkInternalServerError()
	}

	response := operations.NewPostShrinkOK()
	response.Payload = shrinkHandler.responseMapper.PostShrinkOKBody(link)
	return response
}

func (shrinkHandler *ShrinkHandler) Redirect(params operations.GetShortCodeParams) middleware.Responder {
	longUrl, err := shrinkHandler.shrinkUsecase.GetLongUrlByCode(params.HTTPRequest.Context(), params.ShortCode)
	if err != nil {
		shrinkHandler.logger.Error("unable to get long url by code", err, map[string]interface{}{
			"short_code": params.ShortCode,
		})
		return operations.NewGetShortCodeInternalServerError()
	}

	if longUrl == "" {
		return operations.NewGetShortCodeNotFound()
	}

	if err := shrinkHandler.statisticsUsecase.SaveRedirectEvent(params.HTTPRequest.Context(), params.ShortCode, *params.UserAgent); err != nil {
		shrinkHandler.logger.Error("unable to save redirect event", err, map[string]interface{}{
			"short_code": params.ShortCode,
		})
		return operations.NewGetShortCodeInternalServerError()
	}

	response := operations.NewGetShortCodeFound()
	response.Location = longUrl
	response.Payload = &operations.GetShortCodeFoundBody{
		Data: &models.RedirectURL{
			LongURL: longUrl,
		},
		Success: true,
	}
	return response
}
