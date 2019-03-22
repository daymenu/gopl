package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:created_at`
	Body      string    //in Markdown format
}

type User struct {
	Id      int    `json:"id"`
	HTMLURL string `json:"html_url"`
	Login   string `json:"login"`
}
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

func main() {
	result, err := SearchIssues([]string{"vue-element-admin"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.TotalCount)
	for _, v := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", v.Number, v.User.Login, v.Title)
	}
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
