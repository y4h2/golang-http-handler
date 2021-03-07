package api

import (
	"log"
	"net/http"
)

var _ http.Handler = new(Handler)

// NewHandler transforms our customized handler to http.HandleFunc
func NewHandler(handler Handler) func(w http.ResponseWriter, r *http.Request) {
	return handler.ServeHTTP
}

// Handler defines our customized api handler signature
type Handler func(ResponseWriter, *Request) error

// ServeHTTP implements http.Handler interface.
func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writer := NewResponse(w)
	req := NewRequest(r)
	err := fn(writer, req) // Call handler function
	if err == nil {
		return
	}

	if clientErr, ok := ToClientError(err); ok {
		log.Printf("Client side error: %s", clientErr.Error())
		writer.WriteError(clientErr.StatusCode(), clientErr.Error())
		return
	}

	log.Printf("internal handler error: %s", err.Error())
	writer.WriteError(http.StatusInternalServerError, "internal error")
}
