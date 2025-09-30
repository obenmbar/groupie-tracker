package groupino

import (
	"net/http"
	"os"
)

func Style(w http.ResponseWriter, r *http.Request) {
	fileinfo, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		http.NotFound(w, r)
		return
	}
}
