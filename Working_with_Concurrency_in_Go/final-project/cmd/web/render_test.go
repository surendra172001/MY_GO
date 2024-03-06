package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfig_AddDefaultData(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.Session.Put(req.Context(), "flash", "flash")
	testApp.Session.Put(req.Context(), "warning", "warning")
	testApp.Session.Put(req.Context(), "error", "error")

	td := testApp.AddDefaultData(&TemplateData{}, req)

	if td.Flash != "flash" {
		t.Errorf("The flash message must be equal to flash")
	}

	if td.Warning != "warning" {
		t.Errorf("The warning message must be equal to warning")
	}

	if td.Error != "error" {
		t.Errorf("The error message must be equal to error")
	}
}

func TestConfig_IsAuthenticated(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	auth := testApp.isAuthenticated(req)

	if auth {
		t.Errorf("Returns true for authenticated when it should return false")
	}

	testApp.Session.Put(req.Context(), "userID", true)

	auth = testApp.isAuthenticated(req)

	if !auth {
		t.Errorf("Returns false for authenticated when it should return true")
	}
}

func TestConfig_render(t *testing.T) {
	templatesPath = "./templates"

	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	testApp.render(rr, req, "home.page.gohtml", &TemplateData{})

	if rr.Result().StatusCode != 200 {
		t.Errorf("Something went wrong while rendering")
	}
}
