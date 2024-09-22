package Handle

import (
	"fmt"
	"net/http"
	"text/template"

	Func "GroupieTracker/Ressources"
)

func Information(w http.ResponseWriter, r *http.Request) {
	temple, err := template.ParseFiles("templates/InformtionArtist.html")
	if err != nil {
		fmt.Println(";;;;;;;")
		return
	}
	val := r.FormValue("name1")
	if val==""{

	}
	fmt.Println(val)
	var A []Func.Art
	errr := Func.FetchData(c.Artists, &A)
	// for _, v := range A {
	// 	if v.Name == val {
	// 		err = temple.Execute(w, v)
	// 		if err != nil {
	// 			fmt.Println("ok")
	// 			return
	// 		}
	// 		break

	// 	}
	// }

	if errr != nil {
		fmt.Println("ok")
		return
	}
	err = temple.Execute(w, A[1].CreationDate)
	if err != nil {
		fmt.Println("ok")
		return
	}
}
