package main

import (
	"fmt"
	"net/http"

	Handle "GroupieTracker/FuncHandle"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/static", Handle.StaticHandle)
	http.HandleFunc("/", Handle.HomeHandle)
	http.HandleFunc("/Artist", Handle.ArtistsHandle)
	http.HandleFunc("/Search", Handle.SearchHandle)
	http.HandleFunc("/filter", Handle.FilterHandle)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
