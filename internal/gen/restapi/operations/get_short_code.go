// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/dimuska139/urlshortener/internal/gen/models"
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

/* GetShortCode swagger:route GET /{shortCode} getShortCode

Represents a short URL. Tracks the visit and redirects tio the corresponding long URL

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

// GetShortCodeFoundBody get short code found body
//
// swagger:model GetShortCodeFoundBody
type GetShortCodeFoundBody struct {
	models.APIResult

	// data
	Data *models.RedirectURL `json:"data,omitempty"`

	// success
	// Example: true
	Success bool `json:"success,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetShortCodeFoundBody) UnmarshalJSON(raw []byte) error {
	// GetShortCodeFoundBodyAO0
	var getShortCodeFoundBodyAO0 models.APIResult
	if err := swag.ReadJSON(raw, &getShortCodeFoundBodyAO0); err != nil {
		return err
	}
	o.APIResult = getShortCodeFoundBodyAO0

	// GetShortCodeFoundBodyAO1
	var dataGetShortCodeFoundBodyAO1 struct {
		Data *models.RedirectURL `json:"data,omitempty"`

		Success bool `json:"success,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetShortCodeFoundBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetShortCodeFoundBodyAO1.Data

	o.Success = dataGetShortCodeFoundBodyAO1.Success

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetShortCodeFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getShortCodeFoundBodyAO0, err := swag.WriteJSON(o.APIResult)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getShortCodeFoundBodyAO0)
	var dataGetShortCodeFoundBodyAO1 struct {
		Data *models.RedirectURL `json:"data,omitempty"`

		Success bool `json:"success,omitempty"`
	}

	dataGetShortCodeFoundBodyAO1.Data = o.Data

	dataGetShortCodeFoundBodyAO1.Success = o.Success

	jsonDataGetShortCodeFoundBodyAO1, errGetShortCodeFoundBodyAO1 := swag.WriteJSON(dataGetShortCodeFoundBodyAO1)
	if errGetShortCodeFoundBodyAO1 != nil {
		return nil, errGetShortCodeFoundBodyAO1
	}
	_parts = append(_parts, jsonDataGetShortCodeFoundBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get short code found body
func (o *GetShortCodeFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.APIResult
	if err := o.APIResult.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetShortCodeFoundBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getShortCodeFound" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getShortCodeFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get short code found body based on the context it is used
func (o *GetShortCodeFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.APIResult
	if err := o.APIResult.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetShortCodeFoundBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getShortCodeFound" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getShortCodeFound" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetShortCodeFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetShortCodeFoundBody) UnmarshalBinary(b []byte) error {
	var res GetShortCodeFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetShortCodeInternalServerErrorBody get short code internal server error body
//
// swagger:model GetShortCodeInternalServerErrorBody
type GetShortCodeInternalServerErrorBody struct {
	models.APIResult

	// success
	// Example: false
	Success bool `json:"success,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetShortCodeInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// GetShortCodeInternalServerErrorBodyAO0
	var getShortCodeInternalServerErrorBodyAO0 models.APIResult
	if err := swag.ReadJSON(raw, &getShortCodeInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.APIResult = getShortCodeInternalServerErrorBodyAO0

	// GetShortCodeInternalServerErrorBodyAO1
	var dataGetShortCodeInternalServerErrorBodyAO1 struct {
		Success bool `json:"success,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetShortCodeInternalServerErrorBodyAO1); err != nil {
		return err
	}

	o.Success = dataGetShortCodeInternalServerErrorBodyAO1.Success

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetShortCodeInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getShortCodeInternalServerErrorBodyAO0, err := swag.WriteJSON(o.APIResult)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getShortCodeInternalServerErrorBodyAO0)
	var dataGetShortCodeInternalServerErrorBodyAO1 struct {
		Success bool `json:"success,omitempty"`
	}

	dataGetShortCodeInternalServerErrorBodyAO1.Success = o.Success

	jsonDataGetShortCodeInternalServerErrorBodyAO1, errGetShortCodeInternalServerErrorBodyAO1 := swag.WriteJSON(dataGetShortCodeInternalServerErrorBodyAO1)
	if errGetShortCodeInternalServerErrorBodyAO1 != nil {
		return nil, errGetShortCodeInternalServerErrorBodyAO1
	}
	_parts = append(_parts, jsonDataGetShortCodeInternalServerErrorBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get short code internal server error body
func (o *GetShortCodeInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.APIResult
	if err := o.APIResult.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this get short code internal server error body based on the context it is used
func (o *GetShortCodeInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.APIResult
	if err := o.APIResult.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *GetShortCodeInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetShortCodeInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetShortCodeInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetShortCodeNotFoundBody get short code not found body
//
// swagger:model GetShortCodeNotFoundBody
type GetShortCodeNotFoundBody struct {
	models.APIResult

	// success
	// Example: false
	Success bool `json:"success,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetShortCodeNotFoundBody) UnmarshalJSON(raw []byte) error {
	// GetShortCodeNotFoundBodyAO0
	var getShortCodeNotFoundBodyAO0 models.APIResult
	if err := swag.ReadJSON(raw, &getShortCodeNotFoundBodyAO0); err != nil {
		return err
	}
	o.APIResult = getShortCodeNotFoundBodyAO0

	// GetShortCodeNotFoundBodyAO1
	var dataGetShortCodeNotFoundBodyAO1 struct {
		Success bool `json:"success,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetShortCodeNotFoundBodyAO1); err != nil {
		return err
	}

	o.Success = dataGetShortCodeNotFoundBodyAO1.Success

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetShortCodeNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getShortCodeNotFoundBodyAO0, err := swag.WriteJSON(o.APIResult)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getShortCodeNotFoundBodyAO0)
	var dataGetShortCodeNotFoundBodyAO1 struct {
		Success bool `json:"success,omitempty"`
	}

	dataGetShortCodeNotFoundBodyAO1.Success = o.Success

	jsonDataGetShortCodeNotFoundBodyAO1, errGetShortCodeNotFoundBodyAO1 := swag.WriteJSON(dataGetShortCodeNotFoundBodyAO1)
	if errGetShortCodeNotFoundBodyAO1 != nil {
		return nil, errGetShortCodeNotFoundBodyAO1
	}
	_parts = append(_parts, jsonDataGetShortCodeNotFoundBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get short code not found body
func (o *GetShortCodeNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.APIResult
	if err := o.APIResult.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this get short code not found body based on the context it is used
func (o *GetShortCodeNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.APIResult
	if err := o.APIResult.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *GetShortCodeNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetShortCodeNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetShortCodeNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
