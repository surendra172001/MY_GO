package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var templatesPath = "./cmd/web/templates"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	DataMap       map[string]any
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
	// User *data.user
}

func (app *Config) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", templatesPath),
		fmt.Sprintf("%s/header.partial.gohtml", templatesPath),
		fmt.Sprintf("%s/navbar.partial.gohtml", templatesPath),
		fmt.Sprintf("%s/footer.partial.gohtml", templatesPath),
		fmt.Sprintf("%s/alerts.partial.gohtml", templatesPath),
	}

	tempSlice := []string{fmt.Sprintf("%s/%s", templatesPath, t)}

	tempSlice = append(tempSlice, partials...)

	if td == nil {
		td = &TemplateData{}
	}

	temp, err := template.ParseFiles(tempSlice...)
	if err != nil {
		app.ErrorLogger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := temp.Execute(w, app.AddDefaultData(td, r)); err != nil {
		app.ErrorLogger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Config) AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Authenticated = app.isAuthenticated(r)
	// TODO: if user is authenticated get more info about the user
	td.Now = time.Now()
	return td
}

func (app *Config) isAuthenticated(r *http.Request) bool {
	return app.Session.Exists(r.Context(), "userID")
}
