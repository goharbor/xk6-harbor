// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectDeletable project deletable
//
// swagger:model ProjectDeletable
type ProjectDeletable struct {

	// Whether the project can be deleted.
	Deletable bool `json:"deletable,omitempty" js:"deletable"`

	// The detail message when the project can not be deleted.
	Message string `json:"message,omitempty" js:"message"`
}

// Validate validates this project deletable
func (m *ProjectDeletable) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project deletable based on context it is used
func (m *ProjectDeletable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectDeletable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectDeletable) UnmarshalBinary(b []byte) error {
	var res ProjectDeletable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
