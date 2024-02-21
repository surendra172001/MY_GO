package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) Routes() http.Handler {
	// create a new router
	mux := chi.NewRouter()

	// middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.LoadSession)

	// routes
	mux.Get("/", app.HomePage)
	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLogin)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegister)
	mux.Post("/activate-account", app.ActivateAccount)

	return mux
}
