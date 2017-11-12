// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddCommentToTaskParamsBody A comment to create
//
// These values can have github flavored markdown.
//
// swagger:model addCommentToTaskParamsBody
type AddCommentToTaskParamsBody struct {

	// content
	// Required: true
	Content *string `json:"content"`

	// user Id
	// Required: true
	UserID *int64 `json:"userId"`
}

/* polymorph addCommentToTaskParamsBody content false */

/* polymorph addCommentToTaskParamsBody userId false */

// Validate validates this add comment to task params body
func (m *AddCommentToTaskParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddCommentToTaskParamsBody) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *AddCommentToTaskParamsBody) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddCommentToTaskParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddCommentToTaskParamsBody) UnmarshalBinary(b []byte) error {
	var res AddCommentToTaskParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
