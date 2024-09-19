package main

import (
	"net/http"

	Handle "GroupieTracker/FuncHandle"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handle.Home)
	http.ListenAndServe(":8080", mux)
}
