package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListClustersParams creates a new ListClustersParams object
// with the default values initialized.
func NewListClustersParams() *ListClustersParams {
	var ()
	return &ListClustersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListClustersParamsWithTimeout creates a new ListClustersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListClustersParamsWithTimeout(timeout time.Duration) *ListClustersParams {
	var ()
	return &ListClustersParams{

		timeout: timeout,
	}
}

// NewListClustersParamsWithContext creates a new ListClustersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListClustersParamsWithContext(ctx context.Context) *ListClustersParams {
	var ()
	return &ListClustersParams{

		Context: ctx,
	}
}

// NewListClustersParamsWithHTTPClient creates a new ListClustersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListClustersParamsWithHTTPClient(client *http.Client) *ListClustersParams {
	var ()
	return &ListClustersParams{
		HTTPClient: client,
	}
}

/*ListClustersParams contains all the parameters to send to the API endpoint
for the list clusters operation typically these are written to a http.Request
*/
type ListClustersParams struct {

	/*XRequestID
	  A unique UUID for the request

	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list clusters params
func (o *ListClustersParams) WithTimeout(timeout time.Duration) *ListClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list clusters params
func (o *ListClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list clusters params
func (o *ListClustersParams) WithContext(ctx context.Context) *ListClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list clusters params
func (o *ListClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list clusters params
func (o *ListClustersParams) WithHTTPClient(client *http.Client) *ListClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list clusters params
func (o *ListClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list clusters params
func (o *ListClustersParams) WithXRequestID(xRequestID *string) *ListClustersParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list clusters params
func (o *ListClustersParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *ListClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
