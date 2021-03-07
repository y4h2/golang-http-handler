package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Request defines customized http request struct
type Request struct {
	http.Request

	id string
}

// NewRequest is the constructor of Request
func NewRequest(r *http.Request) *Request {
	return &Request{
		Request: r,
		id:      uuid.NewString(),
	}
}

// ID returns request ID
func (req *Request) ID() string {
	return req.id
}

// ReadBody reads request body and write it back
func (req *Request) ReadBody() (buf []byte, err error) {
	defer req.Body.Close()
	buf, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read body")
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(buf))
	return
}

// ValidatePayload validates payload and decode the payload back to val
func (req *Request) ValidatePayload(val PayloadValidator) error {
	bodyBytes, err := req.ReadBody()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bodyBytes, &v); err != nil {
		return NewClientErr("Bad JSON body: " + err.Error())
	}

	if v == nil {
		return NewClientErr("JSON body is missing")
	}
	if err := val.Validate(); err != nil {
		return NewClientErr(http.StatusBadRequest, err.Error())
	}

	return nil
}

// PayloadValidator defines a request payload validator
type PayloadValidator interface {
	Validate() error
}
