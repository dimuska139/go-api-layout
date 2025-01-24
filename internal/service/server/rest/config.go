package rest

import "github.com/dimuska139/urlshortener/internal/service/server/rest/middleware"

type Config struct {
	Port       int               `yaml:"port"`
	Middleware middleware.Config `yaml:"middleware"`
}
