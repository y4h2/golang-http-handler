package api

import (
	"net/http"

	"github.com/pkg/errors"
)

// ToClientError asserts error to ClientError
func ToClientError(err error) (ClientError, bool) {
	e, ok := errors.Cause(err).(ClientError)
	if !ok {
		return nil, false
	}
	return e, true
}

// ClientError defines the functions of an user facing error
type ClientError interface {
	error

	IsClientError()
	StatusCode() int
}

// NewNotFoundErr is the constructor of NotFoundErr
func NewNotFoundErr(message string) ClientErr {
	return NewClientErr(http.StatusNotFound, message)
}

type ClientErr struct {
	statusCode int
	message    string
}

func NewClientErr(statusCode int, message string) ClientErr {
	return ClientErr{
		statusCode: statusCode,
		message:    message,
	}
}

func (err ClientErr) StatusCode() int {
	return err.statusCode
}

// IsClientError marks NotFoundErr as ClientError
func (err ClientErr) IsClientError() {}

func (err ClientErr) Error() string {
	return err.message
}
