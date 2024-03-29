// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery --name API --keeptree --with-expecter --case underscore

// API is the interface of the member client
type API interface {
	/*
	   CreateProjectMember creates project member

	   Create project member relationship, the member can be one of the user_member and group_member,  The user_member need to specify user_id or username. If the user already exist in harbor DB, specify the user_id,  If does not exist in harbor DB, it will SearchAndOnBoard the user. The group_member need to specify id or ldap_group_dn. If the group already exist in harbor DB. specify the user group's id,  If does not exist, it will SearchAndOnBoard the group. */
	CreateProjectMember(ctx context.Context, params *CreateProjectMemberParams) (*CreateProjectMemberCreated, error)
	/*
	   DeleteProjectMember deletes project member*/
	DeleteProjectMember(ctx context.Context, params *DeleteProjectMemberParams) (*DeleteProjectMemberOK, error)
	/*
	   GetProjectMember gets the project member information

	   Get the project member information*/
	GetProjectMember(ctx context.Context, params *GetProjectMemberParams) (*GetProjectMemberOK, error)
	/*
	   ListProjectMembers gets all project member information

	   Get all project member information*/
	ListProjectMembers(ctx context.Context, params *ListProjectMembersParams) (*ListProjectMembersOK, error)
	/*
	   UpdateProjectMember updates project member

	   Update project member relationship*/
	UpdateProjectMember(ctx context.Context, params *UpdateProjectMemberParams) (*UpdateProjectMemberOK, error)
}

// New creates a new member API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for member API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateProjectMember creates project member

Create project member relationship, the member can be one of the user_member and group_member,  The user_member need to specify user_id or username. If the user already exist in harbor DB, specify the user_id,  If does not exist in harbor DB, it will SearchAndOnBoard the user. The group_member need to specify id or ldap_group_dn. If the group already exist in harbor DB. specify the user group's id,  If does not exist, it will SearchAndOnBoard the group.
*/
func (a *Client) CreateProjectMember(ctx context.Context, params *CreateProjectMemberParams) (*CreateProjectMemberCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createProjectMember",
		Method:             "POST",
		PathPattern:        "/projects/{project_name_or_id}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateProjectMemberReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *CreateProjectMemberCreated:
		return value, nil
	case *CreateProjectMemberBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateProjectMemberUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateProjectMemberForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateProjectMemberConflict:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateProjectMemberInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createProjectMember: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteProjectMember deletes project member
*/
func (a *Client) DeleteProjectMember(ctx context.Context, params *DeleteProjectMemberParams) (*DeleteProjectMemberOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProjectMember",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_name_or_id}/members/{mid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteProjectMemberReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *DeleteProjectMemberOK:
		return value, nil
	case *DeleteProjectMemberBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DeleteProjectMemberUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DeleteProjectMemberForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DeleteProjectMemberInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteProjectMember: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetProjectMember gets the project member information

Get the project member information
*/
func (a *Client) GetProjectMember(ctx context.Context, params *GetProjectMemberParams) (*GetProjectMemberOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProjectMember",
		Method:             "GET",
		PathPattern:        "/projects/{project_name_or_id}/members/{mid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetProjectMemberReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetProjectMemberOK:
		return value, nil
	case *GetProjectMemberBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetProjectMemberUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetProjectMemberForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetProjectMemberNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetProjectMemberInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProjectMember: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListProjectMembers gets all project member information

Get all project member information
*/
func (a *Client) ListProjectMembers(ctx context.Context, params *ListProjectMembersParams) (*ListProjectMembersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listProjectMembers",
		Method:             "GET",
		PathPattern:        "/projects/{project_name_or_id}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListProjectMembersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *ListProjectMembersOK:
		return value, nil
	case *ListProjectMembersBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListProjectMembersUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListProjectMembersForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListProjectMembersNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListProjectMembersInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listProjectMembers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateProjectMember updates project member

Update project member relationship
*/
func (a *Client) UpdateProjectMember(ctx context.Context, params *UpdateProjectMemberParams) (*UpdateProjectMemberOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProjectMember",
		Method:             "PUT",
		PathPattern:        "/projects/{project_name_or_id}/members/{mid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateProjectMemberReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *UpdateProjectMemberOK:
		return value, nil
	case *UpdateProjectMemberBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateProjectMemberUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateProjectMemberForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateProjectMemberNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateProjectMemberInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateProjectMember: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}
