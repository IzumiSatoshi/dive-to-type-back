package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Thread struct {
	UUID        string
	Title       string
	Description string
	Created_at  string
}

func main() {
	db, err := sql.Open("postgres", "user=admin password=secret dbname=dive_to_type sslmode=disable")
	handleErr(err)
	rows, err := db.Query("SELECT * FROM threads")
	handleErr(err)

	threads := []Thread{}
	for rows.Next() {
		thread := Thread{}
		var err = rows.Scan(&thread.UUID, &thread.Title, &thread.Description, &thread.Created_at)
		handleErr(err)
		threads = append(threads, thread)
	}
	fmt.Println(threads)
	json_byte, err := json.Marshal(threads)
	handleErr(err)
	fmt.Println(string(json_byte))
}

func handleErr(err error) {
	if err != nil {
		log.Print(err)
	}
}
