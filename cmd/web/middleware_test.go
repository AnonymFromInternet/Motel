package main

import (
	"net/http"
	"testing"
)

func TestCSRFMiddleware(t *testing.T) {
	var handler mockHttpHandler
	csrfHandler := CSRFMiddleware(handler)
	if csrfHandler == nil {
		t.Fatal("csrfHandler is nil")
	}
}

func TestSessionLoadMiddleware(t *testing.T) {
	var handler mockHttpHandler
	sessionLoadMiddleware := SessionLoadMiddleware(handler)
	if sessionLoadMiddleware == nil {
		t.Fatal("sessionLoadMiddleware is nil")
	}
}

type mockHttpHandler struct{}

func (receiver mockHttpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {}
