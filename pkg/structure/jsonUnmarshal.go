package structure

import (
	"encoding/json"
	"net/http"
	"time"
)

func JsonReader() ([]Artists, error) {
	//client := http.Client{Timeout: time.Second * 10}
	//var artists []Artist
	//var relations IndexRelations
	//var concertDates IndexDates
	//var locations IndexLocations
	//
	//art, err := client.Get("https://groupietrackers.herokuapp.com/api/artists")
	//if err != nil {
	//	return nil, err
	//}
	//defer art.Body.Close()
	//
	//rel, err := client.Get("https://groupietrackers.herokuapp.com/api/relation")
	//if err != nil {
	//	return nil, err
	//}
	//defer rel.Body.Close()
	//
	//con, err := client.Get("https://groupietrackers.herokuapp.com/api/dates")
	//if err != nil {
	//	return nil, err
	//}
	//defer con.Body.Close()
	//
	//loc, err := client.Get("https://groupietrackers.herokuapp.com/api/locations")
	//if err != nil {
	//	return nil, err
	//}
	//defer loc.Body.Close()
	//
	//err = json.NewDecoder(art.Body).Decode(&artists)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = json.NewDecoder(con.Body).Decode(&concertDates)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = json.NewDecoder(loc.Body).Decode(&locations)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = json.NewDecoder(rel.Body).Decode(&relations)
	//if err != nil {
	//	return nil, err
	//}
	//
	//for i := range artists {
	//	artists[i].Locations = locations.Locations[i]
	//	artists[i].ConcertDates = concertDates.Dates[i]
	//	artists[i].Concerts = relations.Relations[i]
	//}
	//
	//return artists, err
	netClient := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	rel, err := netClient.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}

	defer rel.Body.Close()

	loc, err := netClient.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}

	defer loc.Body.Close()

	date, err := netClient.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}

	defer date.Body.Close()

	var List []Artists
	json.NewDecoder(resp.Body).Decode(&List)
	if err != nil {
		return nil, err
	}

	var locdates IndexRel
	err = json.NewDecoder(rel.Body).Decode(&locdates)
	if err != nil {
		return nil, err
	}

	var location IndexLoc
	err = json.NewDecoder(loc.Body).Decode(&location)
	if err != nil {
		return nil, err
	}

	var dates IndexDates
	err = json.NewDecoder(date.Body).Decode(&dates)
	if err != nil {
		return nil, err
	}

	for i := range List {
		List[i].Relations = locdates.IndexRel[i]
		List[i].Locations = location.IndexLoc[i]
		List[i].ConcertDates = dates.IndexDates[i]
	}

	return List, nil
}
