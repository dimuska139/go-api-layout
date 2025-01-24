// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotFoundError Not found
//
// swagger:model NotFoundError
type NotFoundError struct {

	// common
	// Example: Link is not found
	Common string `json:"common,omitempty"`
}

// Validate validates this not found error
func (m *NotFoundError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this not found error based on context it is used
func (m *NotFoundError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotFoundError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotFoundError) UnmarshalBinary(b []byte) error {
	var res NotFoundError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
