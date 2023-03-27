package app_handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ufpmp/httpd"
)

type APIInfo struct {
	Version string   `json:"version"`
	Authors []Author `json:"authors"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
}

// HttpHandlers ties the URL path and methods to the correct function.
func HttpHandlers(r *mux.Router) {
	r.HandleFunc("/version", versionHandler)

	log.Print("Registered base handlers.")
}

func DefaultHttpHandler(r *mux.Router) {
	//r.HandleFunc("/").Handler()
}

// testHandler returns a json response when going to /api/test
func versionHandler(w http.ResponseWriter, _ *http.Request) {
	info := APIInfo{
		Version: "< 1.0.0",
		Authors: []Author{
			{"Yovany Molina", "yomole"},
			{"Natalie Valcin", "natalievalcin"},
			{"Sam Barthelemy", "sbarthelemy01"},
			{"Christopher Bursch", "burschc"},
		},
	}

	response, err := json.Marshal(info)
	if err != nil {
		httpd.PipeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	//Return the JSON data to the caller.
	_, err = w.Write(response)
	if err != nil {
		httpd.PipeError(w, err)
	}
}

// angularHandler passes along a blank request to angular.
func angularHandler(r *mux.Router) {
	r.Use(mux.CORSMethodMiddleware(r))
}
