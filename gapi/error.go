package gapi

import "errors"

var (
	ErrValidationUserNameRequired     = errors.New("validation: user name: required")
	ErrValidationUserPasswordRequired = errors.New("validation: user password: required")
	ErrValidationUserRoleRequired     = errors.New("validation: user role: required")
)
