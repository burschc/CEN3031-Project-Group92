package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/browser"
	"log"
	"net/http"
	"os"
	"ufpmp/httpd/sprint1"
)

func main() {
	//Run a command to create a new window using the system's default browser.
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
