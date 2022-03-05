package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raihan2bd/bookings/pkg/config"
	"github.com/raihan2bd/bookings/pkg/handler"
)

func routes(app *config.AppConfig) http.Handler {
	// initializing the router form chi
	router := chi.NewRouter()

	// middleware will start here
	router.Use(SessionLoad)

	// all routes will start here
	router.Get("/", handler.Repo.Home)
	router.Get("/about", handler.Repo.About)

	//serve static files
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}
