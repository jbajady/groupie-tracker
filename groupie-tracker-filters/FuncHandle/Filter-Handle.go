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

	filtred := Func.FilterData(r)
	var buf bytes.Buffer
	err = temple.Execute(&buf, filtred)
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
