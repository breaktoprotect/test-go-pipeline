// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new todos API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for todos API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTodo(params *CreateTodoParams, opts ...ClientOption) (*CreateTodoCreated, error)

	DeleteTodoByID(params *DeleteTodoByIDParams, opts ...ClientOption) (*DeleteTodoByIDOK, error)

	GetTodoByID(params *GetTodoByIDParams, opts ...ClientOption) (*GetTodoByIDOK, error)

	GetTodosList(params *GetTodosListParams, opts ...ClientOption) (*GetTodosListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateTodo adds a todo

  Optional
*/
func (a *Client) CreateTodo(params *CreateTodoParams, opts ...ClientOption) (*CreateTodoCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTodoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createTodo",
		Method:             "POST",
		PathPattern:        "/todos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTodoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateTodoCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createTodo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteTodoByID deletes an existing todo item

  Optional
*/
func (a *Client) DeleteTodoByID(params *DeleteTodoByIDParams, opts ...ClientOption) (*DeleteTodoByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTodoByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteTodoById",
		Method:             "DELETE",
		PathPattern:        "/todos/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTodoByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTodoByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTodoById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTodoByID returns a specific todo item

  Optional extended description in Markdown.
*/
func (a *Client) GetTodoByID(params *GetTodoByIDParams, opts ...ClientOption) (*GetTodoByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTodoByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTodoById",
		Method:             "GET",
		PathPattern:        "/todos/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTodoByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTodoByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTodoById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTodosList returns a list of todos

  Optional extended description in Markdown.
*/
func (a *Client) GetTodosList(params *GetTodosListParams, opts ...ClientOption) (*GetTodosListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTodosListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTodosList",
		Method:             "GET",
		PathPattern:        "/todos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTodosListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTodosListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTodosList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
