package Handle

import (
	"fmt"
	"net/http"
	"text/template"

	Func "GroupieTracker/Ressources"
)

func Home(w http.ResponseWriter, r *http.Request) {
	temple, err := template.ParseFiles("templates/Homepage.html")
	if err != nil {
		return
	}
	errr := Func.FetchData("https://groupietrackers.herokuapp.com/api/artists", &Func.Artists)
	if errr != nil {
		fmt.Println(";;;;;;;")
		return
	}
	err = temple.Execute(w, Func.Artists)
	if err != nil {
		fmt.Println("ok")
		return
	}
}
