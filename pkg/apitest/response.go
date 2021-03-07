package apitest

import (
	"net/http/httptest"

	"github.com/y4h2/golang-http-handler/pkg/api"
)

// ResponseRecorder extend *httptest.ResponseRecorder with api.ResponseWriter
type ResponseRecorder struct {
	*httptest.ResponseRecorder

	w api.ResponseWriter
}

// NewResponseRecorder is the constructor of ResponseRecorder
func NewResponseRecorder() *ResponseRecorder {
	recorder := httptest.NewRecorder()
	return &ResponseRecorder{recorder, api.NewResponse(recorder)}
}

// WriteJSON defines the mock function of api.ResponseWriter.WriteJSON
func (recorder *ResponseRecorder) WriteJSON(body interface{}) {
	recorder.w.WriteJSON(body)
}

// WriteError defines the mock function of api.ResponseWriter.WriteError
func (recorder *ResponseRecorder) WriteError(errorCode int, errorMessage string) {
	recorder.w.WriteError(errorCode, errorMessage)
}
