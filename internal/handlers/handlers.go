package handlers

import (
	"github.com/AnonymFromInternet/Motel/internal/render"
	"log"
	"net/http"
)

func GetHandlerMainPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "main.page.gohtml"
	err := render.Template(writer, fileName)

	if err != nil {
		log.Println("cannot render template with name 'main'")
	}
}

func GetHandlerContactsPage(writer http.ResponseWriter, request *http.Request) {
	const fileName = "contacts.page.gohtml"
	err := render.Template(writer, fileName)

	if err != nil {
		log.Println("cannot render template with name 'contacts'")
	}
}
