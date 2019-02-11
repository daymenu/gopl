package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type dollars float32

type good struct {
	Good  string
	Price dollars
}
type database []struct {
	good
}

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func main() {
	db := database{{"shoes", 50}, {"socks", 5}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/edit", db.edit)
	http.HandleFunc("/save", db.save)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile(strings.Join([]string{"ch7", "ex11", "list.html"}, string(os.PathSeparator)))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
	}
	htmllist := template.Must(template.New("list").Parse(string(html)))

	if err = htmllist.Execute(w, struct{ Goods database }{Goods: db}); err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	goodName := r.URL.Query().Get("item")
	for _, item := range db {
		if item.Good == goodName {
			html, err := ioutil.ReadFile(strings.Join([]string{"ch7", "ex11", "read.html"}, string(os.PathSeparator)))
			if err != nil {
				http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
				return
			}
			htmlread := template.Must(template.New("read").Parse(string(html)))

			if err = htmlread.Execute(w, item); err != nil {
				http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
			}
			return
		}
	}
	http.Error(w, "没有找到该商品", http.StatusNotFound)
}

func (db database) edit(w http.ResponseWriter, r *http.Request) {
	goodName := r.URL.Query().Get("item")
	for _, item := range db {
		if item.Good == goodName {
			html, err := ioutil.ReadFile(strings.Join([]string{"ch7", "ex11", "update.html"}, string(os.PathSeparator)))
			if err != nil {
				http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
				return
			}
			htmlread := template.Must(template.New("read").Parse(string(html)))

			if err = htmlread.Execute(w, item); err != nil {
				http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
			}
			return
		}
	}
	http.Error(w, "没有找到该商品", http.StatusNotFound)
}
func (db database) add(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile(strings.Join([]string{"ch7", "ex11", "create.html"}, string(os.PathSeparator)))
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
		return
	}
	htmladd := template.Must(template.New("read").Parse(string(html)))

	if err = htmladd.Execute(w, nil); err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
	}
}
func (db database) save(w http.ResponseWriter, r *http.Request) {
	goodName := r.URL.Query().Get("item")
	newName := r.FormValue("goodname")
	sprice := r.FormValue("price")
	price, err := strconv.ParseFloat(strings.TrimLeft(sprice, "$"), 10)
	if err != nil {
		fmt.Fprint(w, "请输入正确的价格", err)
		return
	}
	for i, item := range db {
		if item.Good == goodName {
			db[i].Good = newName
			db[i].Price = dollars(price)
			db.list(w, r)
			return
		}
	}
}
