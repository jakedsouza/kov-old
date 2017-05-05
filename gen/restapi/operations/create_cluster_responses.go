///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/supervised-io/kov/gen/models"
)

// CreateClusterAcceptedCode is the HTTP code returned for type CreateClusterAccepted
const CreateClusterAcceptedCode int = 202

/*CreateClusterAccepted create cluster task has been accepted

swagger:response createClusterAccepted
*/
type CreateClusterAccepted struct {

	/*
	  In: Body
	*/
	Payload models.TaskID `json:"body,omitempty"`
}

// NewCreateClusterAccepted creates CreateClusterAccepted with default headers values
func NewCreateClusterAccepted() *CreateClusterAccepted {
	return &CreateClusterAccepted{}
}

// WithPayload adds the payload to the create cluster accepted response
func (o *CreateClusterAccepted) WithPayload(payload models.TaskID) *CreateClusterAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cluster accepted response
func (o *CreateClusterAccepted) SetPayload(payload models.TaskID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateClusterAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// CreateClusterConflictCode is the HTTP code returned for type CreateClusterConflict
const CreateClusterConflictCode int = 409

/*CreateClusterConflict The provided cluster name already exists

swagger:response createClusterConflict
*/
type CreateClusterConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateClusterConflict creates CreateClusterConflict with default headers values
func NewCreateClusterConflict() *CreateClusterConflict {
	return &CreateClusterConflict{}
}

// WithPayload adds the payload to the create cluster conflict response
func (o *CreateClusterConflict) WithPayload(payload *models.Error) *CreateClusterConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cluster conflict response
func (o *CreateClusterConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateClusterConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateClusterDefault Error

swagger:response createClusterDefault
*/
type CreateClusterDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateClusterDefault creates CreateClusterDefault with default headers values
func NewCreateClusterDefault(code int) *CreateClusterDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateClusterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create cluster default response
func (o *CreateClusterDefault) WithStatusCode(code int) *CreateClusterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create cluster default response
func (o *CreateClusterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create cluster default response
func (o *CreateClusterDefault) WithPayload(payload *models.Error) *CreateClusterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create cluster default response
func (o *CreateClusterDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateClusterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
