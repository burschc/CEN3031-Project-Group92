package app_handler

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Main Page")
	if err != nil {
		log.Fatal(err)
	}
}
