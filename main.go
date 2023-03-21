package main

import (
	template2 "html/template"
	"log"
	"net/http"
)

const portNumber = "localhost:8080"

func GetHandlerMainPage(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "main")
}

func GetHandlerContactsPage(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "contacts")
}

func renderTemplate(w http.ResponseWriter, templateRootFileName string) {
	const extension = ".page.tmpl"
	parsedTemplate, err := template2.ParseFiles("./templates/" + templateRootFileName + extension)

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

func main() {
	http.HandleFunc("/main", GetHandlerMainPage)
	http.HandleFunc("/contacts", GetHandlerContactsPage)

	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatal("cannot start server. Error :", err)
	}
}
