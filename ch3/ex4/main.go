package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleSurface)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func HandleSurface(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprint(w, surface())
}
