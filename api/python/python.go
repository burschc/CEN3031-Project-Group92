package python

import (
	"log"
	"os"
	"os/exec"
)

// PythonCmd is the commmand line python executable used for the web app (by default it is "python").
var PythonCmd = "python"

// IgnorePython is a boolean which describes whether to use aspects of the web app that require python.
var IgnorePython = false

// PythonVenv is the folder containing the project's python virtual environment.
const PythonVenv = "python/venv/"

// SetupPythonVenv runs a python script which checks if the virtual environment exists and creates it if it doesn't.
// The python virtual environment is useful for running critical python scripts (like gjf) with no impact or remaining
// files for the user if they later decide to delete/uninstall the application.
func SetupPythonVenv() {
	cmd := exec.Command(PythonCmd, "python/make_venv.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Print(cmd.Run())
}

// RemovePythonVenv removes the Python virtual environment by deleting the venv folder.
func RemovePythonVenv() {
	if err := os.RemoveAll(PythonVenv); err != nil {
		log.Print(err.Error())
	}
}
