package groupino

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type Data struct {
	Artists []Artist
}
var lenn *int
type Location struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}
type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Relation struct {
	Id           int              `json:"id"`
	Daterelation   map[string][]string `json:"datesLocations"`
}
type Global struct {
	Artist   Artist
	Location Location
	Date     Date
	Relation Relation
}

type Errorpage struct {
	Mesege string
	Status int 
}
