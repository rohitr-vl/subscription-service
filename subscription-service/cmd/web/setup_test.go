package main

import (
	"context"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"subscription-service/data"
	"sync"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
)

var testApp Config

func TestMain(m *testing.M) {
	gob.Register(data.User{})

	tmpPath = "./../../tmp"
	pathToManual = "./../../pdf"

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = true
	session.Cookie.SameSite = http.SameSiteLaxMode

	testApp = Config{
		Session:       session,
		DB:            nil,
		InfoLog:       log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrorLog:      log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Wait:          &sync.WaitGroup{},
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
		Models:        data.TestNew(nil),
	}

	//create dummy mailer
	errorChan := make(chan error)
	mailerChan := make(chan Message, 100)
	mailerDoneChan := make(chan bool)

	testApp.Mailer = Mail{
		ErrorChan:  errorChan,
		MailerChan: mailerChan,
		DoneChan:   mailerDoneChan,
		Wait:       testApp.Wait,
	}

	go func() {
		for {
			select {
			case <-testApp.Mailer.MailerChan:
				testApp.Wait.Done()
			case <-testApp.ErrorChan:
			case <-testApp.ErrorChanDone:
				return
			}
		}
	}()

	go func() {
		select {
		case err := <-testApp.ErrorChan:
			testApp.ErrorLog.Println(err)
		case <-testApp.ErrorChanDone:
			return
		}
	}()
	os.Exit(m.Run())
}

// to test renderer there is a func AddDefaultData(), for which we will need session data
func getCtx(req *http.Request) context.Context {
	ctx, err := testApp.Session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
