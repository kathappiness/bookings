package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/kathappiness/bookings/internal/config"
	"github.com/kathappiness/bookings/internal/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	// get the template cache from the app config
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get a template from template cache")
	}
	//Creating a buffer, trying to execute that buffer directly for error checking
	//I can ignore the result, but track down the error if it has occured and it comes
	// from the value stored in the map (cache)
	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// creating a cache variable
	myCache := map[string]*template.Template{}
	// get all files named *.page.tmpl from the ./templates,
	// this will populate pages with a slice of strings
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		// here getting a name of a template from a full path stored in a variable page
		name := filepath.Base(page)
		// ts stands for template set, want to parse it and create a parsed version of template
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// looking for layouts
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

//A more complicated approach

// It's the easier approach with a map:
// Created a pkg level variable with rendered templates,
// then checking if the required template has been already rendered -> get it from the var,
// if not -> render it and add to that var which is cache

// var tmpcache = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error
// 	// check if the template is already in the cache
// 	_, inMap := tmpcache[t]
// 	if !inMap {
// 		log.Println("creating template and addind to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		log.Println("using cached template")
// 	}
// 	tmpl = tmpcache[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}
// 	//parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	//add template to cache (map)
// 	tmpcache[t] = tmpl
// 	return nil
// }
