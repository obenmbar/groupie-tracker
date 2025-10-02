package main

import (
	"fmt"
	groupino "groupino/functions"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",groupino.HomePage)
	http.HandleFunc("/artist/", groupino.ArtistPage)
	http.HandleFunc("/style/",groupino.Style)
	fmt.Println("http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)					
}
}