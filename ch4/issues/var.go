package main

import "time"

var Domain = "https://api.github.com"

var UserName = "daymenu"
var Repos = "phpcap"

var IssuesURL = ""

type User struct {
	Login    string
	Id       int
	ReposURL string `json:"repos_url"`
	HTMLURL  string `json:"html_url"`
}

type Issues struct {
	Url        string
	ReposURL   string `json:"repository_url"`
	ID, Number int
	HTMLURL    string `json:"html_url"`
	Title      string
	State      string
	Locked     bool
	User       *User
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ClosedAt   time.Time `json:"closed_at"`
	Body       string    //in Markdown format
}
