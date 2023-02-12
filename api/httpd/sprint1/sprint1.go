package sprint1

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var sprintOneSite = template.Must(template.ParseGlob("frontMockup/*.html"))

const faviconName = "/icons8-map-marker-material-filled-32.ico"

// RegisterHandlers ties the URL path and methods to the correct function.
func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/sprint1", pageLoad)
	r.HandleFunc("/map/search", searchPostHandler).Methods("POST")
	r.HandleFunc("/filter/pd", filterPostHandler).Methods("POST")

	r.HandleFunc("/favicon"+faviconName, faviconHandler)
	r.HandleFunc("/teapot", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		_, err := w.Write([]byte("Short and stout"))
		if err != nil {
			log.Print(err)
		}
	})
}

// pageLoad loads the mockup of the UF Parking Plus application page made for sprint 1.
func pageLoad(w http.ResponseWriter, _ *http.Request) {

	err := sprintOneSite.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func faviconHandler(w http.ResponseWriter, _ *http.Request) {
	log.Print("Test")
}

// searchPostHandler is responsible for processing user input at the top search bar.
func searchPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Not implemented yet!")

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	searchRequest := r.PostForm.Get("map/search")

	log.Print("Received search request for : " + searchRequest)

	//Reload the initial page.
	//PageLoad(w, r)
	_, err = w.Write([]byte(searchRequest))
	if err != nil {
		log.Print(err)
	}
}

// filterPostHandler is responsible for processing user input for the filters.
func filterPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Not implemented yet!")

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	decal := r.PostForm.Get("filter/pd")

	log.Print("Decal selected was: " + decal)

	_, err = w.Write([]byte(decal))
	if err != nil {
		log.Print(err)
	}
}
