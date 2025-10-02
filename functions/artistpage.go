package groupino

import (
	"fmt"
	"net/http"
	"strconv"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405-methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	Id, erratoi := strconv.Atoi(id)
	if erratoi != nil {
		fmt.Println("iuezgfuyezgf")
	}
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", Id)
	var artist artist
	err1 := FetchData(url, &artist)
	if err1 != nil {
		fmt.Println(err1)
		http.Error(w, "500 internal server error ,, error en fetch data", http.StatusInternalServerError)
		return
	}
	fmt.Println(artist)
	//locurl:=artist.Locations

}
