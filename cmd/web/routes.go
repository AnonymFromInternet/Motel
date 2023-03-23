package main

import (
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/AnonymFromInternet/Motel/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func getHandler(appConfig *app.Config) http.Handler {
	multiplexer := chi.NewRouter()
	multiplexer.Use(middleware.Recoverer)
	multiplexer.Use(CSRFMiddleware)
	multiplexer.Use(SessionLoadMiddleware)

	multiplexer.Get("/main", handlers.Repo.GetHandlerMainPage)
	multiplexer.Get("/contacts", handlers.Repo.GetHandlerContactsPage)

	fileServer := http.FileServer(http.Dir("./static/")) // путь указывается относительно рута
	multiplexer.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return multiplexer
}
