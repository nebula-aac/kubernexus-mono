package errors

import (
	"fmt"
)

var (
	ErrRequestInvalidCode = "1000"
	ErrPanicCode          = "1001"
	ErrGrpcListenerCode   = "1002"
	ErrGrpcServerCode     = "1003"
)

func ErrPanic(r interface{}) error {
	return New(
		ErrPanicCode,
		Alert,
		[]string{fmt.Sprintf("%v", r)},
		[]string{},
		[]string{},
		[]string{},
	)
}

func ErrGrpcListener(err error) error {
	return New(
		ErrGrpcListenerCode,
		Alert,
		[]string{"Error during gRPC listener initialization"},
		[]string{err.Error()},
		[]string{},
		[]string{},
	)
}

func ErrGrpcServer(err error) error {
	return New(
		ErrGrpcServerCode,
		Alert,
		[]string{"Error during gRPC server initialization"},
		[]string{err.Error()},
		[]string{},
		[]string{},
	)
}
