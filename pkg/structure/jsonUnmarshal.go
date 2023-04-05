package structure

import (
	"encoding/json"
	"net/http"
	"time"
)

func JsonReader() ([]Artist, error) {
	client := http.Client{Timeout: time.Second * 10}
	var artists []Artist
	var relations []Relation
	var concertDates []Dates
	var locations []Locations

	art, err := client.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer art.Body.Close()

	rel, err := client.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer rel.Body.Close()

	con, err := client.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}
	defer con.Body.Close()

	loc, err := client.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	defer loc.Body.Close()

	err = json.NewDecoder(art.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(con.Body).Decode(&concertDates)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(loc.Body).Decode(&locations)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(rel.Body).Decode(&relations)
	if err != nil {
		return nil, err
	}

	for i := range artists {
		artists[i].Locations = locations[i]
		artists[i].ConcertDates = concertDates[i]
		artists[i].Concerts = relations[i]
	}

	return artists, err
}
