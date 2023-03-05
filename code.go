package errors

import (
	stderrors "errors"
	"fmt"
)

type withCode struct {
	err   error
	code  int
	cause error
}

func WithCode(code int, format string, args ...interface{}) error {
	return &withCode{
		err:  fmt.Errorf(format, args...),
		code: code,
	}
}

func WrapC(err error, code int, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return &withCode{
		err:   fmt.Errorf(format, args...),
		code:  code,
		cause: err,
	}
}

// Error return the externally-safe error message.
func (w *withCode) Error() string { return w.Error() }

// Cause return the cause of the withCode error.
func (w *withCode) Cause() error { return w.cause }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *withCode) Unwrap() error { return w.cause }

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//	type causer interface {
//	       Cause() error
//	}
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}

		if cause.Cause() == nil {
			break
		}

		err = cause.Cause()
	}
	return err
}

func Is(err, target error) bool { return stderrors.Is(err, target) }
