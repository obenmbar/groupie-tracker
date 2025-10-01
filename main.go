package main

import (
	groupino "groupino/functions"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",groupino.HomePage)
	http.HandleFunc("/artist", groupino.ArtistPage)
	http.HandleFunc("/style",groupino.Style)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)					
}
}