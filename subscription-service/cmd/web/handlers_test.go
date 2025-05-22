package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"subscription-service/data"
	"testing"
)

// setup table tests which will test all pages
var pageTests = []struct {
	name               string
	url                string
	expectedStatusCode int
	handler            http.HandlerFunc
	sessionData        map[string]any
	expectedHTML       string
}{
	{
		name:               "home",
		url:                "/",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.HomePage,
	},
	{
		name:               "login page",
		url:                "/login",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.LoginPage,
		expectedHTML:       `<h1 class="mt-5">Login</h1>`,
	},
	{
		name:               "logout page",
		url:                "/logout",
		expectedStatusCode: http.StatusSeeOther,
		handler:            testApp.Logout,
		sessionData: map[string]any{
			"userID": 1,
			"user":   data.User{},
		},
	},
}

func Test_Pages(t *testing.T) {
	// the existing package level var is overwritten here, as per the location of test file
	pathToTemplates = "./templates"
	// testApp.InfoLog.Println("test_home_route, pathToTemplates:", pathToTemplates)

	for _, e := range pageTests {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", e.url, nil)

		ctx := getCtx(req)
		req = req.WithContext(ctx)

		if len(e.sessionData) > 0 {
			for key, value := range e.sessionData {
				testApp.Session.Put(ctx, key, value)
			}
		}

		e.handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("%s failed; expected %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
		} else {
			t.Logf("%s route test passed", e.name)
		}

		if len(e.expectedHTML) > 0 {
			html := rr.Body.String()
			if !strings.Contains(html, e.expectedHTML) {
				t.Errorf("%s failed: expected to find %s, but did not", e.name, e.expectedHTML)
			} else {
				t.Logf("%s content test passed", e.name)
			}
		}
	}
}

func TestConfig_PostLoginPage(t *testing.T) {
	pathToTemplates = "./templates"

	postedData := url.Values{
		"email":    {"admin@example.com"},
		"password": {"abc123abc123abc123abc123"},
	}
	/*	testApp.InfoLog.Println(postedData.Get("email"))
		testApp.InfoLog.Println(postedData.Get("password"))
	*/
	testApp.InfoLog.Println("Test login post request body:", strings.NewReader(postedData.Encode()))

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(postedData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(testApp.PostLoginPage)
	handler.ServeHTTP(rr, req)
	testApp.InfoLog.Println("Response Body:", rr.Body)
	if rr.Code != http.StatusSeeOther {
		t.Error("wrong code returned")
	}
	if !testApp.Session.Exists(ctx, "userID") {
		t.Error("did not find userID in session")
	}
}

func TestConfig_SubscribeToPlan(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/subscribe?id=1", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.Session.Put(ctx, "user", data.User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "User",
		Active:    1,
	})

	handler := http.HandlerFunc(testApp.SubscribeToPlan)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("Expected status code StatusSeeOther, but got %d", rr.Code)
	}

}
