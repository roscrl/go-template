package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestHandleHome(t *testing.T) {

	is := is.New(t)
	server := newMinimalServer(DEV)
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, r)    // integration test like (middlewares included)
	server.handleHome()(w, r) // unit test like (no middlewares)

	is.Equal(w.Result().StatusCode, http.StatusOK)
}
