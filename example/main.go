package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/y4h2/golang-http-handler/pkg/api"
)

func HelloWorldHandler(w api.ResponseWriter, r *api.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
	return nil
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", api.NewHandler(HelloWorldHandler)).Methods("GET")
}
