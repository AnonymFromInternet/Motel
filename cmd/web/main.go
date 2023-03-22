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
	appConfig.IsDevelopmentMode = true

	repository := handlers.CreateNewRepository(&appConfig)
	handlers.AsksToGetTheRepository(repository)

	render.AsksToGetTheAppConfig(&appConfig)

	server := &http.Server{
		Addr:    portNumber,
		Handler: getHandler(&appConfig),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("cannot start server")
	}
}
