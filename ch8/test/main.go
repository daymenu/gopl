package main

import (
	"fmt"
	"net/http"
)

type db struct{}

func main() {
	// 最便捷
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello"))
	// })
	// http.ListenAndServe("0.0.0.0:80", nil)

	// 自己定义serverMux
	mux := http.NewServeMux()
	d := new(db)
	mux.HandleFunc("/", http.HandlerFunc(d.ServeHTTP))
	mux.Handle("/hello/", d)
	pattern := "/hello/"
	fmt.Println(pattern[len(pattern)-1])
	fmt.Println('/')
	http.ListenAndServe(":80", mux)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("mux handlefunc"))
}

func (d *db) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
