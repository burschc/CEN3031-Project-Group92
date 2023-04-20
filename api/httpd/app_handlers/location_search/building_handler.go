package location_search

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"ufpmp/httpd"
)

// locationsURL is the URL of the locationsJSON file used for building search.
// There are two possible json files for buildings, and this one seems to be the more updated version.
// The other JSON file can be found here: https://campusmap.ufl.edu/library/cmapjson/geo_buildings.json.
// Note that the syntax for using this URL and the existing campusmap API is [...]/searchBldg?origin=*&srch={location}
const locationsURL = "https://campusmap.ufl.edu/library/api/searchBldg"

// onlineBoilerplate is the required information after the locationsURL for online searches minus the search term.
const onlineBoilerplate = "?origin=*&srch="

// locationsJSON is the filename of the json file with the parking data located in the json cache folder.
const locationsJSON = "geo_buildings.json"

// defaultLocationFilter is the name of the field in the location struct that we filter by default when getting
// locations offline.
const defaultLocationFilter = "Name"

// locationWildcard is the symbol used to make the web app return all locations.
const locationWildcard = "*"

// Location is the structure of the locations JSON file using in JSON encoding and decoding.
type Location struct {
	Code         string  `json:"BLDGCODE"`
	ID           string  `json:"BLDG"`
	Name         string  `json:"NAME"`
	Abbreviation string  `json:"ABBREV"`
	OfficialName string  `json:"OFFICIAL_ROOM_NAME"`
	Latitude     float32 `json:"LAT"`
	Longitude    float32 `json:"LON"`
}

// LocationSearchHandlers register the back-end handlers used for the location search feature.
func LocationSearchHandlers(r *mux.Router) {
	r.HandleFunc("/search/offline/{location}", offlineLocationHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/search/offline/{location}/{filter}", offlineLocationHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/search/online/{location}", onlineLocationHandler).Methods(http.MethodGet, http.MethodOptions)

	log.Print("Registered online and offline search handlers.")
}

// offlineLocationHandle encodes the result of the getLocationsOffline function into a JSON response which gets sent back
// to the requester.
func offlineLocationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["filter"] == "" {
		vars["filter"] = defaultLocationFilter
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(getLocationsOffline(vars["location"], vars["filter"])); err != nil {
		httpd.PipeError(w, err)
	}
}

// onlineLocationHandler encodes the result of the getLocationsOnline function into a JSON response which gets sent back
// to the requester.
func onlineLocationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(getLocationsOnline(vars["location"])); err != nil {
		httpd.PipeError(w, err)
	}
}

// getLocationsOnline sends a response to the online campus map API endpoint for location searches and returns the response
// in a pre-written struct format.
func getLocationsOnline(location string) []Location {
	//See if the search term was an asterisk. If so, change the search term to be an empty string to return all.
	if location == locationWildcard {
		location = ""
	}

	//Get the corresponding JSON file from the online API.
	get, err := http.Get(strings.Replace(locationsURL+onlineBoilerplate+location, " ", "%20", -1))
	if err != nil {
		log.Printf("Error getting JSON for location '%v' from online API search: %v", location, err.Error())
		return nil
	}

	defer get.Body.Close()

	//Create the array of location structs used to store the locations found by the string.
	var data []Location

	//Decode the response into the array of location structs.
	decoder := json.NewDecoder(get.Body)
	err = decoder.Decode(&data)
	if err != nil {
		log.Printf("Error in JSON for location '%v' from online API search: %#v.", location, err)

		if err, ok := err.(*json.SyntaxError); ok {
			//Print the offset.
			log.Printf("Syntax error occured at offset %v. This could have been caused by an errant space in the "+
				"request.", err.Offset)
		}
		return nil
	}

	//Close the response like a good little developer.
	err = get.Body.Close()
	if err != nil {
		log.Printf("Error closing response from online API search: %v", err.Error())
		return nil
	}

	return data
}

// getLocationsOffline first tries to get a new JSON file if the locations JSON is not in the cache or is too old. Then,
// it sifts through an encoded version of the file for specified locations and returns the result as a Location struct array.
func getLocationsOffline(location string, filter string) []Location {
	//Get the JSON file if we already do not have it in the cache.
	httpd.GetNewJSON(locationsJSON, locationsURL, func(string) {})

	//Open the file to encode it in JSON format.
	reader, err := os.Open(httpd.JsonCachePath + locationsJSON)
	if err != nil {
		log.Printf("Failed to open the file '%v': %v", locationsJSON, err.Error())
		return nil
	}

	defer reader.Close()

	//Encode the data in the location array format.
	var data []Location
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&data)
	if err != nil {
		log.Printf("Failed to encode '%v' into the locations structure: %v", locationsJSON, err.Error())
		return nil
	}

	//See if the search term was an asterisk. If so, return the entire array.
	if location == locationWildcard {
		return data
	}

	//Sift through the array for any building names that contain the supplied string.
	var filteredData []Location

	for _, v := range data {
		field := reflect.ValueOf(v).FieldByName(filter)

		if strings.Contains(field.String(), location) {
			filteredData = append(filteredData, v)
		}
	}

	//Close the reader like a good little developer.
	if err = reader.Close(); err != nil {
		log.Printf("Failed to close the reader for the locations JSON file: %v", err.Error())
	}

	return filteredData
}
