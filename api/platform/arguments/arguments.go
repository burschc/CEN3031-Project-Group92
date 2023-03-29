package arguments

import (
	"flag"
	"github.com/pkg/browser"
	"log"
	"time"
	"ufpmp/httpd"
	"ufpmp/httpd/mux_functions"
	"ufpmp/httpd/sprint1"
	"ufpmp/python"
)

// GracefulWait is the duration for which the server gracefully wait for existing connections to finish
var GracefulWait time.Duration

func ProcessArguments() {
	flag.DurationVar(&GracefulWait, "graceful-timeout", time.Second*15, "the duration for which the server "+
		"gracefully wait for existing connections to finish - e.g. 15s or 1m")

	flag.StringVar(&mux_functions.ServerProperties.BackendPort, "port", "8080", "Sets the port used "+
		"for the web app backend server.")

	flagClearCache := flag.Bool("clear-cache", false, "Clears the JSON cache by removing the cache "+
		"folder entirely. Useful to refresh all of the JSON files used in the web application ahead of the refresh "+
		"deadline.")

	flagRemoveVenv := flag.Bool("remove-venv", false, "Removes the Python virtual environment by "+
		"deleting the venv folder.")

	flagSprint1 := flag.Bool("sprint1", false, "Registers sprint 1 endpoints and sends the default "+
		"browser to the sprint1 site index.")

	flag.Parse()

	//

	if *flagClearCache {
		httpd.ClearJSONCache()
		log.Print("Cleared JSON Cache.")
	}

	if *flagRemoveVenv {
		python.RemovePythonVenv()
		python.SetupPythonVenv()
		log.Print("Removed python virtual environment.")
	}

	if *flagSprint1 {
		//Run a command to create a new window using the system's default browser.
		err := browser.OpenURL("http://localhost:" + mux_functions.ServerProperties.BackendPort + "/api/sprint1")
		if err != nil {
			log.Fatal(err.Error())
		}

		//Register the URLs associated with the sprint 1 mockup.
		//TODO: Fix to also have /api/ in the path.
		mux_functions.ServerProperties.ContingentRoutes =
			append(mux_functions.ServerProperties.ContingentRoutes, sprint1.RegisterHandlers)

		log.Print("Registered Sprint 1 mockup URLs.")
	}

}
