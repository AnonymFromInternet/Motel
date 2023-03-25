package app

import (
	"github.com/alexedwards/scs/v2"
	"log"
	"text/template"
)

type Config struct {
	IsDevelopmentMode bool
	TemplatesCache    map[string]*template.Template
	InfoLogger        *log.Logger
	ErrorLogger       *log.Logger
	Session           *scs.SessionManager
}
