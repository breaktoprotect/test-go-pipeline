// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetTodosParams creates a new GetTodosParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTodosParams() *GetTodosParams {
	return &GetTodosParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTodosParamsWithTimeout creates a new GetTodosParams object
// with the ability to set a timeout on a request.
func NewGetTodosParamsWithTimeout(timeout time.Duration) *GetTodosParams {
	return &GetTodosParams{
		timeout: timeout,
	}
}

// NewGetTodosParamsWithContext creates a new GetTodosParams object
// with the ability to set a context for a request.
func NewGetTodosParamsWithContext(ctx context.Context) *GetTodosParams {
	return &GetTodosParams{
		Context: ctx,
	}
}

// NewGetTodosParamsWithHTTPClient creates a new GetTodosParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTodosParamsWithHTTPClient(client *http.Client) *GetTodosParams {
	return &GetTodosParams{
		HTTPClient: client,
	}
}

/* GetTodosParams contains all the parameters to send to the API endpoint
   for the get todos operation.

   Typically these are written to a http.Request.
*/
type GetTodosParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get todos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTodosParams) WithDefaults() *GetTodosParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get todos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTodosParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get todos params
func (o *GetTodosParams) WithTimeout(timeout time.Duration) *GetTodosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get todos params
func (o *GetTodosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get todos params
func (o *GetTodosParams) WithContext(ctx context.Context) *GetTodosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get todos params
func (o *GetTodosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get todos params
func (o *GetTodosParams) WithHTTPClient(client *http.Client) *GetTodosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get todos params
func (o *GetTodosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTodosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}