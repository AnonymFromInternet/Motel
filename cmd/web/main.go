package main

import (
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/AnonymFromInternet/Motel/internal/handlers"
	"github.com/AnonymFromInternet/Motel/internal/render"
	"github.com/AnonymFromInternet/Motel/internal/templatesCache"
	"log"
	"net/http"
)

const portNumber = "localhost:8080"

func main() {
	var appConfig app.Config
	var err error

	appConfig.TemplatesCache, err = templatesCache.Create()
	if err != nil {
		log.Fatal("[package main]:[func main] - cannot get app config")
	}

	render.AsksToGetTheAppConfig(&appConfig)

	http.HandleFunc("/main", handlers.GetHandlerMainPage)
	http.HandleFunc("/contacts", handlers.GetHandlerContactsPage)

	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatal("cannot start server. Error :", err)
	}
}
