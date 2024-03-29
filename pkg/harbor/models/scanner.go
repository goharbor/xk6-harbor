// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Scanner scanner
//
// swagger:model Scanner
type Scanner struct {

	// Name of the scanner
	// Example: Trivy
	Name string `json:"name,omitempty" js:"name"`

	// Name of the scanner provider
	// Example: Aqua Security
	Vendor string `json:"vendor,omitempty" js:"vendor"`

	// Version of the scanner adapter
	// Example: v0.9.1
	Version string `json:"version,omitempty" js:"version"`
}

// Validate validates this scanner
func (m *Scanner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this scanner based on context it is used
func (m *Scanner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Scanner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Scanner) UnmarshalBinary(b []byte) error {
	var res Scanner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
