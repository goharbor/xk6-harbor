// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery --name API --keeptree --with-expecter --case underscore

// API is the interface of the user client
type API interface {
	/*
	   CreateUser creates a local user

	   This API can be used only when the authentication mode is for local DB.  When self registration is disabled.*/
	CreateUser(ctx context.Context, params *CreateUserParams) (*CreateUserCreated, error)
	/*
	   DeleteUser marks a registered user as be removed

	   This endpoint let administrator of Harbor mark a registered user as removed.It actually won't be deleted from DB.
	*/
	DeleteUser(ctx context.Context, params *DeleteUserParams) (*DeleteUserOK, error)
	/*
	   GetCurrentUserInfo gets current user info*/
	GetCurrentUserInfo(ctx context.Context, params *GetCurrentUserInfoParams) (*GetCurrentUserInfoOK, error)
	/*
	   GetCurrentUserPermissions gets current user permissions*/
	GetCurrentUserPermissions(ctx context.Context, params *GetCurrentUserPermissionsParams) (*GetCurrentUserPermissionsOK, error)
	/*
	   GetUser gets a user s profile*/
	GetUser(ctx context.Context, params *GetUserParams) (*GetUserOK, error)
	/*
	   ListUsers lists users*/
	ListUsers(ctx context.Context, params *ListUsersParams) (*ListUsersOK, error)
	/*
	   SearchUsers searches users by username

	   This endpoint is to search the users by username.  It's open for all authenticated requests.
	*/
	SearchUsers(ctx context.Context, params *SearchUsersParams) (*SearchUsersOK, error)
	/*
	   SetCliSecret sets c l i secret for a user

	   This endpoint let user generate a new CLI secret for himself.  This API only works when auth mode is set to 'OIDC'. Once this API returns with successful status, the old secret will be invalid, as there will be only one CLI secret for a user.*/
	SetCliSecret(ctx context.Context, params *SetCliSecretParams) (*SetCliSecretOK, error)
	/*
	   SetUserSysAdmin updates a registered user to change to be an administrator of harbor*/
	SetUserSysAdmin(ctx context.Context, params *SetUserSysAdminParams) (*SetUserSysAdminOK, error)
	/*
	   UpdateUserPassword changes the password on a user that already exists

	   This endpoint is for user to update password. Users with the admin role can change any user's password. Regular users can change only their own password.
	*/
	UpdateUserPassword(ctx context.Context, params *UpdateUserPasswordParams) (*UpdateUserPasswordOK, error)
	/*
	   UpdateUserProfile updates user s profile*/
	UpdateUserProfile(ctx context.Context, params *UpdateUserProfileParams) (*UpdateUserProfileOK, error)
}

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateUser creates a local user

This API can be used only when the authentication mode is for local DB.  When self registration is disabled.
*/
func (a *Client) CreateUser(ctx context.Context, params *CreateUserParams) (*CreateUserCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUser",
		Method:             "POST",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *CreateUserCreated:
		return value, nil
	case *CreateUserBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateUserUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateUserForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateUserConflict:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateUserInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteUser marks a registered user as be removed

This endpoint let administrator of Harbor mark a registered user as removed.It actually won't be deleted from DB.
*/
func (a *Client) DeleteUser(ctx context.Context, params *DeleteUserParams) (*DeleteUserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUser",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteUserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *DeleteUserOK:
		return value, nil
	case *DeleteUserUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DeleteUserForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DeleteUserNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DeleteUserInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCurrentUserInfo gets current user info
*/
func (a *Client) GetCurrentUserInfo(ctx context.Context, params *GetCurrentUserInfoParams) (*GetCurrentUserInfoOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCurrentUserInfo",
		Method:             "GET",
		PathPattern:        "/users/current",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCurrentUserInfoReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetCurrentUserInfoOK:
		return value, nil
	case *GetCurrentUserInfoUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetCurrentUserInfoInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCurrentUserInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCurrentUserPermissions gets current user permissions
*/
func (a *Client) GetCurrentUserPermissions(ctx context.Context, params *GetCurrentUserPermissionsParams) (*GetCurrentUserPermissionsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCurrentUserPermissions",
		Method:             "GET",
		PathPattern:        "/users/current/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCurrentUserPermissionsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetCurrentUserPermissionsOK:
		return value, nil
	case *GetCurrentUserPermissionsUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetCurrentUserPermissionsInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCurrentUserPermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUser gets a user s profile
*/
func (a *Client) GetUser(ctx context.Context, params *GetUserParams) (*GetUserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUser",
		Method:             "GET",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetUserOK:
		return value, nil
	case *GetUserUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetUserForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetUserNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetUserInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListUsers lists users
*/
func (a *Client) ListUsers(ctx context.Context, params *ListUsersParams) (*ListUsersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListUsersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *ListUsersOK:
		return value, nil
	case *ListUsersUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListUsersForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListUsersInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchUsers searches users by username

This endpoint is to search the users by username.  It's open for all authenticated requests.
*/
func (a *Client) SearchUsers(ctx context.Context, params *SearchUsersParams) (*SearchUsersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "searchUsers",
		Method:             "GET",
		PathPattern:        "/users/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchUsersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *SearchUsersOK:
		return value, nil
	case *SearchUsersUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *SearchUsersInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetCliSecret sets c l i secret for a user

This endpoint let user generate a new CLI secret for himself.  This API only works when auth mode is set to 'OIDC'. Once this API returns with successful status, the old secret will be invalid, as there will be only one CLI secret for a user.
*/
func (a *Client) SetCliSecret(ctx context.Context, params *SetCliSecretParams) (*SetCliSecretOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setCliSecret",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}/cli_secret",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetCliSecretReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *SetCliSecretOK:
		return value, nil
	case *SetCliSecretBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *SetCliSecretUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *SetCliSecretForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *SetCliSecretNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *SetCliSecretPreconditionFailed:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *SetCliSecretInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setCliSecret: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetUserSysAdmin updates a registered user to change to be an administrator of harbor
*/
func (a *Client) SetUserSysAdmin(ctx context.Context, params *SetUserSysAdminParams) (*SetUserSysAdminOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setUserSysAdmin",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}/sysadmin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetUserSysAdminReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *SetUserSysAdminOK:
		return value, nil
	case *SetUserSysAdminUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *SetUserSysAdminForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *SetUserSysAdminNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *SetUserSysAdminInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setUserSysAdmin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserPassword changes the password on a user that already exists

This endpoint is for user to update password. Users with the admin role can change any user's password. Regular users can change only their own password.
*/
func (a *Client) UpdateUserPassword(ctx context.Context, params *UpdateUserPasswordParams) (*UpdateUserPasswordOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserPassword",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUserPasswordReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *UpdateUserPasswordOK:
		return value, nil
	case *UpdateUserPasswordBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateUserPasswordUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateUserPasswordForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateUserPasswordInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserProfile updates user s profile
*/
func (a *Client) UpdateUserProfile(ctx context.Context, params *UpdateUserProfileParams) (*UpdateUserProfileOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserProfile",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUserProfileReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *UpdateUserProfileOK:
		return value, nil
	case *UpdateUserProfileUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateUserProfileForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateUserProfileNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateUserProfileInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}
