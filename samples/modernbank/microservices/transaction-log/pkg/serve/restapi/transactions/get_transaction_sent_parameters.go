// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTransactionSentParams creates a new GetTransactionSentParams object
// no default values defined in spec.
func NewGetTransactionSentParams() GetTransactionSentParams {

	return GetTransactionSentParams{}
}

// GetTransactionSentParams contains all the bound params for the get transaction sent operation
// typically these are obtained from a http.Request
//
// swagger:parameters getTransactionSent
type GetTransactionSentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Account number that made the transactions
	  Required: true
	  In: path
	*/
	Sender int64
	/*Transaction ID
	  Required: true
	  In: path
	*/
	Transaction string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTransactionSentParams() beforehand.
func (o *GetTransactionSentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSender, rhkSender, _ := route.Params.GetOK("sender")
	if err := o.bindSender(rSender, rhkSender, route.Formats); err != nil {
		res = append(res, err)
	}

	rTransaction, rhkTransaction, _ := route.Params.GetOK("transaction")
	if err := o.bindTransaction(rTransaction, rhkTransaction, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSender binds and validates parameter Sender from path.
func (o *GetTransactionSentParams) bindSender(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("sender", "path", "int64", raw)
	}
	o.Sender = value

	return nil
}

// bindTransaction binds and validates parameter Transaction from path.
func (o *GetTransactionSentParams) bindTransaction(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Transaction = raw

	return nil
}