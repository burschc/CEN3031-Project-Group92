package database

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

var r *mux.Router
var client *http.Client

const username1 = "jdoe"
const username2 = "appledontsue"

const password1 = "password"
const password2 = "ohno"

// init changes to the top level directory for the api (/api/) for each test and creates a new mux router for each test.
func init() {
	if err := os.Chdir("../"); err != nil {
		log.Fatal(err.Error())
	}

	//To avoid circular dependencies, the handler registration has to occur here instead of calling the code that
	//already exists.
	r = mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	DatabaseHandlers(api)

	DeclareDatabase("test.db")
}

func TestSignup(t *testing.T) {

	const endpoint = "/api/signup"

	t.Run("Should create user jdoe", func(t *testing.T) {
		form := url.Values{}
		form.Add("username", username1)
		form.Add("password", password1)
		//form.Add("signup", "Signup")

		request := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(form.Encode()))
		request.Form = form

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		//Get the response and read it as a byte array.
		res := recorder.Result()
		defer res.Body.Close()

		printDatabase()

		if res.StatusCode != http.StatusOK {
			t.Errorf("Signup failed (Expected %v, got %v)!", http.StatusOK, res.StatusCode)
		}

		if err := res.Body.Close(); err != nil {
			log.Printf("Failed to close response body: %v", err.Error())
		}
	})
}

func TestSignin(t *testing.T) {

	const endpoint = "/api/login"

	t.Cleanup(cleanup)

	t.Run("Should sign into jdoe", func(t *testing.T) {
		form := url.Values{}
		form.Add("username", username1)
		form.Add("password", password1)
		//form.Add("login", "Login")

		request := httptest.NewRequest(http.MethodPost, endpoint, strings.NewReader(form.Encode()))
		request.Form = form

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		//Get the response and read it as a byte array.
		res := recorder.Result()
		defer res.Body.Close()

		printDatabase()

		if res.StatusCode != http.StatusOK {
			t.Errorf("Signin failed (Expected %v, got %v)!", http.StatusOK, res.StatusCode)
		}

		if err := res.Body.Close(); err != nil {
			log.Printf("Failed to close response body: %v", err.Error())
		}

	})
}

func cleanup() {
	Database.Close()

	log.Printf("Use Status: %v", Database.Database.Stats().InUse)

	if err := os.Remove("test.db"); err != nil {
		log.Fatalf("Remove the file yourself ya bum! %v", err.Error())
	}
}
