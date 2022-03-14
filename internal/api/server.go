package api

import (
	"github.com/dimuska139/urlshortener/internal/api/middleware"
	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/gen/restapi"
	"github.com/dimuska139/urlshortener/internal/gen/restapi/operations"
	shrinkHhandlers "github.com/dimuska139/urlshortener/internal/handlers"
	"github.com/dimuska139/urlshortener/internal/logging"
	"github.com/go-openapi/loads"
	"net/http"
)

type RestAPI struct {
	config *config.Config
	api    *operations.UrlshortenerAPI
}

func NewRestAPI(
	config *config.Config,
	logger logging.Loggerer,
	middlewareFactory *middleware.MiddlewareFactory,
	shrinkHandler *shrinkHhandlers.ShrinkHandler) (*RestAPI, error) {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		logger.Fatal("unable to analyze swagger spec", err, nil)
	}
	api := operations.NewUrlshortenerAPI(swaggerSpec)
	api.Logger = logger.Printf

	//api.AuthPostAuthLoginBasicHandler = auth.PostAuthLoginBasicHandlerFunc(authHandler.LoginBasic)
	//api.AuthPostAuthRegistrationBasicHandler = auth.PostAuthRegistrationBasicHandlerFunc(authHandler.RegistrationBasic)
	api.PostShrinkHandler = operations.PostShrinkHandlerFunc(shrinkHandler.Shrink)
	api.GetShortCodeHandler = operations.GetShortCodeHandlerFunc(shrinkHandler.Redirect)

	api.PreServerShutdown = func() {}
	api.ServerShutdown = func() {}
	//api.BearerAuth = middlewareFactory.NewJwtAuthMiddleware

	return &RestAPI{config: config, api: api}, nil
}

func (server *RestAPI) setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func (server *RestAPI) Start() error {
	handler := server.api.Serve(server.setupMiddlewares)
	srv := restapi.NewServer(server.api)
	// defer server.Shutdown()
	srv.ConfigureFlags()
	srv.SetHandler(handler)
	srv.Port = server.config.Port
	return srv.Serve()
}
