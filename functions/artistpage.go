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

// ArtistPage handles the display of a specific artist's details page.
// It fetches data from the Groupie Tracker API based on the artist ID provided in the URL,
// retrieves related information (locations, concert dates, relations),
// and renders them using the "pageartist.html" template.
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	// Check if the HTTP request method is GET; if not, return a 405 error
	if r.Method != http.MethodGet {
		RendError(w, "405-methode not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the artist ID from the URL query and validate it
	id := r.URL.Query().Get("id")

	Id, erratoi := strconv.Atoi(id)

	if erratoi != nil || Id < 1 || Id > 52 {
		RendError(w, "404 not found id ", http.StatusNotFound)
		return
	}

	// Fetch the artist data from the API using the given ID
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", Id)

	var artist Artist
	err1 := FetchData(url, &artist)
	if err1 != nil {
		RendError(w, "500 internal server error , error en fetch data", http.StatusInternalServerError)
		return
	}

	// Fetch related data: locations, concert dates, and relations
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

	// Combine all the data into one global struct for the template
	GlobalStruct := Global{
		Artist:   artist,
		Date:     Concertdate,
		Location: Locations,
		Relation: Relations,
	}

	// Parse the artist HTML template file
	tempp, err := template.ParseFiles("templates/pageartist.html")
	if err != nil {
		RendError(w, "500 internale server error en parse file artist ", http.StatusInternalServerError)
		return
	}

	// Execute the template with the data and send it to the HTTP response
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
