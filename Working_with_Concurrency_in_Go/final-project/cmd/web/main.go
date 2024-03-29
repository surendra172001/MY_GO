package main

import (
	"database/sql"
	"encoding/gob"
	"final-project/cmd/web/data"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = 80

// create database connection - postgres
// create session - redis
// create channels
// create a wait group
// setup application config
// setup mail - mailhog
// listen to web connections
func main() {
	fmt.Println("Subscription service")

	db := initDB()
	db.Ping()

	session := initSession()

	wg := sync.WaitGroup{}

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := Config{
		Session:       session,
		DB:            db,
		Wait:          &wg,
		InfoLogger:    infoLogger,
		ErrorLogger:   errorLogger,
		Models:        data.New(db),
		rootDir:       ".",
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
	}

	app.Mailer = app.createMailer()
	go app.listenForMail()

	go app.listenForShutdown()

	go app.listenForError()

	app.serve()

}

func (app *Config) listenForMail() {
	for {

		select {
		case msg := <-app.Mailer.MailerChan:
			app.Mailer.SendMail(msg, app.Mailer.ErrorChan)
		case err := <-app.Mailer.ErrorChan:
			app.ErrorLogger.Println(err.Error())
		case <-app.Mailer.DoneChan:
			return
		}
	}

}

func (app *Config) createMailer() *Mail {
	mailerChan := make(chan Message, 100)
	errorChan := make(chan error)
	mailerDoneChan := make(chan bool)

	mailer := &Mail{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Encryption:  "none",
		FromAddress: "info@mycompany.com",
		FromName:    "info",
		Wait:        app.Wait,
		ErrorChan:   errorChan,
		MailerChan:  mailerChan,
		DoneChan:    mailerDoneChan,
	}
	return mailer
}

func (app *Config) serve() {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: app.Routes(),
	}

	app.InfoLogger.Println("Starting the wev server...")

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

func initSession() *scs.SessionManager {
	gob.Register(data.User{})
	// session initialization
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true
	return session
}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	return redisPool
}

func initDB() *sql.DB {
	connection := connectToDB()
	if connection == nil {
		log.Panic("Database is not up")
	}

	return connection
}

func connectToDB() *sql.DB {

	count := 0

	DSN := "host=localhost port=5432 user=postgres password=password dbname=concurrency sslmode=disable timezone=UTC connect_timeout=5"

	// log.Println("DSN", DSN)

	// for _, x := range os.Environ() {
	// 	fmt.Println(x)
	// }

	// log.Println(DSN)

	for {

		conn, err := openDB(DSN)

		if err != nil {
			log.Println("postgres is not ready yet...", err)
		} else {
			log.Println("Connected to database!")
			return conn
		}

		if count > 10 {
			return nil
		}

		count++

		log.Println("Backing off for 1 second")

		time.Sleep(time.Second)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	// fmt.Println("err", err)
	return conn, nil
}

func (app *Config) listenForError() {
	for {
		select {
		case err := <-app.ErrorChan:
			app.ErrorLogger.Println(err.Error())
		case <-app.ErrorChanDone:
			app.InfoLogger.Println("Stopping to listen to errors")
			return
		}
	}
}

func (app *Config) listenForShutdown() {
	app.InfoLogger.Println("Listening for shutdown..")
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGABRT)

	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {
	app.InfoLogger.Println("doing cleanup tasks...")
	app.Wait.Wait()
	app.InfoLogger.Println("Done waiting")
	app.Mailer.DoneChan <- true
	app.ErrorChanDone <- true
	app.InfoLogger.Println("Done Putting Value")
	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.DoneChan)
	close(app.ErrorChan)
	close(app.ErrorChanDone)
	app.InfoLogger.Println("Closed all channels.", "Done the cleanup tasks and shutting down...")
}
