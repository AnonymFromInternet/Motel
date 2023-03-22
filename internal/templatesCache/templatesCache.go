package templatesCache

import (
	"fmt"
	"path/filepath"
	"text/template"
)

func Create() (map[string]*template.Template, error) {
	templatesCache := map[string]*template.Template{}

	pagesFullFileNames, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		fmt.Println(err)

		return templatesCache, err
	}

	for _, templateFileFullPathWithName := range pagesFullFileNames {
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
