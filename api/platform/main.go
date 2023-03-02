package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
	"ufpmp/database"
	"ufpmp/httpd"
	"ufpmp/httpd/app_handlers"
	"ufpmp/httpd/app_handlers/decal_filter"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/browser"
)

var BackendURL = "localhost:"
var BackendPort = "8080"

func main() {

	//Setup the database, or open it if it already exists.
	database.SetupOrOpenBasicDatabase()

	//Create the GorillaMux router and subrouter for general api calls.
	r := mux.NewRouter()

	//Use the Custom-made CORS header middleware tied to the logging middleware.
	//rCustom := handlers.LoggingHandler(os.Stdout, r)

	//Start the logging middleware.
	rCustom := CORSHeaderMiddleware(r)

	//Process any command line arguments.
	for i, arg := range os.Args[1:] {

		//Mockup command line option
		if arg == "-s1" {
			//Run a command to create a new window using the system's default browser.
			err := browser.OpenURL("http://" + BackendURL + BackendPort + "/api/sprint1")
			if err != nil {
				log.Print(err)
			}

			//Register the URLs associated with the sprint 1 mockup.
			//TODO: Fix to also have /api/ in the path.
			//sprint1.RegisterHandlers(r)

			log.Print("Registered Sprint 1 mockup URLs.")
		}

		//Clear the cache
		if arg == "-cc" {
			httpd.ClearJSONCache()
			log.Print("Cleared JSON Cache.")
		}

		if arg == "-p" {
			newBackendPort := os.Args[i+2]
			if newBackendPort != "" {
				log.Print("Changing port from default " + BackendPort + " to " + newBackendPort)
				BackendPort = newBackendPort
			} else {
				log.Print("Port change flag called but no port was given! Using default port " + BackendPort)
			}
		}

	}

	RegisterHandlers(r)

	log.Print("Checking Python Virtual Environment...\n\n")
	SetupPythonVenv()

	//Create a server with the following properties:
	server := &http.Server{
		Handler: rCustom,
		Addr:    BackendURL + BackendPort,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Starting http server accessible through " + server.Addr)
	log.Fatal(server.ListenAndServe())
}

// SetupPythonVenv runs a python script which checks if the virtual environment exists and creates it if it doesn't.
// The python virtual environment is useful for running critical python scripts (like gjf) with no impact or remaining
// files for the user if they later decide to delete/uninstall the application.
func SetupPythonVenv() {
	cmd := exec.Command("python", "python/make_venv.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Print(cmd.Run())
}

// RegisterHandlers registers all functions for the entire web application. It logs a message confirming that all paths
// in the function have been registered. Handlers should be registered in an order where the default is registered last.
func RegisterHandlers(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	decal_filter.DecalFilterHandlers(api)
	app_handlers.HttpHandlers(api)

	app_handlers.DefaultHttpHandler(r)
}

// CORSHeaderMiddleware prepends a default CORS header that will work pretty much anywhere. For browser security, this
// really shouldn't be used in production.
func CORSHeaderMiddleware(r *mux.Router) http.Handler {
	return handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
			"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
			"Cache-Control", "Content-Range", "Range"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
			"X-Requested-With", "If-Modified-Since", "Cache-Control",
			"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
		handlers.MaxAge(86400),
	)(r))
}
