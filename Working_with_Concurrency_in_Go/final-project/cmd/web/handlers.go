package main

import (
	"final-project/cmd/web/data"
	"fmt"
	"html/template"
	"net/http"
)

func (app *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

func (app *Config) PostLogin(w http.ResponseWriter, r *http.Request) {
	_ = app.Session.RenewToken(r.Context())

	// parse form post
	err := r.ParseForm()
	if err != nil {
		app.ErrorLogger.Println(err)
	}

	// get email and password from form post
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.Models.User.GetByEmail(email)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// check password
	validPassword, err := user.PasswordMatches(password)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !validPassword {
		msg := Message{
			To:      "info@mycompany.com",
			Subject: "Invalid credentials",
			Data:    "Wrong password entered",
		}
		app.SendMail(msg)
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// okay, so log user in
	app.Session.Put(r.Context(), "userID", user.ID)
	app.Session.Put(r.Context(), "user", user)

	app.Session.Put(r.Context(), "flash", "Successful login!")

	// redirect the user
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Config) Logout(w http.ResponseWriter, r *http.Request) {
	app.Session.Destroy(r.Context())
	app.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *Config) RegisterPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *Config) PostRegister(w http.ResponseWriter, r *http.Request) {
	// parse form
	err := r.ParseForm()
	if err != nil {
		app.ErrorLogger.Println("Form Parsing error", err.Error())
		return
	}
	// create a user
	u := data.User{
		Email:     r.FormValue("email"),
		FirstName: r.FormValue("first-name"),
		LastName:  r.FormValue("last-name"),
		Password:  r.FormValue("password"),
		Active:    0,
		IsAdmin:   0,
	}

	_, err = u.Insert(u)
	if err != nil {
		app.Session.Put(r.Context(), "error", "User Creation error")
		http.Redirect(w, r, "/Register", http.StatusSeeOther)
		return
	}

	// send an activation email
	url := fmt.Sprintf("http://localhost/activate?email=%s", u.Email)
	signedUrl := GenerateTokenFromString(url)
	app.InfoLogger.Println("Signed URL", signedUrl)

	msg := Message{
		To:       u.Email,
		Subject:  "Account activation",
		Template: "confirmation-email",
		Data:     template.HTML(signedUrl),
	}

	app.SendMail(msg)
	app.Session.Put(r.Context(), "flash", "Success in user registration, Please check your email for account activation")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {

}
