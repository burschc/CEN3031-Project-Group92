package httpd

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const JsonCachePath = "cache/json/"

// GetJSONFromURL gets a json file from an url and stores it in the local backend cache.
func GetJSONFromURL(jsonURL string, fileName string) {
	//Check that the Json Cache Path exists.
	if _, err := os.Stat(JsonCachePath); err != nil {
		log.Print("cache path " + JsonCachePath + " does not exist. Creating...")
		if err = os.MkdirAll(JsonCachePath, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	//Check if the file already exists in the cache.
	if _, err := os.Stat(JsonCachePath + fileName); err == nil {
		log.Print("file " + fileName + " exists in the cache.")
		return
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
		return
	}

	//Create a file in the Json Cache path and copy over the data to it.
	file, err := os.Create(JsonCachePath + fileName)
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
