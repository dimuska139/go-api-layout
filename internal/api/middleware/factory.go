package middleware

import "github.com/dimuska139/urlshortener/internal/logging"

type MiddlewareFactory struct {
	logger logging.Loggerer
}

func NewMiddlewareFactory(logger logging.Loggerer) (*MiddlewareFactory, error) {
	return &MiddlewareFactory{logger}, nil
}
