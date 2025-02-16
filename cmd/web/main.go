package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/kathappiness/bookings/pkg/config"
	"github.com/kathappiness/bookings/pkg/handlers"
	"github.com/kathappiness/bookings/pkg/render"
)

const portNum = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting a server on port %s", portNum)

	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

// addValues is the about page handler
// func addValues(x, y int) int {
// 	return x + y
// }

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 10.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "Cannot divide by 0!")
// 		return
// 	}
// 	_, _ = fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 10.0, f)
// }

// func divideValues(x, y float32) (float32, error) {
// 	res := x / y
// 	if y <= 0 {
// 		err := errors.New("cannot divide by 0")
// 		return 0, err
// 	}
// 	return res, nil
// }
