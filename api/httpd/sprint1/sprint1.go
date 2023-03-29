package sprint1

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"ufpmp/httpd"
)

var sprintOneSite = template.Must(template.ParseGlob("frontMockup/*.html"))

// RegisterHandlers ties the URL path and methods to the correct function.
func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/sprint1", pageLoad)
	r.HandleFunc("/map/search", searchPostHandler).Methods("POST")
	r.HandleFunc("/filter/pd", filterPostHandler).Methods("POST")

	//Handler for favicon requests. This is also useful to tie any static files we wish to serve to frontend.
	r.PathPrefix("/api/favicon/").Handler(http.StripPrefix("/api/favicon",
		http.FileServer(http.Dir("frontMockup/favicon"))))

	//The most important header code. Handle with extreme caution.
	r.HandleFunc("/teapot", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		_, _ = w.Write([]byte("Short and stout"))
	})
}

// pageLoad loads the mockup of the UF Parking Plus application page made for sprint 1.
func pageLoad(w http.ResponseWriter, _ *http.Request) {
	if err := sprintOneSite.ExecuteTemplate(w, "index.html", nil); err != nil {
		httpd.PipeError(w, err)
	}
}

// searchPostHandler is responsible for processing user input at the top search bar.
func searchPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Not implemented yet!")

	if err := r.ParseForm(); err != nil {
		httpd.PipeError(w, err)
	}

	searchRequest := r.PostForm.Get("map/search")

	log.Print("Received search request for : " + searchRequest)

	_, _ = w.Write([]byte(searchRequest))
}

// filterPostHandler is responsible for processing user input for the filters.
func filterPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Not implemented yet!")

	if err := r.ParseForm(); err != nil {
		httpd.PipeError(w, err)
	}

	decal := r.PostForm.Get("filter/pd")

	log.Print("Decal selected was: " + decal)

	_, _ = w.Write([]byte(decal))
}
