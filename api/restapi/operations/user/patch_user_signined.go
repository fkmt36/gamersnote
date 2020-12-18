// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"gamersnote.com/v1/models"
)

// PatchUserSigninedHandlerFunc turns a function with the right signature into a patch user signined handler
type PatchUserSigninedHandlerFunc func(PatchUserSigninedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchUserSigninedHandlerFunc) Handle(params PatchUserSigninedParams) middleware.Responder {
	return fn(params)
}

// PatchUserSigninedHandler interface for that can handle valid patch user signined params
type PatchUserSigninedHandler interface {
	Handle(PatchUserSigninedParams) middleware.Responder
}

// NewPatchUserSignined creates a new http.Handler for the patch user signined operation
func NewPatchUserSignined(ctx *middleware.Context, handler PatchUserSigninedHandler) *PatchUserSignined {
	return &PatchUserSignined{Context: ctx, Handler: handler}
}

/*PatchUserSignined swagger:route PATCH /users/me/signined user patchUserSignined

PatchUserSignined patch user signined API

*/
type PatchUserSignined struct {
	Context *middleware.Context
	Handler PatchUserSigninedHandler
}

func (o *PatchUserSignined) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchUserSigninedParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PatchUserSigninedBody patch user signined body
//
// swagger:model PatchUserSigninedBody
type PatchUserSigninedBody struct {

	// email
	// Required: true
	Email models.Email `json:"email"`

	// password
	// Required: true
	Password models.Password `json:"password"`
}

// Validate validates this patch user signined body
func (o *PatchUserSigninedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchUserSigninedBody) validateEmail(formats strfmt.Registry) error {

	if err := o.Email.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "email")
		}
		return err
	}

	return nil
}

func (o *PatchUserSigninedBody) validatePassword(formats strfmt.Registry) error {

	if err := o.Password.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "password")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchUserSigninedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchUserSigninedBody) UnmarshalBinary(b []byte) error {
	var res PatchUserSigninedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}