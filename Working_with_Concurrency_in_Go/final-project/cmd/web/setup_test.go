package main

import (
	"context"
	"encoding/gob"
	"final-project/cmd/web/data"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
)

var testApp Config

func TestMain(m *testing.M) {
	gob.Register(data.User{})
	// session initialization
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	testApp = Config{
		Session:       session,
		DB:            nil,
		InfoLogger:    infoLogger,
		ErrorLogger:   errorLogger,
		Wait:          &sync.WaitGroup{},
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
		rootDir:       "../..",
		Models:        data.New_Test(nil),
	}

	mailerChan := make(chan Message, 100)
	mailerErrorChan := make(chan error)
	mailerDoneChan := make(chan bool)
	mailer := Mail{
		Wait:       testApp.Wait,
		MailerChan: mailerChan,
		ErrorChan:  mailerErrorChan,
		DoneChan:   mailerDoneChan,
	}

	testApp.Mailer = &mailer

	go func() {
		for {
			select {
			case <-mailerChan:
				testApp.Wait.Done()
			case <-mailerErrorChan:
			case <-mailerDoneChan:
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case err := <-testApp.ErrorChan:
				testApp.ErrorLogger.Println(err.Error())
			case <-testApp.ErrorChanDone:
				return
			}
		}
	}()

	os.Exit(m.Run())
}

func getCtx(req *http.Request) context.Context {
	ctx, err := testApp.Session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return ctx
}
