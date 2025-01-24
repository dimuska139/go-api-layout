package rest

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/tidwall/sjson"
	"net/http"
	"time"

	"github.com/dimuska139/urlshortener/internal/service/server/rest/gen/restapi"
	"github.com/dimuska139/urlshortener/internal/service/server/rest/gen/restapi/operations"
	shrinkHandler "github.com/dimuska139/urlshortener/internal/service/server/rest/handler/shrink"
	"github.com/dimuska139/urlshortener/internal/service/server/rest/middleware"
	"github.com/dimuska139/urlshortener/pkg/logging"
)

const version = "1.0.0"

type Server struct {
	server            *http.Server
	config            Config
	middlewareFactory *middleware.Factory

	shrinkHandler *shrinkHandler.Handler
}

func NewServer(
	cfg Config,
	middlewareFactory *middleware.Factory,
	shrinkHandler *shrinkHandler.Handler,
) (*Server, error) {
	ctx := context.Background()

	restApi := &Server{
		config:            cfg,
		middlewareFactory: middlewareFactory,
		shrinkHandler:     shrinkHandler,
	}

	api, err := restApi.buildAPI(ctx)
	if err != nil {
		return nil, fmt.Errorf("build API: %w", err)
	}

	restApi.server = &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           api,
	}

	return restApi, nil
}

func (s *Server) buildAPI(ctx context.Context) (*http.ServeMux, error) {
	specJSON, _ := sjson.Set(string(restapi.SwaggerJSON), "info.version", version)

	swaggerSpec, err := loads.Analyzed([]byte(specJSON), "")
	if err != nil {
		return nil, fmt.Errorf("analyze Swagger spec: %w", err)
	}

	api := operations.NewUrlshortenerAPI(swaggerSpec)
	api.Logger = func(format string, v ...any) {
		logging.Info(ctx, fmt.Sprintf(format, v...))
	}

	api.PostShrinkHandler = operations.PostShrinkHandlerFunc(s.shrinkHandler.Shrink)
	api.GetShortCodeHandler = operations.GetShortCodeHandlerFunc(s.shrinkHandler.Redirect)
	api.ServeError = handleErrors

	api.UseSwaggerUI()

	mux := http.DefaultServeMux
	mux.Handle("/", api.Serve(s.setupMiddlewares))

	return mux, nil
}

func (s *Server) setupMiddlewares(handler http.Handler) http.Handler {
	middlewares := []func(handler http.Handler) http.Handler{
		s.middlewareFactory.Cors.GetMiddleware(),
	}

	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown server: %w", err)
	}

	return nil
}
