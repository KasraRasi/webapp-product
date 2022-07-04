package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Floatmap  map[string]float64
	Data      map[string]interface{}
	CsrfToken string
	Flash     string
	Warning   string
	Error     string
}
