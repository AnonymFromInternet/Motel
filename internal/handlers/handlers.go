package handlers

import (
	"github.com/AnonymFromInternet/Motel/internal/render"
	"net/http"
)

func GetHandlerMainPage(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, "main")
}

func GetHandlerContactsPage(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, "contacts")
}
