package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
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
		session:     session,
		DB:          db,
		Wait:        &wg,
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
	}

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
	// session initialization
	session := &scs.SessionManager{}
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
			return redis.Dial("tcp", os.Getenv("REDIS"))
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

	DSN := os.Getenv("DSN")

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
