package handlers

import (
	"github.com/AnonymFromInternet/Motel/internal/render"
	"log"
	"net/http"
)

func GetHandlerMainPage(writer http.ResponseWriter, request *http.Request) {
	err := render.Template(writer, "main")

	if err != nil {
		log.Println("cannot render template with name 'main'")
	}
}

func GetHandlerContactsPage(writer http.ResponseWriter, request *http.Request) {
	err := render.Template(writer, "contacts")

	if err != nil {
		log.Println("cannot render template with name 'contacts'")
	}
}
