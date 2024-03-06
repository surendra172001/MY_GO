package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

var testRoutes = []string{
	"/",
	"/login",
	"/logout",
	"/register",
	"/activate",
	"/members/plans",
	"/members/subscribe",
}

func Test_Routes_Exist(t *testing.T) {
	chiRoutes := testApp.Routes().(chi.Router)

	for _, route := range testRoutes {
		RouteExists(t, chiRoutes, route)
	}
}

func RouteExists(t *testing.T, routes chi.Router, route string) {
	found := false

	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if foundRoute == route {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("did not found the %s route", route)
	}
}
