// errwrapper package contains errWrapper struct which wrap errors + code(httpStatusCode).
// Mainly used in service layer as return type
package errwrapper

import "fmt"

type ErrWrapper interface {
	error

	// getter method for errs field
	Errors() []error

	// getter method for code field
	Code() int
}

type errWrapper struct {
	errs []error
	code    int // httpCode to determine the status (unauthorized, server error, etc)
}

// constructor of apiError struct
func New(httpCode int, errs ...error) ErrWrapper{
	return &errWrapper{
		code : httpCode,
		errs : errs,
	}
}

func (a *errWrapper) Code() int {
	return a.code
}

func (a *errWrapper) Errors() []error {
	return a.errs
}

func (a *errWrapper) Error() string {
	return fmt.Sprint(a.errs)
}