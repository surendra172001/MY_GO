package main

import (
	"final-project/cmd/web/data"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var pages = []struct {
	name               string
	url                string
	expectedStatusCode int
	sessionData        map[string]any
	handler            http.HandlerFunc
	expectedHTML       string
}{
	{
		name:               "Home",
		url:                "/",
		handler:            http.HandlerFunc(testApp.HomePage),
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "Login page",
		url:                "/login",
		handler:            http.HandlerFunc(testApp.LoginPage),
		expectedStatusCode: http.StatusOK,
		expectedHTML:       `<h1 class="mt-5">Login</h1>`,
	},
	{
		name:               "Logout page",
		url:                "/logout",
		handler:            http.HandlerFunc(testApp.Logout),
		expectedStatusCode: http.StatusSeeOther,
		sessionData: map[string]any{
			"userID": 1,
			"user":   data.User{},
		},
	},
}

func Test_Pages(t *testing.T) {
	templatesPath = "./templates"

	for _, e := range pages {
		rr := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/", nil)
		ctx := getCtx(req)
		req = req.WithContext(ctx)

		if len(e.sessionData) > 0 {
			for key, val := range e.sessionData {
				testApp.Session.Put(req.Context(), key, val)
			}
		}

		e.handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("%s failed : expected status code %d and got %d", e.name, e.expectedStatusCode, rr.Code)
		}

		if len(e.expectedHTML) > 0 {
			html := rr.Body.String()
			// log.Print(html)
			if !strings.Contains(html, e.expectedHTML) {
				t.Errorf("%s failed : expected to find %s and but didn't got", e.name, e.expectedHTML)
			}
		}
	}
}

func TestConfig_PostLoginPage(t *testing.T) {
	templatesPath = "./templates"

	rr := httptest.NewRecorder()

	postedData := url.Values{
		"email":    {"admin@example.com"},
		"password": {"abc123abc123abc123abc123"},
	}

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(testApp.PostLogin)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("Wrong status code, expected %d found %d", http.StatusSeeOther, rr.Code)
	}

	if !testApp.Session.Exists(req.Context(), "userID") {
		t.Error("userID population unsuccessful")
	}
}

func TestConfig_SubscribeToPlan(t *testing.T) {
	templatesPath = "./templates"

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/members/subscribe?id=1", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.Session.Put(req.Context(), "user", data.User{
		FirstName: "admin",
		LastName:  "admin",
		Email:     "admin@example.com",
		Active:    1,
		IsAdmin:   1,
	})

	handler := http.HandlerFunc(testApp.SubscribeToPlan)

	handler.ServeHTTP(rr, req)

	testApp.Wait.Wait()

	if rr.Code != http.StatusSeeOther {
		t.Errorf("Wrong status code, expected %d found %d", http.StatusSeeOther, rr.Code)
	}
}
