package model

import "encoding/json"

const (
	UserStatusOnline  = 1
	UserStatusOffline = iota
)

const TableName = "users"

type User struct {
	Id        int    `json:"user_id"`
	Password  string `json:"password"`
	Nickname  string `json:"nick"`
	Sex       string `json:"sex"`
	Header    string `json:"header"`
	LastLogin string `json:"last_login"`
	Status    int    `json:"status"`
}

func UserById(id string) (*User, error) {
	userjson, err := redisCli.HGet(TableName, id).Result()
	if err != nil {
		return nil, err
	}
	var user User
	err = json.Unmarshal([]byte(userjson), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Insert(user *User) (bool, error) {
	userjon, err := json.Marshal(user)
	if err != nil {
		return false, err
	}
	redisCli.HSet(TableName, string(user.Id), userjon)
	return true, nil
}

func Update(user *User) (bool, error) {
	userjon, err := json.Marshal(user)
	if err != nil {
		return false, err
	}
	redisCli.HSet(TableName, string(user.Id), userjon)
	return true, nil
}
