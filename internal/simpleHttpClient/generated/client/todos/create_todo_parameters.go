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

	"test-go-pipeline/internal/simplehttpclient/generated/models"
)

// NewCreateTodoParams creates a new CreateTodoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTodoParams() *CreateTodoParams {
	return &CreateTodoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTodoParamsWithTimeout creates a new CreateTodoParams object
// with the ability to set a timeout on a request.
func NewCreateTodoParamsWithTimeout(timeout time.Duration) *CreateTodoParams {
	return &CreateTodoParams{
		timeout: timeout,
	}
}

// NewCreateTodoParamsWithContext creates a new CreateTodoParams object
// with the ability to set a context for a request.
func NewCreateTodoParamsWithContext(ctx context.Context) *CreateTodoParams {
	return &CreateTodoParams{
		Context: ctx,
	}
}

// NewCreateTodoParamsWithHTTPClient creates a new CreateTodoParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTodoParamsWithHTTPClient(client *http.Client) *CreateTodoParams {
	return &CreateTodoParams{
		HTTPClient: client,
	}
}

/* CreateTodoParams contains all the parameters to send to the API endpoint
   for the create todo operation.

   Typically these are written to a http.Request.
*/
type CreateTodoParams struct {

	// Body.
	Body *models.Item

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create todo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTodoParams) WithDefaults() *CreateTodoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create todo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTodoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create todo params
func (o *CreateTodoParams) WithTimeout(timeout time.Duration) *CreateTodoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create todo params
func (o *CreateTodoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create todo params
func (o *CreateTodoParams) WithContext(ctx context.Context) *CreateTodoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create todo params
func (o *CreateTodoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create todo params
func (o *CreateTodoParams) WithHTTPClient(client *http.Client) *CreateTodoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create todo params
func (o *CreateTodoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create todo params
func (o *CreateTodoParams) WithBody(body *models.Item) *CreateTodoParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create todo params
func (o *CreateTodoParams) SetBody(body *models.Item) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTodoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
