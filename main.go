package main

import (
	"fmt"
	groupino "groupino/functions"
)

func main() {
var slice groupino.Data
err := groupino.FetchData("https://groupietrackers.herokuapp.com/api/artists",&slice.Artists)
if err != nil {
	panic(err)
}
fmt.Println(slice)
}