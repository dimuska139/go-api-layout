package middleware

import (
	"github.com/dimuska139/urlshortener/internal/service/server/rest/middleware/cors"
)

type Factory struct {
	Cors *cors.Middleware
}
