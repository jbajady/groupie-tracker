package Func

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func FilterData(r *http.Request) ([]Artest, error) {
	var filtred []Artest
	checked := false
	r.ParseForm()
	if len(r.Form)==0{
		return nil, nil
	}
	location := strings.ToLower(strings.TrimSpace(r.Form["text"][0]))
	firstAlbumMin, err := strconv.Atoi(r.Form["FirstAlbumngativ"][0])
	firstAlbumMax, er1 := strconv.Atoi(r.Form["FirstAlbumpositive"][0])
	creationDateMin, er2 := strconv.Atoi(r.Form["Creationnigativ"][0])
	creationDateMax, er3 := strconv.Atoi(r.Form["Creationpositive"][0])
	checkbox := r.Form["Members"]
	if err != nil || er1 != nil || er2 != nil || er3 != nil {
		return nil, fmt.Errorf("jhvjhwrg")
	}
	creationDateMIN := creationDateMin * (-1)
	firstAlbumMIN := firstAlbumMin * (-1)
	for _, Art := range Artists {
		Firstalbu := strings.Split(Art.FirstAlbum, "-")
		FirstAlbume, _ := strconv.Atoi(Firstalbu[2])
		if firstAlbumMIN <= firstAlbumMax {
		    if FirstAlbume>= firstAlbumMIN && FirstAlbume<= firstAlbumMax {
			    checked=true
			}
		} else {
					if FirstAlbume<= firstAlbumMIN && FirstAlbume >= firstAlbumMax {
						checked=true
					}
		        }
	if checked{
		checked=false
		if creationDateMIN < creationDateMax {
				if Art.CreationDate >= creationDateMIN && Art.CreationDate <= creationDateMax {
					checked = true
				}
			} else {
				if Art.CreationDate <= creationDateMIN && Art.CreationDate >= creationDateMax {
					checked = true
				}
			}
				for _, loc := range Art.Location {
		if strings.Contains(strings.ToLower(loc), location) && checked {
			checked = true
		} 
		if len(checkbox) != 0 && checked{
					for _, n := range checkbox {
						nm, er := strconv.Atoi(n)
						if er != nil {
							return nil, er
						}
						if len(Art.Members) == nm {
							checked = true
							break
						}else{
							checked=false
						}
					}
				} 
	      }
		  if checked  {
			filtred = append(filtred, Art)
		checked=false
		}


	}
				

	}

	return filtred, nil
}
