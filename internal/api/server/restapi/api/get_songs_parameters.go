// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetSongsParams creates a new GetSongsParams object
//
// There are no default values defined in the spec.
func NewGetSongsParams() GetSongsParams {

	return GetSongsParams{}
}

// GetSongsParams contains all the bound params for the get songs operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetSongs
type GetSongsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Filter Song By Group
	  In: query
	*/
	FilterSongByGroup *string
	/*Filter Song By Link
	  In: query
	*/
	FilterSongByLink *string
	/*Filter Song By Name
	  In: query
	*/
	FilterSongByName *string
	/*Filter Song By Release Date
	  In: query
	*/
	FilterSongByReleaseDate *string
	/*Filter Song By Text
	  In: query
	*/
	FilterSongByText *string
	/*Offset Configs
	  Required: true
	  Minimum: 0
	  In: query
	*/
	Limit int64
	/*Offset Configs
	  Required: true
	  Minimum: 0
	  In: query
	*/
	Offset int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSongsParams() beforehand.
func (o *GetSongsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFilterSongByGroup, qhkFilterSongByGroup, _ := qs.GetOK("FilterSongByGroup")
	if err := o.bindFilterSongByGroup(qFilterSongByGroup, qhkFilterSongByGroup, route.Formats); err != nil {
		res = append(res, err)
	}

	qFilterSongByLink, qhkFilterSongByLink, _ := qs.GetOK("FilterSongByLink")
	if err := o.bindFilterSongByLink(qFilterSongByLink, qhkFilterSongByLink, route.Formats); err != nil {
		res = append(res, err)
	}

	qFilterSongByName, qhkFilterSongByName, _ := qs.GetOK("FilterSongByName")
	if err := o.bindFilterSongByName(qFilterSongByName, qhkFilterSongByName, route.Formats); err != nil {
		res = append(res, err)
	}

	qFilterSongByReleaseDate, qhkFilterSongByReleaseDate, _ := qs.GetOK("FilterSongByReleaseDate")
	if err := o.bindFilterSongByReleaseDate(qFilterSongByReleaseDate, qhkFilterSongByReleaseDate, route.Formats); err != nil {
		res = append(res, err)
	}

	qFilterSongByText, qhkFilterSongByText, _ := qs.GetOK("FilterSongByText")
	if err := o.bindFilterSongByText(qFilterSongByText, qhkFilterSongByText, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFilterSongByGroup binds and validates parameter FilterSongByGroup from query.
func (o *GetSongsParams) bindFilterSongByGroup(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.FilterSongByGroup = &raw

	return nil
}

// bindFilterSongByLink binds and validates parameter FilterSongByLink from query.
func (o *GetSongsParams) bindFilterSongByLink(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.FilterSongByLink = &raw

	return nil
}

// bindFilterSongByName binds and validates parameter FilterSongByName from query.
func (o *GetSongsParams) bindFilterSongByName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.FilterSongByName = &raw

	return nil
}

// bindFilterSongByReleaseDate binds and validates parameter FilterSongByReleaseDate from query.
func (o *GetSongsParams) bindFilterSongByReleaseDate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.FilterSongByReleaseDate = &raw

	return nil
}

// bindFilterSongByText binds and validates parameter FilterSongByText from query.
func (o *GetSongsParams) bindFilterSongByText(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.FilterSongByText = &raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetSongsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("limit", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("limit", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *GetSongsParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", o.Limit, 0, false); err != nil {
		return err
	}

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *GetSongsParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("offset", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("offset", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = value

	if err := o.validateOffset(formats); err != nil {
		return err
	}

	return nil
}

// validateOffset carries on validations for parameter Offset
func (o *GetSongsParams) validateOffset(formats strfmt.Registry) error {

	if err := validate.MinimumInt("offset", "query", o.Offset, 0, false); err != nil {
		return err
	}

	return nil
}
