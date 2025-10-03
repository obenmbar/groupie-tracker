package groupino

import (
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
		http.Error(w, "405-methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	Id, erratoi := strconv.Atoi(id)
	fmt.Println(Id)
	if erratoi != nil || Id < 1 || Id > 52 {
		http.Error(w, "404 not found id ", http.StatusNotFound)
		return
	}
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", Id)
	var artist Artist
	err1 := FetchData(url, &artist)
	if err1 != nil {
		http.Error(w, "500 internal server error ,, error en fetch data", http.StatusInternalServerError)
		return
	}

	location := artist.Locations
	concertdate := artist.ConcertDates
	relatio := artist.Relations
	fmt.Println(relatio)
	err2 := FetchData(location, &Locations)
	if err2 != nil {
		http.Error(w, "500 internal server error en fetch data de location", http.StatusInternalServerError)
		return
	}
	err3 := FetchData(concertdate, &Concertdate)
	if err3 != nil {
		http.Error(w, "500 internal server error en fetch data de concert date ", http.StatusInternalServerError)
		return
	}
	err4 := FetchData(relatio, &Relations)
	if err4 != nil {
		http.Error(w, "500 internal server error en fetch data de relation", http.StatusInternalServerError)
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
		fmt.Println("aaaa")
		http.Error(w, "500 internale server error en parse file artist ", http.StatusInternalServerError)
		return
	}

	err = tempp.Execute(w, GlobalStruct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 internale server error en execute artist ", http.StatusInternalServerError)
		return
	}
}
