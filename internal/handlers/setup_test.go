package handlers

import (
	"bytes"
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"text/template"
)

// Требуется мультиплексор или по крайней мере один обработчик пути. Нет смысла тестировать все пути, ведь сейчас
// они абсолютно одинаковые.

func TESTGetMultiplexer(appConfig *app.Config) http.Handler {
	multiplexer := chi.NewRouter()

	multiplexer.Get("/main", TESTGetHandlerMainPage)

	return multiplexer
}

func TESTGetHandlerMainPage(writer http.ResponseWriter, request *http.Request) {
	err := TESTRenderTemplate(writer, request)

	if err != nil {
		log.Println("cannot render template with name 'main'")
	}
}

func TESTRenderTemplate(writer http.ResponseWriter, request *http.Request) error {
	var templates map[string]*template.Template
	var err error

	templates, err = TESTCreateTemplatesCache()

	templateCache, _ := templates["main.page"]

	buffer := new(bytes.Buffer)

	err = templateCache.Execute(buffer, nil)
	if err != nil {
		return err
	}

	_, err = buffer.WriteTo(writer)
	if err != nil {
		return err
	}

	return nil
}

func TESTCreateTemplatesCache() (map[string]*template.Template, error) {
	var templateMap map[string]*template.Template
	tmpl, err := template.New("main.page.gohtml").ParseFiles("../../templates/main.page.gohtml")
	if err != nil {
		return nil, err
	}

	tmpl, err = tmpl.ParseGlob("../../templates/base.layout.gohtml")

	templateMap["main.page"] = tmpl

	return templateMap, nil
}
