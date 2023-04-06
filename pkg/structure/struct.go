package structure

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    Locations
	ConcertDates ConcertDates
	Relations    Relations
}

type IndexRel struct {
	IndexRel []Relations `json:"index"`
}

type IndexLoc struct {
	IndexLoc []Locations `json:"index"`
}

type IndexDates struct {
	IndexDates []ConcertDates `json:"index"`
}

type Relations struct {
	Id        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}

type Locations struct {
	Id           int      `json:"id"`
	Locations    []string `json:"locations"`
	ConcertDates ConcertDates
}

type ConcertDates struct {
	Id           int      `json:"id"`
	ConcertDates []string `json:"dates"`
}
