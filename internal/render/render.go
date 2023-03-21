package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templatesCache = make(map[string]*template.Template)

func Template(writer http.ResponseWriter, templateFirstName string) error {
	templateFromCache, existsInCache := templatesCache[templateFirstName]

	if !existsInCache {
		err := createTemplateCache(templateFirstName)

		if err != nil {
			log.Println("[package render]:[func Template] - cannot create template cache")

			return err
		}

		templateFromCache = templatesCache[templateFirstName]
	}

	err := templateFromCache.Execute(writer, nil)

	if err != nil {
		log.Println("[package render]:[func Template] - cannot execute template from cache")

		return err
	}

	return nil
}

func createTemplateCache(templateName string) error {
	const layout = "./templates/base.layout.gohtml"

	const pageExtension = "page.gohtml"
	page := fmt.Sprintf("./templates/%s.%s", templateName, pageExtension)

	pageAndLayout := []string{
		page,
		layout,
	}

	// This file is a result of merging page with layout
	parsedMergedTemplate, err := template.ParseFiles(pageAndLayout...)

	if err != nil {
		log.Println("[package render]:[func createTemplateCache] - cannot parse files")

		return err
	}

	templatesCache[templateName] = parsedMergedTemplate

	return nil
}
