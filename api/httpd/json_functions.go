package httpd

import (
	"errors"
	geojson "github.com/paulmach/go.geojson"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"ufpmp/python"
)

// JsonCachePath is the default folder to place json files that the application needs.
const JsonCachePath = "cache/json/"

// PythonScripts is the folder containing all project-specific created python code.
const PythonScripts = "python/"

// DefaultUpdateTime is the default time, in parsable form, to auto-update the geojson file for a non-logged in user.
const DefaultUpdateTime = "24h"

type Validation func(string)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//											PUBLIC UTILITY FUNCTIONS												  //
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetNewJSON(filename string, url string, validation Validation) {
	if !IsFresh(filename) {
		log.Printf("%v is not fresh (older than %v/doesn't exist)! Getting a new one...", filename, DefaultUpdateTime)
		//Try to see if we can remove the current json file.
		if err := os.Remove(JsonCachePath + filename); err != nil {
			log.Printf("Could not remove %v: %v", filename, err.Error())
		}

		//Try to see if we can grab a new version of the json file. If we can't, we shouldn't try to validate anything.
		if err := GetJSONFromURL(url, filename); err == nil {
			log.Printf("Running any additional processing functions...")
			validation(filename)
		}
	}
}

// GetJSONFromURL gets a json file from an url and stores it in the local backend cache.
func GetJSONFromURL(jsonURL string, filename string) error {
	log.Printf("Acquiring JSON file %v from %v...", filename, jsonURL)

	//Check that the Json Cache Path exists.
	if _, err := os.Stat(JsonCachePath); err != nil {
		log.Print("cache path " + JsonCachePath + " does not exist. Creating...")
		if err = os.MkdirAll(JsonCachePath, os.ModePerm); err != nil {
			log.Fatal(err)
			return err
		}
	}

	//Check if the file already exists in the cache.
	if _, err := os.Stat(JsonCachePath + filename); err == nil {
		log.Print("file " + filename + " exists in the cache.")
		return nil
	}

	//Attempt to get the json file from the url.
	get, err := http.Get(jsonURL)

	if err != nil {
		log.Print(err)
	}

	defer get.Body.Close()

	//Double check that the file is a JSON file (through header only).
	//TODO: Possibly implement full JSON verification in the future.
	fileType := get.Header.Values("Content-Type")[0]
	if !strings.Contains(fileType, "application/json") {
		log.Print("URL " + jsonURL + " does not point to a JSON file")
		return errors.New("URL " + jsonURL + " does not point to a JSON file")
	}

	//Create a file in the Json Cache path and copy over the data to it.
	file, err := os.Create(JsonCachePath + filename)
	if err != nil {
		log.Print(err)
	}

	if _, err = io.Copy(file, get.Body); err != nil {
		log.Print(err)
	}

	//Make sure to close the file and connection afterward.
	if err = file.Close(); err != nil {
		log.Print(err)
	}

	if err = get.Body.Close(); err != nil {
		log.Print(err)
	}

	return err
}

// FileToFC searches the json cache for a specified geojson file and returns it in FeatureCollection format.
func FileToFC(filename string) *geojson.FeatureCollection {
	data := geojson.NewFeatureCollection()

	//If the file does not exist, return a blank feature collection.
	if _, err := os.Stat(JsonCachePath + filename); err != nil {
		log.Print("File does not exist. Returning a blank feature collection...")
		return data
	}

	/*Open the geojson file and return its unmarshalled feature collection. the geojson library automatically validates
	  the files presented to it from what I gathered looking at the library's code.*/

	file, err := os.ReadFile(JsonCachePath + "parking_lots.json")
	if err != nil {
		log.Print(err)
		return data
	}

	data, err = geojson.UnmarshalFeatureCollection(file)
	if err != nil {
		log.Print(err)
	}

	return data
}

// FCToFile saves a feature collection to a file in the json cache.
func FCToFile(filename string, fc *geojson.FeatureCollection) error {
	//Detect if a file with the given name exists.
	if _, err := os.Stat(filename); err == nil {
		log.Print("File " + filename + " exists! Overwriting file with new feature collection...")
	}

	//Marhsall the JSON to a byte array and overwrite the json file with it.
	data, err := fc.MarshalJSON()
	if err != nil {
		return err
	}

	if err = os.WriteFile(JsonCachePath+filename, data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// ValidateGeoJson will run gjf (a python script) on the target geojson file to fix any issues it may have.
func ValidateGeoJson(filename string) {
	if python.IgnorePython {
		return
	}

	//Validate the file using gjf (overwrite flag crashes the program when executing from golang).
	log.Print("Calling gjf on " + filename)

	osName := runtime.GOOS

	if osName == "windows" {
		log.Print("Running gjf as a Windows venv Python script file.")
		cmd := exec.Command(python.PythonVenv+"Scripts/gjf", JsonCachePath+filename)
		log.Print(cmd.Run())
	} else if (osName == "darwin") || (osName == "linux") {
		log.Print("Running gjf as a Linux/Darwin venv Python script file.")
		cmd := exec.Command(python.PythonVenv+"bin/gjf", JsonCachePath+filename)
		log.Print(cmd.Run())
	} else {
		log.Print("Unsupported OS for venv Python script. The JSON File will NOT be fixed.")
		return
	}

	//If there is a fixed version of the file, replace the old version with the fixed version.
	if _, err := os.Stat(JsonCachePath + filename); err == nil {
		if err = os.Remove(JsonCachePath + filename); err != nil {
			log.Print(err)
			return
		}

		if err = os.Rename(JsonCachePath+strings.TrimSuffix(filename, filepath.Ext(filename))+"_fixed.json",
			JsonCachePath+filename); err != nil {
			log.Print(err)
		}
	}
}

// ClearJSONCache removes all JSON files from the backend cache folder.
func ClearJSONCache() {
	files, err := filepath.Glob(JsonCachePath + "/*.json")
	if err != nil {
		log.Print(err)
		return
	}

	for _, file := range files {
		if err := os.Remove(file); err != nil {
			log.Print(err)
		}
	}

}

// IsFresh will check the json file to see if it is newer than the automatic update time and will return true if it is
// newer. This is vitally important as gjf WILL BREAK the json file if it tries to fix errors more than once.
func IsFresh(filename string) bool {
	//Todo: Check with frontend if leaflet accepts validated or raw geojson (geojsonlint likes gjf output only)

	//Get the file's statistics. If the file does not exist, return false regardless.
	filestat, err := os.Stat(JsonCachePath + filename)
	if err != nil {
		log.Print(err)
		return false
	}

	age := time.Since(filestat.ModTime())

	//Todo: once accounts exist, request the user's preferred update time and replace defaultUpdateTime with it.
	limit, err := time.ParseDuration(DefaultUpdateTime)
	if err != nil {
		log.Print(err)
		return false
	}

	return age < limit
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//											PRIVATE UTILITY FUNCTIONS												  //
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
