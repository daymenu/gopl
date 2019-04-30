package main

import (
	"log"
	"net/http"

	"github.com/gopl/web/controllers"
)

func main() {
	http.HandleFunc("/user", controllers.User)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
