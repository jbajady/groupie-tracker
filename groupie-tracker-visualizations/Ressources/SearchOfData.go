package Func

import (
	"strconv"
	"strings"
)

// Func  For Get  The  Artists Needs Using  Search
func SearchOfData(text string, SearchArtist []Artest) []Artest {
	checked := false
	if len(text) == 0 {
		SearchArtist = Artists
		return SearchArtist
	}
	if SearchArtist != nil {
		SearchArtist = nil
	}
	for _, Artist := range Artists {
		creationdata := strconv.Itoa(Artist.CreationDate)
		if strings.Contains(strings.ToLower(Artist.Name), text) || strings.Contains(strings.ToLower(Artist.FirstAlbum), text) || strings.Contains(creationdata, text) {
			checked = true
		}
		for _, ch := range Artist.Members {
			if strings.Contains(strings.ToLower(ch), text) {
				checked = true
			}
		}
		for _, locactin := range Artist.Location {
			if strings.Contains(strings.ToLower(locactin), text) {
				checked = true
			}
		}
		for _, data := range Artist.Date {
			if strings.Contains(strings.ToLower(data), text) {
				checked = true
			}
		}
		if checked {
			SearchArtist = append(SearchArtist, Artist)
			checked = false
		}

	}
	return SearchArtist
}
