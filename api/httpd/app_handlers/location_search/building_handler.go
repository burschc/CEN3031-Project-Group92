package location_search

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	r.HandleFunc("/search/online/{location}", onlineLocationHandler).Methods(http.MethodGet, http.MethodOptions)

	log.Print("Registered online and offline search handlers.")
}

// offlineLocationHandler
func offlineLocationHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// onlineLocationHandler
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
