// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetImageParams creates a new GetImageParams object
// no default values defined in spec.
func NewGetImageParams() GetImageParams {

	return GetImageParams{}
}

// GetImageParams contains all the bound params for the get image operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetImage
type GetImageParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the image to retrieve
	  Required: true
	  In: path
	*/
	ImageID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetImageParams() beforehand.
func (o *GetImageParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rImageID, rhkImageID, _ := route.Params.GetOK("image_id")
	if err := o.bindImageID(rImageID, rhkImageID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindImageID binds and validates parameter ImageID from path.
func (o *GetImageParams) bindImageID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ImageID = raw

	return nil
}
