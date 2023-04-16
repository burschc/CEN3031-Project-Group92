package python

import (
	"log"
	"os"
	"os/exec"
	"ufpmp/httpd"
)

// SetupPythonVenv runs a python script which checks if the virtual environment exists and creates it if it doesn't.
// The python virtual environment is useful for running critical python scripts (like gjf) with no impact or remaining
// files for the user if they later decide to delete/uninstall the application.
func SetupPythonVenv() {
	log.Print("Checking Python Virtual Environment...\n\n")

	cmd := exec.Command("python", "python/make_venv.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Print(cmd.Run())
}

// RemovePythonVenv removes the Python virtual environment by deleting the venv folder.
func RemovePythonVenv() {
	if err := os.RemoveAll(httpd.PythonVenv); err != nil {
		log.Print(err.Error())
	}
}
