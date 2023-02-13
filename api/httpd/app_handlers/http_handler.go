package app_handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// RegisterHandlers ties the URL path and methods to the correct function.
func HttpHandlers(r *mux.Router) {
	r.HandleFunc("/test", TestHandler)
}

func TestHandler(w http.ResponseWriter, _ *http.Request) {
	log.Print("LOG!")

	response, err := json.RawMessage(`{"Hello", "World"}`).MarshalJSON()
	if err != nil {
		log.Print("Error when making the json response!")
		log.Print(err)

		//Print an error 503 to let the requester know it can't be done
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	//Return the JSON data to the caller.
	_, err = w.Write(response)
	if err != nil {
		log.Print(err)
	}
}
