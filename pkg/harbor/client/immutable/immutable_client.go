// Code generated by go-swagger; DO NOT EDIT.

package immutable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery --name API --keeptree --with-expecter --case underscore

// API is the interface of the immutable client
type API interface {
	/*
	   CreateImmuRule adds an immutable tag rule to current project

	   This endpoint add an immutable tag rule to the project
	*/
	CreateImmuRule(ctx context.Context, params *CreateImmuRuleParams) (*CreateImmuRuleCreated, error)
	/*
	   DeleteImmuRule deletes the immutable tag rule*/
	DeleteImmuRule(ctx context.Context, params *DeleteImmuRuleParams) (*DeleteImmuRuleOK, error)
	/*
	   ListImmuRules lists all immutable tag rules of current project

	   This endpoint returns the immutable tag rules of a project
	*/
	ListImmuRules(ctx context.Context, params *ListImmuRulesParams) (*ListImmuRulesOK, error)
	/*
	   UpdateImmuRule updates the immutable tag rule or enable or disable the rule*/
	UpdateImmuRule(ctx context.Context, params *UpdateImmuRuleParams) (*UpdateImmuRuleOK, error)
}

// New creates a new immutable API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for immutable API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateImmuRule adds an immutable tag rule to current project

This endpoint add an immutable tag rule to the project
*/
func (a *Client) CreateImmuRule(ctx context.Context, params *CreateImmuRuleParams) (*CreateImmuRuleCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateImmuRule",
		Method:             "POST",
		PathPattern:        "/projects/{project_name_or_id}/immutabletagrules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateImmuRuleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *CreateImmuRuleCreated:
		return value, nil
	case *CreateImmuRuleBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateImmuRuleUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateImmuRuleForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateImmuRuleNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateImmuRuleInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateImmuRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteImmuRule deletes the immutable tag rule
*/
func (a *Client) DeleteImmuRule(ctx context.Context, params *DeleteImmuRuleParams) (*DeleteImmuRuleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteImmuRule",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteImmuRuleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *DeleteImmuRuleOK:
		return value, nil
	case *DeleteImmuRuleBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DeleteImmuRuleUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DeleteImmuRuleForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DeleteImmuRuleInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteImmuRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListImmuRules lists all immutable tag rules of current project

This endpoint returns the immutable tag rules of a project
*/
func (a *Client) ListImmuRules(ctx context.Context, params *ListImmuRulesParams) (*ListImmuRulesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListImmuRules",
		Method:             "GET",
		PathPattern:        "/projects/{project_name_or_id}/immutabletagrules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListImmuRulesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *ListImmuRulesOK:
		return value, nil
	case *ListImmuRulesBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListImmuRulesUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListImmuRulesForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListImmuRulesInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListImmuRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateImmuRule updates the immutable tag rule or enable or disable the rule
*/
func (a *Client) UpdateImmuRule(ctx context.Context, params *UpdateImmuRuleParams) (*UpdateImmuRuleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateImmuRule",
		Method:             "PUT",
		PathPattern:        "/projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateImmuRuleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *UpdateImmuRuleOK:
		return value, nil
	case *UpdateImmuRuleBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateImmuRuleUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateImmuRuleForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateImmuRuleInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateImmuRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}
