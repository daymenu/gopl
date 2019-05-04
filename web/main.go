package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gopl/web/models"

	"github.com/gopl/web/controllers"
)

func main() {
	defer models.Close()
	http.HandleFunc("/user", controllers.User)
	fmt.Println("http server is start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
