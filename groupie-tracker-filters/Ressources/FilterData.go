package Func

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func FilterData(r *http.Request) []Artest {
	var filtred []Artest
	checked := false
	r.ParseForm()
	text := r.Form["text"][0]
	firstAlbumN, _ := strconv.Atoi(r.Form["FirstAlbumngativ"][0])
	firstAlbumP, _ := strconv.Atoi(r.Form["FirstAlbumpositive"][0])
	creationDateN, _ := strconv.Atoi(r.Form["Creationnigativ"][0])
	creationDateP, _ := strconv.Atoi(r.Form["Creationpositive"][0])
	members := r.Form["Members"]
	fmt.Println(firstAlbumN*(-1), firstAlbumP)
	fmt.Println(creationDateN*(-1), creationDateP)
	for _, Art := range Artists {
		Firstalbum, _ := strconv.Atoi(Art.FirstAlbum)
		creationDateN = creationDateN * (-1)
		firstAlbumN = firstAlbumN * (-1)
		if (Art.CreationDate >= creationDateN && Art.CreationDate <= creationDateP) &&
			(Firstalbum >= firstAlbumN && Firstalbum <= firstAlbumP) {

			for _, loc := range Art.Location {

				if strings.Contains(loc, text) {
					checked = true
				}
			}
			if checked {
				for _, n := range members {
					nm, _ := strconv.Atoi(n)
					if len(Art.Members) == nm {
						checked = true
					}
				}
			}
		}

		if checked {
			filtred = append(filtred, Art)
			checked = false
		}

	}
	fmt.Println(filtred)
	return filtred
}
