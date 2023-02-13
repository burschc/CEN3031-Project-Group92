package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"ufpmp/httpd"
	"ufpmp/httpd/app_handlers"
	"ufpmp/httpd/app_handlers/decal_filter"
	"ufpmp/httpd/sprint1"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/browser"
)

func main() {

	//Create the GorillaMux router and subrouter for general api calls.
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	//Start the logging middleware.
	rLogged := handlers.LoggingHandler(os.Stdout, r)

	//Process any command line arguments.
	for _, arg := range os.Args[1:] {

		//Mockup command line option
		if arg == "-m" {
			//Run a command to create a new window using the system's default browser.
			err := browser.OpenURL("http://localhost:8080/api/sprint1")
			if err != nil {
				log.Print(err)
			}

			//Register the URLs associated with the sprint 1 mockup.
			sprint1.RegisterHandlers(api)

			log.Print("Registered Sprint 1 mockup URLs.")
		}

		//Clear the cache
		if arg == "-cc" {
			httpd.ClearJSONCache()
		}

	}

	app_handlers.HttpHandlers(api)
	decal_filter.DecalFilterHandlers(api)

	//Create a server with the following properties:
	server := &http.Server{
		Handler: rLogged,
		Addr:    "localhost:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
