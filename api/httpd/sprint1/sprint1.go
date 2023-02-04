package sprint1

import (
	"html/template"
	"log"
	"net/http"
)

var sprintOneSite = template.Must(template.ParseGlob("frontMockup/*.html"))

// PageLoad loads the mockup of the UF Parking Plus application page made for sprint 1.
func PageLoad(w http.ResponseWriter, r *http.Request) {

	err := sprintOneSite.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// SearchPostHandler is responsible for processing user input at the top search bar.
func SearchPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Not implemented yet!")

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	searchRequest := r.PostForm.Get("/map/search")
	log.Print("Received search request for : " + searchRequest)

	//Reload the initial page.
	PageLoad(w, r)
}

// FilterPostHandler is responsible for processing user input for the filters.
func FilterPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Not implemented yet!")

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	decal := r.PostForm.Get("filter/pd")

	log.Print("Decal selected was: " + decal)

	PageLoad(w, r)
}
