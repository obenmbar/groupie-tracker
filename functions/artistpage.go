package groupino

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var (
	Locations   Location
	Concertdate Date
	Relations   Relation
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RendError(w, "405-methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	Id, erratoi := strconv.Atoi(id)
	if erratoi != nil || Id < 1 || Id > *lenn {
		RendError(w, "404 not found id ", http.StatusNotFound)

		return
	}
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", Id)
	var artist Artist
	err1 := FetchData(url, &artist)
	if err1 != nil {
		RendError(w, "500 internal server error , error en fetch data", http.StatusInternalServerError)
		return
	}

	location := artist.Locations
	concertdate := artist.ConcertDates
	relatio := artist.Relations

	err2 := FetchData(location, &Locations)
	defer Relations.Clear()
	if err2 != nil {
		RendError(w, "500 internal server error en fetch data de location", http.StatusInternalServerError)
		return
	}
	err3 := FetchData(concertdate, &Concertdate)
	if err3 != nil {
		RendError(w, "500 internal server error en fetch data de concert date ", http.StatusInternalServerError)
		return
	}
	err4 := FetchData(relatio, &Relations)
	if err4 != nil {
		RendError(w, "500 internal server error en fetch data de relation", http.StatusInternalServerError)
		return
	}

	GlobalStruct := Global{
		Artist:   artist,
		Date:     Concertdate,
		Location: Locations,
		Relation: Relations,
	}
	tempp, err := template.ParseFiles("templates/pageartist.html")
	if err != nil {
		RendError(w, "500 internale server error en parse file artist ", http.StatusInternalServerError)
		return
	}

	var buffer bytes.Buffer

	err = tempp.Execute(&buffer, GlobalStruct)
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
