package Handle

import (
	"net/http"
	"text/template"
)

// // the error page of the site is being worked on.
func ErrorHandle(w http.ResponseWriter, StatusCodes int) {
	error, er := template.ParseFiles("templtes/Error.html")
	if er != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "internal server error 500", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(StatusCodes)
	va := http.StatusText(StatusCodes)

	er = error.Execute(w, va)
	if er != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "internal server error 500", http.StatusInternalServerError)
		return
	}
}
