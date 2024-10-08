// Code generated by go-swagger; DO NOT EDIT.

package blog_post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/john-k-antony/richpanel-blogposts-golang/models/apimodels"
)

// UpdateBlogpostByIDOKCode is the HTTP code returned for type UpdateBlogpostByIDOK
const UpdateBlogpostByIDOKCode int = 200

/*
UpdateBlogpostByIDOK Blog post updated

swagger:response updateBlogpostByIdOK
*/
type UpdateBlogpostByIDOK struct {

	/*
	  In: Body
	*/
	Payload *apimodels.Blogpost `json:"body,omitempty"`
}

// NewUpdateBlogpostByIDOK creates UpdateBlogpostByIDOK with default headers values
func NewUpdateBlogpostByIDOK() *UpdateBlogpostByIDOK {

	return &UpdateBlogpostByIDOK{}
}

// WithPayload adds the payload to the update blogpost by Id o k response
func (o *UpdateBlogpostByIDOK) WithPayload(payload *apimodels.Blogpost) *UpdateBlogpostByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update blogpost by Id o k response
func (o *UpdateBlogpostByIDOK) SetPayload(payload *apimodels.Blogpost) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBlogpostByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBlogpostByIDBadRequestCode is the HTTP code returned for type UpdateBlogpostByIDBadRequest
const UpdateBlogpostByIDBadRequestCode int = 400

/*
UpdateBlogpostByIDBadRequest Bad request

swagger:response updateBlogpostByIdBadRequest
*/
type UpdateBlogpostByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ResponseError `json:"body,omitempty"`
}

// NewUpdateBlogpostByIDBadRequest creates UpdateBlogpostByIDBadRequest with default headers values
func NewUpdateBlogpostByIDBadRequest() *UpdateBlogpostByIDBadRequest {

	return &UpdateBlogpostByIDBadRequest{}
}

// WithPayload adds the payload to the update blogpost by Id bad request response
func (o *UpdateBlogpostByIDBadRequest) WithPayload(payload *apimodels.ResponseError) *UpdateBlogpostByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update blogpost by Id bad request response
func (o *UpdateBlogpostByIDBadRequest) SetPayload(payload *apimodels.ResponseError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBlogpostByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBlogpostByIDUnauthorizedCode is the HTTP code returned for type UpdateBlogpostByIDUnauthorized
const UpdateBlogpostByIDUnauthorizedCode int = 401

/*
UpdateBlogpostByIDUnauthorized Unauthenticated Request

swagger:response updateBlogpostByIdUnauthorized
*/
type UpdateBlogpostByIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ResponseError `json:"body,omitempty"`
}

// NewUpdateBlogpostByIDUnauthorized creates UpdateBlogpostByIDUnauthorized with default headers values
func NewUpdateBlogpostByIDUnauthorized() *UpdateBlogpostByIDUnauthorized {

	return &UpdateBlogpostByIDUnauthorized{}
}

// WithPayload adds the payload to the update blogpost by Id unauthorized response
func (o *UpdateBlogpostByIDUnauthorized) WithPayload(payload *apimodels.ResponseError) *UpdateBlogpostByIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update blogpost by Id unauthorized response
func (o *UpdateBlogpostByIDUnauthorized) SetPayload(payload *apimodels.ResponseError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBlogpostByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBlogpostByIDNotFoundCode is the HTTP code returned for type UpdateBlogpostByIDNotFound
const UpdateBlogpostByIDNotFoundCode int = 404

/*
UpdateBlogpostByIDNotFound Blog post not found

swagger:response updateBlogpostByIdNotFound
*/
type UpdateBlogpostByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ResponseError `json:"body,omitempty"`
}

// NewUpdateBlogpostByIDNotFound creates UpdateBlogpostByIDNotFound with default headers values
func NewUpdateBlogpostByIDNotFound() *UpdateBlogpostByIDNotFound {

	return &UpdateBlogpostByIDNotFound{}
}

// WithPayload adds the payload to the update blogpost by Id not found response
func (o *UpdateBlogpostByIDNotFound) WithPayload(payload *apimodels.ResponseError) *UpdateBlogpostByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update blogpost by Id not found response
func (o *UpdateBlogpostByIDNotFound) SetPayload(payload *apimodels.ResponseError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBlogpostByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBlogpostByIDUnprocessableEntityCode is the HTTP code returned for type UpdateBlogpostByIDUnprocessableEntity
const UpdateBlogpostByIDUnprocessableEntityCode int = 422

/*
UpdateBlogpostByIDUnprocessableEntity Invalid request

swagger:response updateBlogpostByIdUnprocessableEntity
*/
type UpdateBlogpostByIDUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ResponseError `json:"body,omitempty"`
}

// NewUpdateBlogpostByIDUnprocessableEntity creates UpdateBlogpostByIDUnprocessableEntity with default headers values
func NewUpdateBlogpostByIDUnprocessableEntity() *UpdateBlogpostByIDUnprocessableEntity {

	return &UpdateBlogpostByIDUnprocessableEntity{}
}

// WithPayload adds the payload to the update blogpost by Id unprocessable entity response
func (o *UpdateBlogpostByIDUnprocessableEntity) WithPayload(payload *apimodels.ResponseError) *UpdateBlogpostByIDUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update blogpost by Id unprocessable entity response
func (o *UpdateBlogpostByIDUnprocessableEntity) SetPayload(payload *apimodels.ResponseError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBlogpostByIDUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
