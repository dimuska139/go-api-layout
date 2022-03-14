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

// PostShrinkHandlerFunc turns a function with the right signature into a post shrink handler
type PostShrinkHandlerFunc func(PostShrinkParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostShrinkHandlerFunc) Handle(params PostShrinkParams) middleware.Responder {
	return fn(params)
}

// PostShrinkHandler interface for that can handle valid post shrink params
type PostShrinkHandler interface {
	Handle(PostShrinkParams) middleware.Responder
}

// NewPostShrink creates a new http.Handler for the post shrink operation
func NewPostShrink(ctx *middleware.Context, handler PostShrinkHandler) *PostShrink {
	return &PostShrink{Context: ctx, Handler: handler}
}

/* PostShrink swagger:route POST /shrink postShrink

Сократить ссылку

*/
type PostShrink struct {
	Context *middleware.Context
	Handler PostShrinkHandler
}

func (o *PostShrink) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostShrinkParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostShrinkBadRequestBody post shrink bad request body
//
// swagger:model PostShrinkBadRequestBody
type PostShrinkBadRequestBody struct {
	models.APIResult

	// success
	// Example: false
	Success bool `json:"success,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostShrinkBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PostShrinkBadRequestBodyAO0
	var postShrinkBadRequestBodyAO0 models.APIResult
	if err := swag.ReadJSON(raw, &postShrinkBadRequestBodyAO0); err != nil {
		return err
	}
	o.APIResult = postShrinkBadRequestBodyAO0

	// PostShrinkBadRequestBodyAO1
	var dataPostShrinkBadRequestBodyAO1 struct {
		Success bool `json:"success,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostShrinkBadRequestBodyAO1); err != nil {
		return err
	}

	o.Success = dataPostShrinkBadRequestBodyAO1.Success

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostShrinkBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postShrinkBadRequestBodyAO0, err := swag.WriteJSON(o.APIResult)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postShrinkBadRequestBodyAO0)
	var dataPostShrinkBadRequestBodyAO1 struct {
		Success bool `json:"success,omitempty"`
	}

	dataPostShrinkBadRequestBodyAO1.Success = o.Success

	jsonDataPostShrinkBadRequestBodyAO1, errPostShrinkBadRequestBodyAO1 := swag.WriteJSON(dataPostShrinkBadRequestBodyAO1)
	if errPostShrinkBadRequestBodyAO1 != nil {
		return nil, errPostShrinkBadRequestBodyAO1
	}
	_parts = append(_parts, jsonDataPostShrinkBadRequestBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post shrink bad request body
func (o *PostShrinkBadRequestBody) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this post shrink bad request body based on the context it is used
func (o *PostShrinkBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
func (o *PostShrinkBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostShrinkBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostShrinkBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostShrinkInternalServerErrorBody post shrink internal server error body
//
// swagger:model PostShrinkInternalServerErrorBody
type PostShrinkInternalServerErrorBody struct {
	models.APIResult

	// success
	// Example: false
	Success bool `json:"success,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostShrinkInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// PostShrinkInternalServerErrorBodyAO0
	var postShrinkInternalServerErrorBodyAO0 models.APIResult
	if err := swag.ReadJSON(raw, &postShrinkInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.APIResult = postShrinkInternalServerErrorBodyAO0

	// PostShrinkInternalServerErrorBodyAO1
	var dataPostShrinkInternalServerErrorBodyAO1 struct {
		Success bool `json:"success,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostShrinkInternalServerErrorBodyAO1); err != nil {
		return err
	}

	o.Success = dataPostShrinkInternalServerErrorBodyAO1.Success

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostShrinkInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postShrinkInternalServerErrorBodyAO0, err := swag.WriteJSON(o.APIResult)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postShrinkInternalServerErrorBodyAO0)
	var dataPostShrinkInternalServerErrorBodyAO1 struct {
		Success bool `json:"success,omitempty"`
	}

	dataPostShrinkInternalServerErrorBodyAO1.Success = o.Success

	jsonDataPostShrinkInternalServerErrorBodyAO1, errPostShrinkInternalServerErrorBodyAO1 := swag.WriteJSON(dataPostShrinkInternalServerErrorBodyAO1)
	if errPostShrinkInternalServerErrorBodyAO1 != nil {
		return nil, errPostShrinkInternalServerErrorBodyAO1
	}
	_parts = append(_parts, jsonDataPostShrinkInternalServerErrorBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post shrink internal server error body
func (o *PostShrinkInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this post shrink internal server error body based on the context it is used
func (o *PostShrinkInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
func (o *PostShrinkInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostShrinkInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostShrinkInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostShrinkOKBody post shrink o k body
//
// swagger:model PostShrinkOKBody
type PostShrinkOKBody struct {
	models.APIResult

	// data
	Data *models.ProcessedLink `json:"data,omitempty"`

	// success
	// Example: true
	Success bool `json:"success,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostShrinkOKBody) UnmarshalJSON(raw []byte) error {
	// PostShrinkOKBodyAO0
	var postShrinkOKBodyAO0 models.APIResult
	if err := swag.ReadJSON(raw, &postShrinkOKBodyAO0); err != nil {
		return err
	}
	o.APIResult = postShrinkOKBodyAO0

	// PostShrinkOKBodyAO1
	var dataPostShrinkOKBodyAO1 struct {
		Data *models.ProcessedLink `json:"data,omitempty"`

		Success bool `json:"success,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostShrinkOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostShrinkOKBodyAO1.Data

	o.Success = dataPostShrinkOKBodyAO1.Success

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostShrinkOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postShrinkOKBodyAO0, err := swag.WriteJSON(o.APIResult)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postShrinkOKBodyAO0)
	var dataPostShrinkOKBodyAO1 struct {
		Data *models.ProcessedLink `json:"data,omitempty"`

		Success bool `json:"success,omitempty"`
	}

	dataPostShrinkOKBodyAO1.Data = o.Data

	dataPostShrinkOKBodyAO1.Success = o.Success

	jsonDataPostShrinkOKBodyAO1, errPostShrinkOKBodyAO1 := swag.WriteJSON(dataPostShrinkOKBodyAO1)
	if errPostShrinkOKBodyAO1 != nil {
		return nil, errPostShrinkOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostShrinkOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post shrink o k body
func (o *PostShrinkOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostShrinkOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postShrinkOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postShrinkOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post shrink o k body based on the context it is used
func (o *PostShrinkOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostShrinkOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postShrinkOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postShrinkOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostShrinkOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostShrinkOKBody) UnmarshalBinary(b []byte) error {
	var res PostShrinkOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
