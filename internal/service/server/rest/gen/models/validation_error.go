// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ValidationError validation error
//
// swagger:model ValidationError
type ValidationError map[string]string

// Validate validates this validation error
func (m ValidationError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this validation error based on context it is used
func (m ValidationError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
