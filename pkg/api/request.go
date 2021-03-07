package api

import (
	"net/http"

	"github.com/google/uuid"
)

// Request defines customized http request struct
type Request struct {
	http.Request

	id string
}

// NewRequest is the constructor of Request
func NewRequest() *Request {
	return &Request{
		id: uuid.NewString(),
	}
}

// ID returns request ID
func (req *Request) ID() string {
	return req.id
}

func (req *Request) ValidateBody(val PayloadValidator) error {
	if err := val.Validate(); err != nil {
		return NewClientErr(http.StatusBadRequest, err.Error())
	}

	return nil
}

type PayloadValidator interface {
	Validate() error
}
