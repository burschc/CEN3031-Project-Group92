package httpd

import (
	"log"
	"net/http"
	"runtime"
	"strconv"
)

// PipeError reports an error to the internal console log and the browser console log.
func PipeError(w http.ResponseWriter, err error) {
	_, filename, line, _ := runtime.Caller(1)
	log.Print(err)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error() + "\n" + "in " + filename + " line " + strconv.Itoa(line)))
}
