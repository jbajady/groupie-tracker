package Handle

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	Func "GroupieTracker/Ressources"
)

func Information(w http.ResponseWriter, r *http.Request) {
	temple, err := template.ParseFiles("./templates/InformtionArtist.html")
	if err != nil {
		fmt.Println(";;;;;;;")
		return
	}
	val := r.FormValue("name1")
	IdArtest, err := strconv.Atoi(val)
	if err != nil {
		fmt.Print("atoi")
		return
	}
	Func.GenriateData(IdArtest)

	DATA := Func.DataFinal{
		Artiste:  Func.Artists[IdArtest-1], // Artist data
		Relation: Func.Relation,            // Relation data
		Location: Func.Location,
		Date:     Func.Date,
	}
	fmt.Println(DATA.Date)
	err = temple.Execute(w, DATA)
	if err != nil {
		fmt.Println("ok")
		return
	}
}
