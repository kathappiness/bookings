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
	mux.Post("/make-reservation", http.HandlerFunc(handlers.Repo.PostReservation))
	mux.Get("/reservation-summary", http.HandlerFunc(handlers.Repo.ReservationSummary))

	mux.Get("/majors-suite", http.HandlerFunc(handlers.Repo.Majors))

	mux.Get("/search-availability", http.HandlerFunc(handlers.Repo.Availability))
	mux.Post("/search-availability", http.HandlerFunc(handlers.Repo.PostAvailability))
	mux.Post("/search-availability-json", http.HandlerFunc(handlers.Repo.AvailabilityJSON))
	mux.Get("/choose-room/{id}", http.HandlerFunc(handlers.Repo.ChooseRoom))
	mux.Get("/book-room", http.HandlerFunc(handlers.Repo.BookRoom))

	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
