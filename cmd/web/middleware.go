package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf はすべてのPOSTリクエストにCSRF保護を追加
func NoSurf(next http.Handler) http.Handler{
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		SameSite: http.SameSiteLaxMode,
		Secure: app.InProduction,
	})

	return csrfHandler
}

// SessionLoad は、リクエストごとにセッションをロードして保存
func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}