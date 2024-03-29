// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScannerRegistrationReq scanner registration req
//
// swagger:model ScannerRegistrationReq
type ScannerRegistrationReq struct {

	// An optional value of the HTTP Authorization header sent with each request to the Scanner Adapter API.
	//
	// Example: Bearer: JWTTOKENGOESHERE
	AccessCredential string `json:"access_credential,omitempty" js:"accessCredential"`

	// Specify what authentication approach is adopted for the HTTP communications.
	// Supported types Basic", "Bearer" and api key header "X-ScannerAdapter-API-Key"
	//
	// Example: Bearer
	Auth string `json:"auth,omitempty" js:"auth"`

	// An optional description of this registration.
	// Example: A free-to-use tool that scans container images for package vulnerabilities.\n
	Description string `json:"description,omitempty" js:"description"`

	// Indicate whether the registration is enabled or not
	Disabled *bool `json:"disabled,omitempty" js:"disabled"`

	// The name of this registration
	// Example: Trivy
	// Required: true
	Name *string `json:"name" js:"name"`

	// Indicate if skip the certificate verification when sending HTTP requests
	SkipCertVerify *bool `json:"skip_certVerify,omitempty" js:"skipCertVerify"`

	// A base URL of the scanner adapter.
	// Example: http://harbor-scanner-trivy:8080
	// Required: true
	// Format: uri
	URL *strfmt.URI `json:"url" js:"url"`

	// Indicate whether use internal registry addr for the scanner to pull content or not
	UseInternalAddr *bool `json:"use_internal_addr,omitempty" js:"useInternalAddr"`
}

// Validate validates this scanner registration req
func (m *ScannerRegistrationReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScannerRegistrationReq) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ScannerRegistrationReq) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this scanner registration req based on context it is used
func (m *ScannerRegistrationReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScannerRegistrationReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScannerRegistrationReq) UnmarshalBinary(b []byte) error {
	var res ScannerRegistrationReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
