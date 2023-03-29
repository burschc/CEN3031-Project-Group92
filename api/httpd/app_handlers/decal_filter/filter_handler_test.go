package decal_filter

import (
	"encoding/json"
	"github.com/gorilla/mux"
	geojson "github.com/paulmach/go.geojson"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var r *mux.Router
var client *http.Client

// init changes to the top level directory for the api (/api/) for each test and creates a new mux router for each test.
func init() {
	if err := os.Chdir("../../../"); err != nil {
		log.Fatal(err.Error())
	}

	//To avoid circular dependencies, the handler registration has to occur here instead of calling the code that
	//already exists.
	r = mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	DecalFilterHandlers(api)
}

// TestDecalTypesPresent will test the Decal Types handler using decals that exist in
// "https://campusmap.ufl.edu/assets/parking_polys.json". The result should be a json list containing the given number
// of elements.
func TestDecalDevTypesHandler(t *testing.T) {
	const numExpected = 24 //Can be easily changed by the author of the JSON file we are using.
	const endpoint = "/api/filter/dev/decals"

	var res *http.Response

	//Make the http request to the proper endpoint and set up a recorder to write the response.
	t.Run("Should receive a response from the endpoint.", func(t *testing.T) {
		//Create the request and recorder and pass them to the created Gorilla Mux.
		request := httptest.NewRequest(http.MethodGet, endpoint, nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		//Get the response and read it as a byte array.
		res = recorder.Result()

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d, got status code %d.", http.StatusOK, res.StatusCode)
		}
	})

	t.Run("Response should have expected number of elements.", func(t *testing.T) {
		//Write the result value to a temporary variable.
		raw, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf(err.Error())
		}

		//Decompose the result into an array of strings.
		var data []string

		if err = json.Unmarshal(raw, &data); err != nil {
			t.Fatalf(err.Error())
		}

		//Confirm that the length of the array is equal to the number we expect.
		if len(data) != numExpected {
			t.Errorf("Expected %d elements in array, recieved %d.", numExpected, len(data))
		}
	})

	//Close the response body.
	if err := res.Body.Close(); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestDecalTypesHandler(t *testing.T) {
	//Get the number of existing defined decals.
	var numExpected = len(ParkingDecals)
	const endpoint = "/api/filter/decals"

	var res *http.Response

	//Make the http request to the proper endpoint and set up a recorder to write the response.
	t.Run("Should receive a response from the endpoint.", func(t *testing.T) {
		//Create the request and recorder and pass them to the created Gorilla Mux.
		request := httptest.NewRequest(http.MethodGet, endpoint, nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		//Get the response and read it as a byte array.
		res = recorder.Result()

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d, got status code %d.", http.StatusOK, res.StatusCode)
		}
	})

	t.Run("Response should have expected number of elements.", func(t *testing.T) {
		//Write the result value to a temporary variable.
		raw, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf(err.Error())
		}

		//Decompose the result into an array of strings.
		var data []string

		if err = json.Unmarshal(raw, &data); err != nil {
			t.Fatalf(err.Error())
		}

		//Confirm that the length of the array is equal to the number we expect.
		if len(data) != numExpected {
			t.Errorf("Expected %d elements in array, recieved %d.", numExpected, len(data))
		}
	})

	//Close the response body.
	if err := res.Body.Close(); err != nil {
		t.Fatalf(err.Error())
	}
}

