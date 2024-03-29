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
	mux.Get("/activate", app.ActivateAccount)

	mux.Mount("/members", app.AuthRouter())

	return mux
}

func (app *Config) AuthRouter() http.Handler {
	// create a new router
	mux := chi.NewRouter()

	// middleware
	mux.Use(app.Auth)

	mux.Get("/plans", app.ChooseSubscription)
	mux.Get("/subscribe", app.SubscribeToPlan)

	return mux
}
