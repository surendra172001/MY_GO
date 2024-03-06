package main

import (
	"database/sql"
	"final-project/cmd/web/data"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type Config struct {
	Session       *scs.SessionManager
	DB            *sql.DB
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	Wait          *sync.WaitGroup
	Models        data.Models
	Mailer        *Mail
	ErrorChan     chan error
	ErrorChanDone chan bool
	rootDir       string
}
