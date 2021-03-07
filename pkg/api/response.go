package api

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// ResponseWriter is the extension of
type ResponseWriter interface {
	http.ResponseWriter

	WriteJSON(body interface{}) (int, error)
	WriteError(statusCode int, message string) (int, error)
}

// Response extends http.ResponseWriter to provide handy functions
type Response struct {
	http.ResponseWriter
}

// NewResponse is the constructor of Response
func NewResponse(w http.ResponseWriter) ResponseWriter {
	return Response{w}
}

// WriteHeader wraps http.ResponseWriter WriteHeader
// sets content-type to json
func (resp Response) WriteHeader(statusCode int) {
	resp.Header().Set("Content-Type", "application/json")
	resp.ResponseWriter.WriteHeader(statusCode)
}

// WriteJSON writes JSON format message to response
func (resp Response) WriteJSON(body interface{}) (int, error) {
	jsonResponse, err := json.Marshal(body)
	if err != nil {
		return 0, errors.Wrap(err, "WriteJSON")
	}

	return resp.Write(jsonResponse)
}

// WriteError writes error message to response
func (resp Response) WriteError(statusCode int, message string) (int, error) {
	resp.WriteHeader(statusCode)
	return resp.WriteJSON(map[string]string{
		"error": message,
	})
}
