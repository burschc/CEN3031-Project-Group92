package main

import (
	"github.com/pkg/browser"
	"log"
	"os"
	"os/exec"
	"ufpmp/database"
	"ufpmp/httpd"
	"ufpmp/httpd/mux_functions"
)

// BackendPort is the default backend port for the web app.
var BackendPort = "8080"

func main() {

	//Set up the database, or open it if it already exists.
	database.SetupOrOpenBasicDatabase()

	//Process any command line arguments.
	for i, arg := range os.Args[1:] {

		//Mockup command line option
		if arg == "-s1" {
			//Run a command to create a new window using the system's default browser.
			err := browser.OpenURL("http://localhost:" + BackendPort + "/api/sprint1")
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

	//Make sure Python is installed and the virtual enviornment with all the required scripts is available.
	log.Print("Checking Python Virtual Environment...\n\n")
	SetupPythonVenv()

	//Create the router and server for the web app and register the handlers for the mux_functions.
	r, server := mux_functions.CreateAppServer(mux_functions.AppServerProperties{
		BackendURL:  "localhost",
		BackendPort: ":" + BackendPort,
	})
	mux_functions.RegisterHandlers(r)

	log.Print("Started http server accessible through " + server.Addr)
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
