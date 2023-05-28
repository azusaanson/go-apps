package gapi

import (
	"errors"
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrValidationUserNameRequired     = errors.New("validation: user name: required")
	ErrValidationUserPasswordRequired = errors.New("validation: user password: required")
	ErrValidationUserRoleRequired     = errors.New("validation: user role: required")
	ErrDuplicateUserName              = errors.New("duplicate: user name")
	ErrNotFoundUser                   = errors.New("not found: user")
)

// INVALID_ARGUMENT = 3
func invalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")

	statusDetails, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}

	return statusDetails.Err()
}

/*
NOT_FOUND = 5
ALREADY_EXISTS = 6
PERMISSION_DENIED = 7
RESOURCE_EXHAUSTED = 8
FAILED_PRECONDITION = 9
ABORTED = 10
OUT_OF_RANGE = 11
UNAUTHENTICATED = 16
*/
func clientError(code codes.Code, err error) error {
	if err != nil {
		fmt.Printf("%+v\n", err) // print stack trace
	}

	return status.New(code, err.Error()).Err()
}

/*
DEADLINE_EXCEEDED = 4
UNIMPLEMENTED = 12
INTERNAL = 13
UNAVAILABLE = 14
DATA_LOSS = 15
*/
func serverError(err error) error {
	if err != nil {
		fmt.Printf("%+v\n", err) // print stack trace
	}

	if _, ok := status.FromError(err); !ok {
		status.New(codes.Internal, err.Error()).Err()
	}

	return err
}

func fieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}
