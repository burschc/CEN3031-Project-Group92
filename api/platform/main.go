package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
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
			log.Print("Cleared JSON Cache.")
		}

	}

	app_handlers.HttpHandlers(api)
	decal_filter.DecalFilterHandlers(api)

	log.Print("Checking Python Virtual Environment...\n\n")
	SetupPythonVenv()

	//Create a server with the following properties:
	server := &http.Server{
		Handler: rLogged,
		Addr:    "localhost:8080",

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
	cmd := exec.Command("py", "python/make_venv.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Print(cmd.Run())
}
