package main

import "net/http"

type db struct{}

func main() {
	d := db{}
	http.ListenAndServe(":8080", &d)
}

// ServeHttp hello 服务器
func (d *db) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
	w.Write([]byte(r.URL.Hostname()))
	w.Write([]byte(r.URL.RequestURI()))
}
