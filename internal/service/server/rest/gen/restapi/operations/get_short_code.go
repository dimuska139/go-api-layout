// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetShortCodeHandlerFunc turns a function with the right signature into a get short code handler
type GetShortCodeHandlerFunc func(GetShortCodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetShortCodeHandlerFunc) Handle(params GetShortCodeParams) middleware.Responder {
	return fn(params)
}

// GetShortCodeHandler interface for that can handle valid get short code params
type GetShortCodeHandler interface {
	Handle(GetShortCodeParams) middleware.Responder
}

// NewGetShortCode creates a new http.Handler for the get short code operation
func NewGetShortCode(ctx *middleware.Context, handler GetShortCodeHandler) *GetShortCode {
	return &GetShortCode{Context: ctx, Handler: handler}
}

/*
	GetShortCode swagger:route GET /{shortCode} getShortCode

Represents a short URL. Tracks the visit and redirects to the corresponding long URL
*/
type GetShortCode struct {
	Context *middleware.Context
	Handler GetShortCodeHandler
}

func (o *GetShortCode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetShortCodeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
