package main

import (
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
