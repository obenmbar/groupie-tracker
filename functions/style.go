package groupino

import (
	"net/http"
	"os"
)

// Style serves static files (like CSS, JS, or images) requested by the client.
// It ensures that the requested path is a valid file and not a directory before serving it.
func Style(w http.ResponseWriter, r *http.Request) {
	fileinfo, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		RendError(w, "403 forbidden", http.StatusForbidden)
		return
	}
	if fileinfo.IsDir() {
		RendError(w, "403 forbiden tu ne peux pas acceder a cette file", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
