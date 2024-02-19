package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type Config struct {
	session     *scs.SessionManager
	DB          *sql.DB
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	Wait        *sync.WaitGroup
}
