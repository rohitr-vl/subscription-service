package main

import "net/http"

func (app *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	// if empty, this is called stub handler
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

func (app *Config) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = app.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}
	// get email id and password
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	/*
		app.InfoLog.Println("Received input:", email, password)
		allUsers, err := app.Models.User.GetAll()
		if err != nil {
			app.ErrorLog.Println("error fetching allUsers:", err)
		}
		app.InfoLog.Println("All users:", allUsers)
	*/
	user, err := app.Models.User.GetByEmail(email)
	if err != nil {
		app.ErrorLog.Println("User not found in DB. err:", err)
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// verify password
	validPassword, err := user.PasswordMatches(password)
	if err != nil {
		app.ErrorLog.Println("Error while matching password. err:", err)
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if !validPassword {
		msg := Message{
			To:      email,
			Subject: "Failed Login Attempt",
			Data:    "Invalid login attempt",
		}
		app.sendEmail(msg)
		app.ErrorLog.Println("User gave invalid password. err:", err)
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// log the user in
	app.Session.Put(r.Context(), "userID", user.ID)
	app.Session.Put(r.Context(), "user", user)

	app.Session.Put(r.Context(), "flash", "Successful Login!")
	//redirect to home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Config) Logout(w http.ResponseWriter, r *http.Request) {
	// clean up session
	_ = app.Session.Destroy(r.Context())
	_ = app.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *Config) RegisterPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *Config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	// create a user

	// send activation email

	//subscribe the use to an account
}

func (app *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// validate the url

	// generate an invoice

	// send email with invoice & plan details attachment
}
