package app

import (
	"log"
	"text/template"
)

type Config struct {
	UseCache       bool
	TemplatesCache map[string]*template.Template
	InfoLogger     *log.Logger
}
