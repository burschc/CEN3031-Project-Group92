package httpd

import (
	"log"
	"math"
	"os"
	"strconv"
	"testing"
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
	var expectedSize int64 = 2060 //acquired from the JSON file header.

	GetJSONFromURL(jsonURL, filename)

	//Check to make sure that the file is in the json cache path.
	info, err := os.Stat(JsonCachePath + filename)
	if err != nil {
		t.Fatalf(err.Error())
	}

	//Check to make sure that the file has the right amount of information.
	if info.Size() != expectedSize {
		//Check to make sure that the file is at least greater than 0 bytes.
		if info.Size() == 0 {
			t.Fatalf("Size of JSON file is 0 bytes when expected size was not 0 bytes!")
		}

		t.Logf("Warning: size of JSON file does not match content header, but it is more than 0 bytes.")
		t.Logf("size difference is " + strconv.Itoa(int(math.Abs(float64(info.Size()-expectedSize)))))
	}

}

// TestNBPXML pulls an XML file from the National Bank of Poland. This test *should* fail since it is an XML file.
func TestNBPXML(t *testing.T) {
	t.Cleanup(cleanup)

	jsonURL := "https://api.nbp.pl/api/exchangerates/tables/A?format=xml"
	filename := "temp.json"

	//Grab the file
	GetJSONFromURL(jsonURL, filename)

	//The file should not exist because the function should not have created it.
	_, err := os.Stat(JsonCachePath + filename)
	if err == nil {
		log.Fatalf("GetJSONFromURL should not have copied the XML file to the cache!")
	}
}

// TestGoogleHTML pulls the Google homepage as a html file. This test should fail.
func TestGoogleHTML(t *testing.T) {
	t.Cleanup(cleanup)

	jsonURL := "https://www.google.com/"
	filename := "temp.json"

	//Grab the file
	GetJSONFromURL(jsonURL, filename)

	//The file should not exist because the function should not have created it.
	_, err := os.Stat(JsonCachePath + filename)
	if err == nil {
		log.Fatalf("GetJSONFromURL should not have copied the XML file to the cache!")
	}
}

// cleanup Attempts to remove the temporary json cache directory.
func cleanup() {
	if err := os.RemoveAll("cache"); err != nil {
		log.Print(err)
		log.Fatal("You will have to delete the testing cache/json directory yourself")
	}
}
