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

	models "github.com/Maksim646/TestAssignment/internal/api/definition"
)

// NewCreateSongParams creates a new CreateSongParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSongParams() *CreateSongParams {
	return &CreateSongParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSongParamsWithTimeout creates a new CreateSongParams object
// with the ability to set a timeout on a request.
func NewCreateSongParamsWithTimeout(timeout time.Duration) *CreateSongParams {
	return &CreateSongParams{
		timeout: timeout,
	}
}

// NewCreateSongParamsWithContext creates a new CreateSongParams object
// with the ability to set a context for a request.
func NewCreateSongParamsWithContext(ctx context.Context) *CreateSongParams {
	return &CreateSongParams{
		Context: ctx,
	}
}

// NewCreateSongParamsWithHTTPClient creates a new CreateSongParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSongParamsWithHTTPClient(client *http.Client) *CreateSongParams {
	return &CreateSongParams{
		HTTPClient: client,
	}
}

/*
CreateSongParams contains all the parameters to send to the API endpoint

	for the create song operation.

	Typically these are written to a http.Request.
*/
type CreateSongParams struct {

	/* CreateSong.

	   Create Song Body
	*/
	CreateSong *models.CreateSongBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create song params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSongParams) WithDefaults() *CreateSongParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create song params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSongParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create song params
func (o *CreateSongParams) WithTimeout(timeout time.Duration) *CreateSongParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create song params
func (o *CreateSongParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create song params
func (o *CreateSongParams) WithContext(ctx context.Context) *CreateSongParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create song params
func (o *CreateSongParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create song params
func (o *CreateSongParams) WithHTTPClient(client *http.Client) *CreateSongParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create song params
func (o *CreateSongParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateSong adds the createSong to the create song params
func (o *CreateSongParams) WithCreateSong(createSong *models.CreateSongBody) *CreateSongParams {
	o.SetCreateSong(createSong)
	return o
}

// SetCreateSong adds the createSong to the create song params
func (o *CreateSongParams) SetCreateSong(createSong *models.CreateSongBody) {
	o.CreateSong = createSong
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSongParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CreateSong != nil {
		if err := r.SetBodyParam(o.CreateSong); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
