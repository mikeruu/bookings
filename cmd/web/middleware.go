package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//Nosurf adds csrf protection
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//SessionLoad loads and saves session data on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
