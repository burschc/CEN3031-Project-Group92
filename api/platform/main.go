package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"ufpmp/httpd/sprint1"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/browser"
)

func main() {
	//Run a command to create a new window using the system's default browser.
	database, _ := sql.Open("sqlite3", "./nraboy.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Nic", "Raboy")
	rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}

	err := browser.OpenURL("http://localhost:8080/sprint1")
	if err != nil {
		log.Print(err)
	}

	//Create the GorillaMux router and register some endpoints for the mockup application.
	r := mux.NewRouter()

	r.HandleFunc("/sprint1", sprint1.PageLoad)
	r.HandleFunc("/map/search", sprint1.SearchPostHandler).Methods("POST")
	r.HandleFunc("/filter/pd", sprint1.FilterPostHandler).Methods("POST")

	//Start the logging middleware and start web server on port 8080.
	rlogged := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(":8080", rlogged))
}
