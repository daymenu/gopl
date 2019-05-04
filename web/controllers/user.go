package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gopl/web/models"
)

func Show(w http.ResponseWriter, r *http.Request) {
	user, err := models.GetUserById(2)
	if err != nil {
		log.Fatal(err)
	}
	userJson, err := json.Marshal(user)
	w.Write(userJson)
}

func Post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user models.User
	user.UserName = r.Form.Get("username")
	user.Password = r.Form.Get("password")
	user.Status, _ = strconv.Atoi(r.Form.Get("status"))
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()
	var err error
	user.Id, err = models.AddUser(&user)
	if err != nil {
		log.Fatal(err)
	}
	userJson, _ := json.Marshal(user)
	w.Write(userJson)
}

func Put(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user models.User
	id, _ := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	user.UserName = r.Form.Get("username")
	user.Password = r.Form.Get("password")
	user.Status, _ = strconv.Atoi(r.Form.Get("status"))
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()
	var err error
	_, err = models.UpdateUser(id, &user)
	if err != nil {
		log.Fatal(err)
	}
	userJson, _ := json.Marshal(user)
	w.Write(userJson)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, _ := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	var err error
	_, err = models.DeleteUser(id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(struct {
		Code int
		Msg  string
	}{
		200,
		"删除成功",
	})
}
func User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Show(w, r)
	case "POST":
		Post(w, r)
	case "PUT":
		Put(w, r)
	case "DELETE":
		Delete(w, r)
	}
}
