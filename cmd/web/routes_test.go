package main

import (
	"testing"

	"github.com/deveshkakad/bookings-app/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//Test passes
	default:
		t.Error("type is not *chi.Mux but is: ", v)

	}
}
