// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetTodoByIDParams creates a new GetTodoByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTodoByIDParams() *GetTodoByIDParams {
	return &GetTodoByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTodoByIDParamsWithTimeout creates a new GetTodoByIDParams object
// with the ability to set a timeout on a request.
func NewGetTodoByIDParamsWithTimeout(timeout time.Duration) *GetTodoByIDParams {
	return &GetTodoByIDParams{
		timeout: timeout,
	}
}

// NewGetTodoByIDParamsWithContext creates a new GetTodoByIDParams object
// with the ability to set a context for a request.
func NewGetTodoByIDParamsWithContext(ctx context.Context) *GetTodoByIDParams {
	return &GetTodoByIDParams{
		Context: ctx,
	}
}

// NewGetTodoByIDParamsWithHTTPClient creates a new GetTodoByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTodoByIDParamsWithHTTPClient(client *http.Client) *GetTodoByIDParams {
	return &GetTodoByIDParams{
		HTTPClient: client,
	}
}

/* GetTodoByIDParams contains all the parameters to send to the API endpoint
   for the get todo by Id operation.

   Typically these are written to a http.Request.
*/
type GetTodoByIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get todo by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTodoByIDParams) WithDefaults() *GetTodoByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get todo by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTodoByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get todo by Id params
func (o *GetTodoByIDParams) WithTimeout(timeout time.Duration) *GetTodoByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get todo by Id params
func (o *GetTodoByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get todo by Id params
func (o *GetTodoByIDParams) WithContext(ctx context.Context) *GetTodoByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get todo by Id params
func (o *GetTodoByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get todo by Id params
func (o *GetTodoByIDParams) WithHTTPClient(client *http.Client) *GetTodoByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get todo by Id params
func (o *GetTodoByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get todo by Id params
func (o *GetTodoByIDParams) WithID(id string) *GetTodoByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get todo by Id params
func (o *GetTodoByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTodoByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
