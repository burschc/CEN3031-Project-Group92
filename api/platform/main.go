package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
	"ufpmp/httpd/sprint1"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/browser"
)

func main() {
	//Run a command to create a new window using the system's default browser.
	err := browser.OpenURL("http://localhost:8080/api/sprint1")
	if err != nil {
		log.Print(err)
	}

	//Create the GorillaMux router and register some endpoints for the mockup application.
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("LOG!")

		json, err := json.RawMessage(`{"Hello", "World"}`).MarshalJSON()
		if err != nil {
			log.Print("Error when making the json response!")
			log.Print(err)

			//Print an error 503 to let the requester know it can't be done
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		//Return the JSON data to the caller.
		w.Write(json)
	})

	api.HandleFunc("/sprint1", sprint1.PageLoad)
	api.HandleFunc("/map/search", sprint1.SearchPostHandler).Methods("POST")
	api.HandleFunc("/filter/pd", sprint1.FilterPostHandler).Methods("POST")

	//Start the logging middleware and start web server on port 8080.
	rlogged := handlers.LoggingHandler(os.Stdout, r)

	//Create a server with the following properties:
	server := &http.Server{
		Handler: rlogged,
		Addr:    "localhost:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
