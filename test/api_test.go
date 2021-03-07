package api_test

import (
	"testing"

	"github.com/y4h2/golang-http-handler/pkg/api"
)

func TestAPI(t *testing.T) {
	api.NewNotFoundErr("test")
	t.Log("test")
}
