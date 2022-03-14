// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIResult Стандартизированный ответ API
//
// swagger:model ApiResult
type APIResult struct {

	// Данные
	Data interface{} `json:"data,omitempty"`

	// Список ошибок
	Errors interface{} `json:"errors,omitempty"`

	// Разная опциональная дополнительная информация
	Meta interface{} `json:"meta,omitempty"`

	// Результат выполнения запроса (true - нет ошибки, false - ошибка)
	// Example: true
	Success bool `json:"success,omitempty"`
}

// Validate validates this Api result
func (m *APIResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Api result based on context it is used
func (m *APIResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIResult) UnmarshalBinary(b []byte) error {
	var res APIResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
