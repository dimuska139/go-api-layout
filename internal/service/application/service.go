package application

import (
	"context"

	"github.com/dimuska139/urlshortener/internal/service/server/rest"
	"github.com/dimuska139/urlshortener/pkg/logging"
)

type Application struct {
	restAPI *rest.Server
}

func NewApplication(
	restAPI *rest.Server) *Application {
	return &Application{
		restAPI: restAPI,
	}
}

func (a *Application) Run(ctx context.Context) {
	go func() {
		if err := a.restAPI.Start(); err != nil {
			logging.Fatal(ctx, "Can't start API server",
				"err", err.Error())
		}
	}()
}

func (a *Application) Stop(ctx context.Context) {
	if err := a.restAPI.Stop(ctx); err != nil {
		logging.Fatal(ctx, "Can't stop API server",
			"err", err.Error())
	}
}
