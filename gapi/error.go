package gapi

import (
	"errors"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrValidationUserNameRequired     = errors.New("validation: user name: required")
	ErrValidationUserPasswordRequired = errors.New("validation: user password: required")
	ErrValidationUserRoleRequired     = errors.New("validation: user role: required")
	ErrDuplicateUserName              = errors.New("duplicate: user name")
)

func invalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")

	statusDetails, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}

	return statusDetails.Err()
}

func errorWithCode(code codes.Code, err error) error {
	return status.New(code, err.Error()).Err()
}

func fieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}
