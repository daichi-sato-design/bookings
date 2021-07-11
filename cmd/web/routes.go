package main

import (
	"net/http"

	"github.com/daichi-sato-design/bookings/pkg/config"
	"github.com/daichi-sato-design/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler{
	mux := chi.NewRouter()

	// A good base middleware stack
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}