package rest

import "net/http"

type MuxAdaptor struct {
	fn http.HandlerFunc
}

func NewMuxHandler(fn http.HandlerFunc) *MuxAdaptor {
	return &MuxAdaptor{
		fn: fn,
	}
}

func (h *MuxAdaptor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.fn(w, r)
}
