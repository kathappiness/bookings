package models

import "github.com/kathappiness/bookings/internal/forms"

// TemplateData holds dataset from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	//CSRFoken is a cross-site request forgery token
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
