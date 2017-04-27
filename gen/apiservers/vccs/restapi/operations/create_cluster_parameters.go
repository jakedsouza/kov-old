package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/supervised-io/kov/gen/apiservers/vccs/models"
)

// NewCreateClusterParams creates a new CreateClusterParams object
// with the default values initialized.
func NewCreateClusterParams() CreateClusterParams {
	var ()
	return CreateClusterParams{}
}

// CreateClusterParams contains all the bound params for the create cluster operation
// typically these are obtained from a http.Request
//
// swagger:parameters createCluster
type CreateClusterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*A unique UUID for the request
	  Min Length: 1
	  In: header
	*/
	XRequestID *string
	/*the config of the cluster to be created
	  Required: true
	  In: body
	*/
	ClusterConfig *models.ClusterConfig
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CreateClusterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := o.bindXRequestID(r.Header[http.CanonicalHeaderKey("X-Request-Id")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ClusterConfig
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("clusterConfig", "body"))
			} else {
				res = append(res, errors.NewParseError("clusterConfig", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ClusterConfig = &body
			}
		}

	} else {
		res = append(res, errors.Required("clusterConfig", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateClusterParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XRequestID = &raw

	if err := o.validateXRequestID(formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateClusterParams) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.MinLength("X-Request-Id", "header", (*o.XRequestID), 1); err != nil {
		return err
	}

	return nil
}
