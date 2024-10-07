package main

import (
	"fmt"
	"net/http"

	Handle "GroupieTracker/FuncHandle"
)

func main() {
	http.HandleFunc("/", Handle.HomeHandle)
	http.HandleFunc("/Artist", Handle.ArtistsHandle)
	http.HandleFunc("/Search", Handle.SearchHandle)
	http.HandleFunc("/filter", Handle.Filter)
	fmt.Println("10.1.15.4:8080")
	http.ListenAndServe(":8080", nil)
}
