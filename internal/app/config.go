package app

import (
	"log"
	"text/template"
)

type Config struct {
	IsDevelopmentMode bool
	TemplatesCache    map[string]*template.Template
	InfoLogger        *log.Logger
}
