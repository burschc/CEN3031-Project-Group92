package decal_filter

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ufpmp/httpd"
)

const parkingLots = "https://campusmap.ufl.edu/library/cmapjson/parking_lots.json"

// LotJSON is the structure representing the layout of the parking_lots.json file.
type LotJSON struct {
	JsonType string `json:"type"`
	Features []struct {
		FeatureType       string `json:"feature"`
		FeatureProperties struct {
			Jtype    string `json:"JTYPE"`
			Decal    string `json:"DECAL"`
			ObjectID string `json:"OBJECTID"`
		}
		Geometry []struct {
			GeometryType string          `json:"type"`
			Coordinates  [][][][]float32 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

// DecalFilterHandlers registers the functions which deal with the parking decal filters.
func DecalFilterHandlers(r *mux.Router) {
	r.HandleFunc("/filter/{decal}", findDecal)
}

// findDecal looks through the parking_lots json file for lots which are valid for the request's passed in decal.
func findDecal(w http.ResponseWriter, r *http.Request) {
	//Extract the variables which should include the decal from the request.
	vars := mux.Vars(r)

	//Get the JSON file if we already do not have it in the cache.
	httpd.GetJSONFromURL(parkingLots, "parking_lots.json")

	//Look in the JSON file for the lots that are accepted by the variable passed in.

	//Make a new JSON file with these isolated lots and write it as a response.

	log.Print(vars["decal"])
}
