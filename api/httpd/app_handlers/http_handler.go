package app_handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ufpmp/httpd"
)

// HttpHandlers ties the URL path and methods to the correct function.
func HttpHandlers(r *mux.Router) {
	r.HandleFunc("/test", testHandler)

	log.Print("Registered base handlers.")
}

func DefaultHttpHandler(r *mux.Router) {
	//r.HandleFunc("/").Handler()
}

// testHandler returns a json response when going to /api/test
func testHandler(w http.ResponseWriter, _ *http.Request) {
	log.Print("LOG!")

	response, err := json.RawMessage(`{"Hello", "World"}`).MarshalJSON()
	if err != nil {
		httpd.PipeError(w, err)
		return
	}

	//Return the JSON data to the caller.
	_, _ = w.Write(response)
}

// angularHandler passes along a blank request to angular.
func angularHandler(r *mux.Router) {
	r.Use(mux.CORSMethodMiddleware(r))

}
