package render

import (
	template2 "html/template"
	"log"
	"net/http"
)

func Template(w http.ResponseWriter, templateRootFileName string) {
	const extension = ".page.gohtml"
	parsedTemplate, err := template2.ParseFiles("./templates/"+templateRootFileName+extension, "./templates/base.layout.gohtml")

	if err != nil {
		log.Println("cannot read parsedTemplate, error :", err)

		return
	}

	err = parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Println("cannot execute parsedTemplate to response writer, error :", err)

		return
	}
}
