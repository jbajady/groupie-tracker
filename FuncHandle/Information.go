package Handle

import (
	"fmt"
	"net/http"
	"text/template"

	Func "GroupieTracker/Ressources"
)

func Information(w http.ResponseWriter, r *http.Request) {
	temple, err := template.ParseFiles("templates/Homepage.html")
	data, er := Func.FetchData("https://groupietrackers.herokuapp.com/api")
	if er != nil {
		fmt.Println(";;;;;;;")
		return
	}
	// fmt.Println("+++++++++++++",data)

	// fmt.Fprint(w, data)
	// v, er := json.Marshal(data)
	// var s data
	// er = json.Unmarshal(v, &s)
	// if er != nil {
	// 	return
	// }

	err = temple.Execute(w, data)
	if err != nil {
		fmt.Println("ok")
		return
	}
}
