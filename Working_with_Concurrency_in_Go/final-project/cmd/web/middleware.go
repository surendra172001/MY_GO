package main

import "net/http"

func (app *Config) LoadSession(next http.Handler) http.Handler {
	return app.session.LoadAndSave(next)
}
