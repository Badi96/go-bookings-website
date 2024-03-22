package models

// TemplateData holds data from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string // Cross-site request forgery
	Flash     string
	Warning   string
	Error     string
}
