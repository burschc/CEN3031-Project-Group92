package decal_filter

import (
	"encoding/json"
	"github.com/gorilla/mux"
	geojson "github.com/paulmach/go.geojson"
	"log"
	"net/http"
	"os"
	"ufpmp/httpd"
)

// parkingLots is the URL of the parking lots json file used for the decal filter.
// There are two possible json files for parking lots, and this one seems to be the more updated version.
// The other json file can be found here: https://campusmap.ufl.edu/library/cmapjson/parking_lots.json.
const parkingLots = "https://campusmap.ufl.edu/assets/parking_polys.json"

// parkingJSON is the filename of the json file with the parking data located in the json cache folder.
const parkingJSON = "parking_lots.json"

// decalProperty is the title of the feature property which houses the decal for the lot.
const decalProperty = "Lot_Class"

// DecalFilterHandlers registers the functions which deal with the parking decal filters.
// It logs a message confirming that all paths in the function have been registered.
func DecalFilterHandlers(r *mux.Router) {
	r.HandleFunc("/filter/decal/{decal}", findDecalHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/filter/decals", decalTypesHandler).Methods(http.MethodGet, http.MethodOptions)

	log.Print("Registered filter handlers.")
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//												HTTP HANDLERS														  //
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// decalTypesHandler returns a list of all decal types in the json file as a json array to the requester.
func decalTypesHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(getDecalTypes()); err != nil {
		httpd.PipeError(w, err)
	}
}

// findDecalHandler returns a feature collection consisting of all lots which match the given decal.
func findDecalHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(findDecal(vars["decal"])); err != nil {
		httpd.PipeError(w, err)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//												UTILITY FUNCTIONS													  //
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// findDecal looks through the parking_lots json file for lots which are valid for the request's passed in decal.
// It returns a feature collection.
func findDecal(decal string) *geojson.FeatureCollection {
	//Get the JSON file if we already do not have it in the cache.
	if !httpd.IsFresh(parkingJSON) {
		log.Print(parkingJSON + " is not fresh!")
		getNewJSON()
	}

	/*Make a new feature collection and then sift through the json file's feature collection for entries that
	  have the same decal property value as wwe are looking for.*/
	filteredLots := geojson.NewFeatureCollection()

	//Open the json file as a feature collection and return that feature collection.
	fc := httpd.ConvertToFC(parkingJSON)

	//If the decal is 'any' return the full feature collection.
	if decal == "any" {
		return fc
	}

	//Scan through and isolate those features whose decal property matches our target decal.
	for _, v := range fc.Features {
		if v.Properties[decalProperty] == decal {
			filteredLots.AddFeature(v)
		}
	}

	return filteredLots
}

// getDecalTypes returns an array containing all the unique decals in the parking lots json file.
func getDecalTypes() []interface{} {
	//Get the JSON file if we already do not have it in the cache.
	if !httpd.IsFresh(parkingJSON) {
		log.Print(parkingJSON + " is not fresh!")
		getNewJSON()
	}

	set := make(map[interface{}]bool)
	var decals []interface{}

	//Unmarshall the file into a feature collection we can traverse.
	fc := httpd.ConvertToFC(parkingJSON)

	//Traverse through and find unique decal types using a set.
	for _, v := range fc.Features {
		if !set[v.Properties[decalProperty]] && v.Properties[decalProperty] != nil {
			decals = append(decals, v.Properties[decalProperty])
			set[v.Properties[decalProperty]] = true
		}
	}

	return decals
}

func getNewJSON() {
	//Try to see if we can remove the current json file.
	if err := os.Remove(httpd.JsonCachePath + parkingJSON); err != nil {
		log.Print(err)
	}

	//Try to see if we can grab a new version of the json file. If we can't, we shouldn't try to validate anything.
	if err := httpd.GetJSONFromURL(parkingLots, parkingJSON); err == nil {
		httpd.ValidateGeoJson(parkingJSON)
	}
}
