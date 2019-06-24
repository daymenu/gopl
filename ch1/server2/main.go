package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/count", Count)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// Index is web site index page
func Index(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprint(w, r.URL.Path)
}

// Count is web site index page
func Count(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	mu.Unlock()
	fmt.Fprint(w, count)
}
