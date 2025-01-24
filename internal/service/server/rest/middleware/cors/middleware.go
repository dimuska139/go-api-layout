package cors

import (
	"github.com/rs/cors"
	"net/http"
)

type Middleware struct {
	config Config
}

func NewMiddleware(config Config) *Middleware {
	return &Middleware{
		config: config,
	}
}

func (m *Middleware) GetMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		allowedOrigins := []string{"*"}
		if len(m.config.AllowedOrigins) != 0 {
			allowedOrigins = m.config.AllowedOrigins
		}

		return cors.New(cors.Options{
			AllowedOrigins: allowedOrigins,
			AllowedHeaders: []string{"*"},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodPost,
				http.MethodGet,
				http.MethodPut,
				http.MethodDelete,
			},
		}).Handler(next)
	}
}
