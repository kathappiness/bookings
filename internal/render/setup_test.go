package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/kathappiness/bookings/internal/config"
	"github.com/kathappiness/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	// what I am going to put in the session
	gob.Register(models.Reservation{})
	// change this to true when in production
	testApp.InProduction = false
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog
	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// Instead of "session.Cookie.Secure = testApp.InProduction" put false in it for testing purposes
	session.Cookie.Secure = false

	testApp.Session = session

	// Makes sure that var app has everything you need for testing
	app = &testApp
	// Runs tests
	os.Exit(m.Run())
}

type MyWriter struct{}

func (w *MyWriter) Header() http.Header {
	var h http.Header
	return h
}

func (w *MyWriter) WriteHeader(i int) {}

func (w *MyWriter) Write(b []byte) (int, error) {
	lenght := len(b)
	return lenght, nil
}
