package Handle

import (
	"fmt"
	"net/http"
	"text/template"

	Func "GroupieTracker/Ressources"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		ErrorHandle(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		ErrorHandle(w, http.StatusMethodNotAllowed)
		return
	}
	temple, err := template.ParseFiles("templates/Homepage.html")
	if err != nil {
		fmt.Println("ok")
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
	errr := Func.FetchData("https://groupietrackers.herokuapp.com/api/artists", &Func.Artists)
	if errr != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
	err = temple.Execute(w, Func.Artists)
	if err != nil {
		fmt.Println("ok")
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
}
