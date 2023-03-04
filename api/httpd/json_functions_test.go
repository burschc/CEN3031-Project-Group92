package httpd

import (
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

// init creates a temporary json cache directory for the json function tests.
// This is because golang tests run in the same directory as the test package and changing the working directory is
// a painful ordeal.
func init() {
	if err := os.MkdirAll(JsonCachePath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

// TestGitHubJSON pulls a JSON file from the GitHub api. It checks for JSON in the content header and verifies if the
// file was completely downloaded. It will then delete the file from the cache in the cleanup stage.
func TestGitHubJSON(t *testing.T) {

	t.Cleanup(cleanup)

	//Grab the JSON file.
	jsonURL := "https://api.github.com/users/yomole/repos"
	filename := "temp.json"

	t.Run("Should be able to get the JSON file", func(t *testing.T) {
		if err := GetJSONFromURL(jsonURL, filename); err != nil {
			t.Fatalf("Could not grab JSON!")
		}
	})

	//Check to make sure that the file is in the json cache path.
	t.Run("Should be in the cache path", func(t *testing.T) {
		var err error
		if _, err = os.Stat(JsonCachePath + filename); err != nil {
			t.Fatalf(err.Error())
		}
	})

	t.Run("Should have the right amount of information", func(t *testing.T) {

		var expectedSize int64

		t.Run("Should find the uncompressed expected size from the header", func(t *testing.T) {
			//Disable compression on the response to avoid a content length of -1 from gzip encoded JSON.
			noCompClient := &http.Client{
				Transport: &http.Transport{
					DisableCompression: true,
				},
			}

			get, err := noCompClient.Get(jsonURL)
			if err != nil {
				t.Fatalf(err.Error())
			}

			defer get.Body.Close()

			expectedSize = get.ContentLength

			if err = get.Body.Close(); err != nil {
				t.Fatalf(err.Error())
			}

		})

		info, err := os.Stat(JsonCachePath + filename)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if info.Size() != expectedSize {
			if info.Size() == 0 {
				t.Fatalf("Size of compressed JSON file is 0 bytes when expected size was not 0 bytes!")
			}

			t.Logf("Warning: size of compressed JSON file does not match content header, but it is more than 0 bytes.")
			t.Logf("File is " + strconv.Itoa(int(info.Size())) + " bytes long.")
			t.Logf("Expected size is " + strconv.Itoa(int(info.Size())) + " bytes long.")
			t.Logf("Zize difference is " + strconv.Itoa(int(math.Abs(float64(info.Size()-expectedSize)))) + " bytes.")
			t.Fail()
		}
	})
}

// TestNBPXML pulls an XML file from the National Bank of Poland. This test *should* fail since it is an XML file.
func TestNBPXML(t *testing.T) {
	t.Cleanup(cleanup)

	jsonURL := "https://api.nbp.pl/api/exchangerates/tables/A?format=xml"
	filename := "temp.json"

	//Grab the file
	t.Run("Should NOT be able to get the JSON file", func(t *testing.T) {
		if err := GetJSONFromURL(jsonURL, filename); err == nil {
			t.Fatalf("Grabbed non-JSON file!")
		}
	})

	//The file should not exist because the function should not have created it.
	t.Run("Should NOT be in the cache path", func(t *testing.T) {
		var err error
		if _, err = os.Stat(JsonCachePath + filename); err == nil {
			t.Fatalf("GetJSONFromURL should not have copied the XML file to the cache!")
		}
	})
}

// TestGoogleHTML pulls the Google homepage as a html file. This test should fail.
func TestGoogleHTML(t *testing.T) {
	t.Cleanup(cleanup)

	jsonURL := "https://www.google.com/"
	filename := "temp.json"

	//Grab the file
	t.Run("Should NOT be able to get the JSON file", func(t *testing.T) {
		if err := GetJSONFromURL(jsonURL, filename); err == nil {
			t.Fatalf("Grabbed non-JSON file!")
		}
	})

	//The file should not exist because the function should not have created it.
	t.Run("Should NOT be in the cache path", func(t *testing.T) {
		var err error
		if _, err = os.Stat(JsonCachePath + filename); err == nil {
			t.Fatalf("GetJSONFromURL should not have copied the XML file to the cache!")
		}
	})
}

// TestLotsFC tests the FeatureCollection conversion function on the parking lots json file.
func TestLotsFC(t *testing.T) {
	t.Cleanup(cleanup)

	jsonURL := "https://campusmap.ufl.edu/assets/parking_polys.json"
	filename := "parking_lots.json"

	t.Run("Should be able to get the JSON file", func(t *testing.T) {
		if err := GetJSONFromURL(jsonURL, filename); err != nil {
			t.Fatalf("Could not grab JSON!")
		}
	})

	t.Run("Should be able to create a feature collection with some amount of elements", func(t *testing.T) {
		fc := FileToFC(filename)

		if len(fc.Features) == 0 {
			t.Fatalf("Number of features is 0 when it shouldn't!")
		}
	})
}

// TestLotsFCNoExist tests the FeatureCollection conversion function on a file that does not exist.
func TestLotsFCNoExist(t *testing.T) {

	filename := "programmingSkills.json"

	t.Run("File should NOT exist", func(t *testing.T) {
		if _, err := os.Stat(JsonCachePath + filename); err == nil {
			t.Fatalf("File exists when it shouldn't!?!?")
		}
	})

	t.Run("Should be blank FeatureCollection", func(t *testing.T) {
		fc := FileToFC(filename)

		if len(fc.Features) > 0 {
			t.Fatalf("FeatureCollection has features when it should be blank!")
		}
	})
}

// TestIsFresh tests the IsFresh function on two copies of the same JSON file. One of them has had their metadata edited
// so that it appears older than the default update time.
func TestIsFresh(t *testing.T) {
	t.Cleanup(cleanup)
	jsonURL := "https://api.github.com/users/yomole/repos"

	filename1 := "testFresh.json"
	filename2 := "testOld.json"

	t.Run("Should be able to get the JSON files", func(t *testing.T) {
		if err := GetJSONFromURL(jsonURL, filename1); err != nil {
			t.Fatalf("Could not grab JSON!")
		}

		if err := GetJSONFromURL(jsonURL, filename2); err != nil {
			t.Fatalf("Could not grab JSON!")
		}
	})

	//Edit filename2 to be older than the default update time.
	limit, err := time.ParseDuration(DefaultUpdateTime)
	if err != nil {
		log.Fatalf(err.Error())
	}
	oldTime := time.Now().Add(-limit)

	if err = os.Chtimes(JsonCachePath+filename2, oldTime, oldTime); err != nil {
		log.Fatalf(err.Error())
	}

	t.Run(filename1+" should be fresh", func(t *testing.T) {
		if !IsFresh(filename1) {
			t.Fatalf(filename1 + "is not fresh when it should be!")
		}
	})

	t.Run(filename2+"should NOT be fresh", func(t *testing.T) {
		if IsFresh(filename2) {
			t.Fatalf(filename2 + "is fresh when it shouldn't be!")
		}
	})
}

// cleanup Attempts to remove the temporary json cache directory.
func cleanup() {
	if err := os.RemoveAll("cache"); err != nil {
		log.Print(err)
		log.Fatal("You will have to delete the testing cache/json directory yourself")
	}
}
