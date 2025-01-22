// Code generated by go-swagger; DO NOT EDIT.

package definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SongVerse song verse
//
// swagger:model SongVerse
type SongVerse struct {

	// verse
	// Required: true
	Verse *string `json:"verse"`
}

// Validate validates this song verse
func (m *SongVerse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVerse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SongVerse) validateVerse(formats strfmt.Registry) error {

	if err := validate.Required("verse", "body", m.Verse); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this song verse based on context it is used
func (m *SongVerse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SongVerse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SongVerse) UnmarshalBinary(b []byte) error {
	var res SongVerse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
