package Handle

import (
	"net/http"
	"strconv"
	"text/template"

	Func "GroupieTracker/Ressources"
)

func ArtistHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ArtistHandle" {
		ErrorHandle(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		ErrorHandle(w, http.StatusMethodNotAllowed)
		return
	}
	temple, err := template.ParseFiles("./templates/InformtionArtist.html")
	if err != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
	val := r.FormValue("id")

	IdArtest, err := strconv.Atoi(val)
	if err != nil || IdArtest < 1 || IdArtest > 52 {
		ErrorHandle(w, http.StatusNotFound)
		return
	}
	Func.GenriateData(IdArtest)

	DATA := Func.DataFinal{
		Artiste:  Func.Artists[IdArtest-1], // Artist data
		Relation: Func.Relation,            // Relation data
		Location: Func.Location,
		Date:     Func.Date,
	}
	err = temple.Execute(w, DATA)
	if err != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
}
