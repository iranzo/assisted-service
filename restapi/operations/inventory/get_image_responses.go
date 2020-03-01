// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/filanov/bm-inventory/models"
)

// GetImageOKCode is the HTTP code returned for type GetImageOK
const GetImageOKCode int = 200

/*GetImageOK Image information

swagger:response getImageOK
*/
type GetImageOK struct {

	/*
	  In: Body
	*/
	Payload *models.Image `json:"body,omitempty"`
}

// NewGetImageOK creates GetImageOK with default headers values
func NewGetImageOK() *GetImageOK {

	return &GetImageOK{}
}

// WithPayload adds the payload to the get image o k response
func (o *GetImageOK) WithPayload(payload *models.Image) *GetImageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get image o k response
func (o *GetImageOK) SetPayload(payload *models.Image) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetImageNotFoundCode is the HTTP code returned for type GetImageNotFound
const GetImageNotFoundCode int = 404

/*GetImageNotFound Image not found

swagger:response getImageNotFound
*/
type GetImageNotFound struct {
}

// NewGetImageNotFound creates GetImageNotFound with default headers values
func NewGetImageNotFound() *GetImageNotFound {

	return &GetImageNotFound{}
}

// WriteResponse to the client
func (o *GetImageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
