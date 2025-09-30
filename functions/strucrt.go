package groupino

type artist struct {
	ID           int
	name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type Data struct {
	Artists []artist
}
type location struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}
type date struct {
	Id    int      `json:"id"`
	dates []string `json:"dates"`
}
type relation struct {
	Id           int              `json:"id"`
	daterelation map[string][]int `json:"datesLocations"`
}
