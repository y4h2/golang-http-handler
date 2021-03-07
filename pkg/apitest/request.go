package apitest

import (
	"io"
	"net/http/httptest"

	"github.com/y4h2/golang-http-handler/pkg/api"
)

func NewRequest(method string, target string, body io.Reader) *api.Request {
	req := httptest.NewRequest(method, target, body)

	return &api.Request{
		Request: *req,
	}
}
