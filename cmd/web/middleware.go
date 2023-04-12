package main

import (
	"github.com/AnonymFromInternet/Motel/internal/helpers"
	"github.com/justinas/nosurf"
	"net/http"
)

func CSRFMiddleware(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   !appConfig.IsDevelopmentMode,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func SessionLoadMiddleware(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !helpers.IsLoggedIn(request) {
			http.Redirect(writer, request, "/main", http.StatusSeeOther)
		}

		next.ServeHTTP(writer, request)
	})
}
