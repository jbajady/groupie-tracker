package Handle

import (
	"bytes"
	"net/http"
	"text/template"

	Func "GroupieTracker/Ressources"
)

func Filter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filter" {
		ErrorHandle(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		ErrorHandle(w, http.StatusMethodNotAllowed)
		return
	}
	temple, err := template.ParseFiles("templates/Homepage.html")
	if err != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
	filtred, err1 := Func.FilterData(r)
	if err1 != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
	if filtred==nil {
		filtred=Func.Artists
	}
	resulte := Func.SearchResult{
		Artists:      filtred,
		SearchArtist: Func.Artists,
	}
	var buf bytes.Buffer
	err = temple.Execute(&buf, resulte)
	if err != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
	_, er := w.Write(buf.Bytes())
	if er != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
}
