package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"subscription-service/data"
	"time"

	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

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
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}
	// validate data, mandatory fields, email check, etc.
	// create a user
	u := data.User{
		Email: r.Form.Get("email"),
		FirstName: r.Form.Get("first-name"),
		LastName: r.Form.Get("last-name"),
		Password: r.Form.Get("password"),
		Active: 0,
		IsAdmin: 0,
	}

	_, err = u.Insert(u)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Unable to create user.")
		app.Session.Put(r.Context(), "flash", "Technical error occurred, could not register.")
		http.Redirect(w,r,"/register", http.StatusSeeOther)
		return
	}

	// send activation email
	url := fmt.Sprintf("http://localhost/activate?email=%s", u.Email)
	signedUrl := GenerateTokenFromString(url)
	app.InfoLog.Println(signedUrl)

	msg := Message{
		To: u.Email,
		Subject: "Account Activation",
		Template: "confirmation-email",
		Data: template.HTML(signedUrl),
	}

	app.sendEmail(msg)
	app.Session.Put(r.Context(), "flash", "Confirmation email sent, kindly check your email.")
	http.Redirect(w,r,"/login",http.StatusSeeOther)

}

func (app *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// validate the url
	url := r.RequestURI
	testUrl := fmt.Sprintf("http://localhost%s",url)
	okay := VerifyToken(testUrl)
	if !okay {
		app.Session.Put(r.Context(), "error", "Invalid Token.")
		http.Redirect(w,r,"/", http.StatusSeeOther)
		return
	}

	// activate account
	u, err := app.Models.User.GetByEmail(r.URL.Query().Get("email"))
	if err!= nil {
		app.Session.Put(r.Context(), "error", "Unable to activate user.")
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	u.Active = 1
	err = u.Update()
	if err!= nil {
		app.Session.Put(r.Context(), "error", "Unable to update user.")
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	app.Session.Put(r.Context(), "flash", "Account Activated !")
	http.Redirect(w,r,"/login",http.StatusSeeOther)
	
	
}

func (app *Config) SubscribeToPlan(w http.ResponseWriter, r *http.Request) {
	// get id of the plan that is choosen
	id := r.URL.Query().Get("id")
	planID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error while reading plan id:", err)
	}
	log.Println("Request to subscribe to plan id:", planID)

	// get the plan from db
	plan, err := app.Models.Plan.GetOne(planID)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Unable to find plan.")
		http.Redirect(w,r,"/members/plans", http.StatusSeeOther)
		return
	}

	// get the user from session
	user, ok := app.Session.Get(r.Context(), "user").(data.User)
	if !ok {
		app.Session.Put(r.Context(), "error", "Login again!")
		http.Redirect(w,r,"/login", http.StatusSeeOther)
		return
	}

	// subscribe the user to an account
	err = app.Models.Plan.SubscribeUserToPlan(user, *plan)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Error subscribing to the plan")
		http.Redirect(w,r,"/members/plan", http.StatusSeeOther)
		return
	}

	// since user data has been updated, get latest data from db and update it in session
	u, err := app.Models.User.GetOne(user.ID)
	if err != nil {
		app.ErrorLog.Println("Error getting user from DB!")
		http.Redirect(w,r,"/members/plan", http.StatusSeeOther)
		return
	}
	app.Session.Put(r.Context(), "user", u)

	// generate an invoice PDF and send it in email
	app.Wait.Add(1)

	go func() {
		defer app.Wait.Done()

		invoice, err := app.getInvoice(user, plan)
		if err != nil {
			app.ErrorChan <- err
		}
		msg := Message {
			To: user.Email,
			Subject: "Your Invoice",
			Data: invoice,
			Template: "invoice",
		}
		app.sendEmail(msg)
	}()

	// generate a plan details PDF and send it in email
	app.Wait.Add(1)
	go func ()  {
		defer app.Wait.Done()

		pdf := app.generateManual(user, plan)
		err := pdf.OutputFileAndClose(fmt.Sprintf("./tmp/%d_manual.pdf", user.ID))
		if err != nil {
			app.ErrorChan <- err
			return
		}
		msg := Message{
			To: user.Email,
			Subject: "Your Manual",
			Data: "Your user manual is attached",
			AttachmentMap: map[string]string{
				"Manual.pdf": fmt.Sprintf("./tmp/%d_manual.pdf", user.ID),
			},
		}
		app.sendEmail(msg)

		// test app error chan
		app.ErrorChan <- errors.New("some custom error")
	}()
	// redirect
	app.Session.Put(r.Context(), "flash", "Subscribed!")
	http.Redirect(w,r, "/members/plans", http.StatusSeeOther)
}

func (app *Config) generateManual(u data.User, plan *data.Plan) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.SetMargins(10,13,10)

	importer := gofpdi.NewImporter()

	//simulating pdf creation time
	time.Sleep(5*time.Second)

	t := importer.ImportPage(pdf, "./pdf/manual.pdf",1, "/MediaBox")
	pdf.AddPage()
	importer.UseImportedTemplate(pdf,t,0,0,215.9,0)
	pdf.SetX(75)
	pdf.SetY(150)

	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0,4,fmt.Sprintf("Hi %s %s", u.FirstName, u.LastName), "", "L", false)
	pdf.Ln(5)
	pdf.MultiCell(0,4,fmt.Sprintf("%s User Guide", plan.PlanName), "", "C", false)

	return pdf
}

func (app *Config) getInvoice(u data.User, plan *data.Plan) (string, error) {
	// generate pdf yourself, for now just using string
	app.InfoLog.Println("Plan Formatted Amount:", plan.PlanAmountFormatted)
	app.InfoLog.Println("Plan Amount:", plan.PlanAmount)
	return plan.PlanAmountFormatted, nil
}

func (app *Config) ChooseSubscription(w http.ResponseWriter, r *http.Request) {
	// checking if session is active, is handled by middleware

	// fetch plans data from db
	plans, err := app.Models.Plan.GetAll()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}
	dataMap := make(map[string]any)
	dataMap["plans"] = plans
	// send data to render page
	app.render(w,r,"plans.html.gohtml", &TemplateData{
		Data: dataMap,
	})
}