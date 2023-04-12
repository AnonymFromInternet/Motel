package main

import (
	"github.com/AnonymFromInternet/Motel/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func getHandler() http.Handler { // // &appConfig в него передавать необязательно
	multiplexer := chi.NewRouter()

	multiplexer.Use(middleware.Recoverer)
	multiplexer.Use(SessionLoadMiddleware)
	multiplexer.Use(CSRFMiddleware)

	multiplexer.Get("/main", handlers.Repo.GetHandlerMainPage)
	multiplexer.Get("/about", handlers.Repo.GetHandlerAboutPage)
	multiplexer.Get("/room", handlers.Repo.GetHandlerRoomPage)
	multiplexer.Get("/blue-room", handlers.Repo.GetHandlerBlueRoomPage)
	multiplexer.Get("/contacts", handlers.Repo.GetHandlerContactsPage)
	multiplexer.Get("/availability", handlers.Repo.GetHandlerAvailabilityPage)
	multiplexer.Get("/reservation", handlers.Repo.GetHandlerReservationPage)
	multiplexer.Get("/reservation-confirm", handlers.Repo.GetHandlerReservationConfirmPage)
	multiplexer.Get("/login", handlers.Repo.GetLoginPage)
	multiplexer.Get("/logout", handlers.Repo.Logout)

	multiplexer.Post("/availability", handlers.Repo.PostHandlerAvailabilityPage)
	multiplexer.Post("/availability-json", handlers.Repo.PostHandlerAvailabilityPageJSON)
	multiplexer.Post("/login", handlers.Repo.PostLoginPage)

	multiplexer.Route("/admin", func(mux chi.Router) {
		mux.Use(AuthMiddleware)

		mux.Get("/dashboard", handlers.Repo.GetAdminDashboard)
	})

	fileServer := http.FileServer(http.Dir("./static/")) // путь указывается относительно рута
	multiplexer.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return multiplexer
}
