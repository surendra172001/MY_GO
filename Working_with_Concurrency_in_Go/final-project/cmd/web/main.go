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
		Session:     session,
		DB:          db,
		Wait:        &wg,
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
		Models:      data.New(db),
	}

	go app.listenForShutdown()

	app.serve()

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
	app.InfoLogger.Println("Done the cleanup tasks and shutting down...")
}
