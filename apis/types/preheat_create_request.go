// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PreheatCreateRequest Request option of creating a preheat task in supernode.
//
// swagger:model PreheatCreateRequest
type PreheatCreateRequest struct {

	// URL may contains some changeful query parameters such as authentication parameters. Dragonfly will
	// filter these parameter via 'filter'. The usage of it is that different URL may generate the same
	// download taskID.
	//
	Filter string `json:"filter,omitempty"`

	// If there is any authentication step of the remote server, the headers should contains authenticated information.
	// Dragonfly will sent request taking the headers to remote server.
	//
	Headers map[string]string `json:"headers,omitempty"`

	// This field is used for generating new downloading taskID to identify different downloading task of remote URL.
	//
	Identifier string `json:"identifier,omitempty"`

	// this must be image or file
	//
	// Required: true
	// Enum: [image file]
	Type *string `json:"type"`

	// the image or file location
	// Required: true
	// Min Length: 3
	URL *string `json:"url"`
}

// Validate validates this preheat create request
func (m *PreheatCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var preheatCreateRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["image","file"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		preheatCreateRequestTypeTypePropEnum = append(preheatCreateRequestTypeTypePropEnum, v)
	}
}

const (

	// PreheatCreateRequestTypeImage captures enum value "image"
	PreheatCreateRequestTypeImage string = "image"

	// PreheatCreateRequestTypeFile captures enum value "file"
	PreheatCreateRequestTypeFile string = "file"
)

// prop value enum
func (m *PreheatCreateRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, preheatCreateRequestTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PreheatCreateRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *PreheatCreateRequest) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.MinLength("url", "body", string(*m.URL), 3); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PreheatCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PreheatCreateRequest) UnmarshalBinary(b []byte) error {
	var res PreheatCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
