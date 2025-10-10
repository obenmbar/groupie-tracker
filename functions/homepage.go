package groupino

import (
	"bytes"
	"html/template"
	"net/http"
)

const url string = "https://groupietrackers.herokuapp.com/api/artists"

// HomePage handles the main page of the website.
// It fetches all artists from the API and displays them using the "pagehome.html" template.
func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RendError(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		RendError(w, "4000 bad request en methode", http.StatusBadRequest)
		return
	}
	temp, err := template.ParseFiles("templates/pagehome.html")
	if err != nil {
		RendError(w, "500 internale server error, error en parsfile", http.StatusInternalServerError)
		return
	}
	var Slicedata Data

	err1 := FetchData(url, &Slicedata.Artists)
	lene := len(Slicedata.Artists)
	lenn = &lene
	if err1 != nil {
		RendError(w, "500 internal server error , error en fetch data", http.StatusInternalServerError)
		return
	}
	var buffer bytes.Buffer

	err = temp.Execute(&buffer, Slicedata.Artists)
	if err != nil {
		RendError(w, "500 internal server error ", http.StatusInternalServerError)
		return
	}
	_, err = buffer.WriteTo(w)
	if err != nil {
		RendError(w, "500 internal server error ", http.StatusInternalServerError)
		return
	}
}
