package main

import (
	"fmt"
	"log"
	"net/http"

	"gopl/web/models"

	"gopl/web/controllers"
)

func main() {
	defer models.Close()
	http.HandleFunc("/user", controllers.User)
	fmt.Println("http server is start")
	mux := http.NewServeMux()
	mux.HandleFunc("/muser", controllers.User)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
