package Func

var (
	Artists        []Artest
	Relation RelationST
	Location  Locations
	Date          Dates
)


type Artest struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}
type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type RelationST struct {
	DatesLocation map[string][]string `json:"datesLocations"`
}
type DataFinal struct {
	Artiste        Artest
	Relation RelationST
	Location  Locations
	Date          Dates
}
