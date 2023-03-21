package render

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func Template(writer http.ResponseWriter, templateFirstName string) error {
	templatesCache, err := createTemplatesCache()
	if err != nil {
		log.Fatal(err)
	}

	templateCache, templateExistsInCache := templatesCache[templateFirstName]

	if templateExistsInCache {
		// Хорошей практикой является использовать буфер, и только потом Execute для более точного отлова ошибок
		buf := new(bytes.Buffer)

		err = templateCache.Execute(buf, nil)
		if err != nil {
			fmt.Println(err)

			return err
		}

		_, err = buf.WriteTo(writer)
		if err != nil {
			return err
		}
	} else {
		log.Fatal(errors.New("false file name! template for this file name cannot be exist"))
	}

	return nil
}

func createTemplatesCache() (map[string]*template.Template, error) {
	templatesCache := map[string]*template.Template{}

	pageFullFileNames, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		fmt.Println(err)

		return templatesCache, err
	}

	for _, templateFileFullPathWithName := range pageFullFileNames {
		fileName := filepath.Base(templateFileFullPathWithName)

		tmpl, err := template.New(fileName).ParseFiles(templateFileFullPathWithName)
		if err != nil {
			return templatesCache, err
		}

		layoutFullFileNames, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return templatesCache, err
		}

		if len(layoutFullFileNames) > 0 {
			// Merging already existing data from page.gohtml with layout.gohtml
			// А что будет, если будет несколько файлов .layout.gohtml? Скорее всего ошибка
			tmpl, err = tmpl.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return templatesCache, err
			}
		}

		templatesCache[fileName] = tmpl
	}

	return templatesCache, nil
}
