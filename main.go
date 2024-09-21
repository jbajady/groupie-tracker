package main

import (
	"fmt"
	"net/http"

	Handle "GroupieTracker/FuncHandle"  
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handle.Home)
	mux.HandleFunc("/Information", Handle.Information)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
