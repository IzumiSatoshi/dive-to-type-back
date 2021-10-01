package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "user=admin password=secret dbname=dive_to_type sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	DB.Close()
}
