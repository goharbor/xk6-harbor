// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery --name API --keeptree --with-expecter --case underscore

// API is the interface of the permissions client
type API interface {
	/*
	   GetPermissions gets system or project level permissions info

	   This endpoint is for retrieving resource and action info that only provides for admin user(system admin and project admin).
	*/
	GetPermissions(ctx context.Context, params *GetPermissionsParams) (*GetPermissionsOK, error)
}

// New creates a new permissions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for permissions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetPermissions gets system or project level permissions info

This endpoint is for retrieving resource and action info that only provides for admin user(system admin and project admin).
*/
func (a *Client) GetPermissions(ctx context.Context, params *GetPermissionsParams) (*GetPermissionsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPermissions",
		Method:             "GET",
		PathPattern:        "/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPermissionsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetPermissionsOK:
		return value, nil
	case *GetPermissionsUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetPermissionsForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetPermissionsNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetPermissionsInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}
