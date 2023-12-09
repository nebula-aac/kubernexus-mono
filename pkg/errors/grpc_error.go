package errors

import (
	"fmt"

	"github.com/nebula-aac/kubernexus-mono/pkg/errors"
)

var (
	ErrRequestInvalidCode = "1000"
	ErrPanicCode          = "1001"
	ErrGrpcListenerCode   = "1002"
	ErrGrpcServerCode     = "1003"
)

func ErrPanic(r interface{}) error {
	return errors.New(
		ErrPanicCode,
		errors.Alert,
		[]string{fmt.Sprintf("%v", r)},
		[]string{},
		[]string{},
		[]string{},
	)
}

func ErrGrpcListener(err error) error {
	return errors.New(
		ErrGrpcListenerCode,
		errors.Alert,
		[]string{"Error during gRPC listener initialization"},
		[]string{err.Error()},
		[]string{},
		[]string{},
	)
}

func ErrGrpcServer(err error) error {
	return errors.New(
		ErrGrpcServerCode,
		errors.Alert,
		[]string{"Error during gRPC server initialization"},
		[]string{err.Error()},
		[]string{},
		[]string{},
	)
}
