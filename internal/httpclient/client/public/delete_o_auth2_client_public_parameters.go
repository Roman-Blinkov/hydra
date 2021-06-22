// Code generated by go-swagger; DO NOT EDIT.

package public

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

// NewDeleteOAuth2ClientPublicParams creates a new DeleteOAuth2ClientPublicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOAuth2ClientPublicParams() *DeleteOAuth2ClientPublicParams {
	return &DeleteOAuth2ClientPublicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOAuth2ClientPublicParamsWithTimeout creates a new DeleteOAuth2ClientPublicParams object
// with the ability to set a timeout on a request.
func NewDeleteOAuth2ClientPublicParamsWithTimeout(timeout time.Duration) *DeleteOAuth2ClientPublicParams {
	return &DeleteOAuth2ClientPublicParams{
		timeout: timeout,
	}
}

// NewDeleteOAuth2ClientPublicParamsWithContext creates a new DeleteOAuth2ClientPublicParams object
// with the ability to set a context for a request.
func NewDeleteOAuth2ClientPublicParamsWithContext(ctx context.Context) *DeleteOAuth2ClientPublicParams {
	return &DeleteOAuth2ClientPublicParams{
		Context: ctx,
	}
}

// NewDeleteOAuth2ClientPublicParamsWithHTTPClient creates a new DeleteOAuth2ClientPublicParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOAuth2ClientPublicParamsWithHTTPClient(client *http.Client) *DeleteOAuth2ClientPublicParams {
	return &DeleteOAuth2ClientPublicParams{
		HTTPClient: client,
	}
}

/* DeleteOAuth2ClientPublicParams contains all the parameters to send to the API endpoint
   for the delete o auth2 client public operation.

   Typically these are written to a http.Request.
*/
type DeleteOAuth2ClientPublicParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete o auth2 client public params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOAuth2ClientPublicParams) WithDefaults() *DeleteOAuth2ClientPublicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete o auth2 client public params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOAuth2ClientPublicParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete o auth2 client public params
func (o *DeleteOAuth2ClientPublicParams) WithTimeout(timeout time.Duration) *DeleteOAuth2ClientPublicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete o auth2 client public params
func (o *DeleteOAuth2ClientPublicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete o auth2 client public params
func (o *DeleteOAuth2ClientPublicParams) WithContext(ctx context.Context) *DeleteOAuth2ClientPublicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete o auth2 client public params
func (o *DeleteOAuth2ClientPublicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete o auth2 client public params
func (o *DeleteOAuth2ClientPublicParams) WithHTTPClient(client *http.Client) *DeleteOAuth2ClientPublicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete o auth2 client public params
func (o *DeleteOAuth2ClientPublicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOAuth2ClientPublicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