// TestFindDecalHandlerPresent tests the find decal handler on decals that are present (exist). each run should return
// a feature collection that is not blank. Note that URLs containing spaces must have spaces replaced with %20 for the
// tests to work.
func TestFindDecalHandlerPresent(t *testing.T) {
	t.Run("Should Find Decal Orange.", func(t *testing.T) {
		const decalInput = "Orange"

		//Create the request and recorder and pass them to the created Gorilla Mux.
		request := httptest.NewRequest(http.MethodGet, "/api/filter/decal/"+decalInput, nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		//Get the response and read it as a byte array.
		res := recorder.Result()
		raw, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf(err.Error())
		}

		//Check the status code to make sure we received a valid response.
		if res.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d, got status code %d", http.StatusOK, res.StatusCode)
		}

		//Get the data as byte[] and turn it into a feature collection for analysis
		fc, err := geojson.UnmarshalFeatureCollection(raw)
		if err != nil {
			t.Fatalf(err.Error())
		}

		//Figure out how long the feature collection is. There should be at least one element.
		if len(fc.Features) == 0 {
			t.Errorf("Expected at least one feature for %v decal, got 0.", decalInput)
		}

		//Close the response body.
		if err = res.Body.Close(); err != nil {
			t.Fatalf(err.Error())
		}
	})

	t.Run("Should Find Decal Disabled Student.", func(t *testing.T) {
		const decalInput = "Disabled%20Student"

		//Create the request and recorder and pass them to the created Gorilla Mux.
		request := httptest.NewRequest(http.MethodGet, "/api/filter/decal/"+decalInput, nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		//Get the response and read it as a byte array.
		res := recorder.Result()
		raw, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf(err.Error())
		}

		//Check the status code to make sure we received a valid response.
		if res.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d, got status code %d", http.StatusOK, res.StatusCode)
		}

		//Get the data as byte[] and turn it into a feature collection for analysis
		fc, err := geojson.UnmarshalFeatureCollection(raw)
		if err != nil {
			t.Fatalf(err.Error())
		}

		//Figure out how long the feature collection is. There should be at least one element.
		if len(fc.Features) == 0 {
			t.Errorf("Expected at least one feature for %v decal, got 0.", decalInput)
		}

		//Close the response body.
		if err = res.Body.Close(); err != nil {
			t.Fatalf(err.Error())
		}
	})
}

// TestFindDecalHandlerAbsent tests the find decal handler on decals that are absent (don't exist). each run should
// return a blank feature collection. Note that URLs containing spaces must have spaces replaced with %20 for the tests
// to work.
func TestFindDecalHandlerAbsent(t *testing.T) {
	t.Run("Should NOT Find Decal Decal All Decals (No Park and Ride).", func(t *testing.T) {
		const decalInput = "All%20Decals%20(No%20Park%20and%20Ride"

		//Create the request and recorder and pass them to the created Gorilla Mux.
		request := httptest.NewRequest(http.MethodGet, "/api/filter/decal/"+decalInput, nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		res := recorder.Result()
		raw, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf(err.Error())
		}

		//Check the status code to make sure we received a valid response (blank feature collection).
		if res.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d, got status code %d", http.StatusOK, res.StatusCode)
		}

		//Get the data as byte[] and turn it into a feature collection for analysis
		fc, err := geojson.UnmarshalFeatureCollection(raw)
		if err != nil {
			t.Fatalf(err.Error())
		}

		//Figure out how long the feature collection is. There should be no elements.
		if len(fc.Features) > 0 {
			t.Errorf("Expected zero features for nonexistent decal '%v', got %d.", decalInput, len(fc.Features))
		}

		//Close the response body.
		if err = res.Body.Close(); err != nil {
			t.Fatalf(err.Error())
		}
	})

	t.Run("Should NOT Find Decal Purple.", func(t *testing.T) {
		const decalInput = "Purple"

		//Create the request and recorder and pass them to the created Gorilla Mux.
		request := httptest.NewRequest(http.MethodGet, "/api/filter/decal/"+decalInput, nil)
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		res := recorder.Result()
		raw, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf(err.Error())
		}

		//Check the status code to make sure we received a valid response (blank feature collection).
		if res.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d, got status code %d", http.StatusOK, res.StatusCode)
		}

		//Get the data as byte[] and turn it into a feature collection for analysis
		fc, err := geojson.UnmarshalFeatureCollection(raw)
		if err != nil {
			t.Fatalf(err.Error())
		}

		//Figure out how long the feature collection is. There should be no elements.
		if len(fc.Features) > 0 {
			t.Errorf("Expected zero features for nonexistent decal '%v', got %d.", decalInput, len(fc.Features))
		}

		//Close the response body.
		if err = res.Body.Close(); err != nil {
			t.Fatalf(err.Error())
		}
	})
}
