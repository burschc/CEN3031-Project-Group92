package decal_filter

import (
	"compress/gzip"
	"encoding/json"
	"github.com/gorilla/mux"
	geojson "github.com/paulmach/go.geojson"
	"log"
	"net/http"
	"os"
	"ufpmp/httpd"
)

// parkingLots is the URL of the parkings lots json file used for the decal filter.
const parkingLots = "https://campusmap.ufl.edu/library/cmapjson/parking_lots.json"

// DecalFilterHandlers registers the functions which deal with the parking decal filters.
func DecalFilterHandlers(r *mux.Router) {
	r.HandleFunc("/filter/decal/{decal}", findDecal)
	r.HandleFunc("/filter/decals/", getDecalTypes)
	r.HandleFunc("/filter/error/", func(w http.ResponseWriter, r *http.Request) {
		httpd.PipeError(w, gzip.ErrChecksum)
	})

	log.Print("Registered filter handlers.")
}

// findDecal looks through the parking_lots json file for lots which are valid for the request's passed in decal.
func findDecal(w http.ResponseWriter, r *http.Request) {
	//Extract the variables which should include the decal from the request.
	vars := mux.Vars(r)

	//Get the JSON file if we already do not have it in the cache.
	httpd.GetJSONFromURL(parkingLots, "parking_lots.json")

	//Look in the JSON file for the lots that are accepted by the variable passed in.
	if vars["decal"] == "any" {
		file, err := os.ReadFile(httpd.JsonCachePath + "parking_lots.json")
		if err != nil {
			httpd.PipeError(w, err)
			return
		}

		_, err = w.Write(file)
		if err != nil {
			httpd.PipeError(w, err)
		}
	} else {
		//Make a new JSON file with these isolated lots and write it as a response.
		file, err := os.ReadFile(httpd.JsonCachePath + "parking_lots.json")
		if err != nil {
			httpd.PipeError(w, err)
			return
		}

		fc, err := geojson.UnmarshalFeatureCollection(file)
		if err != nil {
			httpd.PipeError(w, err)
			return
		}

		filteredLots := geojson.NewFeatureCollection()

		for _, v := range fc.Features {
			if v.Properties["DECAL"] == vars["decal"] {
				filteredLots.AddFeature(v)
			}
		}

		toWrite, err := filteredLots.MarshalJSON()
		if err != nil {
			httpd.PipeError(w, err)
		}

		_, err = w.Write(toWrite)
		if err != nil {
			httpd.PipeError(w, err)
		}
	}

	log.Print(vars["decal"])
}

// getDecalTypes returns an array containing all the unique decals in the parking lots json file.
func getDecalTypes(w http.ResponseWriter, _ *http.Request) {
	//Get the JSON file if we already do not have it in the cache.
	httpd.GetJSONFromURL(parkingLots, "parking_lots.json")

	file, err := os.ReadFile(httpd.JsonCachePath + "parking_lots.json")
	if err != nil {
		httpd.PipeError(w, err)
		return
	}

	fc, err := geojson.UnmarshalFeatureCollection(file)
	if err != nil {
		httpd.PipeError(w, err)
		return
	}

	set := make(map[interface{}]bool)
	var decals []interface{}

	for _, v := range fc.Features {
		if !set[v.Properties["DECAL"]] {
			decals = append(decals, v.Properties["DECAL"])
			set[v.Properties["DECAL"]] = true
		}
	}

	err = json.NewEncoder(w).Encode(decals)
	if err != nil {
		httpd.PipeError(w, err)
	}
}
