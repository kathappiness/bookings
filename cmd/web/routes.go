package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kathappiness/bookings/internal/config"
	"github.com/kathappiness/bookings/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/generals-quaters", handlers.Repo.Generals)
	mux.Get("/make-reservation", http.HandlerFunc(handlers.Repo.Reservation))
	mux.Get("/majors-suite", http.HandlerFunc(handlers.Repo.Majors))

	mux.Get("/search-availability", http.HandlerFunc(handlers.Repo.Availability))
	mux.Post("/search-availability", http.HandlerFunc(handlers.Repo.PostAvailability))
	mux.Post("/search-availability-json", http.HandlerFunc(handlers.Repo.AvailabilityJSON))

	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))

	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux.Get("/make-reservation", http.HandlerFunc(handlers.Repo.Reservation))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
