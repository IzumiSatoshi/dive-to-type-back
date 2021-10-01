package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"dive-to-type-back/database"
)

type Thread struct {
	UUID        string
	Title       string
	Description string
	Created_at  string
}

func main() {
	database.InitDB()
	defer database.CloseDB()

	http.HandleFunc("/", ThreadsRequestHandler)

	log.Print(http.ListenAndServe("localhost:8080", nil))
}

func ThreadsRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetThreads(w, r)
	case "POST":
		InsertThread(w, r)
	}
}

func GetThreads(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM threads")
	handleErr(err)

	threads := []Thread{}
	for rows.Next() {
		thread := Thread{}
		var err = rows.Scan(&thread.UUID, &thread.Title, &thread.Description, &thread.Created_at)
		handleErr(err)
		threads = append(threads, thread)
	}

	json_byte, err := json.Marshal(threads)
	handleErr(err)
	fmt.Println(string(json_byte))

	w.Write(json_byte)

}
func InsertThread(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var thread Thread
	err := decoder.Decode(&thread)
	if err != nil {
		log.Print(err)
	}
	log.Println(thread.Title)
	log.Println(thread.Description)
}

func handleErr(err error) {
	if err != nil {
		log.Print(err)
	}
}
