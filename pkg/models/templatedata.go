package models

// TemplateData holds dataset from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	//CFRSToken is a cross-site request forgery token
	CFRSToken string
	Flash     string
	Warning   string
	Error     string
}
