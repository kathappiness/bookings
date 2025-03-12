package render

import (
	"net/http"
	"testing"

	"github.com/kathappiness/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	// Putting something in the session to test it
	session.Put(r.Context(), "flash", "123")

	res := AddDefaultData(&td, r)

	if res.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	// creating a context fron the request that has been just built
	ctx := r.Context()
	// Putting session data in the context
	// r.Header.Get("X-Session") - a part that makes it an active session
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	// Putting context back in the request
	r = r.WithContext(ctx)
	return r, nil
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	// creating template cache before rendering smth
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	// putting template cache into app variable
	app.TemplateCache = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	var ww MyWriter
	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}
	err = RenderTemplate(&ww, r, "non-existing.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered a template that doesn't exist")
	}

}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
