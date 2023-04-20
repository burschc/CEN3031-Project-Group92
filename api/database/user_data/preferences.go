package user_data

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"ufpmp/database"
	"ufpmp/httpd"
	"ufpmp/httpd/app_handlers/decal_filter"
	"ufpmp/httpd/cookies"
)

// DecalResponse is how the JSON passtype/get response is formatted.
type DecalResponse struct {
	Name string
	Type int
}

// PreferencesHandlers registers the endpoints used for getting and setting user preferences.
func PreferencesHandlers(r *mux.Router) {
	r.HandleFunc("/account/set/passtype/{passtype}", setPassType).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/account/get/passtype", getPassType).Methods(http.MethodGet, http.MethodOptions)

	log.Print("Registered preferences handlers.")
}

// GetPassType returns the user's parking pass type. Returns -1 if there are errors while getting the pass type.
func getPassType(w http.ResponseWriter, r *http.Request) {
	//Ensure that the login cookie is valid.
	cookie := cookies.GetLoginCookie(r)
	if cookie == nil {
		log.Printf("Cannot get user preferences since cookie does not exist!")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("No user logged in."))
		return
	}

	//Create a decal response object to populate.
	var response DecalResponse

	//Get the user's account using the username stored in the cookie.
	username := cookies.GetLoginCookie(r).Value
	var passType int

	account := database.Database.Database.QueryRow("SELECT passtype FROM users WHERE username=?", username)
	err := account.Scan(&passType)

	switch {
	case err == sql.ErrNoRows:
		{
			log.Printf("User does not exist in the system.")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	case err != nil:
		{
			httpd.PipeError(w, err)
			return
		}
	}

	response.Type = passType
	response.Name = decal_filter.GetName(passType)
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error when encoding decal response.")
		w.Write([]byte(response.Name))
	}
}

func setPassType(w http.ResponseWriter, r *http.Request) {
	//Ensure that the login cookie is valid.
	cookie := cookies.GetLoginCookie(r)
	if cookie == nil {
		log.Printf("Cannot get user preferences since cookie does not exist!")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("No user logged in."))
		return
	}
	//Process the passType sent in.
	passType := mux.Vars(r)["passtype"]
	passNum := decal_filter.GetIndex(passType)

	//Get the user's account using the username stored in the cookie and make the query.
	username := cookies.GetLoginCookie(r).Value

	//Create a transaction for the database.
	txn, err := database.Database.Database.Begin()
	if err != nil {
		log.Printf("Error encountered when creating database transaction: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		txn.Rollback()
		return
	}

	stmt, err := txn.Prepare("UPDATE users SET passtype=? WHERE username=?")
	if err != nil {
		log.Printf("Error encountered when creating database statement: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		txn.Rollback()
		return
	}

	_, err = stmt.Exec(passNum, username)

	switch {
	case err == sql.ErrNoRows:
		{
			log.Printf("User does not exist in the system.")
			w.WriteHeader(http.StatusUnauthorized)
			txn.Rollback()
			return
		}
	case err != nil:
		{
			httpd.PipeError(w, err)
			txn.Rollback()
			return
		}
	}

	if err = txn.Commit(); err != nil {
		log.Printf("Error commiting the database transaction: %v", err.Error())
		txn.Rollback()
		return
	}

	w.Write([]byte("Updated passType to " + decal_filter.GetName(passNum) + " (" + strconv.Itoa(passNum) + ") for user " + username))
}
