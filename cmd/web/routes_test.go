package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/kathappiness/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)
	//Checking if the received object is of the type http.Handler
	switch v := mux.(type) {
	case *chi.Mux: // do nothing:)
	default:
		t.Error(fmt.Printf("type is *chi.Mux, it's a type %T", v))
	}
}
