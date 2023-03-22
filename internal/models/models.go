package models

type TemplatesData struct {
	BasicData   map[string]interface{}
	Error       string
	Warning     string
	CSRFToken   string
	ShowMessage string
}
