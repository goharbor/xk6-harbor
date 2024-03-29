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

// ChartMetadata The metadata of chart version
//
// swagger:model ChartMetadata
type ChartMetadata struct {

	// The API version of this chart
	// Required: true
	APIVersion *string `json:"apiVersion" js:"apiVersion"`

	// The version of the application enclosed in the chart
	// Required: true
	AppVersion *string `json:"appVersion" js:"appVersion"`

	// Whether or not this chart is deprecated
	Deprecated bool `json:"deprecated,omitempty" js:"deprecated"`

	// A one-sentence description of chart
	Description string `json:"description,omitempty" js:"description"`

	// The name of template engine
	// Required: true
	Engine *string `json:"engine" js:"engine"`

	// The URL to the relevant project page
	Home string `json:"home,omitempty" js:"home"`

	// The URL to an icon file
	// Required: true
	Icon *string `json:"icon" js:"icon"`

	// A list of string keywords
	Keywords []string `json:"keywords" js:"keywords"`

	// The name of the chart
	// Required: true
	Name *string `json:"name" js:"name"`

	// The URL to the source code of chart
	Sources []string `json:"sources" js:"sources"`

	// A SemVer 2 version of chart
	// Required: true
	Version *string `json:"version" js:"version"`
}

// Validate validates this chart metadata
func (m *ChartMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChartMetadata) validateAPIVersion(formats strfmt.Registry) error {

	if err := validate.Required("apiVersion", "body", m.APIVersion); err != nil {
		return err
	}

	return nil
}

func (m *ChartMetadata) validateAppVersion(formats strfmt.Registry) error {

	if err := validate.Required("appVersion", "body", m.AppVersion); err != nil {
		return err
	}

	return nil
}

func (m *ChartMetadata) validateEngine(formats strfmt.Registry) error {

	if err := validate.Required("engine", "body", m.Engine); err != nil {
		return err
	}

	return nil
}

func (m *ChartMetadata) validateIcon(formats strfmt.Registry) error {

	if err := validate.Required("icon", "body", m.Icon); err != nil {
		return err
	}

	return nil
}

func (m *ChartMetadata) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ChartMetadata) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this chart metadata based on context it is used
func (m *ChartMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChartMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChartMetadata) UnmarshalBinary(b []byte) error {
	var res ChartMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
