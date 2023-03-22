package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/AnonymFromInternet/Motel/internal/models"
	"github.com/AnonymFromInternet/Motel/internal/templatesCache"
	"log"
	"net/http"
	"text/template"
)

var appConfiguration *app.Config

func addDataToTemplate(templateData *models.TemplatesData) *models.TemplatesData {
	// it is possible to add data here
	return templateData
}

func Template(writer http.ResponseWriter, templateFirstName string, templateData *models.TemplatesData) error {
	var templates map[string]*template.Template
	var err error

	if appConfiguration.IsDevelopmentMode {
		templates, err = templatesCache.Create()
		if err != nil {
			log.Fatal("[package render]:[func Template] - cannot create template cache")
		}
	} else {
		templates = appConfiguration.TemplatesCache
	}

	templateCache, templateExistsInCache := templates[templateFirstName]

	if templateExistsInCache {
		// Хорошей практикой является использовать буфер, и только потом Execute для более точного отлова ошибок
		buffer := new(bytes.Buffer)

		templateData = addDataToTemplate(templateData)

		err := templateCache.Execute(buffer, templateData)
		if err != nil {
			fmt.Println(err)

			return err
		}

		_, err = buffer.WriteTo(writer)
		if err != nil {
			return err
		}
	} else {
		log.Fatal(errors.New("false file name! template for this file name cannot be exist"))
	}

	return nil
}

func AsksToGetTheAppConfig(config *app.Config) {
	appConfiguration = config
}
