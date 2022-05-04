package http

import "net/http"

func NewHttpClient() *http.Client {
	return http.DefaultClient
}
