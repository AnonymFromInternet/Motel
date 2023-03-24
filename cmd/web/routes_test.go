package main

import (
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestGetHandler(t *testing.T) {
	var appConfig *app.Config
	multiplexer := getHandler(appConfig)

	switch dt := multiplexer.(type) {
	case *chi.Mux:
	// ok
	default:
		t.Error("data type of handler is not *chi.Mux, but :", dt)
	}
}
