package groupino

import (
	"html/template"
	"net/http"
)

const url string = "https://groupietrackers.herokuapp.com/api/artists"

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "4000 bad request en methode", http.StatusBadRequest)
		return
	}
	temp, err := template.ParseFiles("templates/pagehome.html")
	if err != nil {
		http.Error(w, "500 internale server error en parsfile", http.StatusInternalServerError)
		return
	}
	var Slicedata Data
	err1 := FetchData(url, &Slicedata.Artists)
	if err1 != nil {
		http.Error(w, "500 internal server error ,, error en fetch data", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, Slicedata.Artists)
}
