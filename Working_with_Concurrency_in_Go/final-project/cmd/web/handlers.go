package main

import (
	"final-project/cmd/web/data"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

func (app *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(w http.ResponseWriter, r *http.Request) {
	if app.Session.Exists(r.Context(), "userID") {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
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
	validPassword, err := app.Models.User.PasswordMatches(password)
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

	_, err = app.Models.User.Insert(u)
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
	// validate url
	url := fmt.Sprintf("http://localhost%s", r.URL)
	app.InfoLogger.Println(url)
	okay := VerifyToken(url)

	if !okay {
		app.Session.Put(r.Context(), "error", "Invalid Activation url")
		http.Redirect(w, r, "/Register", http.StatusSeeOther)
		return
	}

	email := r.URL.Query().Get("email")
	user, err := app.Models.User.GetByEmail(email)

	if err != nil {
		app.Session.Put(r.Context(), "error", "No user found with the specified email")
		app.ErrorLogger.Println(err.Error())
		http.Redirect(w, r, "/Register", http.StatusSeeOther)
		return
	}

	user.Active = 1
	err = app.Models.User.Update(user)

	if err != nil {
		app.Session.Put(r.Context(), "error", "User activation failed")
		app.ErrorLogger.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	app.Session.Put(r.Context(), "flash", "User activation Successful")
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (app *Config) ChooseSubscription(w http.ResponseWriter, r *http.Request) {
	// authorize user
	if !app.Session.Exists(r.Context(), "userID") {
		app.Session.Put(r.Context(), "warning", "You must login to see this page")
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	plans, err := app.Models.Plan.GetAll()
	if err != nil {
		app.ErrorLogger.Println("Not able to fetch plan")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dataMap := map[string]any{
		"plans": plans,
	}

	app.render(w, r, "plan.page.gohtml", &TemplateData{
		Data: dataMap,
	})
}

func (app *Config) SubscribeToPlan(w http.ResponseWriter, r *http.Request) {
	// get the id of the plan that is choosen =
	planID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.ErrorLogger.Println("Invalid plan id -", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get the plan from the database
	plan, err := app.Models.Plan.GetOne(planID)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Please choose a valid plan")
		http.Redirect(w, r, "/members/plans", http.StatusTemporaryRedirect)
		return
	}
	// get the user from the session
	user := app.Session.Get(r.Context(), "user").(data.User)
	// generate an invoice
	app.Wait.Add(1)
	go func() {
		defer app.Wait.Done()
		invoice, err := app.GetInvoice(*plan)
		if err != nil {
			app.ErrorChan <- err
			return
		}
		// send an email with the invoice attached
		msg := Message{
			To:       user.Email,
			Subject:  "Invoice for plan subscription",
			Data:     invoice,
			Template: "invoice",
		}
		app.SendMail(msg)
	}()

	// generate a manual
	app.Wait.Add(1)
	go func() {
		defer app.Wait.Done()
		manual := app.GenerateManual(user, plan)
		manualPath := fmt.Sprintf("%s/tmp/%d_manual.pdf", app.rootDir, user.ID)
		manual.OutputFileAndClose(manualPath)
		// send an email with the invoice attached
		msg := Message{
			To:      user.Email,
			Subject: "Your manual",
			Data:    "Please Get your manual from the attachments",
			AttachmentMap: map[string]string{
				"Manual.pdf": manualPath,
			},
		}

		app.SendMail(msg)
	}()

	// subscribe the user to an account
	err = app.Models.Plan.SubscribeUserToPlan(user, *plan)
	if err != nil {
		// app.ErrorLogger.Println("Subscription Error:", err.Error())
		app.Session.Put(r.Context(), "error", err.Error())
		http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
		return
	}

	app.Session.Put(r.Context(), "flash", "Subscribed")
	// redirect
	u, err := app.Models.User.GetOne(user.ID)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Error getting user from database")
		http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
		return
	}
	app.Session.Put(r.Context(), "user", u)
	http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
}

func (app *Config) GenerateManual(user data.User, plan *data.Plan) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "letter", "")
	pdf.SetMargins(10, 13, 10)

	time.Sleep(5 * time.Second)

	importer := gofpdi.NewImporter()
	t := importer.ImportPage(pdf, fmt.Sprintf("%s/pdf/manual.pdf", app.rootDir), 1, "/MediaBox")

	pdf.AddPage()
	importer.UseImportedTemplate(pdf, t, 0, 0, 215.9, 0)

	pdf.SetXY(75, 150)

	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 4, fmt.Sprintf("%s %s", user.FirstName, user.LastName), "", "C", false)
	pdf.Ln(5)
	pdf.MultiCell(0, 4, fmt.Sprintf("%s - User Guide", plan.PlanName), "", "C", false)

	return pdf
}

func (app *Config) GetInvoice(plan data.Plan) (string, error) {
	return plan.PlanAmountFormatted, nil
}

func (app *Config) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.Session.Exists(r.Context(), "userID") {
			app.Session.Put(r.Context(), "error", "Please login first")
			http.Redirect(w, r, "login", http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}
