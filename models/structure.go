package structure

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	Creationdate int      `json:"creationDate"`
	Firstalbum   string   `json:"firstAlbum"`
}

type Relation struct {
	Id           int                 `json:"id"`
	DateLocation map[string][]string `json:"datesLocations"`
}

type Concert struct {
    Location string
    Date     string
}
