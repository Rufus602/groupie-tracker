package structure

type Artist struct {
	Id           int       `json:"id,omitempty"`
	Image        string    `json:"image,omitempty"`
	Name         string    `json:"name,omitempty"`
	Members      []string  `json:"members,omitempty"`
	CreationDate int       `json:"creationDate,omitempty"`
	FirstAlbum   string    `json:"firstAlbum,omitempty"`
	Locations    Locations `json:"-"`
	ConcertDates Dates     `json:"-"`
	Concerts     Relation  `json:"-"`
}

type Locations struct {
	Id        int      `json:"id,omitempty"`
	Locations []string `json:"locations,omitempty"`
}

type Dates struct {
	Id    int      `json:"id,omitempty"`
	Dates []string `json:"dates,omitempty"`
}
type Relation struct {
	Id             int                 `json:"id,omitempty"`
	DatesLocations map[string][]string `json:"datesLocations,omitempty"`
}
