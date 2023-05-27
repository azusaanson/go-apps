package db

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
DEADLINE_EXCEEDED = 4
UNIMPLEMENTED = 12
INTERNAL = 13
UNAVAILABLE = 14
DATA_LOSS = 15
*/
func errorWithStatus(code codes.Code, err error) error {
	if err != nil {
		fmt.Printf("%+v\n", err) // print stack trace
	}

	return status.New(code, err.Error()).Err()
}
