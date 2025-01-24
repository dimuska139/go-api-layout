package middleware

import "github.com/dimuska139/urlshortener/internal/service/server/rest/middleware/cors"

type Config struct {
	Cors cors.Config `yaml:"cors"`
}
