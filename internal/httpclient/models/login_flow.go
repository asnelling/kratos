// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LoginFlow Login Flow
//
// This object represents a login flow. A login flow is initiated at the "Initiate Login API / Browser Flow"
// endpoint by a client.
//
// Once a login flow is completed successfully, a session cookie or session token will be issued.
//
// swagger:model loginFlow
type LoginFlow struct {

	// active
	Active CredentialsType `json:"active,omitempty"`

	// ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to log in,
	// a new flow has to be initiated.
	// Required: true
	// Format: date-time
	ExpiresAt *strfmt.DateTime `json:"expires_at"`

	// Forced stores whether this login flow should enforce re-authentication.
	Forced bool `json:"forced,omitempty"`

	// id
	// Required: true
	// Format: uuid4
	ID UUID `json:"id"`

	// IssuedAt is the time (UTC) when the flow started.
	// Required: true
	// Format: date-time
	IssuedAt *strfmt.DateTime `json:"issued_at"`

	// messages
	Messages Messages `json:"messages,omitempty"`

	// List of login methods
	//
	// This is the list of available login methods with their required form fields, such as `identifier` and `password`
	// for the password login method. This will also contain error messages such as "password can not be empty".
	// Required: true
	Methods map[string]LoginFlowMethod `json:"methods"`

	// RequestURL is the initial URL that was requested from ORY Kratos. It can be used
	// to forward information contained in the URL's path or query for example.
	// Required: true
	RequestURL *string `json:"request_url"`

	// type
	Type Type `json:"type,omitempty"`
}

// Validate validates this login flow
func (m *LoginFlow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginFlow) validateActive(formats strfmt.Registry) error {

	if swag.IsZero(m.Active) { // not required
		return nil
	}

	if err := m.Active.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("active")
		}
		return err
	}

	return nil
}

func (m *LoginFlow) validateExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("expires_at", "body", m.ExpiresAt); err != nil {
		return err
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LoginFlow) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *LoginFlow) validateIssuedAt(formats strfmt.Registry) error {

	if err := validate.Required("issued_at", "body", m.IssuedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("issued_at", "body", "date-time", m.IssuedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LoginFlow) validateMessages(formats strfmt.Registry) error {

	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	if err := m.Messages.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("messages")
		}
		return err
	}

	return nil
}

func (m *LoginFlow) validateMethods(formats strfmt.Registry) error {

	for k := range m.Methods {

		if err := validate.Required("methods"+"."+k, "body", m.Methods[k]); err != nil {
			return err
		}
		if val, ok := m.Methods[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *LoginFlow) validateRequestURL(formats strfmt.Registry) error {

	if err := validate.Required("request_url", "body", m.RequestURL); err != nil {
		return err
	}

	return nil
}

func (m *LoginFlow) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoginFlow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginFlow) UnmarshalBinary(b []byte) error {
	var res LoginFlow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
